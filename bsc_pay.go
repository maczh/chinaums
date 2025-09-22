package chinaums

//商户扫客户的条码/二维码

const (
	CNY = "156"
	// 支付模式
	PayModeScan  = "CODE_SCAN"
	PayModeECash = "E_CASH"
	PayModeSound = "SOUNDWAVE"
	PayModeNFC   = "NFC"
	PayModeFace  = "FACE_SCAN"
)

type BSCPayReq struct {
	MerchantCode            string       `json:"merchantCode"`
	TerminalCode            string       `json:"terminalCode"`
	TransactionAmount       int64        `json:"transactionAmount"`
	TransactionCurrencyCode string       `json:"transactionCurrencyCode"`
	MerchantOrderId         string       `json:"merchantOrderId"`
	MerchantRemark          string       `json:"merchantRemark"`
	PayMode                 string       `json:"payMode"`
	PayCode                 string       `json:"payCode"`
	SrcReserved             string       `json:"srcReserved"`
	OrderDesc               string       `json:"orderDesc"`
	StoreId                 string       `json:"storeId"`
	LimitCreditCard         bool         `json:"limitCreditCard"`
	OperatorId              string       `json:"operatorId"`
	Goods                   []OrderGoods `json:"goods"`
	SubAppId                string       `json:"subAppId"` //子商户微信公众号ID
	DeviceType              string       `json:"deviceType"`
	SerialNum               string       `json:"serialNum"`
}

type OrderGoods struct {
	GoodsId       string `json:"goodsId"`
	GoodsName     string `json:"goodsName"`
	Quantity      int64  `json:"quantity"`
	Price         int64  `json:"price"`
	GoodsCategory string `json:"goodsCategory"`
	Body          string `json:"body"`
	Discount      int64  `json:"discount"`
}

type BSCPayResp struct {
	ErrCode                 string `json:"errCode"`
	ErrInfo                 string `json:"errInfo"`
	TransactionTime         string `json:"transactionTime"`
	TransactionDateWithYear string `json:"transactionDateWithYear"`
	TransactionAmount       int64  `json:"transactionAmount"`
	ActualTransactionAmount int64  `json:"actualTransactionAmount"`
	Amount                  int64  `json:"amount"`
	OrderId                 string `json:"orderId"`
	CardAttr                string `json:"cardAttr"`
	MchntName               string `json:"mchntName"`
}
