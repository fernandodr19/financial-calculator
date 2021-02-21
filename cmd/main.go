package main

import (
	"fmt"
	"main/pkg/investments"

	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	s1 := investments.Scenario{
		AnnualYield:    10,
		InitialBalance: 10000.00,
		// MonthApply:     1000.00,
		Years: 1,
	}
	log.Infoln(s1.String())
	log.Infoln(s1.GetMonthYield())

	records := s1.Run()
	for i, r := range records {
		log.WithFields(logrus.Fields{"Month": i + 1, "Balance": fmt.Sprintf("%.2f", r)}).Infoln("aa")
	}
}
