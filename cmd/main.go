package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	annualYield := Profit{}
	log.Infoln(annualYield)
	log.Infoln(calculateMonthYield(annualYield))
}
