package serial

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/tarm/serial"

	po "github.com/gregsidelinger/pulse-oximeter/pkg/pulse-oximeter"
)

var Config = &serial.Config{}

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
			po.Spo2.Set(0)
			po.Bpm.Set(0)
			po.Pa.Set(0)
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
					po.Spo2.Set(s)
				}
				if s, err := strconv.ParseFloat(matches["BPM"], 64); err == nil {
					po.Bpm.Set(s)
				}

				if s, err := strconv.ParseFloat(matches["PA"], 64); err == nil {
					po.Pa.Set(s)
				}
			}

		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("%v\n", err)
			po.Spo2.Set(0)
			po.Bpm.Set(0)
			po.Pa.Set(0)
		}
	}

}
