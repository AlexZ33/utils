package ip

import (
	"github.com/AlexZ33/utils/errors"
	"net"
)

// GetLocalIP 获取本机IP地址
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
