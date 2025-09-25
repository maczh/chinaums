package chinaums

import (
	"testing"
	"time"
)

func TestGenNonce(t *testing.T) {
	t.Log(generateNonce())
}

func TestBSCPay(t *testing.T) {
	req := BSCPayReq{
		TransactionAmount: 2,
		MerchantOrderId:   "JHTEST" + time.Now().Format("20060102150405"),
		MerchantRemark:    "测试订单",
		PayMode:           PayModeScan,
		PayCode:           "133880443140539914",
		SrcReserved:       "备注信息",
		OrderDesc:         "火锅消费订单(寄海火锅福州东街口店)",
		StoreId:           "ae48ee12-f901-4cbb-ba0d-4c7497a26c23",
		OperatorId:        "",
		Goods:             nil,
		Longitude:         "+119.299312",
		Latitude:          "+26.087587",
	}
	resp, err := Client.BscPay.Pay(&req)
	if err != nil {
		t.Log("支付接口调用失败:", err)
		return
	}
	t.Log("支付结果:", ToJSONPretty(resp))
}

func TestBSCVoidPay(t *testing.T) {
	req := BSCVoidPayReq{
		MerchantOrderId: "JHTEST20250923160600",
		OriginalOrderId: "20250923160600406370168105",
		MerchantRemark:  "测试取消订单",
	}
	resp, err := Client.BscPay.VoidPay(&req)
	if err != nil {
		return
	}
	t.Log("取消支付结果:", ToJSONPretty(resp))
}

func TestBSCRefund(t *testing.T) {
	requestId, _ := generateNonce()
	req := BSCRefundReq{
		MerchantOrderId:   "JHTEST20250923164207",
		OriginalOrderId:   "20250923164207406370624529",
		MerchantRemark:    "测试退款订单",
		RefundRequestId:   requestId,
		TransactionAmount: 1,
		Goods:             nil,
		StoreId:           "",
		OperatorId:        "",
		RefundDesc:        "测试退款",
	}
	resp, err := Client.BscPay.Refund(&req)
	if err != nil {
		t.Log("退款接口调用失败:", err)
		return
	}
	t.Log("退款结果:", ToJSONPretty(resp))
}

func TestBSCQuery(t *testing.T) {
	req := BSCQueryReq{
		MerchantOrderId: "JHTEST20250923164207",
		OriginalOrderId: "20250923164207406370624529",
		OriTransDate:    "20250923",
	}
	resp, err := Client.BscPay.Query(&req)
	if err != nil {
		t.Log("查询接口调用失败:", err)
		return
	}
	t.Log("查询结果:", ToJSONPretty(resp))
}

func TestBSCRefundQuery(t *testing.T) {
	req := BSCRefundQueryReq{
		MerchantOrderId: "JHTEST20250923164207",
		OriginalOrderId: "20250923164207406370624529",
		OriTransDate:    "20250923",
		RefundRequestId: "63a9055b56945b5864f051735f678751",
	}
	resp, err := Client.BscPay.RefundQuery(&req)
	if err != nil {
		t.Log("退款查询接口调用失败:", err)
		return
	}
	t.Log("退款查询结果:", ToJSONPretty(resp))
}

func TestWxAppPay(t *testing.T) {
	req := WxAppPayReq{
		TotalAmount:    1,
		OriginalAmount: 1,
		MerOrderId:     "3GETJH" + time.Now().Format("20060102150405"),
		SubAppId:       "wx014ad2ef80a147a7",
		SubOpenId:      "o_GKc6wW1I3AbnCI9rXCLpPwd-R8",
		NotifyUrl:      "https://www.baidu.com",
		ReturnUrl:      "https://www.baidu.com",
		OrderDesc:      "火锅消费订单(寄海火锅福州东街口店)",
		Goods:          nil,
	}
	resp, err := Client.WxAppPay.Pay(&req)
	if err != nil {
		t.Log("小程序支付接口调用失败:", err)
		return
	}
	t.Log("微信小程序支付结果:", ToJSONPretty(resp))
}

func TestQRPay(t *testing.T) {
	req := QrPayReq{
		TotalAmount: 1,
		BillNo:      "3GETJH" + time.Now().Format("20060102150405"),
		NotifyUrl:   "https://www.baidu.com",
		ReturnUrl:   "https://www.baidu.com",
		Goods:       nil,
	}
	resp, err := Client.QRPay.Pay(&req)
	if err != nil {
		t.Log("二维码支付接口调用失败:", err)
		return
	}
	t.Log("二维码支付结果:", ToJSONPretty(resp))
}

func TestQRPayQuery(t *testing.T) {
	req := QrPayQueryReq{
		BillNo:   "3GETJH20250925110003",
		BillDate: "20250925",
	}
	resp, err := Client.QRPay.Query(&req)
	if err != nil {
		t.Log("二维码支付查询接口调用失败:", err)
		return
	}
	t.Log("二维码支付查询结果:", ToJSONPretty(resp))
}

func TestQRPayRefund(t *testing.T) {
	req := QrPayRefundReq{
		BillNo:       "3GETJH20250925110003",
		BillDate:     "20250925",
		RefundAmount: 1,
	}
	resp, err := Client.QRPay.Refund(&req)
	if err != nil {
		t.Log("二维码支付退款接口调用失败:", err)
		return
	}
	t.Log("二维码支付退款结果:", ToJSONPretty(resp))
}
