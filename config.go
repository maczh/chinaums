package chinaums

type sdkConfig struct {
	AppId        string `json:"appId"`
	AppKey       string `json:"appKey"`
	SourceId     string `json:"sourceId"`
	EncryptKey   string `json:"encryptKey"`
	ApiHost      string `json:"apiHost"`
	MerchantId   string `json:"merchantId"`
	TerminalId   string `json:"terminalId"`
	MerchantIdCB string `json:"merchantIdCB"`
	TerminalIdCB string `json:"terminalIdCB"`
}

var config = sdkConfig{
	AppId:        "8a81c1dc96cf8e4c01996126a99b12b0",
	AppKey:       "78ed50b460494c889fff3c2def6b5268",
	SourceId:     "3GET",
	EncryptKey:   "DrynX3P3GZQjwbENGYpFemQK6CMxxyx4kcSCp4G4CkzEdG58",
	ApiHost:      "https://api-mop.chinaums.com",
	MerchantId:   "898350100003517",
	TerminalId:   "C5G7YL4R",
	MerchantIdCB: "898350100003516",
	TerminalIdCB: "TVMH1T1D",
}

type gis struct {
	Longitude string `json:"longitude"` //门店经度
	Latitude  string `json:"latitude"`  //门店纬度
}

const (
	URI_BC_PAY              = "/v4/poslink/transaction/pay"
	URI_BC_VOID             = "/v2/poslink/transaction/voidpayment"
	URI_BC_REFUND           = "/v2/poslink/transaction/refund"
	URI_BC_QUERY            = "/v2/poslink/transaction/query"
	URI_BC_REFUND_QUERY     = "/v2/poslink/transaction/query-refund"
	URI_WX_APP_PAY          = "/v1/netpay/wx/unified-order"
	URI_WX_APP_QUERY        = "/v1/netpay/order-query"
	URI_WX_APP_REFUND       = "/v1/netpay/refund"
	URI_WX_APP_REFUND_QUERY = "/v1/netpay/refund-query"
	URI_WX_APP_CLOSE        = "/v1/netpay/close"
	URI_CB_PAY              = "/v1/netpay/bills/get-qrcode"
)

var StoreGISMap = map[string]gis{
	"fe79430f-d455-4bd7-999a-b0cdf4315f17": {
		Longitude: "+119.29129",
		Latitude:  "+26.063591",
	},
	"ae48ee12-f901-4cbb-ba0d-4c7497a26c23": {
		Longitude: "+119.29932",
		Latitude:  "+26.087781",
	},
	"0c0a703c-1285-4d6f-bc31-0ff93bb11ca3": {
		Longitude: "+119.29181",
		Latitude:  "+26.062159",
	},
}
