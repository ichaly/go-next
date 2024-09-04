package main

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"os"
	"testing"
)

func TestWechat(t *testing.T) {
	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式
	// 注册消息处理函数
	dispatcher := openwechat.NewMessageMatchDispatcher()
	// 只处理消息类型为文本类型的消息
	dispatcher.OnText(func(ctx *openwechat.MessageContext) {
		msg := ctx.Message
		if msg.IsText() && msg.Content == "ping" {
			//msg.ReplyText("pong")
			file, _ := os.Open("b304a810dcc8d6c607f2ed9a5eb62b7b.mp4")
			defer file.Close()
			msg.ReplyVideo(file)
		}
	})
	bot.MessageHandler = dispatcher.AsMessageHandler()

	// 登陆
	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")
	defer reloadStorage.Close()
	if err := bot.HotLogin(reloadStorage, openwechat.NewRetryLoginOption()); err != nil {
		fmt.Println(err)
		return
	}

	// 获取登陆的用户
	self, err := bot.GetCurrentUser()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 获取所有的好友
	friends, err := self.Friends()
	fmt.Println(friends, err)

	// 获取所有的群组
	groups, err := self.Groups()
	fmt.Println(groups, err)

	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	bot.Block()
}
