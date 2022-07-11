package simulate

import (
	"fmt"
	"math/rand"
	"time"

	po "github.com/gregsidelinger/pulse-oximeter/pkg/pulse-oximeter"
)

func fudgeValue(v int, min int, max int) int {
	v = v + (rand.Intn(10) - 5)
	if v > max {
		v = max
	}
	if v < min {
		v = min
	}
	return v

}

func Read(minSpo2 int, maxSpo2 int, minBpa int, maxBpa int, minPa int, maxPa int) {
	// Get our random starting values
	rand.Seed(time.Now().UnixNano())
	spo2 := rand.Intn(maxSpo2-minSpo2) + minSpo2 // range is min to max
	bpa := rand.Intn(maxBpa-minBpa) + minBpa     // range is min to max
	pa := rand.Intn(maxPa-minPa) + minPa         // range is min to max

	for {
		fmt.Printf("spo2 %d bpa %d pa %d\n", spo2, bpa, pa)
		po.Spo2.Set(float64(spo2))
		po.Bpm.Set(float64(bpa))
		po.Pa.Set(float64(pa))
		time.Sleep(time.Second * 5)

		spo2 = fudgeValue(spo2, minSpo2, maxSpo2)
		bpa = fudgeValue(bpa, minBpa, maxBpa)
		pa = fudgeValue(pa, minPa, maxPa)

	}

}
