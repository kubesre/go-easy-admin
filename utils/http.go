/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/9
*/

package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"go-easy-admin/common/global"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

// 封装一下http请求

func DoRequest(callMethod, url string, header map[string]string, data interface{}) (string, error) {
	var httpClient = http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   time.Second * 2,
				KeepAlive: time.Second * 2,
			}).DialContext,
		},
		Timeout: time.Second * 2,
	}
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	reqMsg, _ := json.Marshal(data)
	payLoad := bytes.NewReader(reqMsg)
	res, err := http.NewRequest(callMethod, url, payLoad)
	if err != nil {
		global.TPLogger.Error("初始化http失败：", err)
		return "", err
	}
	if header != nil && len(header) > 0 {
		for k, v := range header {
			res.Header.Add(k, v)
		}
	}
	res = res.WithContext(ctx)
	response, err := httpClient.Do(res)
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	if err != nil || response.StatusCode != http.StatusOK {
		global.TPLogger.Error("http请求失败：", err, "responseBody: ", string(bodyBytes))
		return "", err
	}
	return string(bodyBytes), nil
}
