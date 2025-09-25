package chinaums

import (
	"fmt"
	"github.com/levigross/grequests"
	"strings"
	"time"
)

type qrPay struct{}

const (
	InstMidQRPay = "QRPAYDEFAULT"
)

func (b *qrPay) Pay(req *QrPayReq) (*QrPayResp, error) {
	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}
	if req.Mid == "" {
		req.Mid = config.MerchantIdCB
	}
	if req.Tid == "" {
		req.Tid = config.TerminalIdCB
	}
	if req.BillNo == "" {
		return nil, fmt.Errorf("bill no is empty")
	}
	if req.TotalAmount <= 0 {
		return nil, fmt.Errorf("total amount is empty")
	}
	req.InstMid = InstMidQRPay
	req.RequestTimestamp = time.Now().Format("2006-01-02 15:04:05")
	req.BillDate = time.Now().Format("2006-01-02")
	jsonstr := ToJSON(req)
	fmt.Printf("请求JSON: %s\n", jsonstr)
	sign, err := getOpenBodySig(config.AppId, config.AppKey, jsonstr)
	if err != nil {
		log.Errorf("get open body sig failed: %v", err)
		return nil, fmt.Errorf("get open body sig failed: %w", err)
	}
	fmt.Printf("接口地址: %s\n", config.ApiHost+URI_QR_PAY)
	resp, err := grequests.DoRegularRequest("POST", config.ApiHost+URI_QR_PAY, &grequests.RequestOptions{
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
	var res QrPayResp
	FromJSON(resp.String(), &res)
	return &res, nil
}

func (b *qrPay) Query(req *QrPayQueryReq) (*QrPayQueryResp, error) {
	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}
	if req.Mid == "" {
		req.Mid = config.MerchantIdCB
	}
	if req.Tid == "" {
		req.Tid = config.TerminalIdCB
	}
	if req.BillNo == "" {
		return nil, fmt.Errorf("bill no is empty")
	}
	if req.BillDate == "" {
		return nil, fmt.Errorf("bill date is empty")
	}
	req.InstMid = InstMidQRPay
	req.RequestTimestamp = time.Now().Format("2006-01-02 15:04:05")
	jsonstr := ToJSON(req)
	fmt.Printf("请求JSON: %s\n", jsonstr)
	sign, err := getOpenBodySig(config.AppId, config.AppKey, jsonstr)
	if err != nil {
		log.Errorf("get open body sig failed: %v", err)
		return nil, fmt.Errorf("get open body sig failed: %w", err)
	}
	fmt.Printf("接口地址: %s\n", config.ApiHost+URI_QR_QUERY)
	resp, err := grequests.DoRegularRequest("POST", config.ApiHost+URI_QR_QUERY, &grequests.RequestOptions{
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
	var res QrPayQueryResp
	FromJSON(resp.String(), &res)
	return &res, nil
}

func (b *qrPay) Refund(req *QrPayRefundReq) (*QrPayRefundResp, error) {
	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}
	if req.Mid == "" {
		req.Mid = config.MerchantIdCB
	}
	if req.Tid == "" {
		req.Tid = config.TerminalIdCB
	}
	if req.BillNo == "" {
		return nil, fmt.Errorf("bill no is empty")
	}
	if req.RefundAmount <= 0 {
		return nil, fmt.Errorf("refund amount is empty")
	}
	req.InstMid = InstMidQRPay
	req.RequestTimestamp = time.Now().Format("2006-01-02 15:04:05")
	jsonstr := ToJSON(req)
	fmt.Printf("请求JSON: %s\n", jsonstr)
	sign, err := getOpenBodySig(config.AppId, config.AppKey, jsonstr)
	if err != nil {
		log.Errorf("get open body sig failed: %v", err)
		return nil, fmt.Errorf("get open body sig failed: %w", err)
	}
	fmt.Printf("接口地址: %s\n", config.ApiHost+URI_QR_REFUND)
	resp, err := grequests.DoRegularRequest("POST", config.ApiHost+URI_QR_REFUND, &grequests.RequestOptions{
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
	var res QrPayRefundResp
	FromJSON(resp.String(), &res)
	return &res, nil
}
