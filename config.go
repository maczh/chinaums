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
	URI_QR_PAY              = "/v1/netpay/bills/get-qrcode"
	URI_QR_QUERY            = "/v1/netpay/bills/query"
	URI_QR_REFUND           = "/v1/netpay/bills/refund"

	// 交易码
	TRANS_CODE_BUM_BALANCE_QUERY      = "202017"
	TRANS_CODE_BUM_ORDER_QUERY        = "202018"
	TRANS_CODE_BUM_TRANSFER           = "202002"
	TRANS_CODE_BUM_ALLOCATION         = "202004"
	TRANS_CODE_MERCHANT_BALANCE_QUERY = "202006"
	TRANS_CODE_TRADE_DETAIL_QUERY     = "202007"
	TRANS_CODE_OPERATE_QUERY          = "202008"

	BUM_VER        = "100"
	BUM_CHANNEL_ID = "043"
	BUM_GROUP_ID   = "123456"
	BUM_HOST       = "https://mobl-test.chinaums.com"
	URI_BUM        = "/channel/Business/UnifyMulti/"
)
