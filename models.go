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
	RequestTimestamp string       `json:"requestTimestamp"` // 请求时间戳，格式为yyyy-MM-dd HH:mm:ss，自动填充为当前时间
	BillNo           string       `json:"billNo"`           // 订单号，商户订单号唯一
	BillDesc         string       `json:"billDesc"`         // 订单描述,可以加上具体门店名称
	BillDate         string       `json:"billDate"`         // 订单日期，格式为yyyy-MM-dd，自动填充为当前日期
	Goods            []OrderGoods `json:"goods"`            // 订单商品列表，选填
	Mid              string       `json:"mid"`              // 商户ID，必填，默认值为配置文件中的商户ID
	Tid              string       `json:"tid"`              // 终端ID，必填，默认值为配置文件中的终端ID
	MsgId            string       `json:"msgId"`            // 消息ID，选填
	InstMid          string       `json:"instMid"`          // 机构ID，必填，默认值为配置文件中的机构ID
	TotalAmount      int64        `json:"totalAmount"`      // 订单金额，必填，单位分
	NotifyUrl        string       `json:"notifyUrl"`        // 回调通知URL，选填
	ReturnUrl        string       `json:"returnUrl"`        // 支付完成后跳转URL，选填
	SubOrders        []SubOrder   `json:"subOrders"`        // 子订单列表，选填
}

// QrPayResp 二维码支付响应参数
type QrPayResp struct {
	QrCodeId          string `json:"qrCodeId"`          // 二维码ID
	SystemId          string `json:"systemId"`          // 系统ID
	ErrMsg            string `json:"errMsg"`            // 错误信息
	Mid               string `json:"mid"`               // 商户ID
	MsgId             string `json:"msgId"`             // 消息ID
	BillDate          string `json:"billDate"`          // 订单日期，格式为yyyy-MM-dd
	Tid               string `json:"tid"`               // 终端ID
	InstMid           string `json:"instMid"`           // 机构ID
	ResponseTimestamp string `json:"responseTimestamp"` // 响应时间戳，格式为yyyy-MM-dd HH:mm:ss
	ErrCode           string `json:"errCode"`           // 错误码,SUCCESS表示成功
	BillNo            string `json:"billNo"`            // 商户订单号
	BillQRCode        string `json:"billQRCode"`        // 订单二维码，用于扫码支付，前端用些url生成一个二维码
}

// QrPayQueryReq 二维码支付查询请求参数
type QrPayQueryReq struct {
	RequestTimestamp string `json:"requestTimestamp"` // 请求时间戳，格式为yyyy-MM-dd HH:mm:ss，自动填充为当前时间
	Mid              string `json:"mid"`              // 商户ID
	Tid              string `json:"tid"`              // 终端ID
	BillDate         string `json:"billDate"`         // 订单日期，格式为yyyy-MM-dd
	BillNo           string `json:"billNo"`           // 订单号
	MsgId            string `json:"msgId"`            // 消息ID
	InstMid          string `json:"instMid"`          // 机构ID,自动填充为配置文件中的机构ID
}

// QrPayQueryResp 二维码支付查询响应参数
type QrPayQueryResp struct {
	BillPayment struct {
		PayTime         string `json:"payTime"`         // 支付时间，格式为yyyy-MM-dd HH:mm:ss
		BuyerCashPayAmt int    `json:"buyerCashPayAmt"` // 买家支付金额，单位分
		ConnectSys      string `json:"connectSys"`      // 连接系统
		PaySeqId        string `json:"paySeqId"`        // 支付序列ID
		InvoiceAmount   int    `json:"invoiceAmount"`   // 发票金额，单位分
		SettleDate      string `json:"settleDate"`      // 结算日期，格式为yyyy-MM-dd
		BuyerId         string `json:"buyerId"`         // 买家ID
		ReceiptAmount   int    `json:"receiptAmount"`   // 发票金额，单位分
		TotalAmount     int    `json:"totalAmount"`     // 订单金额，单位分
		CouponAmount    int    `json:"couponAmount"`    // 优惠券金额，单位分
		BillBizType     string `json:"billBizType"`
		BuyerPayAmount  int    `json:"buyerPayAmount"` // 买家支付金额，单位分
		TargetOrderId   string `json:"targetOrderId"`  // 目标订单ID
		PayDetail       string `json:"payDetail"`      // 支付详情
		MerOrderId      string `json:"merOrderId"`     // 商户订单ID
		Status          string `json:"status"`         // 订单状态
		TargetSys       string `json:"targetSys"`      // 目标系统，微信、支付宝、云闪付
	} `json:"billPayment"`
	BillDesc          string `json:"billDesc"`          // 订单描述
	MerName           string `json:"merName"`           // 商户名称
	Mid               string `json:"mid"`               // 商户ID
	MsgId             string `json:"msgId"`             // 消息ID
	BillDate          string `json:"billDate"`          // 订单日期，格式为yyyy-MM-dd
	Tid               string `json:"tid"`               // 终端ID
	InstMid           string `json:"instMid"`           // 机构ID,自动填充为配置文件中的机构ID
	TotalAmount       int    `json:"totalAmount"`       // 订单金额，单位分
	CreateTime        string `json:"createTime"`        // 创建时间，格式为yyyy-MM-dd HH:mm:ss
	ResponseTimestamp string `json:"responseTimestamp"` // 响应时间戳，格式为yyyy-MM-dd HH:mm:ss
	ErrCode           string `json:"errCode"`           // 错误码,SUCCESS表示成功
	BillStatus        string `json:"billStatus"`        // 订单状态
	CardAttr          string `json:"cardAttr"`          // 银行卡属性
	BillNo            string `json:"billNo"`            // 订单号
	BillQRCode        string `json:"billQRCode"`        // 订单二维码，用于扫码支付，前端用些url生成一个二维码
}

// QrPayRefundReq 二维码支付退款请求参数
type QrPayRefundReq struct {
	RequestTimestamp string `json:"requestTimestamp"` // 请求时间戳，格式为yyyy-MM-dd HH:mm:ss，自动填充为当前时间
	MsgId            string `json:"msgId"`            // 消息ID
	Mid              string `json:"mid"`              // 商户ID
	Tid              string `json:"tid"`              // 终端ID
	InstMid          string `json:"instMid"`          // 机构ID,自动填充为配置文件中的机构ID
	RefundAmount     int64  `json:"refundAmount"`     // 退款金额，单位分
	BillNo           string `json:"billNo"`           // 订单号
	BillDate         string `json:"billDate"`         // 订单日期，格式为yyyy-MM-dd
}

// QrPayRefundResp 二维码支付退款响应参数
type QrPayRefundResp struct {
	Mid                 string `json:"mid"`                 // 商户ID
	MsgId               string `json:"msgId"`               // 消息ID
	RefundStatus        string `json:"refundStatus"`        // 退款状态
	BillDate            string `json:"billDate"`            // 订单日期，格式为yyyy-MM-dd
	SettleDate          string `json:"settleDate"`          // 结算日期，格式为yyyy-MM-dd
	Tid                 string `json:"tid"`                 // 终端ID
	InstMid             string `json:"instMid"`             // 机构ID,自动填充为配置文件中的机构ID
	RefundOrderId       string `json:"refundOrderId"`       // 退款订单ID
	RefundTargetOrderId string `json:"refundTargetOrderId"` // 退款目标订单ID
	RefundInvoiceAmount int    `json:"refundInvoiceAmount"` // 退款发票金额，单位分
	ResponseTimestamp   string `json:"responseTimestamp"`   // 响应时间戳，格式为yyyy-MM-dd HH:mm:ss
	ErrCode             string `json:"errCode"`             // 错误码,SUCCESS表示成功
	BillStatus          string `json:"billStatus"`          // 订单状态
	CardAttr            string `json:"cardAttr"`            // 银行卡属性
	RefundPayTime       string `json:"refundPayTime"`       // 退款支付时间，格式为yyyy-MM-dd HH:mm:ss
	BillNo              string `json:"billNo"`              // 订单号
	BillQRCode          string `json:"billQRCode"`          // 订单二维码，用于扫码支付，前端用些url生成一个二维码
	MerOrderId          string `json:"merOrderId"`          // 商户订单ID
	RefundAmount        int    `json:"refundAmount"`        // 退款金额，单位分
	TargetSys           string `json:"targetSys"`           // 目标系统，微信、支付宝、云闪付
}

type QrPayNotifyReq struct {
	Mid            string `json:"mid"  form:"mid"`                      // 商户ID
	Tid            string `json:"tid"  form:"tid"`                      // 终端ID
	InstMid        string `json:"instMid" form:"instMid"`               // 机构ID,自动填充为配置文件中的机构ID
	BillNo         string `json:"billNo" form:"billNo"`                 // 订单号
	BillQRCode     string `json:"billQRCode" form:"billQRcode"`         // 订单二维码，用于扫码支付，前端用些url生成一个二维码
	BillDate       string `json:"billDate" form:"billDate"`             // 订单日期，格式为yyyy-MM-dd
	CreateTime     string `json:"createTime" form:"createTime"`         // 创建时间，格式为yyyy-MM-dd HH:mm:ss
	BillStatus     string `json:"billStatus" form:"billStatus"`         // 订单状态
	BillDesc       string `json:"billDesc" form:"billDesc"`             // 订单描述
	TotalAmount    int64  `json:"totalAmount" form:"totalAmount"`       // 订单金额，单位分
	MemberId       string `json:"memberId" form:"memberId"`             // 商户会员ID
	MerName        string `json:"merName" form:"merName"`               // 商户名称
	Memo           string `json:"memo" form:"memo"`                     // 备注
	NotifyId       string `json:"notifyId" form:"notifyId"`             // 通知ID
	SecureStatus   string `json:"secureStatus" form:"secureStatus"`     // 担保状态
	CompleteAmount int64  `json:"completeAmount" form:"completeAmount"` // 担保完成金额，单位分
	BillPayment    string `json:"billPayment" form:"billPayment"`       // 账单支付状态,JSON串
	Sign           string `json:"sign" form:"sign"`                     // 签名
	BankInfo       string `json:"bankInfo" form:"bankInfo"`             // 银行卡信息
	BankCardNo     string `json:"bankCardNo" form:"bankCardNo"`         // 银行卡号
	SeqId          string `json:"seqId" form:"seqId"`                   // 序列号
	ReceiptAmount  int64  `json:"receiptAmount" form:"receiptAmount"`   // 实收金额，单位分
	ExtraBuyerInfo string `json:"extraBuyerInfo" form:"extraBuyerInfo"` // 额外买家信息
	RefundAmount   int64  `json:"refundAmount" form:"refundAmount"`     // 退款金额，单位分
	RefundDesc     string `json:"refundDesc" form:"refundDesc"`         // 退款描述
	MchntUuid      string `json:"mchntUuId" form:"mchntUuId"`           // 商户统一商号
	SubInst        string `json:"subInst" form:"subInst"`               // 子机构ID
	SrcReserved    string `json:"srcReserved" form:"srcReserved"`       // 备注
	QrCodeType     string `json:"qrCodeType" form:"qrCodeType"`         // 二维码类型
}

type QrPayNotifyBillPayment struct {
	TotalAmount   int    `json:"totalAmount"`   // 订单金额，单位分
	PayTime       string `json:"payTime"`       // 支付时间，格式为yyyy-MM-dd HH:mm:ss
	BillBizType   string `json:"billBizType"`   // 账单业务类型
	TargetOrderId string `json:"targetOrderId"` // 支付应用内部订单ID
	PaySeqId      string `json:"paySeqId"`      // 支付流水ID
	InvoiceAmount int    `json:"invoiceAmount"` // 开票金额，单位分
	PayDetail     string `json:"payDetail"`     // 支付明细
	SettleDate    string `json:"settleDate"`    // 结算日期，格式为yyyy-MM-dd
	BuyerId       string `json:"buyerId"`       // 买家ID,对应支付应用的会员ID
	MerOrderId    string `json:"merOrderId"`    // 商户订单ID
	Status        string `json:"status"`        // 状态
	TargetSys     string `json:"targetSys"`     // 支付应用
}
