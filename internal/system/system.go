package system

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/run-bigpig/conduit/internal/types"
	"github.com/run-bigpig/conduit/internal/utils"
	"os"
	"runtime"
	"strings"
)

// SyncHost 将域名添加到windows或则linux或者mac的host中
func SyncHost() error {
	var hostsFilePath string
	//判断当前系统
	switch runtime.GOOS {
	case "windows":
		hostsFilePath = "C:\\Windows\\System32\\drivers\\etc\\hosts"
	case "linux":
		hostsFilePath = "/etc/hosts"
	case "darwin":
		hostsFilePath = "/private/etc/hosts"
	default:
		return errors.New("not support os")
	}
	for _, host := range genHosts() {
		if err := writeHostsFile(hostsFilePath, host.Domain, host.Ip); err != nil {
			return err
		}
	}
	return nil
}

// writeHostsFile 写入hosts文件
func writeHostsFile(hostsFilePath string, domain string, ip string) error {
	// 检查hosts文件是否存在
	if _, err := os.Stat(hostsFilePath); os.IsNotExist(err) {
		log.Errorf("Hosts file not found at %s\n", hostsFilePath)
		return err
	}

	// 读取hosts文件内容
	content, err := os.ReadFile(hostsFilePath)
	if err != nil {
		log.Errorf("Failed to read hosts file: %v\n", err)
		return err
	}

	// 检查是否已经存在该域名
	if strings.Contains(string(content), domain) {
		log.Infof("Domain %s already exists in hosts file\n", domain)
		return nil
	}

	// 追加新的域名到hosts文件
	entry := fmt.Sprintf("\n%s\t%s\n", ip, domain)
	file, err := os.OpenFile(hostsFilePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Errorf("Failed to open hosts file: %v\n", err)
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(entry); err != nil {
		log.Errorf("Failed to write to hosts file: %v\n", err)
		return err
	}
	log.Infof("Successfully added %s -> %s to hosts file\n", domain, ip)
	return nil
}

// 生成hosts
func genHosts() []*types.Hosts {
	records := []string{"api", "copilot-proxy", "copilot-telemetry-service"}
	hosts := make([]*types.Hosts, 0, len(records))
	// 添加服务域名
	hosts = append(hosts, &types.Hosts{
		Domain: utils.GetDomain(),
		Ip:     "127.0.0.1",
	})
	// 添加子域名
	for _, record := range records {
		hosts = append(hosts, &types.Hosts{
			Domain: fmt.Sprintf("%s.%s", record, utils.GetDomain()),
			Ip:     "127.0.0.1",
		})
	}
	return hosts
}
