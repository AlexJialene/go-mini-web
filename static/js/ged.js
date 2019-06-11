
$('#queryPoint').click(function () {
    var nick = $('#nick').val()
    if (nick == null || nick == ""){
        dialog("<span class=\"nes-text is-error\">请输入昵称！</span>")
    }else{
        $.ajax({
            url:'/queryValue',
            type:'POST',
            data:{"nick":nick},
            dataType:"json",
            success:function (res) {
                if (res.ret == 'success'){
                    dialog("昵称："+res.vo.nick+",贡献值：<span class=\"nes-text is-error\">"+res.vo.value+"</span>")
                }else{
                    dialog("<span class=\"nes-text is-error\">未查询到相关数据</span>")
                }
            },
            error :function (res) {
                dialog('<span class="nes-text is-error">未知错误！</span>')
            }
        })
    }

})

$('#subPoint').click(function () {
    var nick = $('#nick').val()
    var content = $('#content').val()

    if (nick == null || nick == ""){
        dialog('<span class="nes-text is-error">请输入昵称！</span>')
    }else if (content == null || content == ""){
        dialog('<span class="nes-text is-error">请输入内容！</span>')
    }else{
        $.ajax({
            url:'/subContent',
            type:'POST',
            data:{"nick":nick,"content":content},
            dataType:"json",
            success:function (res) {
                if (res.ret == 'success'){
                    dialog("<span class=\"nes-text is-error\">已经收到您的建议！</span>")
                }else{
                    dialog('<span class="nes-text is-error">失败，请重试！</span>')
                }
            },
            error :function (res) {
                dialog('<span class="nes-text is-error">未知错误！</span>')
            }
        })
    }

})

function dialog(str) {
    $('#dialogMsg').html("")
    $('#dialogMsg').html(str)
    $('#dialog').css('display', 'block')
}

$('#closeDialog').click(function () {
    $('#dialogMsg').html("")
    $('#dialog').css('display', 'none')
})

$('#closeAboutDialog').click(function () {
    $('#dialogMsg').html("")
    $('#dialog').css('display', 'none')
    location.reload()
})