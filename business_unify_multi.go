package chinaums

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/levigross/grequests"
	"golang.org/x/crypto/pkcs12"
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
	fmt.Printf("原文件MD5: %x", md5Bytes)
	hashed := sha256.Sum256(md5Bytes)
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
		hashed[:],
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
	pfxData, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, fmt.Errorf("read private key file error: %w", err)
	}

	// 首先尝试直接解析
	privateKey, _, err := pkcs12.Decode(pfxData, password)
	var priKey *rsa.PrivateKey
	if err != nil {
		// 如果直接解析失败，尝试解析所有证书
		blocks, err := pkcs12.ToPEM(pfxData, password)
		if err != nil {
			return nil, fmt.Errorf("decode PKCS12 file error: %w", err)
		}

		// 查找私钥
		for _, block := range blocks {
			if block.Type == "PRIVATE KEY" || block.Type == "RSA PRIVATE KEY" {
				privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
				if err == nil {
					priKey = privateKey.(*rsa.PrivateKey)
					break
				}
				// 尝试 PKCS8 格式
				pk, err := x509.ParsePKCS8PrivateKey(block.Bytes)
				if err == nil {
					var ok bool
					priKey, ok = pk.(*rsa.PrivateKey)
					if ok {
						break
					}
				}
			}
		}
		if priKey == nil {
			return nil, fmt.Errorf("no valid RSA private key found in PKCS12 file")
		}
	} else {
		var ok bool
		priKey, ok = privateKey.(*rsa.PrivateKey)
		if !ok {
			return nil, errors.New("not rsa private key")
		}
	}
	return priKey, nil
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

// 加载RSA公钥（CER格式）
func loadPublicKey(keyPath string) (*rsa.PublicKey, error) {
	keyData, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read public key file: %v", err)
	}

	// 首先尝试直接解析为 DER 格式的证书
	cert, err := x509.ParseCertificate(keyData)
	if err == nil {
		// 成功解析为 DER 格式证书
		if rsaPubKey, ok := cert.PublicKey.(*rsa.PublicKey); ok {
			return rsaPubKey, nil
		}
		return nil, errors.New("certificate does not contain RSA public key")
	}

	// 如果 DER 解析失败，尝试 PEM 解码
	block, _ := pem.Decode(keyData)
	if block == nil {
		return nil, errors.New("not a valid PEM format and not a valid DER certificate")
	}

	// 处理 PEM 格式
	switch block.Type {
	case "CERTIFICATE":
		// PEM 格式的证书
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("failed to parse certificate: %v", err)
		}
		rsaPubKey, ok := cert.PublicKey.(*rsa.PublicKey)
		if !ok {
			return nil, errors.New("certificate does not contain RSA public key")
		}
		return rsaPubKey, nil

	case "PUBLIC KEY":
		// PKIX 格式的公钥
		pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("failed to parse PKIX public key: %v", err)
		}
		rsaPubKey, ok := pubKey.(*rsa.PublicKey)
		if !ok {
			return nil, errors.New("public key is not of type RSA")
		}
		return rsaPubKey, nil

	case "RSA PUBLIC KEY":
		// PKCS1 格式的公钥
		return x509.ParsePKCS1PublicKey(block.Bytes)

	default:
		return nil, fmt.Errorf("unsupported PEM block type: %s", block.Type)
	}
}
