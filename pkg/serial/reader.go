package serial

import (
	"bufio"
	"fmt"
	"github.com/tarm/serial"
	"log"
	"regexp"
	"strconv"

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
	s, err := serial.OpenPort(Config)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(s)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		re := regexp.MustCompile(`(?P<date>[a-zA-Z0-9\-]+)\s+(?P<time>[0-9:]+)\s+(?P<SpO2>\d+)\s+(?P<BPM>\d+)\s+(?P<PA>\d+)(?P<Status>.*)`)
		//match := re.FindStringSubmatch(scanner.Text())

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
		log.Fatal(err)
	}

}
