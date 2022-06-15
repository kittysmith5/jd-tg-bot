package http_helper

import (
    "encoding/json"
    "io"
    "io/ioutil"
    "log"
    "net/http"
)

func CheckByX1a0He(cookie string) bool {
    client := http.Client{}
    url := "https://plogin.m.jd.com/cgi-bin/ml/islogin"
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return false
    }
    userAgent := `jdapp;iPhone;9.4.4;14.3;network/4g;Mozilla/5.0 (iPhone; CPU iPhone OS 14_3 like Mac OS X) ` +
        `AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148;supportJDSHWK/1")`
    req.Header.Add("User-Agent", userAgent)
    req.Header.Add("Referer", "https://h5.m.jd.com/")

    req.Header.Add("Cookie", cookie)
    resp, err := client.Do(req)
    if err != nil {
        return false
    }
    defer func(Body io.ReadCloser) {
        err := Body.Close()
        if err != nil {
            log.Println("checkByByX1a0He:", err)
        }
    }(resp.Body)
    bytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return false
    }
    respTxt := string(bytes)
    var jsonMap map[string]string
    err = json.Unmarshal([]byte(respTxt), &jsonMap)
    if err != nil {
        return false
    }
    if jsonMap["islogin"] == "1" {
        return true
    } else {
        return false
    }
}
