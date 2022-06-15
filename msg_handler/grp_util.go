package msg_handler

import (
    api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func isGrp(update *api.Update) bool {
    return isGrpType(update.Message.Chat.Type)
}

func isGrpType(typeStr string) bool {
    if typeStr == "supergroup" || typeStr == "group" {
        return true
    } else {
        return false
    }
}
