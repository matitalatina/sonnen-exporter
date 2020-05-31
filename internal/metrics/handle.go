package metrics

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"mattianatali.it/sonnen-exporter/internal/sonnen"
)

var (
	batteryPowerW = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "sonnen_battery_power_w",
		Help: "Sonnen battery power [W]",
	})
	consumptionW = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "sonnen_consumption_w",
		Help: "Sonnen battery consumption [W]",
	})
	productionW = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "sonnen_production_w",
		Help: "Sonnen battery production [W]",
	})
	gridRepoW = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "sonnen_grid_feed_in_w",
		Help: "Sonnen grid feed in [W]",
	})
	chargePct = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "sonnen_charge_pct",
		Help: "Sonnen Charge [Percentage]",
	})
)

func HandleMetrics() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := sonnen.GetStats()

		if err != nil {
			fmt.Printf("Error encountered: %+v", err)
		}

		batteryPowerW.Set(float64(resp.PacTotalW))
		consumptionW.Set(float64(resp.ConsumptionW))
		productionW.Set(float64(resp.ProductionW))
		gridRepoW.Set(float64(resp.GridFeedInW))
		chargePct.Set(float64(resp.USOC))

		promhttp.Handler().ServeHTTP(w, r)
	}
}
