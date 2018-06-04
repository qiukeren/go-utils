package third

import (
	"encoding/json"

	common "github.com/qiukeren/go-utils/common"
)

type Ip struct {
	City      string `json:"city"`
	Country   string `json:"country"`
	IP        string `json:"ip"`
	IPDecimal int    `json:"ip_decimal"`
}

func OutIp() (string, error) {
	ip := Ip{}
	data, err := common.Get("https://ifconfig.co/json")

	if err == nil {
		err = json.Unmarshal(data, &ip)
		if err == nil {
			return ip.IP, nil
		}
	}

	data, err = common.Get("http://whatismyip.akamai.com")
	if err == nil {
		return string(data), nil
	}

	data, err = common.Get("http://ifconfig.io")
	return string(data), err
}
