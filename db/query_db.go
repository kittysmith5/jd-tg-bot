package db

import (
    "database/sql"
    "log"
)

func GetCkByTgID(tgID string) string {
    db := GetDB()

    defer func(db *sql.DB) {
        err := db.Close()
        if err != nil {
            log.Println("关闭数据库错误： " + err.Error())
        }
    }(db)

    cookies := ""
    sqlStr := "select `value` from `Envs` where `remarks` = ?;"
    err := db.QueryRow(sqlStr, tgID).Scan(&cookies)
    if err != nil {
        log.Println("get cookie by qq number error:", err)
    }
    return cookies
}
