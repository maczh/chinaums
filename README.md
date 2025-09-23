## 银联商务支付SDK包



### 一、调用说明

1. 导入sdk包 `github.com/maczh/chinaums`

2. 入口为 chinaums.Client对象



### 二、商户扫码BSC

商户扫码对象提供支付、撤消支付、退款、订单查询、退款查询等5个方法

#### 2.1 条码支付

PayCode支持微信、支付宝、云闪付的条码

**方法定义**

```go
func (b *bscPay) Pay(req *BSCPayReq) (*BSCPayResp, error)
```

**请求体**

```json
{
	"transactionAmount": 2,
	"transactionCurrencyCode": "156",
	"merchantOrderId": "JHTEST20250923164207",
	"merchantRemark": "测试订单",
	"payCode": "133880443140539914",
	"srcReserved": "备注信息",
	"orderDesc": "火锅消费订单(寄海火锅福州东街口店)",
	"storeId": "ae48ee12-f901-4cbb-ba0d-4c7497a26c23",
}
```

**返回体**

```json
{
	"errCode": "00",
	"errInfo": "10000成功响应码",
	"transactionTime": "164207",
	"transactionDate": "0923",
	"settlementDate": "0923",
	"transactionDateWithYear": "20250923",
	"settlementDateWithYear": "20250923",
	"retrievalRefNum": "15005163609N",
	"transactionAmount": 2,
	"actualTransactionAmount": 2,
	"amount": 2,
	"orderId": "20250923164207406370624529",
	"thirdPartyDiscountInstrution": "微信钱包支付0.02元",
	"thirdPartyDiscountInstruction": "微信钱包支付0.02元",
	"thirdPartyName": "微信钱包",
	"userId": "otdJ_uA6R4R0beWem2hMly22xLzo",
	"thirdPartyBuyerId": "otdJ_uA6R4R0beWem2hMly22xLzo",
	"thirdPartyOrderId": "4200002792202509233617110129",
	"thirdPartyPayInformation": "现金:2",
	"cardAttr": "03",
	"mchntName": "福州市鼓楼区寄海餐饮管理有限公司",
	"chnlType": "WXPay",
	"promotionAmt": "0",
	"systemTraceNum": "000015"
}
```



#### 2.2 撤消支付

全单撤消退款

**方法定义**

```go
func (b *bscPay) VoidPay(req *BSCVoidPayReq) (*BSCVoidPayResp, error)
```

**请求体**

```json
{
	"merchantOrderId": "JHTEST20250923160600",
	"originalOrderId": "20250923160600406370168105",
	"merchantRemark": "测试取消订单"
}
```

**返回体**

```json
{
	"errCode": "00",
	"errInfo": "10000成功响应码",
	"transactionTime": "162421",
	"transactionDate": "0923",
	"settlementDate": "0923",
	"transactionDateWithYear": "20250923",
	"settlementDateWithYear": "20250923",
	"retrievalRefNum": "15005095104N",
	"thirdPartyName": "微信钱包",
	"transactionAmount": 1,
	"cardAttr": "03",
	"orderId": "20250923160600406370168105"
}
```



#### 2.3 退款(部分退款)

订单部分退款，退款金额不允许超过实付金额

**方法定义**

```go
func (b *bscPay) Refund(req *BSCRefundReq) (*BSCRefundResp, error)
```

**请求体**

```json
{
	"merchantOrderId": "JHTEST20250923164207",
	"originalOrderId": "20250923164207406370624529",
	"merchantRemark": "测试退款订单",
	"refundRequestId": "63a9055b56945b5864f051735f678751",
	"transactionAmount": 1,
	"refundDesc": "测试退款"
}
```

**返回体**

```json
{
	"errCode": "00",
	"errInfo": "10000成功响应码",
	"transactionTime": "164325",
	"transactionDate": "0923",
	"settlementDate": "0923",
	"transactionDateWithYear": "20250923",
	"settlementDateWithYear": "20250923",
	"retrievalRefNum": "15005163609N",
	"thirdPartyName": "微信钱包",
	"cardAttr": "03",
	"refundInvoiceAmount": "1",
	"transactionAmount": 1,
	"refundPromotionList": "0|0|0",
	"systemTraceNum": "000016",
	"chnlType": "WXPay",
	"acqInstCode": "UNIONPAY",
	"orderId": "20250923164207406370624529",
	"refundOrderId": "20250923164325406370591285",
	"refundTargetOrderId": "50102804702025092336590324930"
}
```



#### 2.4 支付订单查询

查询支付订单详情

**方法定义**

```go
func (b *bscPay) Query(req *BSCQueryReq) (*BSCQueryResp, error)
```

**请求体**

```json
{
	"merchantOrderId": "JHTEST20250923164207",
	"originalOrderId": "20250923164207406370624529",
	"oriTransDate": "20250923"
}
```

**返回体**

```json
{
	"errCode": "00",
	"errInfo": "00000成功响应码",
	"originalTransactionTime": "20250923164213",
	"queryResCode": "3",
	"queryResDesc": "已退货",
	"originalPayCode": "133880443140539914",
	"originalBatchNo": "000001",
	"originalSystemTraceNum": "000015",
	"originalRetrievalRefNum": "15005163609N",
	"origialRetrievalRefNum": "15005163609N",
	"originalTransactionAmount": 2,
	"amount": 2,
	"refundedAmount": 0,
	"refundAmonunt": 0,
	"actualTransactionAmount": 2,
	"thirdPartyName": "微信钱包",
	"thirdPartyBuyerId": "otdJ_uA6R4R0beWem2hMly22xLzo",
	"thirdPartyOrderId": "4200002792202509233617110129",
	"thirdPartyPayInformation": "现金:2",
	"orderId": "20250923164207406370624529",
	"merchantOrderId": "JHTEST20250923164207",
	"orderStatus": "TRADE_REFUND",
	"orderCloseTime": "2025-09-23 16:43:26",
	"originalSettlementDate": "20250923",
	"mchntName": "寄海海鲜火锅",
	"issInstCode": "交通银行(信用卡)"
}
```



#### 2.5 退款查询

查询订单退款记录

**方法定义**

```go
func (b *bscPay) RefundQuery(req *BSCRefundQueryReq) (*BSCRefundQueryResp, error)
```

**请求体**

```json
{
	"merchantOrderId": "JHTEST20250923164207",
	"originalOrderId": "20250923164207406370624529",
	"refundRequestId": "63a9055b56945b5864f051735f678751",
	"oriTransDate": "20250923"
}
```

**返回体**

```json
{
	"errCode": "00",
	"errInfo": "00000成功响应码",
	"transactionTime": "170149",
	"transactionDate": "0923",
	"settlementDate": "0923",
	"retrievalRefNum": "15005163609N",
	"queryResCode": "00",
	"payCode": "133880443140539914",
	"dealDate": "20250923",
	"dealTime": "164325",
	"originalAmount": "1",
	"dealType": "refund",
	"dealSystemTraceNum": "000016",
	"dealRetrievalRefNum": "15005163609N",
	"batchNo": "000001",
	"originalTransactionDate": "0923",
	"origialRetrievalRefNum": "15005163609N",
	"originalSettlementDate": "20250923",
	"refundInvoiceAmount": "1",
	"refundPromotionList": "0|0|0",
	"thirdPartyName": "微信钱包",
	"refundOrderId": "20250923164325406370591285",
	"refundTargetOrderId": "50102804702025092336590324930"
}
```
