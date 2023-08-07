package pingan

import (
	"bytes"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io"
	"log"
	"net/http"
	"pinganv2/pkg/biz/consts"
	"pinganv2/pkg/dal/rao"
)

// TraverseGetAppointmentList 循环获取预约列表数据直到有数据为止。
func TraverseGetAppointmentList(client *http.Client) (*rao.AppointmentListResp, error) {
	for {
		log.Println("traverse ing...")
		//获取预约列表数据
		appData, err := GetAppointmentList(client)
		if err != nil {
			log.Fatalf("获取信息失败:%v\n", err)
		}

		if len(appData.Data) == 0 {
			log.Println("获取信息列表失败,数据未空。重载...")
			continue
		}

		log.Println("traverse end...")
		return appData, nil
	}
}

// GetAppointmentList 获取列表
func GetAppointmentList(client *http.Client) (*rao.AppointmentListResp, error) {
	req, err := http.NewRequest("GET", consts.LIST_URI, nil)
	if err != nil {
		return nil, fmt.Errorf("请求AppointmentList失败! err:%v", err)
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 Edg/107.0.1418.23")
	req.Header.Set("sec-ch-ua", `"Edge";v="107", "Chromium";v="107", "Not=A?Brand";v="24"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求AppointmentList失败,Do! err:%v", err)
	}
	defer resp.Body.Close()

	//读取response.body
	bodyByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("请求AppointmentList失败,Do! 读取response.body! err:%v", err)
	}
	var jsons = jsoniter.ConfigCompatibleWithStandardLibrary
	var respData = rao.AppointmentListResp{}
	if err = jsons.Unmarshal(bodyByte, &respData); err != nil {
		log.Printf("请求AppointmentList失败,读取response.body! err:%v\n", err)
		return nil, err
	}
	return &respData, nil
}

// TaskReserve 预约

// RunnerGo 执行预约请求
func RunnerGo(client *http.Client, b []byte, c *rao.Conf) {
	s := bytes.NewBuffer(b)
	req, err := http.NewRequest("POST", consts.YUYUE_URI, s)
	if err != nil {
		log.Printf("请求RunnerGo失败! err:%v\n", err)
		return
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) MicroMessenger/7.0.0(0x17000000) MacWechat/3.6.2(0x13060211) MiniProgramEnv/Mac MiniProgram")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Host", "newretail.pingan.com.cn")
	req.Header.Set("Accept-Language", "zh-CN,zh-Hans;q=0.9")
	req.Header.Set("signature", c.Header.Signature)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("sessionId", c.Header.SessionID)
	req.Header.Set("Origin", "https://newretail.pingan.com.cn")
	req.Header.Set("Referer", "https://newretail.pingan.com.cn/ydt/newretail/")
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("请求RunnerGo失败,Do! err:%v\n", err)
		return
	}
	defer resp.Body.Close()

	//读取response.body
	bodyByte, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("请求RunnerGo失败,Do! 读取response.body! err:%v\n", err)
		return
	}
	var jsons = jsoniter.ConfigCompatibleWithStandardLibrary
	var respData map[string]interface{}
	if err = jsons.Unmarshal(bodyByte, &respData); err != nil {
		log.Printf("请求RunnerGo失败,读取response.body! err:%v\n", err)
		return
	}
	fmt.Printf("请求RunnerGo成功结果:%v\n", respData)
	return
}
