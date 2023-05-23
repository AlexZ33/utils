package ip

import (
	"github.com/AlexZ33/utils/errors"
	"net"
	"net/http"
	"strings"
)

// GetLocalIP 获取本机IP地址
// 本代码未考虑IPv6
func GetLocalIP() (ip string, error error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				return ip, nil
			}
		}
	}
	return "", errors.New("获取本机IP地址失败")
}

// IsLocalIP 判断是否是内网ip
// 关于内网地址的判断，请不要忽略IPv6的回环地址和IPv6的唯一本地地址
// 因为URL形式多样，可以使用DNS解析获取规范的IP，从而判断是否是内网资源。
func IsLocalIP(ip net.IP) bool {
	if ip == nil {
		return false
	}
	// 判断是否是回环地址,ipv4时是127.0.0.1; ipv6时是::1
	if ip.IsLoopback() {
		return true
	}

	// 判断ipv4是否是内网
	if ip4 := ip.To4(); ip4 != nil {
		return ip4[0] == 10 || // 10.0.0.0/8
			(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31) || // 172.16.0.0/12
			(ip4[0] == 192 && ip4[1] == 168) // 192.168.0.0/16
	}

	// 判断ipv6是否是内网
	if ip16 := ip.To16(); ip16 != nil {
		// 参考 https://tools.ietf.org/html/rfc4193#section-3
		// 参考 https://en.wikipedia.org/wiki/Private_network#Private_IPv6_addresses
		// 判断ipv6唯一本地地址
		return 0xfd == ip16[0]
	}
	// 不是ip直接返回false
	return false
}

// GetRequestIP 尽最大努力实现获取客户端 IP 的算法
// 解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作
func GetRequestIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}
