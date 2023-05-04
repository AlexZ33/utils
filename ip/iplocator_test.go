package ip

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	InitIpLocator("github.com/lionsoul2014/ip2region/data/ip2region.xdb")
	ip := "47.52.26.78"
	fmt.Println(Search(ip))
	fmt.Println(IpLocation(ip))
}
