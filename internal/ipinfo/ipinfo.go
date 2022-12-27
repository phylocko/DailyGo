/*
Copyright © 2022 Vladislav Pavkin [phylocko@gmail.com]
*/
package ipinfo

import (
	"context"
	"dailygo/internal/log"
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

var httpClient = &http.Client{}

func GetInfo(ip string) (IpInfoResponse, error) {

	var target IpInfoResponse

	url := fmt.Sprintf("http://ipinfo.io/%s?token=be6de701410dd8", ip)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.L.Sugar().Warn(err)
		return target, err
	}

	response, err := httpClient.Do(req)

	if err != nil {
		log.L.Sugar().Warn(err)
		return target, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&target)
	if err != nil {
		return target, err
	}

	return target, nil
}
