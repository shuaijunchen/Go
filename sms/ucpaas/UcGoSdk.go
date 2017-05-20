// UcGoSdk.go
package ucpaas

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// 定义常量
const (
	softVersion = "2014-06-30"             //
	apiUrl      = "https://api.ucpaas.com" //
	appId       = "appid"                  // APPID
	accountSid  = "accountSid"             // accountSid
	authToken   = "authToken"              // authToken
)

var timestamp string

type UcPaaS struct {
	TempSMS templateSMS `json:"templateSMS"`
}

type templateSMS struct {
	AppId      string `json:"appId"`
	TemplateId string `json:"templateId"`
	PhoneNum   string `json:"to"`
	Param      string `json:"param"`
}

type result struct {
	Resp response `json:"resp"`
}

type response struct {
	RespCode        string         `json:"respCode"`
	Failure         int            `json:"failure,omitempty"`
	ReturenTemplSMS returnTemplSMS `json:"returnTemplSMS"`
}

type returnTemplSMS struct {
	CreateDate string `json:"createDate"`
	SMSId      string `json:"smsId`
}

// 初始化信息
func ucinit(phoneNum, tempId string, params ...string) *UcPaaS {
	curTime := time.Now().Format("20060102150405000")
	intCurTime, _ := strconv.Atoi(curTime)
	timestamp = strconv.Itoa(intCurTime)
	catParam := ""
	for index, param := range params {
		if index == 0 {
			catParam = param
		} else {
			catParam = catParam + "," + param
		}
	}

	return &UcPaaS{
		TempSMS: templateSMS{
			AppId:      appId,
			TemplateId: tempId,
			PhoneNum:   phoneNum,
			Param:      catParam,
		},
	}
}

func (res *result) auth(accountSid string) string {
	src := accountSid + ":" + timestamp
	return base64.StdEncoding.EncodeToString([]byte(src))
}

func (res *result) sign(accountSid, authToken string) string {
	sign := accountSid + authToken + timestamp
	return strings.ToUpper(fmt.Sprintf("%x", md5.Sum([]byte(sign))))
}

func (res *result) respResult(url, dataType, method string, data []byte) {
	resp := res.connection(url, dataType, method, data)
	_ = json.Unmarshal([]byte(resp), res)
}

func (res *result) connection(url, dataType, method string, data []byte) string {
	mine := ""
	if dataType == "json" {
		mine = "application/json"
	} else {
		mine = "application/xml"
	}

	client := &http.Client{}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return err.Error()
	}
	req.Header.Add("Accept", mine)
	req.Header.Add("Content-Type", mine)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", res.auth(accountSid))
	req.Header.Add("Content-Length", strconv.Itoa(len(data)))

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return ""
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(respBody)
}

func SendSMS(phoneNum, tempId string, params ...string) (respCode string, err error) {
	var rs result
	uc := ucinit(phoneNum, tempId, params...)
	resouce := "/" + softVersion + "/Accounts/" + accountSid + "/Messages/templateSMS"
	parseUrl, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		return "", errors.New("请求地址出错")
	}
	parseUrl.Path = resouce
	urlStr := fmt.Sprintf("%v?sig=%s", parseUrl, rs.sign(accountSid, authToken))
	data, err := json.MarshalIndent(uc, "", "  ")
	if err != nil {
		return "", err
	}
	fmt.Println(urlStr)
	fmt.Println(string(data))
	rs.respResult(urlStr, "json", "POST", data)
	respCode = rs.Resp.RespCode
	if respCode != "000000" {
		return respCode, errors.New("发送短信失败")
	}
	return "ok", nil
}
