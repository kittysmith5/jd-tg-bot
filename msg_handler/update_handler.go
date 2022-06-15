package msg_handler

import (
    api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    "jd-bot/db"
    "strconv"
)

func commandHandler(update *api.Update) {
    command := update.Message.Command()
    arg := update.Message.CommandArguments()
    userID := update.Message.From.ID
    chatID := update.Message.Chat.ID
    switch command {
    case "add":
        if arg == "" {
            go sendTxtMsgAndDel(chatID, "格式错误，应该为\n /add pt_key=**********; "+
                "pt_pin=********;\n/add和cookie之间使用一个空格隔开！且一次只能添加一个cookie！")
        } else {
            //fmt.Println(arg)
            if isJDCookie(arg) {
                oldCks := db.GetCkByTgID(strconv.FormatInt(userID, 10))
                if oldCks != "" {
                    newCks := setUpdateCookie(oldCks, arg)
                    if newCks == "duplicated" {
                        go sendTxtMsgAndDel(chatID, "对不起，cookie已存在，不需要添加！")
                    } else {
                        if ckIsInDate(arg) {
                            go sendTxtMsgAndDel(chatID, "cookie有效，正在添加到数据库中...")
                            updateFlag := db.UpdateCkByTgID(strconv.FormatInt(userID, 10), newCks)
                            if updateFlag {
                                go sendTxtMsgAndDel(chatID, "添加或者更新cookie成功！")
                            } else {
                                go sendTxtMsgAndDel(chatID, "对不起，添加或者更新cookie失败！请联系管理员或者稍后重试！")
                            }
                        } else {
                            go sendTxtMsgAndDel(chatID, "对不起，cookie无效或者已经过期！请确定cookie有效后再重新输入！")
                        }
                    }
                } else {
                    if ckIsInDate(arg) {
                        go sendTxtMsgAndDel(chatID, "cookie有效，正在添加到数据库中...")
                        newCks := arg
                        insertFlag := db.InsertCookie(strconv.FormatInt(userID, 10), newCks)
                        if insertFlag {
                            go sendTxtMsgAndDel(chatID, "添加cookie成功！")
                        } else {
                            go sendTxtMsgAndDel(chatID, "对不起，添加cookie失败！请联系管理员或者稍后重试！")
                        }
                    } else {
                        go sendTxtMsgAndDel(chatID, "对不起，cookie无效或者已经过期！请确定cookie有效后再重新输入！")
                    }
                }
            } else {
                go sendTxtMsgAndDel(update.Message.Chat.ID, "对不起，cookie格式错误，应该为pt_key="+
                    "开头，中间有个分号，然后是pt_pin,最后以分号结尾！")
            }
        }
    }
}
