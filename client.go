// Package wechatwork 企业微信api的封装
//
// https://work.weixin.qq.com/api/doc#90000/90003/90556
package wechatwork

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"sync"

	"gopkg.in/resty.v1"
)

// WechatWork 企业微信客户端
type WechatWork struct {
	// CorpID 企业 ID，必填
	CorpID string
}

// App 企业微信客户端（分应用）
type App struct {
	*WechatWork

	// CorpSecret 应用的凭证密钥，其实应该叫AgentSecret更好，必填
	CorpSecret string

	// AgentID 应用 ID，必填
	AgentID int64

	tokenMu     *sync.RWMutex
	AccessToken string
	// tokenExpiresIn time.Duration
	// lastRefresh    time.Time

	// refreshTokenRequestChan chan string // chan currentToken
	// refreshTokenResponseChan chan refreshTokenResult // chan {token, err}

	// Token     string `json:"access_token"`
	ExpiresIn int64 `json:"expires_in"`
}

// New 构造一个 WechatWork 对象，需要提供企业 ID
func New(corpID string) *WechatWork {
	return &WechatWork{
		CorpID: corpID,
	}
}

// WithApp 构造本企业下某自建 app 的对象
func (app *WechatWork) WithApp(corpSecret string, agentID int64) *App {
	return &App{
		WechatWork: app,

		CorpSecret: corpSecret,
		AgentID:    agentID,

		tokenMu:     &sync.RWMutex{},
		AccessToken: "",
		// lastRefresh: time.Time{},
	}
}

// NewRestyClient 返回一个resty 的client
func (c *App) NewRestyClient() *resty.Client {
	client := resty.New()
	client.SetDebug(true)
	client.SetHostURL("https://qyapi.weixin.qq.com")
	client.SetDebug(true)
	return client
}

// NewRequest return resty.Request with right url
func (c *App) NewRequest(path string, qs urlValuer, withAccessToken bool) *resty.Request {
	client := resty.New()
	client.SetDebug(true)
	client.SetLogger(os.Stdout)
	client.SetHostURL("https://qyapi.weixin.qq.com")

	values := url.Values{}
	if valuer, ok := qs.(urlValuer); ok {
		values = valuer.IntoURLValues()
	}

	if withAccessToken {
		c.SyncAccessToken()
		// c.SpawnAccessTokenRefresher()
		if c.AccessToken != "" {
			if values.Get("access_token") != "" {
				values.Set("access_token", c.AccessToken)
			} else {
				values.Add("access_token", c.AccessToken)
			}
		}
	}

	url := path + "?" + values.Encode()
	// client.R().URL = url
	req := client.NewRequest()
	req.URL = url
	return req
}

// Get 一切get请求的api调用可使用此方法
func (c *App) Get(path string, qs urlValuer, respObj interface{}, withAccessToken bool) error {
	client := resty.New()
	// client.SetDebug(true)
	client.SetHostURL("https://qyapi.weixin.qq.com")

	values := url.Values{}
	if valuer, ok := qs.(urlValuer); ok {
		values = valuer.IntoURLValues()
	}

	if withAccessToken {
		c.SyncAccessToken()
		// c.SpawnAccessTokenRefresher()
		if c.AccessToken != "" {
			if values.Get("access_token") != "" {
				values.Set("access_token", c.AccessToken)
			} else {
				values.Add("access_token", c.AccessToken)
			}
		}
	}

	url := path + "?" + values.Encode()
	resp, err := client.R().SetResult(respObj).Get(url)
	if err != nil {
		fmt.Fprintln(os.Stdout, resp.Body())
		panic(err)
	}
	return nil
}

// Post 一切Post请求的api调用使用此方法
//
// 企业微信中，删除操作一般都是GET请求，更新操作、批量删除成员是POST请求，没有PUT、PATCH、DELETE
func (c *App) Post(path string, qs urlValuer, body bodyer, respObj interface{}, withAccessToken bool) (interface{}, error) {
	// url := c.composeQyapiURLWithToken(path, req, withAccessToken)
	// urlStr := url.String()
	client := resty.New()
	client.SetDebug(true)
	client.SetHostURL("https://qyapi.weixin.qq.com")

	values := url.Values{}
	if valuer, ok := qs.(urlValuer); ok {
		values = valuer.IntoURLValues()
	}

	if withAccessToken {
		c.SyncAccessToken()
		// c.SpawnAccessTokenRefresher()
		if c.AccessToken != "" {
			if values.Get("access_token") != "" {
				values.Set("access_token", c.AccessToken)
			} else {
				values.Add("access_token", c.AccessToken)
			}
		}
	}

	url := path + "?" + values.Encode()

	b, _ := body.IntoBody()
	// TODO

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(b).
		SetResult(&respObj).
		Post(url)

	if err != nil {
		panic(err)
	}
	// defer resp.Close()

	decoder := json.NewDecoder(bytes.NewReader(resp.Body()))
	err = decoder.Decode(&respObj)
	if err != nil {
		return respObj, err
	}
	return respObj, nil
}
