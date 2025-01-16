package utils

import (
	"crypto/x509"
	"encoding/pem"
	"github.com/gofiber/fiber/v2/log"
)

// ParseCertificateDomain 解析证书内容并验证域名是否包含在证书的域名列表中
func ParseCertificateDomain(certContent string, domains []string) bool {
	// 解码 PEM 格式的证书
	block, _ := pem.Decode([]byte(certContent))
	if block == nil || block.Type != "CERTIFICATE" {
		log.Errorf("无效的证书内容")
		return false
	}
	// 解析证书
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Errorf("无法解析证书: %v", err)
		return false
	}
	for _, domain := range domains {
		exists := false
		for _, d := range cert.DNSNames {
			if d == domain {
				exists = true
			}
		}
		if !exists {
			return false
		}
	}
	return true
}
