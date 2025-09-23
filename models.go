package chinaums

type BSCPayReq struct {
	MerchantCode            string       `json:"merchantCode"`            //商户ID，必填
	TerminalCode            string       `json:"terminalCode"`            //终端ID，必填
	TransactionAmount       int64        `json:"transactionAmount"`       //交易金额，必填，单位分
	TransactionCurrencyCode string       `json:"transactionCurrencyCode"` //交易币种，必填，默认156，自动填
	MerchantOrderId         string       `json:"merchantOrderId"`         //商户订单号，必填
	MerchantRemark          string       `json:"merchantRemark"`          //商户备注，选填
	PayMode                 string       `json:"payMode"`                 //支付模式，必填，默认CODE_SCAN
	PayCode                 string       `json:"payCode"`                 //支付条码，必填，支持微信、支付宝、云闪付
	SrcReserved             string       `json:"srcReserved"`             //保留字段，选填
	OrderDesc               string       `json:"orderDesc"`               //订单描述，选填，建议注明门店名称
	StoreId                 string       `json:"storeId"`                 //门店StoreId，必填,UUID，用于自动匹配经纬度
	OperatorId              string       `json:"operatorId,omitempty"`    //操作员ID，选填
	Goods                   []OrderGoods `json:"goods,omitempty"`         //订单商品详情，选填
	SubAppId                string       `json:"subAppId,omitempty"`      //子商户微信公众号ID
	DeviceType              string       `json:"deviceType"`              //设备类型，必填，默认11
	SerialNum               string       `json:"serialNum,omitempty"`     //设备序列号，选填
	Longitude               string       `json:"longitude"`               //门店经度,10位
	Latitude                string       `json:"latitude"`                //门店纬度,10位
	IP                      string       `json:"ip,omitempty"`            //门店IP，与经纬度必填一个
}

type OrderGoods struct {
	GoodsId        string `json:"goodsId"`                  //商品ID，选填
	GoodsName      string `json:"goodsName"`                //商品名称，选填
	Quantity       int64  `json:"quantity"`                 //商品数量，选填，默认1
	Price          int64  `json:"price"`                    //商品单价，必填，单位分
	GoodsCategory  string `json:"goodsCategory"`            //商品分类，选填
	Body           string `json:"body"`                     //商品描述，选填
	Unit           string `json:"unit,omitempty"`           //商品单位，选填
	Discount       int64  `json:"discount"`                 //商品优惠金额，选填，单位分
	SubMerchantId  string `json:"subMerchantId,omitempty"`  //子商户ID，选填
	MerOrderId     string `json:"merOrderId,omitempty"`     //子商户订单号，选填
	SubOrderAmount int64  `json:"subOrderAmount,omitempty"` //子商户订单金额，选填，单位分
}

type BSCPayResp struct {
	ErrCode                       string `json:"errCode"`                       //错误码
	ErrInfo                       string `json:"errInfo"`                       //错误信息
	TransactionTime               string `json:"transactionTime"`               //交易时间，格式HHmmss
	TransactionDateWithYear       string `json:"transactionDateWithYear"`       //交易日期，格式yyyyMMdd
	TransactionAmount             int64  `json:"transactionAmount"`             //交易金额，单位分
	ActualTransactionAmount       int64  `json:"actualTransactionAmount"`       //营销联盟优惠后实际交易金额，单位分
	Amount                        int64  `json:"amount"`                        //实际支付金额
	OrderId                       string `json:"orderId"`                       //银商订单号,需保存
	CardAttr                      string `json:"cardAttr"`                      //银行卡属性
	MchntName                     string `json:"mchntName"`                     //商户名称
	TransactionDate               string `json:"transactionDate"`               //交易日期，格式yyyyMMdd
	SettlementDate                string `json:"settlementDate"`                //结算日期，格式MMdd
	SettlementDateWithYear        string `json:"settlementDateWithYear"`        //结算日期，格式yyyyMMdd
	RetrievalRefNum               string `json:"retrievalRefNum"`               //检索参考号
	SystemTraceNum                string `json:"systemTraceNum"`                //系统跟踪号
	ChnlType                      string `json:"chnlType"`                      //渠道类型
	ThirdPartyDiscountInstrution  string `json:"thirdPartyDiscountInstrution"`  //第三方优惠说明
	ThirdPartyDiscountInstruction string `json:"thirdPartyDiscountInstruction"` //营销联盟优惠说明
	ThirdPartyName                string `json:"thirdPartyName"`                //第三方名称(微信、支付宝、云闪付)
	UserId                        string `json:"userId"`                        //第三方用户ID
	ThirdPartyBuyerId             string `json:"thirdPartyBuyerId"`             //第三方买家ID
	ThirdPartyOrderId             string `json:"thirdPartyOrderId"`             //第三方订单ID
	ThirdPartyPayInformation      string `json:"thirdPartyPayInformation"`      //第三方支付信息
	PromotionAmt                  string `json:"promotionAmt"`                  //营销联盟优惠金额
}

// BSCVoidPayReq 撤消支付请求
type BSCVoidPayReq struct {
	MerchantCode    string `json:"merchantCode"`             //商户ID
	TerminalCode    string `json:"terminalCode"`             //终端ID
	MerchantOrderId string `json:"merchantOrderId"`          //商户订单号
	OriginalOrderId string `json:"originalOrderId"`          //银商订单号
	MerchantRemark  string `json:"merchantRemark,omitempty"` //商户备注，选填
}

type BSCVoidPayResp struct {
	ErrCode                 string `json:"errCode"`                 //错误码
	ErrInfo                 string `json:"errInfo"`                 //错误信息
	TransactionTime         string `json:"transactionTime"`         //交易时间，格式HHmmss
	TransactionDate         string `json:"transactionDate"`         //交易日期，格式yyyyMMdd
	SettlementDate          string `json:"settlementDate"`          //结算日期，格式MMdd
	TransactionDateWithYear string `json:"transactionDateWithYear"` //交易日期，格式yyyyMMdd
	SettlementDateWithYear  string `json:"settlementDateWithYear"`  //结算日期，格式yyyyMMdd
	RetrievalRefNum         string `json:"retrievalRefNum"`         //检索参考号
	ThirdPartyName          string `json:"thirdPartyName"`          //第三方名称(微信、支付宝、云闪付)
	TransactionAmount       int64  `json:"transactionAmount"`       //交易金额，单位分
	CardAttr                string `json:"cardAttr"`                //银行卡属性
	OrderId                 string `json:"orderId"`                 //银商订单号
}

// BSCRefundReq 退款请求
type BSCRefundReq struct {
	MerchantCode      string       `json:"merchantCode"`    //商户ID
	TerminalCode      string       `json:"terminalCode"`    //终端ID
	MerchantOrderId   string       `json:"merchantOrderId"` //商户订单号
	OriginalOrderId   string       `json:"originalOrderId"` //银商订单号
	MerchantRemark    string       `json:"merchantRemark,omitempty"`
	RefundRequestId   string       `json:"refundRequestId"`      //退款请求号,必填，唯一，小于50位
	TransactionAmount int64        `json:"transactionAmount"`    //退款金额,必填，单位分
	Goods             []OrderGoods `json:"goods,omitempty"`      //商品列表，选填
	StoreId           string       `json:"storeId,omitempty"`    //门店StoreId
	OperatorId        string       `json:"operatorId,omitempty"` //操作员ID
	RefundDesc        string       `json:"refundDesc,omitempty"` //退款描述
}

// BSCRefundResp 退款响应
type BSCRefundResp struct {
	ErrCode                 string `json:"errCode"`                       //错误码
	ErrInfo                 string `json:"errInfo"`                       //错误信息
	TransactionTime         string `json:"transactionTime"`               //交易时间，格式HHmmss
	TransactionDate         string `json:"transactionDate"`               //交易日期，格式yyyyMMdd
	SettlementDate          string `json:"settlementDate"`                //结算日期，格式MMdd
	TransactionDateWithYear string `json:"transactionDateWithYear"`       //交易日期，格式yyyyMMdd
	SettlementDateWithYear  string `json:"settlementDateWithYear"`        //结算日期，格式yyyyMMdd
	RetrievalRefNum         string `json:"retrievalRefNum"`               //检索参考号
	ThirdPartyName          string `json:"thirdPartyName"`                //第三方名称(微信、支付宝、云闪付)
	CardAttr                string `json:"cardAttr"`                      //银行卡属性
	RefundInvoiceAmount     string `json:"refundInvoiceAmount"`           //退款发票金额，单位分
	TransactionAmount       int64  `json:"transactionAmount"`             //退款金额，单位分
	RefundPromotionList     string `json:"refundPromotionList,omitempty"` //退款营销联盟优惠列表
	SystemTraceNum          string `json:"systemTraceNum"`                //系统跟踪号
	ChnlType                string `json:"chnlType"`                      //渠道类型
	AcqInstCode             string `json:"acqInstCode"`                   // acquiring institution code
	OrderId                 string `json:"orderId"`                       //银商订单号
	RefundOrderId           string `json:"refundOrderId"`                 //退款订单号
	RefundTargetOrderId     string `json:"refundTargetOrderId"`           //退款目标订单号
}

// BSCQueryReq 支付查询请求
type BSCQueryReq struct {
	MerchantCode    string `json:"merchantCode"`    //商户ID
	TerminalCode    string `json:"terminalCode"`    //终端ID
	MerchantOrderId string `json:"merchantOrderId"` //商户订单号
	OriginalOrderId string `json:"originalOrderId"` //银商订单号
	OriTransDate    string `json:"oriTransDate"`    //原交易日期,格式YYYYMMDD
}

// BSCQueryResp 支付查询响应
type BSCQueryResp struct {
	ErrCode                   string `json:"errCode"`                 //错误码
	ErrInfo                   string `json:"errInfo"`                 //错误信息
	OriginalTransactionTime   string `json:"originalTransactionTime"` //原交易时间，格式HHmmss
	QueryResCode              string `json:"queryResCode"`            //查询结果码
	QueryResDesc              string `json:"queryResDesc"`            //查询结果描述
	OriginalPayCode           string `json:"originalPayCode"`         //原支付条码
	OriginalBatchNo           string `json:"originalBatchNo"`         //原批次号
	OriginalSystemTraceNum    string `json:"originalSystemTraceNum"`  //原系统跟踪号
	OriginalRetrievalRefNum   string `json:"originalRetrievalRefNum"`
	OrigialRetrievalRefNum    string `json:"origialRetrievalRefNum"`
	OriginalTransactionAmount int64  `json:"originalTransactionAmount"` //原交易金额，单位分
	Amount                    int64  `json:"amount"`                    //交易金额，单位分
	RefundedAmount            int64  `json:"refundedAmount"`            //已退款金额，单位分
	RefundAmonunt             int64  `json:"refundAmonunt"`             //退款金额，单位分
	ActualTransactionAmount   int64  `json:"actualTransactionAmount"`   //实际交易金额，单位分
	ThirdPartyName            string `json:"thirdPartyName"`            //第三方名称(微信、支付宝、云闪付)
	ThirdPartyBuyerId         string `json:"thirdPartyBuyerId"`         //第三方买家ID
	ThirdPartyOrderId         string `json:"thirdPartyOrderId"`         //第三方订单号
	ThirdPartyPayInformation  string `json:"thirdPartyPayInformation"`  //第三方支付信息
	OrderId                   string `json:"orderId"`                   //银商订单号
	MerchantOrderId           string `json:"merchantOrderId"`           //商户订单号
	OrderStatus               string `json:"orderStatus"`               //订单状态
	OrderCloseTime            string `json:"orderCloseTime"`            //订单关闭时间
	OriginalSettlementDate    string `json:"originalSettlementDate"`    //原结算日期，格式MMdd
	MchntName                 string `json:"mchntName"`                 //商户名称
	IssInstCode               string `json:"issInstCode"`               //发卡机构代码
}

// BSCRefundQueryReq 退款查询请求
type BSCRefundQueryReq struct {
	MerchantCode    string `json:"merchantCode"`    //商户ID
	TerminalCode    string `json:"terminalCode"`    //终端ID
	MerchantOrderId string `json:"merchantOrderId"` //商户订单号
	OriginalOrderId string `json:"originalOrderId"` //银商订单号
	RefundRequestId string `json:"refundRequestId"` //退款请求号
	OriTransDate    string `json:"oriTransDate"`    //原交易日期,格式YYYYMMDD
}

// BSCRefundQueryResp 退款查询响应
type BSCRefundQueryResp struct {
	ErrCode                 string `json:"errCode"`                 //错误码
	ErrInfo                 string `json:"errInfo"`                 //错误信息
	TransactionTime         string `json:"transactionTime"`         //交易时间，格式HHmmss
	TransactionDate         string `json:"transactionDate"`         //交易日期，格式YYYYMMDD
	SettlementDate          string `json:"settlementDate"`          //结算日期，格式MMdd
	RetrievalRefNum         string `json:"retrievalRefNum"`         //检索参考号
	QueryResCode            string `json:"queryResCode"`            //查询结果码
	PayCode                 string `json:"payCode"`                 //支付条码
	DealDate                string `json:"dealDate"`                //交易日期，格式YYYYMMDD
	DealTime                string `json:"dealTime"`                //交易时间，格式HHmmss
	OriginalAmount          int64  `json:"originalAmount"`          //原交易金额，单位分
	DealType                string `json:"dealType"`                //交易类型
	DealSystemTraceNum      string `json:"dealSystemTraceNum"`      //交易系统跟踪号
	DealRetrievalRefNum     string `json:"dealRetrievalRefNum"`     //交易检索参考号
	BatchNo                 string `json:"batchNo"`                 //批次号
	OriginalTransactionDate string `json:"originalTransactionDate"` //原交易日期，格式YYYYMMDD
	OrigialRetrievalRefNum  string `json:"origialRetrievalRefNum"`
	OriginalSettlementDate  string `json:"originalSettlementDate"` //原结算日期，格式MMdd
	RefundInvoiceAmount     string `json:"refundInvoiceAmount"`    //退款发票金额，单位分
	RefundPromotionList     string `json:"refundPromotionList"`    //退款优惠列表
	ThirdPartyName          string `json:"thirdPartyName"`         //第三方名称(微信、支付宝、云闪付)
	RefundOrderId           string `json:"refundOrderId"`          //退款订单号
	RefundTargetOrderId     string `json:"refundTargetOrderId"`    //退款目标订单号
}

type WxAppPayReq struct {
	RequestTimestamp   string       `json:"requestTimestamp"` //请求时间戳，格式yyyy-MM-dd HH:mm:ss
	MerOrderId         string       `json:"merOrderId"`       //商户订单号，必填，长度32位，字母数字下划线组合
	Mid                string       `json:"mid"`              //商户ID，必填，长度32位，字母数字下划线组合
	Tid                string       `json:"tid"`              //终端ID，必填，长度32位，字母数字下划线组合
	TotalAmount        int64        `json:"totalAmount"`      //订单金额，单位分，必填，长度11位，整数位10位，小数位2位
	SubAppId           string       `json:"subAppId"`         //微信子商户APPID
	TradeType          string       `json:"tradeType"`        //交易类型，必填，值为MINI
	SubOpenId          string       `json:"subOpenId"`        //微信子商户用户openid
	MsgId              string       `json:"msgId"`            //消息ID，必填，长度64位，字母数字下划线组合
	SrcReserve         string       `json:"srcReserve,omitempty"`
	InstMid            string       `json:"instMid,omitempty"` //机构商户号，选填，值MINIDEFAULT
	Goods              []OrderGoods `json:"goods,omitempty"`
	ExpireTime         string       `json:"expireTime,omitempty"` //订单失效时间，格式yyyy-MM-dd HH:mm:ss，选填，默认30分钟后失效
	GoodsTag           string       `json:"goodsTag,omitempty"`
	OrderDesc          string       `json:"orderDesc,omitempty"` //订单描述，选填
	OriginalAmount     int64        `json:"originalAmount"`      //订单金额，单位分，必填，单位为分
	ProductId          string       `json:"productId,omitempty"` //商品ID
	NotifyUrl          string       `json:"notifyUrl,omitempty"` //通知地址，选填
	ReturnUrl          string       `json:"returnUrl,omitempty"` //网页跳转地址，选填
	ShowUrl            string       `json:"showUrl,omitempty"`   //网页跳转地址，选填
	SecureTransaction  string       `json:"secureTransaction,omitempty"`
	UserId             string       `json:"userId,omitempty"`
	LimitCreditCard    string       `json:"limitCreditCard,omitempty"`
	InstallmentNumber  string       `json:"installmentNumber,omitempty"`
	Name               string       `json:"name,omitempty"` //姓名，选填
	Mobile             string       `json:"mobile,omitempty"`
	CertType           string       `json:"certType,omitempty"`
	CertNo             string       `json:"certNo,omitempty"`
	FixBuyer           string       `json:"fixBuyer,omitempty"` //是否固定买家，选填，默认N
	FeeRatio           string       `json:"feeRatio,omitempty"` //手续费率，选填，默认0
	CostSubsidy        string       `json:"costSubsidy,omitempty"`
	PreauthTransaction string       `json:"preauthTransaction,omitempty"`
	ClientIp           string       `json:"clientIp,omitempty"`
}

type SubOrder struct {
	Mid         string `json:"mid "`
	MerOrderId  string `json:"merOrderId "`
	TotalAmount int64  `json:"totalAmount "`
}

type WxAppPayResp struct {
	MerName        string `json:"merName"` //商户名称
	Mid            string `json:"mid"`     //商户ID
	MiniPayRequest struct {
		Package   string `json:"package"`
		AppId     string `json:"appId"`     //微信APPID
		Sign      string `json:"sign"`      //签名
		PartnerId string `json:"partnerId"` //商户号
		PrepayId  string `json:"prepayId"`  //预支付交易会话ID
		NonceStr  string `json:"nonceStr"`  //随机字符串
		TimeStamp string `json:"timeStamp"` //时间戳
	} `json:"miniPayRequest"` //小程序支付请求参数
	SettleRefId       string `json:"settleRefId"`
	Tid               string `json:"tid"`               //终端ID
	TotalAmount       int64  `json:"totalAmount"`       //订单金额，单位分
	QrCode            string `json:"qrCode"`            //二维码
	TargetMid         string `json:"targetMid"`         //目标商户ID
	ResponseTimestamp string `json:"responseTimestamp"` //响应时间戳，格式yyyy-MM-dd HH:mm:ss
	ErrCode           string `json:"errCode"`           //错误码
	ErrMsg            string `json:"errMsg"`            //错误信息
	PrepayId          string `json:"prepayId"`          //预支付交易会话ID
	TargetStatus      string `json:"targetStatus"`      //目标状态
	SeqId             string `json:"seqId"`             //序列ID
	MerOrderId        string `json:"merOrderId"`        //商户订单号
	Status            string `json:"status"`            //状态
	TargetSys         string `json:"targetSys"`         //目标系统
}
