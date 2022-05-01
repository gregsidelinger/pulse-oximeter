package serial

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/tarm/serial"

	"github.com/prometheus/client_golang/prometheus"
)

var Config = &serial.Config{}

var (
	spo2 = prometheus.NewGauge(prometheus.GaugeOpts{
		//Namespace: "our_company",
		//Subsystem: "blob_storage",
		Name: "spo2",
		Help: "o2 Percentage",
	})

	bpm = prometheus.NewGauge(prometheus.GaugeOpts{
		//Namespace: "our_company",
		//Subsystem: "blob_storage",
		Name: "bpm",
		Help: "beats per minite",
	})

	pa = prometheus.NewGauge(prometheus.GaugeOpts{
		//Namespace: "our_company",
		//Subsystem: "blob_storage",
		Name: "pa",
		Help: "pa",
	})

	status = prometheus.NewGauge(prometheus.GaugeOpts{
		//Namespace: "our_company",
		//Subsystem: "blob_storage",
		Name: "status",
		Help: "status",
	})
)

func init() {
	prometheus.MustRegister(spo2)
	prometheus.MustRegister(bpm)
	prometheus.MustRegister(pa)
}

func findNamedMatches(regex *regexp.Regexp, str string) map[string]string {
	match := regex.FindStringSubmatch(str)

	results := map[string]string{}
	for i, name := range match {
		results[regex.SubexpNames()[i]] = name
	}
	return results
}

func Read() {
	Config.Parity = serial.ParityNone
	Config.StopBits = serial.Stop1
	Config.ReadTimeout = time.Second * 5
	for {
		s, err := serial.OpenPort(Config)
		if err != nil {
			fmt.Printf("%v\n", err)
			spo2.Set(0)
			bpm.Set(0)
			pa.Set(0)
			time.Sleep(time.Second * 5)
			continue
		}

		scanner := bufio.NewScanner(s)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
			re := regexp.MustCompile(`(?P<date>[a-zA-Z0-9\-]+)\s+(?P<time>[0-9:]+)\s+(?P<SpO2>\d+)(\*)?\s+(?P<BPM>\d+)(\*)?\s+(?P<PA>\d+)(?P<Status>.*)`)

			matches := findNamedMatches(re, scanner.Text())
			if matches != nil {
				if s, err := strconv.ParseFloat(matches["SpO2"], 64); err == nil {
					spo2.Set(s)
				}
				if s, err := strconv.ParseFloat(matches["BPM"], 64); err == nil {
					bpm.Set(s)
				}

				if s, err := strconv.ParseFloat(matches["PA"], 64); err == nil {
					pa.Set(s)
				}
			}

		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("%v\n", err)
			spo2.Set(0)
			bpm.Set(0)
			pa.Set(0)
		}
	}

}
