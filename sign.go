package chinaums

import (
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
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

// 私钥签名
func BumSign(data any, rsaPrivateKeyFile string) (string, error) {
	privateKeyPEM, err := os.ReadFile(rsaPrivateKeyFile)
	if err != nil {
		return "", err
	}
	block, _ := pem.Decode(privateKeyPEM)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return "", errors.New("无效的私钥格式")
	}
	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	m := struct2Map(data)
	str := mapToKV(m)
	hash := sha256.Sum256([]byte(str))
	sig, err := rsa.SignPKCS1v15(rand.Reader, pri, crypto.SHA256, hash[:])
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(sig), nil
}

// 公钥验签
func BumVerify(data any, signature string, publicKeyPEM []byte) error {
	block, _ := pem.Decode(publicKeyPEM)
	if block == nil {
		return errors.New("pem decode failed")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return errors.New("not rsa public key")
	}
	sig, _ := hex.DecodeString(signature)
	m := struct2Map(data)
	str := mapToKV(m)
	hash := sha256.Sum256([]byte(str))
	return rsa.VerifyPKCS1v15(rsaPub, crypto.SHA256, hash[:], sig)
}

// 将 map 按 key 字典序拼接成 key=value&...
func mapToKV(m map[string]string) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		if k == "signature" { // 签名本身不参与
			continue
		}
		if m[k] == "" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var b strings.Builder
	for i, k := range keys {
		if i > 0 {
			b.WriteString("&")
		}
		b.WriteString(k)
		b.WriteString("=")
		b.WriteString(m[k])
	}
	return b.String()
}

func struct2Map(data any) map[string]string {
	rs := make(map[string]string)
	str := ToJSON(data)
	FromJSON(str, &rs)
	return rs
}
