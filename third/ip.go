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

func OutIp() (*Ip, error) {
	ip := Ip{}
	data, err := common.Get("https://ifconfig.co/json")

	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &ip)

	if err != nil {
		return nil, err
	}
	return &ip, nil
}
