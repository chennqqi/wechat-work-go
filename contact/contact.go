// Package contact 提供通讯录管理相关的接口
//
// 注意: 关于创建成员（客服答复）
//
// 目前只能使用通讯录的secret 获取token进行创建  其他的secret是没有创建成员的权限的
//
// 获取路径：通讯录管理secret。在“管理工具”-“通讯录同步”里面查看（需开启“API接口同步”）
//
// https://work.weixin.qq.com/api/doc#90000/90135/90193
package contact

import (
	"fmt"

	wechatwork "github.com/dfang/wechat-work-go"
)

// Contact 通讯录
type Contact struct {
	App *wechatwork.App
}

// WithApp 返回 Contact的 实例
//
// 所有通讯录相关API 通过此方法返回的实例调用
func WithApp(app *wechatwork.App) *Contact {
	return &Contact{
		App: app,
	}
}

// CreateMember 创建成员详情
//
// https://work.weixin.qq.com/api/doc#90000/90135/90195
func (contact *Contact) CreateMember(req ReqCreateMember) (RespMemberCreate, error) {
	apiPath := "cgi-bin/user/create"
	uri := fmt.Sprintf("%s?access_token=%s", apiPath, contact.App.GetAccessToken())
	var result RespMemberCreate
	err := contact.App.SimplePost(uri, req, &result)
	if err != nil {
		return RespMemberCreate{}, err
	}
	return result, nil
}

// GetMember 获取成员详情
//
// https://work.weixin.qq.com/api/doc#90000/90135/90196
func (contact *Contact) GetMember(userID string) (RespMemberGet, error) {
	apiPath := "/cgi-bin/user/get"
	uri := fmt.Sprintf("%s?access_token=%s&userid=%s", apiPath, contact.App.GetAccessToken(), userID)
	var result RespMemberGet
	err := contact.App.SimpleGet(uri, &result)
	if err != nil {
		return RespMemberGet{}, err
	}
	return result, nil
}

// DeleteMember 删除成员详情
//
// https://work.weixin.qq.com/api/doc#90000/90135/90198
func (contact *Contact) DeleteMember(userID string) (RespCommon, error) {
	apiPath := "/cgi-bin/user/delete"
	uri := fmt.Sprintf("%s?access_token=%s&userid=%s", apiPath, contact.App.GetAccessToken(), userID)
	var result RespCommon
	err := contact.App.SimpleGet(uri, &result)
	if err != nil {
		return RespCommon{}, err
	}
	return result, nil
}

// DeleteMembers 批量删除成员
//
// https://work.weixin.qq.com/api/doc#90000/90135/90199
func (contact *Contact) DeleteMembers(req ReqBatchDeleteMembers) (RespCommon, error) {
	apiPath := "cgi-bin/user/batchdelete"
	uri := fmt.Sprintf("%s?access_token=%s", apiPath, contact.App.GetAccessToken())
	var result RespCommon
	err := contact.App.SimplePost(uri, req, &result)
	if err != nil {
		return RespCommon{}, err
	}
	return result, nil
}

// UpdateMember 更新成员详情
//
// https://work.weixin.qq.com/api/doc#90000/90135/90197
func (contact *Contact) UpdateMember(body Member) (RespCommon, error) {
	apiPath := "/cgi-bin/user/update"
	uri := fmt.Sprintf("%s?access_token=%s", apiPath, contact.App.GetAccessToken())
	var result RespCommon
	err := contact.App.SimplePost(uri, body, &result)
	if err != nil {
		return RespCommon{}, err
	}
	return result, nil
}

// UserIDToOpenID userid转openid
//
// https://work.weixin.qq.com/api/doc#90000/90135/90202
func (contact *Contact) UserIDToOpenID(userID string) (RespOpenIDInfo, error) {
	apiPath := "/cgi-bin/user/convert_to_openid"
	uri := fmt.Sprintf("%s?access_token=%s", apiPath, contact.App.GetAccessToken())
	var result RespOpenIDInfo
	body := struct {
		UserID string `json:"userid"`
	}{
		UserID: userID,
	}
	err := contact.App.SimplePost(uri, body, &result)
	if err != nil {
		return RespOpenIDInfo{}, err
	}
	return result, nil
}

// OpenIDToUserID userid转openid
//
// https://work.weixin.qq.com/api/doc#90000/90135/90202
func (contact *Contact) OpenIDToUserID(openID string) (RespUserIDInfo, error) {
	apiPath := "/cgi-bin/user/convert_to_userid"
	uri := fmt.Sprintf("%s?access_token=%s", apiPath, contact.App.GetAccessToken())
	var result RespUserIDInfo
	body := struct {
		OpenID string `json:"openid"`
	}{
		OpenID: openID,
	}
	err := contact.App.SimplePost(uri, body, &result)
	if err != nil {
		return RespUserIDInfo{}, err
	}
	return result, nil
}

type RespOpenIDInfo struct {
	RespCommon
	OpenID string `json:"openid"`
}

type RespUserIDInfo struct {
	RespCommon
	UserID string `json:"userid"`
}

// TwoFactorAuth 二次验证
//
// https://work.weixin.qq.com/api/doc#90000/90135/90203
func (contact *Contact) TwoFactorAuth(userID string) (RespCommon, error) {
	apiPath := "/cgi-bin/user/authsucc"
	uri := fmt.Sprintf("%s?access_token=%s&userid=%s", apiPath, contact.App.GetAccessToken(), userID)
	var result RespCommon
	err := contact.App.SimpleGet(uri, &result)
	if err != nil {
		return RespCommon{}, err
	}
	return result, nil
}

// InviteMembers 邀请成员
//
// https://work.weixin.qq.com/api/doc#90000/90135/90975
func (contact *Contact) InviteMembers(body interface{}) (ResqInviteMembers, error) {
	apiPath := "/cgi-bin/cgi-bin/batch/invite"
	uri := fmt.Sprintf("%s?access_token=%s", apiPath, contact.App.GetAccessToken())
	var result ResqInviteMembers
	err := contact.App.SimplePost(uri, body, &result)
	if err != nil {
		return ResqInviteMembers{}, err
	}
	return result, nil
}

type ReqInviteMembers struct {
	User  []string `json:"user"`
	Party []string `json:"party"`
	Tag   []string `json:"tag"`
}

type ResqInviteMembers struct {
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
	Invaliduser  string `json:"invaliduser"`
	Invalidparty string `json:"invalidparty"`
	Invalidtag   string `json:"invalidtag"`
}
