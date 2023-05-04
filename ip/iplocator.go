package ip

import (
	"fmt"
	str "github.com/AlexZ33/utils/string"
	date "github.com/AlexZ33/utils/time"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"github.com/sirupsen/logrus"
	"strings"
	"sync"
)

var (
	once     sync.Once
	searcher *xdb.Searcher
)

// InitIpLocator 使用时，需要加载xdb文件
// github.com/lionsoul2014/ip2region/data/ip2region.xdb
func InitIpLocator(dbPath string) {
	once.Do(func() {
		if str.IsBlank(dbPath) {
			dbPath = "ip2region.xdb"
		}
		start := date.NowTimestamp()
		data, err := xdb.LoadContentFromFile(dbPath)
		if err != nil {
			logrus.Errorf("failed to load content from `%s`: %s\n", dbPath, err)
			return
		}

		if searcher, err = xdb.NewWithBuffer(data); err != nil {
			logrus.Errorf("failed to create searcher with content: %s\n", err)
			return
		}

		fmt.Printf("Load ip2region.xdb success, eslapsed %d ms\n", date.NowTimestamp()-start)
	})
}

func Search(ip string) string {
	if searcher == nil || str.IsBlank(ip) {
		return ""
	}
	region, _ := searcher.SearchByStr(ip)
	return region
}

func IpLocation(ip string) string {
	region := Search(ip) // eg. 中国|0|湖北省|武汉市|电信
	if str.IsBlank(region) {
		return ""
	}

	ss := strings.Split(region, "|")
	if len(ss) != 5 {
		return ""
	}
	var (
		nation   = ss[0]
		province = ss[2]
	)

	if str.IsNotBlank(province) && province != "0" {
		return province
	}

	if str.IsNotBlank(nation) && nation != "0" {
		return nation
	}
	return ""
}
