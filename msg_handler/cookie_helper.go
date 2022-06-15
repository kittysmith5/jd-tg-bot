package msg_handler

import (
    "jd-bot/http_helper"
    "regexp"
    "strings"
)

func isJDCookie(arg string) bool {
    if strings.Contains(arg, "&") {
        return false
    }
    compile, err := regexp.Compile(`^pt_key=.*?;\s*pt_pin=.*?;`)
    if err != nil {
        return false
    }
    return compile.MatchString(arg)
}

func ckIsInDate(cookie string) bool {
    ckInDate := http_helper.CheckByX1a0He(cookie)
    return ckInDate
}

func setUpdateCookie(oldCookies, cookie string) (newCookies string) {
    cookieArr := strings.Split(oldCookies, "&")
    newCookies = ""
    for _, ck := range cookieArr {
        if ck == cookie {
            return "duplicated"
        }
        ckIsOk := http_helper.CheckByX1a0He(ck)
        if ckIsOk {
            newCookies += ck + "&"
        }
    }
    newCookies += cookie
    return newCookies
}
