package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"dcas/config"
)


type a struct {
	ass string
}

var logger *zap.Logger



func init() {
	logFilePath,err := config.Get("logFilePath")
	if err == nil {
		logger = getLogger(logFilePath, zapcore.DebugLevel, 1, 30, 7, true, "Main")
	}
}

func Info(msg string, args ...interface{}) {

	logger.Info(msg)


}



/**
 * 获取日志
 * filePath 日志文件路径
 * level 日志级别
 * maxSize 每个日志文件保存的最大尺寸 单位：M
 * maxBackups 日志文件最多保存多少个备份
 * maxAge 文件最多保存多少天
 * compress 是否压缩
 * serviceName 服务名
 */
func getLogger(filePath string, level zapcore.Level, maxSize int, maxBackups int, maxAge int, compress bool, serviceName string) *zap.Logger {
	hook := lumberjack.Logger{
		Filename:   filePath,   // 日志文件路径
		MaxSize:    maxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: maxBackups, // 日志文件最多保存多少个备份
		MaxAge:     maxAge,     // 文件最多保存多少天
		Compress:   compress,   // 是否压缩
	}
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)
	//公用编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 编码器配置
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	// 输出
	writeSyncer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook))

	core := zapcore.NewCore(encoder, writeSyncer, atomicLevel)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	filed := zap.Fields(zap.String("type", "default"))
	// 构造日志
	logger := zap.New(core, caller, development, filed)

	return logger
}
