package monitoring

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Rackspace cloud monitoring (MaaS) client routines

const (
	URL = "https://monitoring.api.rackspacecloud.com/v1.0"

	AUTH_HEADER = "X-Auth-Token"
)

type GetMonitoringZonesResponse struct {
	Zones []Zone `json:"values"`
	// ignore metadata
}

type Zone struct {
	Id          string   `json:"id"`
	Label       string   `json:"label"`
	CountryCode string   `json:"country_code"`
	SourceIps   []string `json:"source_ips"`
}

func baseURL(tenantId string) string {
	return fmt.Sprintf("%s/%s", URL, tenantId)
}

// given an identity token, get all the public monitoring zones
func GetZones(tenantID string, tokenID string) ([]Zone, error) {

	url := fmt.Sprintf("%s/monitoring_zones", baseURL(tenantID))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set(AUTH_HEADER, tokenID)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var getZoneResponse GetMonitoringZonesResponse
	err = json.Unmarshal(buf, &getZoneResponse)
	if err != nil {
		return nil, err
	}

	return getZoneResponse.Zones, nil
}
