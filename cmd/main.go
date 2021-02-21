package main

import (
	"main/pkg/investments"

	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.Infof("Build info: time[%s] git_hash[%s]", BuildTime, BuildGitCommit)

	s1 := investments.Scenario{
		AnnualYield:    6,
		InitialBalance: 10000.00,
		// MonthApply:     1000.00,
		Years: 20,
	}
	// log.Infoln(s1.String())
	// log.Infoln(s1.GetMonthYield())

	records := s1.Run()
	for i, r := range records {
		log.WithFields(logrus.Fields{
			"balance":               r,
			"month income (.6 a.a)": (r * .06) / 12,
		}).Infoln("Investment simulation for month", i+1)
	}

	// TODO generate chart https://github.com/gonum/plot
}
