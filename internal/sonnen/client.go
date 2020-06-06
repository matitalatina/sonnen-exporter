package sonnen

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Stats struct {
	ApparentOutput            int         `json:"Apparent_output"`
	BackupBuffer              string      `json:"BackupBuffer"`
	BatteryCharging           bool        `json:"BatteryCharging"`
	BatteryDischarging        bool        `json:"BatteryDischarging"`
	ConsumptionW              int         `json:"Consumption_W"`
	Fac                       float64     `json:"Fac"`
	FlowConsumptionBattery    bool        `json:"FlowConsumptionBattery"`
	FlowConsumptionGrid       bool        `json:"FlowConsumptionGrid"`
	FlowConsumptionProduction bool        `json:"FlowConsumptionProduction"`
	FlowGridBattery           bool        `json:"FlowGridBattery"`
	FlowProductionBattery     bool        `json:"FlowProductionBattery"`
	FlowProductionGrid        bool        `json:"FlowProductionGrid"`
	GridFeedInW               int         `json:"GridFeedIn_W"`
	IsSystemInstalled         int         `json:"IsSystemInstalled"`
	OperatingMode             string      `json:"OperatingMode"`
	PacTotalW                 int         `json:"Pac_total_W"`
	ProductionW               int         `json:"Production_W"`
	RSOC                      int         `json:"RSOC"`
	Sac1                      int         `json:"Sac1"`
	Sac2                      interface{} `json:"Sac2"`
	Sac3                      interface{} `json:"Sac3"`
	SystemStatus              string      `json:"SystemStatus"`
	Timestamp                 string      `json:"Timestamp"`
	USOC                      int         `json:"USOC"`
	Uac                       int         `json:"Uac"`
	Ubat                      int         `json:"Ubat"`
}

func GetStats() (*Stats, error) {
	//url := "http://SB-62312.local:8080/api/v1/status"
	url := "http://192.168.1.125:8080/api/v1/status"
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var stats Stats
	err = json.Unmarshal(body, &stats)

	if err != nil {
		return nil, err
	}

	return &stats, nil
}
