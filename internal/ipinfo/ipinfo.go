package ipinfo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type IpInfoResponse struct {
	Ip       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
	Bogon    bool   `json:"bogon"`
}

var httpClient = &http.Client{Timeout: 3 * time.Second}

func GetInfo(ip string) (IpInfoResponse, error) {

	var target IpInfoResponse

	url := fmt.Sprintf("http://ipinfo.io/%s?token=be6de701410dd8", ip)

	r, err := httpClient.Get(url)
	if err != nil {
		return target, err
	}
	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(&target)
	if err != nil {
		return target, err
	}

	return target, nil
}
