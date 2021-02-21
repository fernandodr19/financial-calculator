package main

import (
	"math"

	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	annualYield := 10.0
	log.Infoln(annualYield)
	log.Infoln(calculateMonthYield(annualYield))
}

func calculateMonthYield(annualYield float64) float64 {
	return math.Pow(annualYield, 1.0/12)
}
