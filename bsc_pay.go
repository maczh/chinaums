package chinaums

import (
	"fmt"
	"github.com/levigross/grequests"
	"strings"
)

type bscPay struct{}

//商户扫客户的条码/二维码

const (
	CNY = "156"
	// 支付模式
	PayModeScan  = "CODE_SCAN"
	PayModeECash = "E_CASH"
	PayModeSound = "SOUNDWAVE"
	PayModeNFC   = "NFC"
	PayModeFace  = "FACE_SCAN"

	DeviceTypeBC = "11"
)

// Pay 商户扫码支付(条码)
func (b *bscPay) Pay(req *BSCPayReq) (*BSCPayResp, error) {
	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}
	if req.TransactionAmount <= 0 {
		return nil, fmt.Errorf("transaction amount is invalid")
	}
	if req.MerchantOrderId == "" {
		return nil, fmt.Errorf("merchant order id is empty")
	}
	if req.PayCode == "" {
		return nil, fmt.Errorf("pay code is empty")
	}
	if req.StoreId == "" {
		return nil, fmt.Errorf("store id is empty")
	}
	req.MerchantCode = config.MerchantId
	req.TerminalCode = config.TerminalId
	req.TransactionCurrencyCode = CNY
	req.PayMode = PayModeScan
	req.DeviceType = DeviceTypeBC
	req.Longitude = StoreGISMap[req.StoreId].Longitude
	req.Latitude = StoreGISMap[req.StoreId].Latitude
	//req.StoreId = strings.ReplaceAll(req.StoreId, "-", "")
	req.StoreId = ""
	jsonstr := ToJSON(req)
	fmt.Printf("请求JSON: %s\n", jsonstr)
	sign, err := getOpenBodySig(config.AppId, config.AppKey, jsonstr)
	if err != nil {
		log.Errorf("get open body sig failed: %v", err)
		return nil, fmt.Errorf("get open body sig failed: %w", err)
	}
	fmt.Printf("接口地址: %s\n", config.ApiHost+URI_BC)
	resp, err := grequests.DoRegularRequest("POST", config.ApiHost+URI_BC, &grequests.RequestOptions{
		Headers: map[string]string{
			"Content-Type":  "application/json; charset=UTF-8",
			"Authorization": sign,
		},
		RequestBody: strings.NewReader(jsonstr),
	})
	if err != nil {
		fmt.Printf("post failed: %v\n", err)
		return nil, fmt.Errorf("post failed: %w", err)
	}
	if resp == nil {
		return nil, fmt.Errorf("resp is nil")
	}
	fmt.Printf("resp: %s\n", resp.String())
	var res BSCPayResp
	FromJSON(resp.String(), &res)
	return &res, nil
}

// VoidPay 撤销支付
func (b *bscPay) VoidPay(req *BSCVoidPayReq) (*BSCVoidPayResp, error) {
	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}
	if req.MerchantOrderId == "" {
		return nil, fmt.Errorf("merchant order id is empty")
	}
	if req.OriginalOrderId == "" {
		return nil, fmt.Errorf("original order id is empty")
	}
	req.MerchantCode = config.MerchantId
	req.TerminalCode = config.TerminalId
	jsonstr := ToJSON(req)
	fmt.Printf("请求JSON: %s\n", jsonstr)
	sign, err := getOpenBodySig(config.AppId, config.AppKey, jsonstr)
	if err != nil {
		log.Errorf("get open body sig failed: %v", err)
		return nil, fmt.Errorf("get open body sig failed: %w", err)
	}
	fmt.Printf("接口地址: %s\n", config.ApiHost+URI_BC_VOID)
	resp, err := grequests.DoRegularRequest("POST", config.ApiHost+URI_BC_VOID, &grequests.RequestOptions{
		Headers: map[string]string{
			"Content-Type":  "application/json; charset=UTF-8",
			"Authorization": sign,
		},
		RequestBody: strings.NewReader(jsonstr),
	})
	if err != nil {
		fmt.Printf("post failed: %v\n", err)
		return nil, fmt.Errorf("post failed: %w", err)
	}
	if resp == nil {
		return nil, fmt.Errorf("resp is nil")
	}
	fmt.Printf("resp: %s\n", resp.String())
	var res BSCVoidPayResp
	FromJSON(resp.String(), &res)
	return &res, nil
}

// Refund 退款
func (b *bscPay) Refund(req *BSCRefundReq) (*BSCRefundResp, error) {
	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}
	if req.MerchantOrderId == "" {
		return nil, fmt.Errorf("merchant order id is empty")
	}
	if req.OriginalOrderId == "" {
		return nil, fmt.Errorf("original order id is empty")
	}
	if req.RefundRequestId == "" {
		return nil, fmt.Errorf("refund request id is empty")
	}
	if req.TransactionAmount <= 0 {
		return nil, fmt.Errorf("transaction amount is invalid")
	}
	req.MerchantCode = config.MerchantId
	req.TerminalCode = config.TerminalId
	jsonstr := ToJSON(req)
	fmt.Printf("请求JSON: %s\n", jsonstr)
	sign, err := getOpenBodySig(config.AppId, config.AppKey, jsonstr)
	if err != nil {
		log.Errorf("get open body sig failed: %v", err)
		return nil, fmt.Errorf("get open body sig failed: %w", err)
	}
	fmt.Printf("接口地址: %s\n", config.ApiHost+URI_BC_REFUND)
	resp, err := grequests.DoRegularRequest("POST", config.ApiHost+URI_BC_REFUND, &grequests.RequestOptions{
		Headers: map[string]string{
			"Content-Type":  "application/json; charset=UTF-8",
			"Authorization": sign,
		},
		RequestBody: strings.NewReader(jsonstr),
	})
	if err != nil {
		fmt.Printf("post failed: %v\n", err)
		return nil, fmt.Errorf("post failed: %w", err)
	}
	if resp == nil {
		return nil, fmt.Errorf("resp is nil")
	}
	fmt.Printf("resp: %s\n", resp.String())
	var res BSCRefundResp
	FromJSON(resp.String(), &res)
	return &res, nil
}

// Query 查询订单
func (b *bscPay) Query(req *BSCQueryReq) (*BSCQueryResp, error) {
	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}
	if req.MerchantOrderId == "" {
		return nil, fmt.Errorf("merchant order id is empty")
	}
	if req.OriginalOrderId == "" {
		return nil, fmt.Errorf("original order id is empty")
	}
	req.MerchantCode = config.MerchantId
	req.TerminalCode = config.TerminalId
	jsonstr := ToJSON(req)
	fmt.Printf("请求JSON: %s\n", jsonstr)
	sign, err := getOpenBodySig(config.AppId, config.AppKey, jsonstr)
	if err != nil {
		log.Errorf("get open body sig failed: %v", err)
		return nil, fmt.Errorf("get open body sig failed: %w", err)
	}
	fmt.Printf("接口地址: %s\n", config.ApiHost+URI_BC_QUERY)
	resp, err := grequests.DoRegularRequest("POST", config.ApiHost+URI_BC_QUERY, &grequests.RequestOptions{
		Headers: map[string]string{
			"Content-Type":  "application/json; charset=UTF-8",
			"Authorization": sign,
		},
		RequestBody: strings.NewReader(jsonstr),
	})
	if err != nil {
		fmt.Printf("post failed: %v\n", err)
		return nil, fmt.Errorf("post failed: %w", err)
	}
	if resp == nil {
		return nil, fmt.Errorf("resp is nil")
	}
	fmt.Printf("resp: %s\n", resp.String())
	var res BSCQueryResp
	FromJSON(resp.String(), &res)
	return &res, nil
}

// RefundQuery 退款查询
func (b *bscPay) RefundQuery(req *BSCRefundQueryReq) (*BSCRefundQueryResp, error) {
	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}
	if req.MerchantOrderId == "" {
		return nil, fmt.Errorf("merchant order id is empty")
	}
	if req.OriginalOrderId == "" {
		return nil, fmt.Errorf("original order id is empty")
	}
	if req.RefundRequestId == "" {
		return nil, fmt.Errorf("refund request id is empty")
	}
	if req.OriTransDate == "" {
		return nil, fmt.Errorf("ori trans date is empty")
	}
	req.MerchantCode = config.MerchantId
	req.TerminalCode = config.TerminalId
	jsonstr := ToJSON(req)
	fmt.Printf("请求JSON: %s\n", jsonstr)
	sign, err := getOpenBodySig(config.AppId, config.AppKey, jsonstr)
	if err != nil {
		log.Errorf("get open body sig failed: %v", err)
		return nil, fmt.Errorf("get open body sig failed: %w", err)
	}
	fmt.Printf("接口地址: %s\n", config.ApiHost+URI_BC_REFUND_QUERY)
	resp, err := grequests.DoRegularRequest("POST", config.ApiHost+URI_BC_REFUND_QUERY, &grequests.RequestOptions{
		Headers: map[string]string{
			"Content-Type":  "application/json; charset=UTF-8",
			"Authorization": sign,
		},
		RequestBody: strings.NewReader(jsonstr),
	})
	if err != nil {
		fmt.Printf("post failed: %v\n", err)
		return nil, fmt.Errorf("post failed: %w", err)
	}
	if resp == nil {
		return nil, fmt.Errorf("resp is nil")
	}
	fmt.Printf("resp: %s\n", resp.String())
	var res BSCRefundQueryResp
	FromJSON(resp.String(), &res)
	return &res, nil
}
