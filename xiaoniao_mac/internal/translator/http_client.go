package translator

import (
	"net"
	"net/http"
	"time"
)

// 全局共享的HTTP客户端，优化连接池配置
var sharedHTTPClient = &http.Client{
	Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   10,  // 每个host保持10个空闲连接
		IdleConnTimeout:       90 * time.Second,  // 空闲连接90秒后关闭
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		ResponseHeaderTimeout: 60 * time.Second,
		DisableCompression:    false,  // 保持压缩
	},
	Timeout: 60 * time.Second,
}

// GetSharedHTTPClient 获取共享的HTTP客户端
func GetSharedHTTPClient() *http.Client {
	return sharedHTTPClient
}

// PrewarmConnection 预热连接（用于启动时或切换模型时）
func PrewarmConnection(trans *Translator) error {
	// 发送一个简单的测试请求来预热连接
	_, err := trans.Translate("test", "Direct translation:")
	return err
}