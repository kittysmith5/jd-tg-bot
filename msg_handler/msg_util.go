package msg_handler

import (
    api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    "log"
    "time"
)

func sendTxtMsg(cid int64, txt string) (sentMsg api.Message) {
    msg := api.NewMessage(cid, txt)
    sentMsg, err := bot.Send(msg)
    if err != nil {
        log.Println("发送文字消息错误： " + err.Error())
    }
    return
}

func sendTxtMsgAndDel(cid int64, txt string) {
    msg := api.NewMessage(cid, txt)
    sentMsg, err := bot.Send(msg)
    if err != nil {
        log.Println("发送文字消息错误： " + err.Error())
    }
    time.Sleep(time.Second * 30)
    delMsg(cid, sentMsg.MessageID)
}

func delMsg(chatID int64, msgID int) {
    delMsgConf := api.DeleteMessageConfig{
        ChannelUsername: "",
        ChatID:          chatID,
        MessageID:       msgID,
    }
    _, _ = bot.Send(delMsgConf)
}
