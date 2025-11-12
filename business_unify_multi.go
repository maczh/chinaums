package chinaums

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
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

// GenerateChkFile 生成文件的.chk校验文件
// 参数:
//
//	filePath: 待签名的原文件路径（如XX_XXXXXX_YYYYMMDDHHMMSS.txt）
//	privateKeyPath: 商户私钥文件路径（PEM格式）
//	privateKeyPassword: 私钥解密密码（测试环境为"123456"）
//	chkFilePath: 输出的.chk文件路径
//
// 返回:
//
//	错误信息
func GenerateChkFile(filePath, privateKeyPath, privateKeyPassword, chkFilePath string) error {
	// 1. 读取原文件内容
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("读取原文件失败: %v", err)
	}

	// 2. 计算文件MD5哈希值（字节数组）
	md5Hash := md5.Sum(fileContent)
	md5Bytes := md5Hash[:] // MD5结果为16字节数组

	// 3. 加载私钥
	privateKey, err := loadPrivateKey(privateKeyPath, privateKeyPassword)
	if err != nil {
		return fmt.Errorf("加载私钥失败: %v", err)
	}

	// 4. 使用SHA256withRSA签名MD5字节数组
	signature, err := rsa.SignPKCS1v15(
		rand.Reader,
		privateKey,
		crypto.SHA256, // 文档要求SHA256withRSA算法
		md5Bytes,
	)
	if err != nil {
		return fmt.Errorf("签名失败: %v", err)
	}

	// 5. 将签名结果转为十六进制字符串
	signHex := hex.EncodeToString(signature)

	// 6. 写入.chk文件
	if err := ioutil.WriteFile(chkFilePath, []byte(signHex), 0644); err != nil {
		return fmt.Errorf("写入.chk文件失败: %v", err)
	}

	return nil
}

// 加载带密码的RSA私钥（PEM格式）
func loadPrivateKey(keyPath, password string) (*rsa.PrivateKey, error) {
	keyData, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}

	// 解码PEM块
	block, _ := pem.Decode(keyData)
	if block == nil {
		return nil, errors.New("私钥PEM格式错误")
	}

	// 若私钥加密，使用密码解密
	var keyBytes []byte
	if x509.IsEncryptedPEMBlock(block) {
		keyBytes, err = x509.DecryptPEMBlock(block, []byte(password))
		if err != nil {
			return nil, fmt.Errorf("私钥解密失败: %v", err)
		}
	} else {
		keyBytes = block.Bytes
	}

	// 解析PKCS#1格式私钥
	privateKey, err := x509.ParsePKCS1PrivateKey(keyBytes)
	if err != nil {
		// 尝试解析PKCS#8格式
		pkcs8Key, err := x509.ParsePKCS8PrivateKey(keyBytes)
		if err != nil {
			return nil, fmt.Errorf("解析私钥失败: %v", err)
		}
		privateKey, ok := pkcs8Key.(*rsa.PrivateKey)
		if !ok {
			return nil, errors.New("私钥非RSA类型")
		}
		return privateKey, nil
	}
	return privateKey, nil
}

// VerifyRetChkFile 验证回盘文件的.chk签名
// 参数:
//
//	retFilePath: 回盘文件路径（如XX_XXXXXX_YYYYMMDDHHMMSS.txt.ret）
//	chkFilePath: .chk校验文件路径（XX_XXXXXX_YYYYMMDDHHMMSS.txt.ret.chk）
//	publicKeyPath: 银商公钥文件路径（PEM格式）
//
// 返回:
//
//	验签是否成功，错误信息
func VerifyRetChkFile(retFilePath, chkFilePath, publicKeyPath string) (bool, error) {
	// 1. 读取回盘文件内容，计算MD5
	retContent, err := ioutil.ReadFile(retFilePath)
	if err != nil {
		return false, fmt.Errorf("读取回盘文件失败: %v", err)
	}
	md5Hash := md5.Sum(retContent)
	md5Bytes := md5Hash[:]

	// 2. 读取.chk文件内容（十六进制签名字符串）
	chkContent, err := ioutil.ReadFile(chkFilePath)
	if err != nil {
		return false, fmt.Errorf("读取.chk文件失败: %v", err)
	}
	signature, err := hex.DecodeString(string(chkContent))
	if err != nil {
		return false, fmt.Errorf("chk文件内容不是有效十六进制: %v", err)
	}

	// 3. 加载银商公钥
	publicKey, err := loadPublicKey(publicKeyPath)
	if err != nil {
		return false, fmt.Errorf("加载公钥失败: %v", err)
	}

	// 4. 使用公钥验签（SHA256withRSA）
	err = rsa.VerifyPKCS1v15(
		publicKey,
		crypto.SHA256,
		md5Bytes,
		signature,
	)
	if err != nil {
		return false, fmt.Errorf("验签失败: %v", err)
	}

	return true, nil
}

// 加载RSA公钥（PEM格式）
func loadPublicKey(keyPath string) (*rsa.PublicKey, error) {
	keyData, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyData)
	if block == nil {
		return nil, errors.New("公钥PEM格式错误")
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaPubKey, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("公钥非RSA类型")
	}
	return rsaPubKey, nil
}
