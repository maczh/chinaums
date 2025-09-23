package chinaums

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"time"
)

// 计算HMAC-SHA256
func hmacSHA256(data, key []byte) ([]byte, error) {
	mac := hmac.New(sha256.New, key)
	_, err := mac.Write(data)
	if err != nil {
		return nil, err
	}
	return mac.Sum(nil), nil
}

// 生成随机nonce
func generateNonce() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// 计算字符串的SHA256摘要
func sha256Hex(s string) (string, error) {
	h := sha256.New()
	_, err := io.WriteString(h, s)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// 生成OPEN-BODY-SIG认证的Authorization值
func getOpenBodySig(appId, appKey, body string) (string, error) {
	// 生成时间戳
	timestamp := time.Now().Format("20060102150405")

	// 生成nonce
	nonce, err := generateNonce()
	if err != nil {
		return "", fmt.Errorf("生成nonce失败: %v", err)
	}

	// 计算body的SHA256摘要
	bodyDigest, err := sha256Hex(body)
	if err != nil {
		return "", fmt.Errorf("计算body摘要失败: %v", err)
	}
	//fmt.Println("bodyDigest:")
	//fmt.Println(bodyDigest)

	// 构建待签名字符串
	str1C := appId + timestamp + nonce + bodyDigest
	//fmt.Println("str1_C:" + str1C)

	// 计算HMAC-SHA256签名
	signature, err := hmacSHA256([]byte(str1C), []byte(appKey))
	if err != nil {
		return "", fmt.Errorf("计算HMAC-SHA256签名失败: %v", err)
	}

	// Base64编码签名
	signatureBase64 := base64.StdEncoding.EncodeToString(signature)

	// 构建Authorization字符串
	authorization := fmt.Sprintf("OPEN-BODY-SIG AppId=\"%s\", Timestamp=\"%s\", Nonce=\"%s\", Signature=\"%s\"",
		appId, timestamp, nonce, signatureBase64)

	fmt.Println("Authorization:" + authorization)
	return authorization, nil
}
