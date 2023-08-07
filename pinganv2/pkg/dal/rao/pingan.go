package rao

type Conf struct {
	User   ConfUser   `json:"user" yaml:"user" mapstructure:"user"`
	Header ConfHeader `json:"header" yaml:"header" mapstructure:"header"`
}

type ConfUser struct {
	VehicleNo        string `json:"vehicle_no" yaml:"vehicle_no" mapstructure:"vehicle_no"`
	ContactName      string `json:"contact_name" yaml:"contact_name" mapstructure:"contact_name"`
	ContactTelephone string `json:"contact_telephone" yaml:"contact_telephone" mapstructure:"contact_telephone"`
}

type ConfHeader struct {
	Signature string `json:"signature" yaml:"signature" mapstructure:"signature"`
	SessionID string `json:"session_id" yaml:"session_id" mapstructure:"session_id"`
}

type ReserveReq struct {
	BookingDate         string        `json:"bookingDate"`                                 //预约日期(列表获取): 2023年07月11日 星期二
	BookingTime         string        `json:"bookingTime"`                                 //预约时间段(列表获取): 10:00-11:00 开始时间和结束时间根据列表时间获取时间段
	IDBookingSurvey     string        `json:"idBookingSurvey"`                             //预约id(列表获取): 01ED97C5B7C08930E0633106A8C01FA5
	VehicleNo           string        `json:"vehicleNo"`                                   //车牌号: 京B-ZH523
	ContactName         string        `json:"contactName"`                                 //预约人姓名: 张飞 投保人姓名
	ContactTelephone    string        `json:"contactTelephone"`                            //预约手机号: 166011633xx
	ApplicantIDCard     string        `json:"applicantIdCard"`                             //身份证号: 可以不填
	BusinessName        string        `json:"businessName,default=承保验车"`                   //这个固定的
	StorefrontName      string        `json:"storefrontName,default=摩托车投保预约"`              //这个固定的
	Detailaddress       string        `json:"detailaddress,default=北京市朝阳区世纪财富中心2号楼2层平安门店"` //预约门店地址: 北京市朝阳区世纪财富中心2号楼2层平安门店
	BookingType         int64         `json:"bookingType,default=1"`                       //预约类型: 1 其他不知道
	Storefrontseq       string        `json:"storefrontseq,default=39807"`                 //门店id
	StorefrontTelephone int64         `json:"storefrontTelephone,default=95511"`
	BusinessType        string        `json:"businessType,default=14"` //业务类型 14就好
	BookContent         string        `json:"bookContent"`
	DeptCode            int64         `json:"deptCode,default=39807"`
	ApplicantName       string        `json:"applicantName"`
	BookingSource       string        `json:"bookingSource,default=miniApps"`
	BusinessKey         interface{}   `json:"businessKey"`
	AgentFlag           int64         `json:"agentFlag,default=0"`
	NewCarFlag          int64         `json:"newCarFlag,default=0"`
	NoPolicyFlag        int64         `json:"noPolicyFlag,default=0"`
	InputPolicyNo       string        `json:"inputPolicyNo"`
	Latitude            string        `json:"latitude"`
	Longitude           string        `json:"longitude"`
	OfflineItemList     []interface{} `json:"offlineItemList"`
}

type AppointmentListResp struct {
	Code int64                     `json:"code"`
	Msg  string                    `json:"msg"`
	Data []AppointmentListRespData `json:"data"`
}

type AppointmentListRespData struct {
	StorefrontSeq    string         `json:"storefrontSeq"`
	BookingDate      string         `json:"bookingDate"`
	BusinessType     string         `json:"businessType"`
	TotalBookableNum int64          `json:"totalBookableNum"`
	TotalBookable    int64          `json:"totalBookable"`
	TotalBooked      int64          `json:"totalBooked"`
	BookingRules     []BookingRules `json:"bookingRules"`
}

type BookingRules struct {
	IDBookingSurvey string `json:"idBookingSurvey"`
	StartTime       string `json:"startTime"`
	EndTime         string `json:"endTime"`
	BookableNum     int64  `json:"bookableNum"`
	BookedNum       int64  `json:"bookedNum"`
}
