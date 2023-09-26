package wx_msg

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "strings"
)

// SendAlertToWechat 发送企业微信消息
func send(wxKey, message string) error {
    wechatURL := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="

    msg := map[string]interface{}{
        "msgtype": "markdown",
        "markdown": map[string]string{
            "content": message,
        },
    }

    messageData, err := json.Marshal(msg)
    if err != nil {
        return err
    }

    keys := strings.Split(wxKey, ",")
    for _, key := range keys {
        wechatRobotURL := wechatURL + strings.TrimSpace(key)
        resp, err := http.Post(wechatRobotURL, "application/json", bytes.NewBuffer(messageData))
        if err != nil {
            return err
        }
        defer resp.Body.Close()

        if resp.StatusCode != http.StatusOK {
            return fmt.Errorf("HTTP 请求失败，状态码：%d", resp.StatusCode)
        }
    }

    return nil
}
