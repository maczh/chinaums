package chinaums

import (
	"time"

	"github.com/levigross/grequests"
)

type businessUnifyMulti struct {
	PrivateKeyFile string
	PrivateKeyPwd  string
	PublicKeyFile  string
}

// 202017
func (b *businessUnifyMulti) BalanceQuery(req BumPendingBalanceReq) (*BumPendingBalanceResp, error) {
	req.BumReqHeader = newBumReqHeader(TRANS_CODE_BUM_BALANCE_QUERY)
	sign, err := BumSign(req, b.PrivateKeyFile, b.PrivateKeyPwd)
	if err != nil {
		return nil, err
	}
	req.BumReqHeader.Signature = sign
	log.Debug("请求JSON: " + ToJSON(req))
	resp, err := grequests.DoRegularRequest("POST", config.BumHost+URI_BUM+TRANS_CODE_BUM_BALANCE_QUERY, &grequests.RequestOptions{
		Headers: map[string]string{
			"Content-Type": "application/json; charset=UTF-8",
		},
		JSON: req,
	})
	if err != nil {
		return nil, err
	}
	if !resp.Ok {
		return nil, resp.Error
	}
	log.Debugf("接口返回结果: %s", resp.String())
	var res BumPendingBalanceResp
	err = resp.JSON(&res)
	return &res, err
}

func (b *businessUnifyMulti) MerchantQuery(req BumQueryMerchantReq) (*BumQueryMerchantResp, error) {
	req.BumReqHeader = newBumReqHeader(TRANS_CODE_BUM_MERCHANT_QUERY)
	sign, err := BumSign(req, b.PrivateKeyFile, b.PrivateKeyPwd)
	if err != nil {
		return nil, err
	}
	req.BumReqHeader.Signature = sign
	log.Debug("请求JSON: " + ToJSON(req))
	resp, err := grequests.DoRegularRequest("POST", config.BumHost+URI_BUM+TRANS_CODE_BUM_MERCHANT_QUERY, &grequests.RequestOptions{
		Headers: map[string]string{
			"Content-Type": "application/json; charset=UTF-8",
		},
		JSON: req,
	})
	if err != nil {
		return nil, err
	}
	if !resp.Ok {
		return nil, resp.Error
	}
	log.Debugf("接口返回结果: %s", resp.String())
	var res BumQueryMerchantResp
	err = resp.JSON(&res)
	return &res, err
}

// 202018
func (b *businessUnifyMulti) OrderQuery(req BumRechargeOrderQueryReq) (*QueryRechargeOrderResp, error) {
	req.BumReqHeader = newBumReqHeader(TRANS_CODE_BUM_ORDER_QUERY)
	sign, err := BumSign(req, b.PrivateKeyFile, b.PrivateKeyPwd)
	if err != nil {
		return nil, err
	}
	req.BumReqHeader.Signature = sign
	log.Debug("请求JSON: " + ToJSON(req))
	resp, err := grequests.DoRegularRequest("POST", config.BumHost+URI_BUM+TRANS_CODE_BUM_ORDER_QUERY, &grequests.RequestOptions{
		Headers: map[string]string{
			"Content-Type": "application/json; charset=UTF-8",
		},
		JSON: req,
	})
	if err != nil {
		return nil, err
	}
	if !resp.Ok {
		return nil, resp.Error
	}
	log.Debugf("接口返回结果: %s", resp.String())
	var res QueryRechargeOrderResp
	err = resp.JSON(&res)
	return &res, err
}

// 按订单金额划付请求(202001)
func (b *businessUnifyMulti) TransferByOrder(req *BumOrderTransferReq) (*BumRespHeader, error) {
	req.BumReqHeader = newBumReqHeader(TRANS_CODE_BUM_ORDER_TRANSFER)
	sign, err := BumSign(req, b.PrivateKeyFile, b.PrivateKeyPwd)
	if err != nil {
		return nil, err
	}
	req.BumReqHeader.Signature = sign
	log.Debug("请求JSON: " + ToJSON(req))
	resp, err := grequests.DoRegularRequest("POST", config.BumHost+URI_BUM+TRANS_CODE_BUM_ORDER_TRANSFER, &grequests.RequestOptions{
		Headers: map[string]string{
			"Content-Type": "application/json; charset=UTF-8",
		},
		JSON: req,
	})
	if err != nil {
		return nil, err
	}
	if !resp.Ok {
		return nil, resp.Error
	}
	log.Debugf("接口返回结果: %s", resp.String())
	var res BumRespHeader
	err = resp.JSON(&res)
	return &res, err
}

// 按金额划付请求(202002)
func (b *businessUnifyMulti) Transfer(req *BumTransferReq) (*BumRespHeader, error) {
	req.BumReqHeader = newBumReqHeader(TRANS_CODE_BUM_TRANSFER)
	sign, err := BumSign(req, b.PrivateKeyFile, b.PrivateKeyPwd)
	if err != nil {
		return nil, err
	}
	req.BumReqHeader.Signature = sign
	log.Debug("请求JSON: " + ToJSON(req))
	resp, err := grequests.DoRegularRequest("POST", config.BumHost+URI_BUM+TRANS_CODE_BUM_TRANSFER, &grequests.RequestOptions{
		Headers: map[string]string{
			"Content-Type": "application/json; charset=UTF-8",
		},
		JSON: req,
	})
	if err != nil {
		return nil, err
	}
	if !resp.Ok {
		return nil, resp.Error
	}
	log.Debugf("接口返回结果: %s", resp.String())
	var res BumRespHeader
	err = resp.JSON(&res)
	return &res, err
}

// 按订单金额分账请求(202003)
func (b *businessUnifyMulti) AllocationByOrder(req *BumOrderAllocationReq) (*BumRespHeader, error) {
	req.BumReqHeader = newBumReqHeader(TRANS_CODE_BUM_ORDER_ALLOCATION)
	sign, err := BumSign(req, b.PrivateKeyFile, b.PrivateKeyPwd)
	if err != nil {
		return nil, err
	}
	req.BumReqHeader.Signature = sign
	log.Debug("请求JSON: " + ToJSON(req))
	resp, err := grequests.DoRegularRequest("POST", config.BumHost+URI_BUM+TRANS_CODE_BUM_ORDER_ALLOCATION, &grequests.RequestOptions{
		Headers: map[string]string{
			"Content-Type": "application/json; charset=UTF-8",
		},
		JSON: req,
	})
	if err != nil {
		return nil, err
	}
	if !resp.Ok {
		return nil, resp.Error
	}
	log.Debugf("接口返回结果: %s", resp.String())
	var res BumRespHeader
	err = resp.JSON(&res)
	return &res, err
}

// 按金额分账请求(202004)
func (b *businessUnifyMulti) Allocation(req *BumAllocationReq) (*BumRespHeader, error) {
	req.BumReqHeader = newBumReqHeader(TRANS_CODE_BUM_ALLOCATION)
	sign, err := BumSign(req, b.PrivateKeyFile, b.PrivateKeyPwd)
	if err != nil {
		return nil, err
	}
	req.BumReqHeader.Signature = sign
	log.Debug("请求JSON: " + ToJSON(req))
	resp, err := grequests.DoRegularRequest("POST", config.BumHost+URI_BUM+TRANS_CODE_BUM_ALLOCATION, &grequests.RequestOptions{
		Headers: map[string]string{
			"Content-Type": "application/json; charset=UTF-8",
		},
		JSON: req,
	})
	if err != nil {
		return nil, err
	}
	if !resp.Ok {
		return nil, resp.Error
	}
	log.Debugf("接口返回结果: %s", resp.String())
	var res BumRespHeader
	err = resp.JSON(&res)
	return &res, err
}

func newBumReqHeader(transCode string) BumReqHeader {
	reqId, _ := generateNonce()
	return BumReqHeader{
		TransCode:  transCode,
		VerNo:      BUM_VER,
		SrcReqDate: time.Now().Format("20060102"),
		SrcReqTime: time.Now().Format("150405"),
		SrcReqId:   reqId,
		ChannelId:  BUM_CHANNEL_ID,
		GroupId:    config.BumGroupId,
		Signature:  "",
	}
}
