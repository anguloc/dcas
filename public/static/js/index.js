; (function () {
  var Dcas = {
    lockStatus: false,
    subDom: $(".sub"),
    urlDom: $("[name='url']"),

    init: function () {
      let that = this
      that.subDom.on("click", that.sub)
      that.bindCopy()
    },
    bindCopy: function(){
      var clipboard = new ClipboardJS('.copy');
      clipboard.on('success', function (e) {
        console.info('Action:', e.action);
        console.info('Text:', e.text);
        console.info('Trigger:', e.trigger);
        // Swal.fire("Copied!", "The text has been copied.", "success");
        Swal.fire({
          title: "Copied!",
          text:"The text has been copied.",
          icon:"success",
          timer:"1500"
        })
        e.clearSelection();
      });

      clipboard.on('error', function (e) {
        console.error('Action:', e.action);
        console.error('Trigger:', e.trigger);
        Swal.fire("Copied Fail!", "", "error");
      });
    },
    sub: function () {
      let that = Dcas
      if (!that.lock()) {
        return
      }
      that.subDom.popover("destroy")
      let url = that.urlDom.val()
      url = url.replace(/(^\s*)|(\s*$)/g, "")
      if (url === undefined || url.length === 0) {
        that.unlock()
        return
      }
      $.post("/", {url:url}, function (resp) {
        that.unlock()
        if (resp === undefined || resp.code === undefined || resp.code != 0) {
          Swal.fire('Error', resp.data ? resp.data : "service error", 'warning');
          return
        }

        Swal.fire({
          title: 'Generated successfully',
          showCancelButton: true,
          confirmButtonText: `Copy`,
          html: `<div class="copy" data-clipboard-text="${resp.data.url}">${resp.data.url}</div>`,
        }).then((res) => {
          if (res.isConfirmed) {
            $(".copy").click()
          }
        })
      })
    },
    lock: function () {
      if (this.lockStatus) {
        return false
      }
      this.lockStatus = true
      this.subDom.button('loading')
      this.urlDom.attr("disabled", true)
      return true
    },
    unlock: function () {
      this.lockStatus = false
      this.subDom.button('reset')
      this.urlDom.attr("disabled", false)
    },
  }


  Dcas.init()
})();
