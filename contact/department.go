package contact

import (
	"fmt"

	. "github.com/dfang/wechat-work-go/models"
)

// ListMembers 获取部门下成员概要
//
// https://work.weixin.qq.com/api/doc#90000/90135/90200
func (contact Contact) ListMembers(departmentID string, fetchChild int) (RespListMembers, error) {
	apiPath := "/cgi-bin/user/simplelist"
	uri := fmt.Sprintf("%s?access_token=%s&department_id=%s&fetch_child=%d", apiPath, contact.App.GetAccessToken(), departmentID, fetchChild)
	var result RespListMembers
	err := contact.App.SimpleGet(uri, &result)
	if err != nil {
		return RespListMembers{}, err
	}
	return result, nil
}

// ListMembers2 获取部门下成员详情
//
// https://work.weixin.qq.com/api/doc#90000/90135/90201
func (contact Contact) ListMembers2(departmentID string, fetchChild int) (RespListMembers2, error) {
	apiPath := "/cgi-bin/user/list"
	uri := fmt.Sprintf("%s?access_token=%s&department_id=%s&fetch_child=%d", apiPath, contact.App.GetAccessToken(), departmentID, fetchChild)
	var result RespListMembers2
	err := contact.App.SimpleGet(uri, &result)
	if err != nil {
		return RespListMembers2{}, err
	}
	return result, nil
}

// CreateDepartment 创建部门
//
// https://work.weixin.qq.com/api/doc#90000/90135/90205
func (contact Contact) CreateDepartment(body ReqCreateDepartment) (RespCreateDepartment, error) {
	apiPath := "cgi-bin/department/create"
	uri := fmt.Sprintf("%s?access_token=%s", apiPath, contact.App.GetAccessToken())
	var result RespCreateDepartment
	err := contact.App.SimplePost(uri, body, &result)
	if err != nil {
		return RespCreateDepartment{}, err
	}
	return result, nil
}

// UpdateDepartment 更新部门
//
// https://work.weixin.qq.com/api/doc#90000/90135/90206
func (contact Contact) UpdateDepartment(body ReqUpdateDepartment) (RespCommon, error) {
	apiPath := "/cgi-bin/department/update"
	uri := fmt.Sprintf("%s?access_token=%s", apiPath, contact.App.GetAccessToken())
	var result RespCommon
	err := contact.App.SimplePost(uri, body, &result)
	if err != nil {
		return RespCommon{}, err
	}
	return result, nil
}

// DeleteDepartment 删除部门
//
// https://work.weixin.qq.com/api/doc#90000/90135/90207
func (contact Contact) DeleteDepartment(departmentID string) (RespCommon, error) {
	apiPath := "/cgi-bin/department/delete"
	uri := fmt.Sprintf("%s?access_token=%s&id=%s", apiPath, contact.App.GetAccessToken(), departmentID)
	var result RespCommon
	err := contact.App.SimpleGet(uri, &result)
	if err != nil {
		return RespCommon{}, err
	}
	return result, nil
}

// ListDepartments 获取部门列表
//
// https://work.weixin.qq.com/api/doc#90000/90135/90208
func (contact Contact) ListDepartments(departmentID string) (RespListDepartments, error) {
	apiPath := "/cgi-bin/department/list"
	uri := fmt.Sprintf("%s?access_token=%s&id=%s", apiPath, contact.App.GetAccessToken(), departmentID)
	var result RespListDepartments
	err := contact.App.SimpleGet(uri, &result)
	if err != nil {
		return RespListDepartments{}, err
	}
	return result, nil
}
