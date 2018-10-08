package midd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type openIdStruct struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
}

func GetOpenId(code string) (openId string, unionId string, sessionKey string) {
	client := &http.Client{}

	wechatUrl = "https://api.weixin.qq.com/sns/jscode2session"
	params := url.Values{
		"appid":      "tm",
		"secret":     "tm",
		"js_code":    code,
		"grant_type": "authorization_code",
	}
	request, err := http.NewRequest("GET", wechatUrl, params)
	if err != nil {
		panic(err)
	}
	response, rep_err := client.Do(reqest)
	defer response.Body.Close()

	if rep_err != nil {
		panic(rep_err)
	}
	bytes, read_err := ioutil.ReadAll(response.Body)
	if read_err != nil {
		panic(read_err)
	}
	ptr = &openIdStruct{}
	err = json.Unmarshal(bytes, ptr)
	if err != nil {
		panic(err)
	}
	return ptr.OpenId, ptr.UnionId, ptr.SessionKey
}
