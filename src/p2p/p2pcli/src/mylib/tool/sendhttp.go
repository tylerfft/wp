package tool

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

func SendHttp(method, url string, reqHeader http.Header, ReqBody string) (
	respHeader http.Header,
	respBody string,
	Status string,
	StatusCode int,
	ret int) {

	var client *http.Client

	Prefix1 := strings.ToUpper(url[:6])

	if strings.HasPrefix(Prefix1, "HTTPS") {
		tlsconfig := &tls.Config{InsecureSkipVerify: true}
		client = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig:     tlsconfig,
				DisableKeepAlives:   true,
				Dial:                DialTimeout10,
				TLSHandshakeTimeout: time.Second * 10},
			Timeout: time.Second * 30}
	} else {
		client = &http.Client{
			Transport: &http.Transport{
				DisableKeepAlives: true,
				Dial:              DialTimeout10},
			Timeout: time.Second * 30}
	}

	reqBody := strings.NewReader(ReqBody) //   strings.NewReader(body) //
	req, _ := http.NewRequest(strings.ToUpper(method), url, reqBody)
	req.Header = reqHeader

	resp, err2 := client.Do(req)

	if resp != nil {
		StatusCode = resp.StatusCode
		Status = resp.Status

		Body, err3 := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err3 != nil {
			fmt.Sprintf("err3", err3)
			ret = -1
			return
		}
		respBody = string(Body)
	}

	if err2 != nil || StatusCode != 200 {
		fmt.Sprintf(` send to: %v  resp, StatusCode=%v `, url, StatusCode)
		ret = -1
		return
	}
	if err2 != nil || resp == nil {
		fmt.Sprintf("err2", err2)
		ret = -1
		return
	}

	respHeader = resp.Header
	return
}
func DialTimeout10(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, time.Second*10)
}

//允许跨域访问 调试使用
func SetHeader(w http.ResponseWriter) {

	w.Header().Add("Access-Control-Allow-Origin", "*")  //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "*") //header的类型
	w.Header().Add("Access-Control-Allow-Headers", "accept")
	w.Header().Add("Access-Control-Allow-Headers", `filename, Content-Type, Content-Range, Content-Disposition,`+
		`Content-Description,x-requested-with, reqUserId, reqUserSession, bizCode, reqBizGroup, `+
		`bizsign, timestamp,signature,ver,sessionid,from `)
	w.Header().Add("Access-Control-Allow-Methods", "POST")           //header的类型
	w.Header().Add("Access-Control-Allow-Methods", "OPTIONS")        //header的类型
	w.Header().Add("content-type", "application/json;charset=utf-8") //返回数据格式是json

}
