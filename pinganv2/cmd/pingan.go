// Package cmd /*
package cmd

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"pinganv2/pkg/dal/rao"
	"pinganv2/pkg/logic/pingan"
	"sync"
	"time"
)

// pinganCmd represents the pingan command
var pinganCmd = &cobra.Command{
	Use:   "pingan",
	Short: "这个是运行pingan预约的入口",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("pingan called")
	//},
	Run: handle,
}

func init() {
	rootCmd.AddCommand(pinganCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pinganCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pinganCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func handle(cmd *cobra.Command, args []string) {

	// 定义一个时间字符串和对应的时间格式
	for {
		now := time.Now()
		t := time.Date(now.Year(), now.Month(), now.Day(), 16, 59, 30, 0, now.Location())
		if now.Unix() > t.Unix() {
			for {
				SelectReserve()
			}
		} else {
			time.Sleep(time.Second * 1)
		}
	}
}

func SelectReserve() {
	log.Println("执行开始...")

	//获取配置信息
	var conf rao.Conf
	if err := viper.Unmarshal(&conf); err != nil {
		panic(fmt.Errorf("unmarshal error config file: %w", err))
	}

	//禁用tls证书
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	}
	client := &http.Client{Transport: tr}

	//死循环遍历获取预约列表数据
	appData, err := pingan.TraverseGetAppointmentList(client)
	if err != nil {
		log.Fatalf("获取信息失败:%v\n", err)
	}

	//读取预约数据
	var wg sync.WaitGroup
	for _, v := range appData.Data {
		log.Printf("appData.Data.v:%v\n", v)
		wg.Add(1)
		go TaskReserve(&wg, client, &v, &conf)
	}
	wg.Wait() // 等待所有组的所有 goroutine 完成任务
	log.Println("执行结束.")
}

// TaskReserve 预约任务
func TaskReserve(wg *sync.WaitGroup, client *http.Client, v *rao.AppointmentListRespData, conf *rao.Conf) {

	// 每个 goroutine 的代码
	var groupWg sync.WaitGroup
	for _, r := range v.BookingRules {

		p := rao.ReserveReq{
			BookingDate:         v.BookingDate,
			BookingTime:         r.StartTime + "-" + r.EndTime,
			IDBookingSurvey:     r.IDBookingSurvey,
			VehicleNo:           conf.User.VehicleNo,        //车牌号 京B-ZH523
			ContactName:         conf.User.ContactName,      //预约人姓名: 张飞
			ContactTelephone:    conf.User.ContactTelephone, //预约手机号: 166011633xx
			BusinessType:        v.BusinessType,
			Storefrontseq:       v.StorefrontSeq,
			ApplicantIDCard:     "",
			BusinessName:        "承保验车",
			StorefrontName:      "摩托车投保预约",
			Detailaddress:       "北京市朝阳区世纪财富中心2号楼2层平安门店",
			BookingType:         1,
			StorefrontTelephone: 95511,
			BookingSource:       "miniApps",
		}

		b, err := json.Marshal(p)
		if err != nil {
			log.Printf("jsonBytes:json.Marshal失败:%v\n", err)
		} else {
			//并发处理
			log.Printf("执行预约并发处理开始... %v\n", p)
			groupWg.Add(1)
			go PackerReserve(&groupWg, client, b, conf)
			log.Printf("执行预约并发处理结束.\n")
		}
	}
	groupWg.Wait() // 等待当前组的所有 goroutine 完成任务
	wg.Done()
	return
}

func PackerReserve(groupWg *sync.WaitGroup, client *http.Client, b []byte, conf *rao.Conf) {

	////时间段可预约数为0时也跳过
	//if r.BookableNum == 0 {
	//	continue
	//}
	//执行请求
	pingan.RunnerGo(client, b, conf)
	groupWg.Done()
	return
}
