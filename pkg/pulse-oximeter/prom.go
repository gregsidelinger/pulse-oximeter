package pulseoximeter

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	Spo2 = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "spo2",
		Help: "o2 Percentage",
	})

	Bpm = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "bpm",
		Help: "beats per minite",
	})

	Pa = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pa",
		Help: "pa",
	})

	Status = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "status",
		Help: "status",
	})
)

func init() {
	prometheus.MustRegister(Spo2)
	prometheus.MustRegister(Bpm)
	prometheus.MustRegister(Pa)
}
