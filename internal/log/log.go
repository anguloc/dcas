package log

import (
	"dcas/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var logger *zap.Logger

func init() {
	if config.Conf.Log.FilePath == "" {
		return
	}
	logger = getLogger(config.Conf.Log.FilePath, zapcore.DebugLevel, 4, 3, 3, true)
}

func Debug(msg string, args ...interface{}) {
	if logger == nil {
		return
	}
	logger.Debug(msg)
}

func Info(msg string, args ...interface{}) {
	if logger == nil {
		return
	}
	logger.Info(msg)
}

func Error(msg string, args ...interface{}) {
	if logger == nil {
		return
	}
	logger.Error(msg)
}

func Panic(msg string, args ...interface{}) {
	if logger == nil {
		return
	}
	logger.Panic(msg)
}

func Fatal(msg string, args ...interface{}) {
	if logger == nil {
		return
	}
	logger.Fatal(msg)
}

/**
 * 获取日志
 * filePath 日志文件路径
 * level 日志级别
 * maxSize 每个日志文件保存的最大尺寸 单位：M
 * maxBackups 日志文件最多保存多少个备份
 * maxAge 文件最多保存多少天
 * compress 是否压缩
 */
func getLogger(filePath string, level zapcore.Level, maxSize int, maxBackups int, maxAge int, compress bool) *zap.Logger {
	hook := lumberjack.Logger{
		Filename:   filePath,   // 日志文件路径
		MaxSize:    maxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: maxBackups, // 日志文件最多保存多少个备份
		MaxAge:     maxAge,     // 文件最多保存多少天
		Compress:   compress,   // 是否压缩
	}

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 时间格式换下
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder // 调用路径从根目录起

	encoder := zapcore.NewJSONEncoder(encoderConfig) // json 格式
	//encoder := zapcore.NewConsoleEncoder(encoderConfig) // 单行格式

	// 输出
	writeSyncer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook))

	core := zapcore.NewCore(encoder, writeSyncer, atomicLevel)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号 跳过封装层
	development := zap.AddCallerSkip(1)
	// 构造日志
	logger := zap.New(core, caller, development)

	return logger
}
