package chinaums

import (
	"fmt"
	"github.com/levigross/grequests"
	"strings"
	"time"
)

type wxAppPay struct{}

const (
	InstMidDefault = "MINIDEFAULT"
	TradeTypeMini  = "MINI"
)

// Pay 小程序支付
func (w *wxAppPay) Pay(req *WxAppPayReq) (*WxAppPayResp, error) {
	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}

	if req.MerOrderId == "" {
		return nil, fmt.Errorf("merchant order id is empty")
	}
	if req.TotalAmount <= 0 {
		return nil, fmt.Errorf("total amount is invalid")
	}
	if req.Mid == "" {
		req.Mid = config.MerchantIdCB
	}
	if req.Tid == "" {
		req.Tid = config.TerminalIdCB
	}
	if req.SubOpenId == "" {
		return nil, fmt.Errorf("wechat openid is empty")
	}
	if req.SubAppId == "" {
		return nil, fmt.Errorf("wechat appid is empty")
	}
	req.RequestTimestamp = time.Now().Format("2006-01-02 15:04:05")
	req.InstMid = InstMidDefault
	req.TradeType = TradeTypeMini
	jsonstr := ToJSON(req)
	fmt.Printf("请求JSON: %s\n", jsonstr)
	sign, err := getOpenBodySig(config.AppId, config.AppKey, jsonstr)
	if err != nil {
		log.Errorf("get open body sig failed: %v", err)
		return nil, fmt.Errorf("get open body sig failed: %w", err)
	}
	fmt.Printf("接口地址: %s\n", config.ApiHost+URI_WX_APP_PAY)
	resp, err := grequests.DoRegularRequest("POST", config.ApiHost+URI_WX_APP_PAY, &grequests.RequestOptions{
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
	var res WxAppPayResp
	FromJSON(resp.String(), &res)
	return &res, nil
}

func (w *wxAppPay) Query(req *WxAppQueryReq) (*WxAppQueryResp, error) {
	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}

	if req.MerOrderId == "" {
		return nil, fmt.Errorf("merchant order id is empty")
	}
	if req.TargetOrderId == "" {
		return nil, fmt.Errorf("target order id is empty")
	}
	if req.Mid == "" {
		req.Mid = config.MerchantIdCB
	}
	if req.Tid == "" {
		req.Tid = config.TerminalIdCB
	}
	req.RequestTimestamp = time.Now().Format("2006-01-02 15:04:05")
	req.InstMid = InstMidDefault
	jsonstr := ToJSON(req)
	fmt.Printf("请求JSON: %s\n", jsonstr)
	sign, err := getOpenBodySig(config.AppId, config.AppKey, jsonstr)
	if err != nil {
		log.Errorf("get open body sig failed: %v", err)
		return nil, fmt.Errorf("get open body sig failed: %w", err)
	}
	fmt.Printf("接口地址: %s\n", config.ApiHost+URI_WX_APP_QUERY)
	resp, err := grequests.DoRegularRequest("POST", config.ApiHost+URI_WX_APP_QUERY, &grequests.RequestOptions{
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
	var res WxAppQueryResp
	FromJSON(resp.String(), &res)
	return &res, nil
}

func (w *wxAppPay) Refund(req *WxAppRefundReq) (*WxAppRefundResp, error) {
	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}

	if req.MerOrderId == "" {
		return nil, fmt.Errorf("merchant order id is empty")
	}
	if req.TargetOrderId == "" {
		return nil, fmt.Errorf("target order id is empty")
	}
	if req.RefundOrderId == "" {
		return nil, fmt.Errorf("refund order id is empty")
	}
	if req.RefundAmount <= 0 {
		return nil, fmt.Errorf("refund amount is invalid")
	}
	if req.Mid == "" {
		req.Mid = config.MerchantIdCB
	}
	if req.Tid == "" {
		req.Tid = config.TerminalIdCB
	}
	req.RequestTimestamp = time.Now().Format("2006-01-02 15:04:05")
	req.InstMid = InstMidDefault
	jsonstr := ToJSON(req)
	fmt.Printf("请求JSON: %s\n", jsonstr)
	sign, err := getOpenBodySig(config.AppId, config.AppKey, jsonstr)
	if err != nil {
		log.Errorf("get open body sig failed: %v", err)
		return nil, fmt.Errorf("get open body sig failed: %w", err)
	}
	fmt.Printf("接口地址: %s\n", config.ApiHost+URI_WX_APP_REFUND)
	resp, err := grequests.DoRegularRequest("POST", config.ApiHost+URI_WX_APP_REFUND, &grequests.RequestOptions{
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
	var res WxAppRefundResp
	FromJSON(resp.String(), &res)
	return &res, nil
}

func (w *wxAppPay) RefundQuery(req *WxAppRefundQueryReq) (*WxAppRefundQueryResp, error) {
	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}

	if req.MerOrderId == "" {
		return nil, fmt.Errorf("merchant order id is empty")
	}
	if req.Mid == "" {
		req.Mid = config.MerchantIdCB
	}
	if req.Tid == "" {
		req.Tid = config.TerminalIdCB
	}
	req.RequestTimestamp = time.Now().Format("2006-01-02 15:04:05")
	req.InstMid = InstMidDefault
	jsonstr := ToJSON(req)
	fmt.Printf("请求JSON: %s\n", jsonstr)
	sign, err := getOpenBodySig(config.AppId, config.AppKey, jsonstr)
	if err != nil {
		log.Errorf("get open body sig failed: %v", err)
		return nil, fmt.Errorf("get open body sig failed: %w", err)
	}
	fmt.Printf("接口地址: %s\n", config.ApiHost+URI_WX_APP_REFUND_QUERY)
	resp, err := grequests.DoRegularRequest("POST", config.ApiHost+URI_WX_APP_REFUND_QUERY, &grequests.RequestOptions{
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
	var res WxAppRefundQueryResp
	FromJSON(resp.String(), &res)
	return &res, nil
}

func (w *wxAppPay) Close(req *WxAppCloseReq) (*WxAppCloseResp, error) {
	if req == nil {
		return nil, fmt.Errorf("req is nil")
	}

	if req.MerOrderId == "" {
		return nil, fmt.Errorf("merchant order id is empty")
	}
	if req.Mid == "" {
		req.Mid = config.MerchantIdCB
	}
	if req.Tid == "" {
		req.Tid = config.TerminalIdCB
	}
	req.RequestTimestamp = time.Now().Format("2006-01-02 15:04:05")
	req.InstMid = InstMidDefault
	jsonstr := ToJSON(req)
	fmt.Printf("请求JSON: %s\n", jsonstr)
	sign, err := getOpenBodySig(config.AppId, config.AppKey, jsonstr)
	if err != nil {
		log.Errorf("get open body sig failed: %v", err)
		return nil, fmt.Errorf("get open body sig failed: %w", err)
	}
	fmt.Printf("接口地址: %s\n", config.ApiHost+URI_WX_APP_CLOSE)
	resp, err := grequests.DoRegularRequest("POST", config.ApiHost+URI_WX_APP_CLOSE, &grequests.RequestOptions{
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
	var res WxAppCloseResp
	FromJSON(resp.String(), &res)
	return &res, nil
}
