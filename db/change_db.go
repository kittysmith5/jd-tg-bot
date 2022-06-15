package db

import (
    "database/sql"
    "log"
    "time"
)

func InsertCookie(tgID, cookie string) bool {
    db := GetDB()

    defer func(db *sql.DB) {
        err := db.Close()
        if err != nil {
            log.Println("关闭数据库错误： " + err.Error())
        }
    }(db)
    createAt := string(time.Now().Format("2006-01-02 15:04:05"))

    sqlStr := "insert into Envs (`value`,`status`,`name`,`remarks`,`createdAt`,`updatedAt`) values (?,?,?,?,?,?);"
    res, err := db.Exec(sqlStr, cookie, 0, "JD_COOKIE", tgID, createAt, createAt)
    if err != nil {
        log.Println("insert cookie error:", err)
        return false
    }

    affected, err := res.RowsAffected()
    if err != nil {
        return false
    }
    if affected != 1 {
        log.Println("insert cookie error:", err)
        return false
    }

    return true
}

func UpdateCkByTgID(tgID, newCks string) bool {
    db := GetDB()
    defer func(db *sql.DB) {
        err := db.Close()
        if err != nil {
            log.Println("关闭数据库错误： " + err.Error())
        }
    }(db)

    updatedAt := string(time.Now().Format("2006-01-02 15:04:05"))
    //update Envs set `value` = ? where `remarks`=?;
    sqlStr := "update Envs set `value`=? ,`updatedAt`=? where `remarks`=?;"

    res, err := db.Exec(sqlStr, newCks, updatedAt, tgID)
    if err != nil {
        log.Println("update cookie error:", err)
        return false
    }

    affected, err := res.RowsAffected()
    if err != nil {
        return false
    }
    if affected != 1 {
        log.Println("update cookie error with affected num:", err)
        return false
    }
    return true
}
