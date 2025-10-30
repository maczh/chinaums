package chinaums

import (
	"strings"
	"time"

	"github.com/levigross/grequests"
)

type businessUnifyMulti struct {
	RsaPrivateKeyFile string
}

func (b *businessUnifyMulti) BalanceQuery(req BumPendingBalanceReq) (*BumPendingBalanceResp, error) {
	req.BumReqHeader = newBumReqHeader(TRANS_CODE_BUM_BALANCE_QUERY)
	sign, err := BumSign(req, b.RsaPrivateKeyFile)
	if err != nil {
		return nil, err
	}
	req.BumReqHeader.Signature = sign
	log.Debug("请求JSON: %s", ToJSON(req))
	resp, err := grequests.DoRegularRequest("POST", BUM_HOST+URI_BUM+TRANS_CODE_BUM_BALANCE_QUERY, &grequests.RequestOptions{
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
	var res BumPendingBalanceResp
	err = resp.JSON(&res)
	return &res, err
}

func (b *businessUnifyMulti) OrderQuery(req *BumOrderQueryReq) (*BumOrderQueryResp, error) {
	req.BumReqHeader = newBumReqHeader(TRANS_CODE_BUM_ORDER_QUERY)
	sign, err := BumSign(req, b.RsaPrivateKeyFile)
	if err != nil {
		return nil, err
	}
	req.BumReqHeader.Signature = sign
	log.Debug("请求JSON: %s", ToJSON(req))
	resp, err := grequests.DoRegularRequest("POST", BUM_HOST+URI_BUM+TRANS_CODE_BUM_ORDER_QUERY, &grequests.RequestOptions{
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
	var res BumPendingBalanceResp
	err = resp.JSON(&res)
	return &res, err
}

func (b *businessUnifyMulti) Transfer(req *BumTransferReq) (*BumTransferResp, error) {
	return nil, nil
}

func (b *businessUnifyMulti) Allocation(req *BumAllocationReq) (*BumAllocationResp, error) {
	return nil, nil
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
		GroupId:    BUM_GROUP_ID,
		Signature:  "",
	}
}
