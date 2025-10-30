package chinaums

type BSCPayReq struct {
	MerchantCode            string       `json:"merchantCode" bson:"merchantCode"`                       //商户ID，必填
	TerminalCode            string       `json:"terminalCode" bson:"terminalCode"`                       //终端ID，必填
	TransactionAmount       int64        `json:"transactionAmount" bson:"transactionAmount"`             //交易金额，必填，单位分
	TransactionCurrencyCode string       `json:"transactionCurrencyCode" bson:"transactionCurrencyCode"` //交易币种，必填，默认156，自动填
	MerchantOrderId         string       `json:"merchantOrderId" bson:"merchantOrderId"`                 //商户订单号，必填
	MerchantRemark          string       `json:"merchantRemark" bson:"merchantRemark"`                   //商户备注，选填
	PayMode                 string       `json:"payMode" bson:"payMode"`                                 //支付模式，必填，默认CODE_SCAN
	PayCode                 string       `json:"payCode" bson:"payCode"`                                 //支付条码，必填，支持微信、支付宝、云闪付
	SrcReserved             string       `json:"srcReserved" bson:"srcReserved"`                         //保留字段，选填
	OrderDesc               string       `json:"orderDesc" bson:"orderDesc"`                             //订单描述，选填，建议注明门店名称
	StoreId                 string       `json:"storeId" bson:"storeId"`                                 //门店StoreId，必填,UUID，用于自动匹配经纬度
	OperatorId              string       `json:"operatorId,omitempty" bson:"operatorId"`                 //操作员ID，选填
	Goods                   []OrderGoods `json:"goods,omitempty" bson:"goods"`                           //订单商品详情，选填
	SubAppId                string       `json:"subAppId,omitempty" bson:"subAppId"`                     //子商户微信公众号ID
	DeviceType              string       `json:"deviceType" bson:"deviceType"`                           //设备类型，必填，默认11
	SerialNum               string       `json:"serialNum,omitempty" bson:"serialNum"`                   //设备序列号，选填
	Longitude               string       `json:"longitude" bson:"longitude"`                             //门店经度,10位
	Latitude                string       `json:"latitude" bson:"latitude"`                               //门店纬度,10位
	IP                      string       `json:"ip,omitempty" bson:"IP"`                                 //门店IP，与经纬度必填一个
}

type OrderGoods struct {
	GoodsId        string `json:"goodsId" bson:"goodsId"`                         //商品ID，选填
	GoodsName      string `json:"goodsName" bson:"goodsName"`                     //商品名称，选填
	Quantity       int64  `json:"quantity" bson:"quantity"`                       //商品数量，选填，默认1
	Price          int64  `json:"price" bson:"price"`                             //商品单价，必填，单位分
	GoodsCategory  string `json:"goodsCategory" bson:"goodsCategory"`             //商品分类，选填
	Body           string `json:"body" bson:"body"`                               //商品描述，选填
	Unit           string `json:"unit,omitempty" bson:"unit"`                     //商品单位，选填
	Discount       int64  `json:"discount" bson:"discount"`                       //商品优惠金额，选填，单位分
	SubMerchantId  string `json:"subMerchantId,omitempty" bson:"subMerchantId"`   //子商户ID，选填
	MerOrderId     string `json:"merOrderId,omitempty" bson:"merOrderId"`         //子商户订单号，选填
	SubOrderAmount int64  `json:"subOrderAmount,omitempty" bson:"subOrderAmount"` //子商户订单金额，选填，单位分
}

type BSCPayResp struct {
	ErrCode                       string `json:"errCode" bson:"errCode"`                                             //错误码
	ErrInfo                       string `json:"errInfo" bson:"errInfo"`                                             //错误信息
	TransactionTime               string `json:"transactionTime" bson:"transactionTime"`                             //交易时间，格式HHmmss
	TransactionDateWithYear       string `json:"transactionDateWithYear" bson:"transactionDateWithYear"`             //交易日期，格式yyyyMMdd
	TransactionAmount             int64  `json:"transactionAmount" bson:"transactionAmount"`                         //交易金额，单位分
	ActualTransactionAmount       int64  `json:"actualTransactionAmount" bson:"actualTransactionAmount"`             //营销联盟优惠后实际交易金额，单位分
	Amount                        int64  `json:"amount" bson:"amount"`                                               //实际支付金额
	OrderId                       string `json:"orderId" bson:"orderId"`                                             //银商订单号,需保存
	CardAttr                      string `json:"cardAttr" bson:"cardAttr"`                                           //银行卡属性
	MchntName                     string `json:"mchntName" bson:"mchntName"`                                         //商户名称
	TransactionDate               string `json:"transactionDate" bson:"transactionDate"`                             //交易日期，格式yyyyMMdd
	SettlementDate                string `json:"settlementDate" bson:"settlementDate"`                               //结算日期，格式MMdd
	SettlementDateWithYear        string `json:"settlementDateWithYear" bson:"settlementDateWithYear"`               //结算日期，格式yyyyMMdd
	RetrievalRefNum               string `json:"retrievalRefNum" bson:"retrievalRefNum"`                             //检索参考号
	SystemTraceNum                string `json:"systemTraceNum" bson:"systemTraceNum"`                               //系统跟踪号
	ChnlType                      string `json:"chnlType" bson:"chnlType"`                                           //渠道类型
	ThirdPartyDiscountInstrution  string `json:"thirdPartyDiscountInstrution" bson:"thirdPartyDiscountInstrution"`   //第三方优惠说明
	ThirdPartyDiscountInstruction string `json:"thirdPartyDiscountInstruction" bson:"thirdPartyDiscountInstruction"` //营销联盟优惠说明
	ThirdPartyName                string `json:"thirdPartyName" bson:"thirdPartyName"`                               //第三方名称(微信、支付宝、云闪付)
	UserId                        string `json:"userId" bson:"userId"`                                               //第三方用户ID
	BankCardNo                    string `json:"bankCardNo" bson:"bankCardNo"`                                       //银行卡号
	ThirdPartyBuyerId             string `json:"thirdPartyBuyerId" bson:"thirdPartyBuyerId"`                         //第三方买家ID
	ThirdPartyOrderId             string `json:"thirdPartyOrderId" bson:"thirdPartyOrderId"`                         //第三方订单ID
	ThirdPartyPayInformation      string `json:"thirdPartyPayInformation" bson:"thirdPartyPayInformation"`           //第三方支付信息
	PromotionAmt                  string `json:"promotionAmt" bson:"promotionAmt"`                                   //营销联盟优惠金额
}

// BSCVoidPayReq 撤消支付请求
type BSCVoidPayReq struct {
	MerchantCode    string `json:"merchantCode" bson:"merchantCode"`               //商户ID
	TerminalCode    string `json:"terminalCode" bson:"terminalCode"`               //终端ID
	MerchantOrderId string `json:"merchantOrderId" bson:"merchantOrderId"`         //商户订单号
	OriginalOrderId string `json:"originalOrderId" bson:"originalOrderId"`         //银商订单号
	MerchantRemark  string `json:"merchantRemark,omitempty" bson:"merchantRemark"` //商户备注，选填
}

type BSCVoidPayResp struct {
	ErrCode                 string `json:"errCode" bson:"errCode"`                                 //错误码
	ErrInfo                 string `json:"errInfo" bson:"errInfo"`                                 //错误信息
	TransactionTime         string `json:"transactionTime" bson:"transactionTime"`                 //交易时间，格式HHmmss
	TransactionDate         string `json:"transactionDate" bson:"transactionDate"`                 //交易日期，格式yyyyMMdd
	SettlementDate          string `json:"settlementDate" bson:"settlementDate"`                   //结算日期，格式MMdd
	TransactionDateWithYear string `json:"transactionDateWithYear" bson:"transactionDateWithYear"` //交易日期，格式yyyyMMdd
	SettlementDateWithYear  string `json:"settlementDateWithYear" bson:"settlementDateWithYear"`   //结算日期，格式yyyyMMdd
	RetrievalRefNum         string `json:"retrievalRefNum" bson:"retrievalRefNum"`                 //检索参考号
	ThirdPartyName          string `json:"thirdPartyName" bson:"thirdPartyName"`                   //第三方名称(微信、支付宝、云闪付)
	TransactionAmount       int64  `json:"transactionAmount" bson:"transactionAmount"`             //交易金额，单位分
	CardAttr                string `json:"cardAttr" bson:"cardAttr"`                               //银行卡属性
	OrderId                 string `json:"orderId" bson:"orderId"`                                 //银商订单号
}

// BSCRefundReq 退款请求
type BSCRefundReq struct {
	MerchantCode      string       `json:"merchantCode" bson:"merchantCode"`               //商户ID
	TerminalCode      string       `json:"terminalCode" bson:"terminalCode"`               //终端ID
	MerchantOrderId   string       `json:"merchantOrderId" bson:"merchantOrderId"`         //商户订单号
	OriginalOrderId   string       `json:"originalOrderId" bson:"originalOrderId"`         //银商订单号
	MerchantRemark    string       `json:"merchantRemark,omitempty" bson:"merchantRemark"` //商户备注，选填
	RefundRequestId   string       `json:"refundRequestId" bson:"refundRequestId"`         //退款请求号,必填，唯一，小于50位
	TransactionAmount int64        `json:"transactionAmount" bson:"transactionAmount"`     //退款金额,必填，单位分
	Goods             []OrderGoods `json:"goods,omitempty" bson:"goods"`                   //商品列表，选填
	StoreId           string       `json:"storeId,omitempty" bson:"storeId"`               //门店StoreId
	OperatorId        string       `json:"operatorId,omitempty" bson:"operatorId"`         //操作员ID
	RefundDesc        string       `json:"refundDesc,omitempty" bson:"refundDesc"`         //退款描述
}

// BSCRefundResp 退款响应
type BSCRefundResp struct {
	ErrCode                 string `json:"errCode" bson:"errCode"`                                   //错误码
	ErrInfo                 string `json:"errInfo" bson:"errInfo"`                                   //错误信息
	TransactionTime         string `json:"transactionTime" bson:"transactionTime"`                   //交易时间，格式HHmmss
	TransactionDate         string `json:"transactionDate" bson:"transactionDate"`                   //交易日期，格式yyyyMMdd
	SettlementDate          string `json:"settlementDate" bson:"settlementDate"`                     //结算日期，格式MMdd
	TransactionDateWithYear string `json:"transactionDateWithYear" bson:"transactionDateWithYear"`   //交易日期，格式yyyyMMdd
	SettlementDateWithYear  string `json:"settlementDateWithYear" bson:"settlementDateWithYear"`     //结算日期，格式yyyyMMdd
	RetrievalRefNum         string `json:"retrievalRefNum" bson:"retrievalRefNum"`                   //检索参考号
	ThirdPartyName          string `json:"thirdPartyName" bson:"thirdPartyName"`                     //第三方名称(微信、支付宝、云闪付)
	CardAttr                string `json:"cardAttr" bson:"cardAttr"`                                 //银行卡属性
	RefundInvoiceAmount     string `json:"refundInvoiceAmount" bson:"refundInvoiceAmount"`           //退款发票金额，单位分
	TransactionAmount       int64  `json:"transactionAmount" bson:"transactionAmount"`               //退款金额，单位分
	RefundPromotionList     string `json:"refundPromotionList,omitempty" bson:"refundPromotionList"` //退款营销联盟优惠列表
	SystemTraceNum          string `json:"systemTraceNum" bson:"systemTraceNum"`                     //系统跟踪号
	ChnlType                string `json:"chnlType" bson:"chnlType"`                                 //渠道类型
	AcqInstCode             string `json:"acqInstCode" bson:"acqInstCode"`                           // acquiring institution code
	OrderId                 string `json:"orderId" bson:"orderId"`                                   //银商订单号
	RefundOrderId           string `json:"refundOrderId" bson:"refundOrderId"`                       //退款订单号
	RefundTargetOrderId     string `json:"refundTargetOrderId" bson:"refundTargetOrderId"`           //退款目标订单号
}

// BSCQueryReq 支付查询请求
type BSCQueryReq struct {
	MerchantCode    string `json:"merchantCode" bson:"merchantCode"`       //商户ID
	TerminalCode    string `json:"terminalCode" bson:"terminalCode"`       //终端ID
	MerchantOrderId string `json:"merchantOrderId" bson:"merchantOrderId"` //商户订单号
	OriginalOrderId string `json:"originalOrderId" bson:"originalOrderId"` //银商订单号
	OriTransDate    string `json:"oriTransDate" bson:"oriTransDate"`       //原交易日期,格式YYYYMMDD
}

// BSCQueryResp 支付查询响应
type BSCQueryResp struct {
	ErrCode                   string `json:"errCode" bson:"errCode"`                                 //错误码
	ErrInfo                   string `json:"errInfo" bson:"errInfo"`                                 //错误信息
	OriginalTransactionTime   string `json:"originalTransactionTime" bson:"originalTransactionTime"` //原交易时间，格式HHmmss
	QueryResCode              string `json:"queryResCode" bson:"queryResCode"`                       //查询结果码
	QueryResDesc              string `json:"queryResDesc" bson:"queryResDesc"`                       //查询结果描述
	OriginalPayCode           string `json:"originalPayCode" bson:"originalPayCode"`                 //原支付条码
	OriginalBatchNo           string `json:"originalBatchNo" bson:"originalBatchNo"`                 //原批次号
	OriginalSystemTraceNum    string `json:"originalSystemTraceNum" bson:"originalSystemTraceNum"`   //原系统跟踪号
	OriginalRetrievalRefNum   string `json:"originalRetrievalRefNum" bson:"originalRetrievalRefNum"`
	OrigialRetrievalRefNum    string `json:"origialRetrievalRefNum" bson:"origialRetrievalRefNum"`
	OriginalTransactionAmount int64  `json:"originalTransactionAmount" bson:"originalTransactionAmount"` //原交易金额，单位分
	Amount                    int64  `json:"amount" bson:"amount"`                                       //交易金额，单位分
	RefundedAmount            int64  `json:"refundedAmount" bson:"refundedAmount"`                       //已退款金额，单位分
	RefundAmonunt             int64  `json:"refundAmonunt" bson:"refundAmonunt"`                         //退款金额，单位分
	ActualTransactionAmount   int64  `json:"actualTransactionAmount" bson:"actualTransactionAmount"`     //实际交易金额，单位分
	ThirdPartyName            string `json:"thirdPartyName" bson:"thirdPartyName"`                       //第三方名称(微信、支付宝、云闪付)
	ThirdPartyBuyerId         string `json:"thirdPartyBuyerId" bson:"thirdPartyBuyerId"`                 //第三方买家ID
	ThirdPartyOrderId         string `json:"thirdPartyOrderId" bson:"thirdPartyOrderId"`                 //第三方订单号
	ThirdPartyPayInformation  string `json:"thirdPartyPayInformation" bson:"thirdPartyPayInformation"`   //第三方支付信息
	OrderId                   string `json:"orderId" bson:"orderId"`                                     //银商订单号
	MerchantOrderId           string `json:"merchantOrderId" bson:"merchantOrderId"`                     //商户订单号
	OrderStatus               string `json:"orderStatus" bson:"orderStatus"`                             //订单状态
	OrderCloseTime            string `json:"orderCloseTime" bson:"orderCloseTime"`                       //订单关闭时间
	OriginalSettlementDate    string `json:"originalSettlementDate" bson:"originalSettlementDate"`       //原结算日期，格式MMdd
	MchntName                 string `json:"mchntName" bson:"mchntName"`                                 //商户名称
	IssInstCode               string `json:"issInstCode" bson:"issInstCode"`                             //发卡机构代码
}

// BSCRefundQueryReq 退款查询请求
type BSCRefundQueryReq struct {
	MerchantCode    string `json:"merchantCode" bson:"merchantCode"`       //商户ID
	TerminalCode    string `json:"terminalCode" bson:"terminalCode"`       //终端ID
	MerchantOrderId string `json:"merchantOrderId" bson:"merchantOrderId"` //商户订单号
	OriginalOrderId string `json:"originalOrderId" bson:"originalOrderId"` //银商订单号
	RefundRequestId string `json:"refundRequestId" bson:"refundRequestId"` //退款请求号
	OriTransDate    string `json:"oriTransDate" bson:"oriTransDate"`       //原交易日期,格式YYYYMMDD
}

// BSCRefundQueryResp 退款查询响应
type BSCRefundQueryResp struct {
	ErrCode                 string `json:"errCode" bson:"errCode"`                                 //错误码
	ErrInfo                 string `json:"errInfo" bson:"errInfo"`                                 //错误信息
	TransactionTime         string `json:"transactionTime" bson:"transactionTime"`                 //交易时间，格式HHmmss
	TransactionDate         string `json:"transactionDate" bson:"transactionDate"`                 //交易日期，格式YYYYMMDD
	SettlementDate          string `json:"settlementDate" bson:"settlementDate"`                   //结算日期，格式MMdd
	RetrievalRefNum         string `json:"retrievalRefNum" bson:"retrievalRefNum"`                 //检索参考号
	QueryResCode            string `json:"queryResCode" bson:"queryResCode"`                       //查询结果码
	PayCode                 string `json:"payCode" bson:"payCode"`                                 //支付条码
	DealDate                string `json:"dealDate" bson:"dealDate"`                               //交易日期，格式YYYYMMDD
	DealTime                string `json:"dealTime" bson:"dealTime"`                               //交易时间，格式HHmmss
	OriginalAmount          int64  `json:"originalAmount" bson:"originalAmount"`                   //原交易金额，单位分
	DealType                string `json:"dealType" bson:"dealType"`                               //交易类型
	DealSystemTraceNum      string `json:"dealSystemTraceNum" bson:"dealSystemTraceNum"`           //交易系统跟踪号
	DealRetrievalRefNum     string `json:"dealRetrievalRefNum" bson:"dealRetrievalRefNum"`         //交易检索参考号
	BatchNo                 string `json:"batchNo" bson:"batchNo"`                                 //批次号
	OriginalTransactionDate string `json:"originalTransactionDate" bson:"originalTransactionDate"` //原交易日期，格式YYYYMMDD
	OrigialRetrievalRefNum  string `json:"origialRetrievalRefNum" bson:"origialRetrievalRefNum"`
	OriginalSettlementDate  string `json:"originalSettlementDate" bson:"originalSettlementDate"` //原结算日期，格式MMdd
	RefundInvoiceAmount     string `json:"refundInvoiceAmount" bson:"refundInvoiceAmount"`       //退款发票金额，单位分
	RefundPromotionList     string `json:"refundPromotionList" bson:"refundPromotionList"`       //退款优惠列表
	ThirdPartyName          string `json:"thirdPartyName" bson:"thirdPartyName"`                 //第三方名称(微信、支付宝、云闪付)
	RefundOrderId           string `json:"refundOrderId" bson:"refundOrderId"`                   //退款订单号
	RefundTargetOrderId     string `json:"refundTargetOrderId" bson:"refundTargetOrderId"`       //退款目标订单号
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
	Mid           string `json:"mid "`
	MerOrderId    string `json:"merOrderId "`
	TotalAmount   int64  `json:"totalAmount "`
	RefundOrderId string `json:"refundOrderId"`
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

type WxAppQueryReq struct {
	RequestTimestamp string `json:"requestTimestamp"`
	Mid              string `json:"mid"`
	MerOrderId       string `json:"merOrderId"`
	TargetOrderId    string `json:"targetOrderId"`
	Tid              string `json:"tid"`
	MsgId            string `json:"msgId"`
	InstMid          string `json:"instMid"`
}

type WxAppQueryResp struct {
	BuyerUsername     string `json:"buyerUsername"`
	PayTime           string `json:"payTime"`
	SeqId             string `json:"seqId"`
	InvoiceAmount     int64  `json:"invoiceAmount"`
	SettleDate        string `json:"settleDate"`
	BuyerId           string `json:"buyerId"`
	TotalAmount       int64  `json:"totalAmount"`
	CouponAmount      int64  `json:"couponAmount"`
	BuyerPayAmount    int64  `json:"buyerPayAmount"`
	TargetOrderId     string `json:"targetOrderId"`
	MerOrderId        string `json:"merOrderId"`
	Status            string `json:"status"`
	TargetSys         string `json:"targetSys"`
	MerName           string `json:"merName"`
	Mid               string `json:"mid"`
	Tid               string `json:"tid"`
	InstMid           string `json:"instMid"`
	ResponseTimestamp string `json:"responseTimestamp"`
	ErrCode           string `json:"errCode"`
}

type WxAppRefundReq struct {
	RequestTimestamp string     `json:"requestTimestamp"`
	MsgId            string     `json:"msgId"`
	Mid              string     `json:"mid"`
	Tid              string     `json:"tid"`
	SubOrders        []SubOrder `json:"subOrders"`
	InstMid          string     `json:"instMid"`
	PlatformAmount   int64      `json:"platformAmount"`
	RefundAmount     int64      `json:"refundAmount"`
	RefundOrderId    string     `json:"refundOrderId"`
	MerOrderId       string     `json:"merOrderId"`
	TargetOrderId    string     `json:"targetOrderId"`
	BillDate         string     `json:"billDate"`
	RefundDesc       string     `json:"refundDesc"`
}

type WxAppRefundResp struct {
	Mid                 string `json:"mid"`
	RefundStatus        string `json:"refundStatus"`
	BillDate            string `json:"billDate"`
	Tid                 string `json:"tid"`
	InstMid             string `json:"instMid"`
	RefundOrderId       string `json:"refundOrderId"`
	RefundTargetOrderId string `json:"refundTargetOrderId"`
	ResponseTimestamp   string `json:"responseTimestamp"`
	ErrCode             string `json:"errCode"`
	Status              string `json:"status"`
	MerOrderId          string `json:"merOrderId"`
	RefundAmount        int64  `json:"refundAmount"`
}

type WxAppRefundQueryReq struct {
	RequestTimestamp string `json:"requestTimestamp"`
	MsgId            string `json:"msgId"`
	Mid              string `json:"mid"`
	Tid              string `json:"tid"`
	InstMid          string `json:"instMid"`
	MerOrderId       string `json:"merOrderId"`
}

type WxAppRefundQueryResp struct {
	PayTime             string `json:"payTime"`
	ConnectSys          string `json:"connectSys"`
	ErrMsg              string `json:"errMsg"`
	MerName             string `json:"merName"`
	Mid                 string `json:"mid"`
	RefundStatus        string `json:"refundStatus"`
	SettleDate          string `json:"settleDate"`
	SettleRefId         string `json:"settleRefId"`
	Tid                 string `json:"tid"`
	RefundOrderId       string `json:"refundOrderId"`
	RefundTargetOrderId string `json:"refundTargetOrderId"`
	TotalAmount         int64  `json:"totalAmount"`
	TargetMid           string `json:"targetMid"`
	ResponseTimestamp   string `json:"responseTimestamp"`
	ErrCode             string `json:"errCode"`
	TargetStatus        string `json:"targetStatus"`
	SeqId               string `json:"seqId"`
	MerOrderId          string `json:"merOrderId"`
	Status              string `json:"status"`
	TargetSys           string `json:"targetSys"`
}

type WxAppCloseReq struct {
	RequestTimestamp string `json:"requestTimestamp"`
	SrcReserve       string `json:"srcReserve"`
	MerOrderId       string `json:"merOrderId"`
	Mid              string `json:"mid"`
	Tid              string `json:"tid"`
	InstMid          string `json:"instMid"`
}

type WxAppCloseResp struct {
	ErrMsg            string `json:"errMsg"`
	Mid               string `json:"mid"`
	Tid               string `json:"tid"`
	InstMid           string `json:"instMid"`
	ResponseTimestamp string `json:"responseTimestamp"`
	ErrCode           string `json:"errCode"`
	MerOrderId        string `json:"merOrderId"`
	MerName           string `json:"merName"`
	SeqId             string `json:"seqId"`
	Status            string `json:"status"`
	TargetSys         string `json:"targetSys"`
	TargetStatus      string `json:"targetStatus"`
	TotalAmount       int64  `json:"totalAmount"`
}

// QrPayReq 二维码支付请求参数
type QrPayReq struct {
	RequestTimestamp string       `json:"requestTimestamp" bson:"requestTimestamp"` // 请求时间戳，格式为yyyy-MM-dd HH:mm:ss，自动填充为当前时间
	BillNo           string       `json:"billNo" bson:"billNo"`                     // 订单号，商户订单号唯一
	BillDesc         string       `json:"billDesc" bson:"billDesc"`                 // 订单描述,可以加上具体门店名称
	BillDate         string       `json:"billDate" bson:"billDate"`                 // 订单日期，格式为yyyy-MM-dd，自动填充为当前日期
	Goods            []OrderGoods `json:"goods" bson:"goods"`                       // 订单商品列表，选填
	Mid              string       `json:"mid" bson:"mid"`                           // 商户ID，必填，默认值为配置文件中的商户ID
	Tid              string       `json:"tid" bson:"tid"`                           // 终端ID，必填，默认值为配置文件中的终端ID
	MsgId            string       `json:"msgId" bson:"msgId"`                       // 消息ID，选填
	InstMid          string       `json:"instMid" bson:"instMid"`                   // 机构ID，必填，默认值为配置文件中的机构ID
	TotalAmount      int64        `json:"totalAmount" bson:"totalAmount"`           // 订单金额，必填，单位分
	NotifyUrl        string       `json:"notifyUrl" bson:"notifyUrl"`               // 回调通知URL，选填
	ReturnUrl        string       `json:"returnUrl" bson:"returnUrl"`               // 支付完成后跳转URL，选填
	SubOrders        []SubOrder   `json:"subOrders" bson:"subOrders"`               // 子订单列表，选填
}

// QrPayResp 二维码支付响应参数
type QrPayResp struct {
	QrCodeId          string `json:"qrCodeId" bson:"qrCodeId"`                   // 二维码ID
	SystemId          string `json:"systemId" bson:"systemId"`                   // 系统ID
	ErrMsg            string `json:"errMsg" bson:"errMsg"`                       // 错误信息
	Mid               string `json:"mid" bson:"mid"`                             // 商户ID
	MsgId             string `json:"msgId" bson:"msgId"`                         // 消息ID
	BillDate          string `json:"billDate" bson:"billDate"`                   // 订单日期，格式为yyyy-MM-dd
	Tid               string `json:"tid" bson:"tid"`                             // 终端ID
	InstMid           string `json:"instMid" bson:"instMid"`                     // 机构ID
	ResponseTimestamp string `json:"responseTimestamp" bson:"responseTimestamp"` // 响应时间戳，格式为yyyy-MM-dd HH:mm:ss
	ErrCode           string `json:"errCode" bson:"errCode"`                     // 错误码,SUCCESS表示成功
	BillNo            string `json:"billNo" bson:"billNo"`                       // 商户订单号
	BillQRCode        string `json:"billQRCode" bson:"billQRCode"`               // 订单二维码，用于扫码支付，前端用些url生成一个二维码
}

// QrPayQueryReq 二维码支付查询请求参数
type QrPayQueryReq struct {
	RequestTimestamp string `json:"requestTimestamp" bson:"requestTimestamp"` // 请求时间戳，格式为yyyy-MM-dd HH:mm:ss，自动填充为当前时间
	Mid              string `json:"mid" bson:"mid"`                           // 商户ID
	Tid              string `json:"tid" bson:"tid"`                           // 终端ID,自动填充为配置文件中的终端ID
	BillDate         string `json:"billDate" bson:"billDate"`                 // 订单日期，格式为yyyy-MM-dd
	BillNo           string `json:"billNo" bson:"billNo"`                     // 订单号
	MsgId            string `json:"msgId" bson:"msgId"`                       // 消息ID
	InstMid          string `json:"instMid" bson:"instMid"`                   // 机构ID,自动填充为配置文件中的机构ID
}

// QrPayQueryResp 二维码支付查询响应参数
type QrPayQueryResp struct {
	BillPayment struct {
		PayTime         string `json:"payTime" bson:"payTime"`                 // 支付时间，格式为yyyy-MM-dd HH:mm:ss
		BuyerCashPayAmt int64  `json:"buyerCashPayAmt" bson:"buyerCashPayAmt"` // 买家支付金额，单位分
		ConnectSys      string `json:"connectSys" bson:"connectSys"`           // 连接系统
		PaySeqId        string `json:"paySeqId" bson:"paySeqId"`               // 支付序列ID
		InvoiceAmount   int64  `json:"invoiceAmount" bson:"invoiceAmount"`     // 发票金额，单位分
		SettleDate      string `json:"settleDate" bson:"settleDate"`           // 结算日期，格式为yyyy-MM-dd
		BuyerId         string `json:"buyerId" bson:"buyerId"`                 // 买家ID
		ReceiptAmount   int64  `json:"receiptAmount" bson:"receiptAmount"`     // 发票金额，单位分
		TotalAmount     int64  `json:"totalAmount" bson:"totalAmount"`         // 订单金额，单位分
		CouponAmount    int64  `json:"couponAmount" bson:"couponAmount"`       // 优惠券金额，单位分
		BillBizType     string `json:"billBizType" bson:"billBizType"`
		BuyerPayAmount  int64  `json:"buyerPayAmount" bson:"buyerPayAmount"` // 买家支付金额，单位分
		TargetOrderId   string `json:"targetOrderId" bson:"targetOrderId"`   // 目标订单ID
		PayDetail       string `json:"payDetail" bson:"payDetail"`           // 支付详情
		MerOrderId      string `json:"merOrderId" bson:"merOrderId"`         // 商户订单ID
		Status          string `json:"status" bson:"status"`                 // 订单状态
		TargetSys       string `json:"targetSys" bson:"targetSys"`           // 目标系统，微信、支付宝、云闪付
	} `json:"billPayment"`
	BillDesc          string `json:"billDesc" bson:"billDesc"`                   // 订单描述
	MerName           string `json:"merName" bson:"merName"`                     // 商户名称
	Mid               string `json:"mid" bson:"mid"`                             // 商户ID
	MsgId             string `json:"msgId" bson:"msgId"`                         // 消息ID
	BillDate          string `json:"billDate" bson:"billDate"`                   // 订单日期，格式为yyyy-MM-dd
	Tid               string `json:"tid" bson:"tid"`                             // 终端ID
	InstMid           string `json:"instMid" bson:"instMid"`                     // 机构ID,自动填充为配置文件中的机构ID
	TotalAmount       int64  `json:"totalAmount" bson:"totalAmount"`             // 订单金额，单位分
	CreateTime        string `json:"createTime" bson:"createTime"`               // 创建时间，格式为yyyy-MM-dd HH:mm:ss
	ResponseTimestamp string `json:"responseTimestamp" bson:"responseTimestamp"` // 响应时间戳，格式为yyyy-MM-dd HH:mm:ss
	ErrCode           string `json:"errCode" bson:"errCode"`                     // 错误码,SUCCESS表示成功
	BillStatus        string `json:"billStatus" bson:"billStatus"`               // 订单状态
	CardAttr          string `json:"cardAttr" bson:"cardAttr"`                   // 银行卡属性
	BillNo            string `json:"billNo" bson:"billNo"`                       // 订单号
	BillQRCode        string `json:"billQRCode" bson:"billQRCode"`               // 订单二维码，用于扫码支付，前端用些url生成一个二维码
}

// QrPayRefundReq 二维码支付退款请求参数
type QrPayRefundReq struct {
	RequestTimestamp string `json:"requestTimestamp" bson:"requestTimestamp"` // 请求时间戳，格式为yyyy-MM-dd HH:mm:ss，自动填充为当前时间
	MsgId            string `json:"msgId" bson:"msgId"`                       // 消息ID
	Mid              string `json:"mid" bson:"mid"`                           // 商户ID
	Tid              string `json:"tid" bson:"tid"`                           // 终端ID,自动填充为配置文件中的终端ID
	InstMid          string `json:"instMid" bson:"instMid"`                   // 机构ID,自动填充为配置文件中的机构ID
	RefundAmount     int64  `json:"refundAmount" bson:"refundAmount"`         // 退款金额，单位分
	BillNo           string `json:"billNo" bson:"billNo"`                     // 订单号
	BillDate         string `json:"billDate" bson:"billDate"`                 // 订单日期，格式为yyyy-MM-dd
}

// QrPayRefundResp 二维码支付退款响应参数
type QrPayRefundResp struct {
	Mid                 string `json:"mid" bson:"mid"`                                 // 商户ID
	MsgId               string `json:"msgId" bson:"msgId"`                             // 消息ID
	RefundStatus        string `json:"refundStatus" bson:"refundStatus"`               // 退款状态
	BillDate            string `json:"billDate" bson:"billDate"`                       // 订单日期，格式为yyyy-MM-dd
	SettleDate          string `json:"settleDate" bson:"settleDate"`                   // 结算日期，格式为yyyy-MM-dd
	Tid                 string `json:"tid" bson:"tid"`                                 // 终端ID,自动填充为配置文件中的终端ID
	InstMid             string `json:"instMid" bson:"instMid"`                         // 机构ID,自动填充为配置文件中的机构ID
	RefundOrderId       string `json:"refundOrderId" bson:"refundOrderId"`             // 退款订单ID
	RefundTargetOrderId string `json:"refundTargetOrderId" bson:"refundTargetOrderId"` // 退款目标订单ID
	RefundInvoiceAmount int    `json:"refundInvoiceAmount" bson:"refundInvoiceAmount"` // 退款发票金额，单位分
	ResponseTimestamp   string `json:"responseTimestamp" bson:"responseTimestamp"`     // 响应时间戳，格式为yyyy-MM-dd HH:mm:ss
	ErrCode             string `json:"errCode" bson:"errCode"`                         // 错误码,SUCCESS表示成功
	BillStatus          string `json:"billStatus" bson:"billStatus"`                   // 订单状态
	CardAttr            string `json:"cardAttr" bson:"cardAttr"`                       // 银行卡属性
	RefundPayTime       string `json:"refundPayTime" bson:"refundPayTime"`             // 退款支付时间，格式为yyyy-MM-dd HH:mm:ss
	BillNo              string `json:"billNo" bson:"billNo"`                           // 订单号
	BillQRCode          string `json:"billQRCode" bson:"billQRCode"`                   // 订单二维码，用于扫码支付，前端用些url生成一个二维码
	MerOrderId          string `json:"merOrderId" bson:"merOrderId"`                   // 商户订单ID
	RefundAmount        int    `json:"refundAmount" bson:"refundAmount"`               // 退款金额，单位分
	TargetSys           string `json:"targetSys" bson:"targetSys"`                     // 目标系统，微信、支付宝、云闪付
}

type QrPayNotifyReq struct {
	Mid            string `json:"mid"  form:"mid" bson:"mid"`                                 // 商户ID
	Tid            string `json:"tid"  form:"tid" bson:"tid"`                                 // 终端ID
	InstMid        string `json:"instMid" form:"instMid" bson:"instMid"`                      // 机构ID,自动填充为配置文件中的机构ID
	BillNo         string `json:"billNo" form:"billNo" bson:"billNo"`                         // 订单号
	BillQRCode     string `json:"billQRCode" form:"billQRcode" bson:"billQRCode"`             // 订单二维码，用于扫码支付，前端用些url生成一个二维码
	BillDate       string `json:"billDate" form:"billDate" bson:"billDate"`                   // 订单日期，格式为yyyy-MM-dd
	CreateTime     string `json:"createTime" form:"createTime" bson:"createTime"`             // 创建时间，格式为yyyy-MM-dd HH:mm:ss
	BillStatus     string `json:"billStatus" form:"billStatus" bson:"billStatus"`             // 订单状态
	BillDesc       string `json:"billDesc" form:"billDesc" bson:"billDesc"`                   // 订单描述
	TotalAmount    int64  `json:"totalAmount" form:"totalAmount" bson:"totalAmount"`          // 订单金额，单位分
	MemberId       string `json:"memberId" form:"memberId" bson:"memberId"`                   // 商户会员ID
	MerName        string `json:"merName" form:"merName" bson:"merName"`                      // 商户名称
	Memo           string `json:"memo" form:"memo" bson:"memo"`                               // 备注
	NotifyId       string `json:"notifyId" form:"notifyId" bson:"notifyId"`                   // 通知ID
	SecureStatus   string `json:"secureStatus" form:"secureStatus" bson:"secureStatus"`       // 担保状态
	CompleteAmount int64  `json:"completeAmount" form:"completeAmount" bson:"completeAmount"` // 担保完成金额，单位分
	BillPayment    string `json:"billPayment" form:"billPayment" bson:"billPayment"`          // 账单支付状态,JSON串
	Sign           string `json:"sign" form:"sign" bson:"sign"`                               // 签名
	BankInfo       string `json:"bankInfo" form:"bankInfo" bson:"bankInfo"`                   // 银行卡信息
	BankCardNo     string `json:"bankCardNo" form:"bankCardNo" bson:"bankCardNo"`             // 银行卡号
	SeqId          string `json:"seqId" form:"seqId" bson:"seqId"`                            // 序列号
	ReceiptAmount  int64  `json:"receiptAmount" form:"receiptAmount" bson:"receiptAmount"`    // 实收金额，单位分
	ExtraBuyerInfo string `json:"extraBuyerInfo" form:"extraBuyerInfo" bson:"extraBuyerInfo"` // 额外买家信息
	RefundAmount   int64  `json:"refundAmount" form:"refundAmount" bson:"refundAmount"`       // 退款金额，单位分
	RefundDesc     string `json:"refundDesc" form:"refundDesc" bson:"refundDesc"`             // 退款描述
	MchntUuid      string `json:"mchntUuId" form:"mchntUuId" bson:"mchntUuid"`                // 商户统一商号
	SubInst        string `json:"subInst" form:"subInst" bson:"subInst"`                      // 子机构ID
	SrcReserved    string `json:"srcReserved" form:"srcReserved" bson:"srcReserved"`          // 备注
	QrCodeType     string `json:"qrCodeType" form:"qrCodeType" bson:"qrCodeType"`             // 二维码类型
}

type QrPayNotifyBillPayment struct {
	TotalAmount   int64  `json:"totalAmount"`   // 订单金额，单位分
	PayTime       string `json:"payTime"`       // 支付时间，格式为yyyy-MM-dd HH:mm:ss
	BillBizType   string `json:"billBizType"`   // 账单业务类型
	TargetOrderId string `json:"targetOrderId"` // 支付应用内部订单ID
	PaySeqId      string `json:"paySeqId"`      // 支付流水ID
	InvoiceAmount int64  `json:"invoiceAmount"` // 开票金额，单位分
	PayDetail     string `json:"payDetail"`     // 支付明细
	SettleDate    string `json:"settleDate"`    // 结算日期，格式为yyyy-MM-dd
	BuyerId       string `json:"buyerId"`       // 买家ID,对应支付应用的会员ID
	MerOrderId    string `json:"merOrderId"`    // 商户订单ID
	Status        string `json:"status"`        // 状态
	TargetSys     string `json:"targetSys"`     // 支付应用
}

// 异步核销模型BusinessUnifyMulti(Bum)
// 公共请求头
type BumReqHeader struct {
	TransCode  string `json:"transCode,omitempty"`  // 交易码
	VerNo      string `json:"verNo,omitempty"`      // 版本
	SrcReqDate string `json:"srcReqDate,omitempty"` // YYYYMMDD
	SrcReqTime string `json:"srcReqTime,omitempty"` // HHmmss
	SrcReqId   string `json:"srcReqId,omitempty"`   // 36 位 UUID
	ChannelId  string `json:"channelId,omitempty"`  // 043
	GroupId    string `json:"groupId,omitempty"`    // 6 位集团号
	Signature  string `json:"signature,omitempty"`  // RSA 签名
}

// 公共响应头
type BumRespHeader struct {
	TransCode  string `json:"transCode"`
	VerNo      string `json:"verNo"`
	SrcReqDate string `json:"srcReqDate"`
	SrcReqTime string `json:"srcReqTime"`
	SrcReqId   string `json:"srcReqId"`
	RespCode   string `json:"respCode"`
	RespMsg    string `json:"respMsg"`
	Signature  string `json:"signature"`
}

// 202017 充值暂挂余额查询
type BumPendingBalanceReq struct {
	BumReqHeader
	MerNo string `json:"merNo,omitempty"` // 14
}
type BumPendingBalanceResp struct {
	BumRespHeader
	MerNo           string `json:"merNo"`
	MerName         string `json:"merName"`
	GroupId         string `json:"groupId"`
	PendingAmt      string `json:"pendingAmt"` // 分
	InTransitAmt    string `json:"inTransitAmt"`
	VerifiedAmt     string `json:"verifiedAmt"`
	OnlineRefundAmt string `json:"onlineRefundAmt"`
	OnlineManRefAmt string `json:"onlineManRefundAmt"`
	OnlineSetRefAmt string `json:"onlineSettleRefundAmt"`
	Reserve1        string `json:"reserve1,omitempty"`
	Reserve2        string `json:"reserve2,omitempty"`
	Reserve3        string `json:"reserve3,omitempty"`
}

// BumOrderQueryReq 异步支付充值订单查询请求(202018)
type BumRechargeOrderQueryReq struct {
	BumReqHeader
	MerNo      string `json:"merNo,omitempty"`      // 企业用户号(14位)
	TransDate  string `json:"transDate,omitempty"`  // 交易日期(YYYYMMDD)
	QueryItem  string `json:"queryItem,omitempty"`  // 查询项(A-商户订单号,B-银商流水号)
	QueryValue string `json:"queryValue,omitempty"` // 查询值
}

type QueryRechargeOrderResp struct {
	BumRespHeader
	MerNo        string `json:"merNo"`
	StlDate      string `json:"stlDate"`
	OrderNo      string `json:"orderNo"`
	MerFee       string `json:"merFee"`
	SettleStatus string `json:"settleStatus"` // "00","01","Special"
	TxnAmt       string `json:"txnAmt"`
	PendingAmt   string `json:"pendingAmt"`
	VerifiedAmt  string `json:"verifiedAmt"`
	InTransitAmt string `json:"inTransitAmt"`
	PndFlg       string `json:"pndFlg"` // 1/2/3
	TransDt      string `json:"transDt"`
	CardType     string `json:"cardType"`
	Pan          string `json:"pan"`
	CardIssuer   string `json:"cardIssuer"`
	OriRefNo     string `json:"oriRefNo,omitempty"`
	OriTransDt   string `json:"oriTransDt,omitempty"`
	OriTxnAmt    string `json:"oriTxnAmt,omitempty"`
}

// BumTransferReq 按金额划付请求(202002)
type BumTransferReq struct {
	BumReqHeader
	MerNo  string `json:"merNo,omitempty"`  // 企业用户号(14位)
	PayAmt string `json:"payAmt,omitempty"` // 划付金额(单位:分,字符串类型)
}

// BumAllocationReq 按金额分账请求(202004)
type BumAllocationReq struct {
	BumReqHeader
	MerNo   string `json:"merNo,omitempty"`   // 企业用户号(14位)
	PayType string `json:"payType,omitempty"` // 分账类型(0-指定金额分账)
	CardNo  string `json:"cardNo,omitempty"`  // 加密卡号(SHA-256哈希)
	Ps      string `json:"ps,omitempty"`      // 分账附言(最长30字符)
	PayAmt  string `json:"payAmt,omitempty"`  // 分账金额(单位:分,字符串类型)
}

// ====== 202006: 商户信息查询 ======
type QueryMerchantInfoReq struct {
	BumReqHeader
	MerNo string `json:"merNo"`
}

type QueryMerchantInfoResp struct {
	BumRespHeader
	MerNo     string `json:"merNo"`
	MerName   string `json:"merName"`
	GroupId   string `json:"groupId"`
	TopMer    string `json:"topMer"`
	MerLevel  string `json:"merLevel"`
	CanPayAmt string `json:"canPayAmt"`
}

// QueryTxnDetailReq 202007
type QueryTxnDetailReq struct {
	BumReqHeader
	MerNo      string `json:"merNo"`
	TransDate  string `json:"transDate"` // yyyyMMdd
	QueryItem  string `json:"queryItem"` // merchant order no
	QueryValue string `json:"queryValue"`
}

type QueryTxnDetailResp struct {
	BumRespHeader
	MerNo        string `json:"merNo"`
	StlDate      string `json:"stlDate"`
	MerOrderNo   string `json:"merOrderNo"`
	CanPayAmt    string `json:"canPayAmt"`
	SettleStatus string `json:"settleStatus"` // "00","01","Special"
	MerFee       string `json:"merFee"`
	TxnAmt       string `json:"txnAmt"`
	PaidAmt      string `json:"paidAmt"`
	AllottedAmt  string `json:"allottedAmt"`
	TransDt      string `json:"transDt"`
	CardType     string `json:"cardType"`
	Pan          string `json:"pan"`
	IssInstId    string `json:"issInstId"`
	OriRefNo     string `json:"oriRefNo,omitempty"`
	OriTransDt   string `json:"oriTransDt,omitempty"`
	OriTxnAmt    string `json:"oriTxnAmt,omitempty"`
}

// ====== 202008: 操作记录查询 ======

type QueryOperationLogReq struct {
	BumReqHeader
	MerNo        string `json:"merNo"`
	ReqDate      string `json:"reqDate"`      // yyyyMMdd
	ReqJournalNo string `json:"reqJournalNo"` // srcReqId
}

type OperationOrder struct {
	OrderNo   string `json:"orderNo"`
	OrderType string `json:"orderType"` // 核销/分账
	Status    string `json:"status"`    // 1/2/3/5/6
	OrderAmt  string `json:"orderAmt"`
}

type QueryOperationLogResp struct {
	BumRespHeader
	MerNo        string           `json:"merNo"`
	ReqDate      string           `json:"reqDate"`
	ReqJournalNo string           `json:"reqJournalNo"`
	OrderSet     []OperationOrder `json:"orderSet"`
	Status       string           `json:"status"` // overall 0/1
}
