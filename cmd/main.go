package main

import (
	"main/pkg/investments"

	"github.com/sirupsen/logrus"
)

// Injected on build time by Makefile
var (
	BuildGitCommit = "undefined"
	BuildTime      = "undefined"
)

func main() {
	log := logrus.New()
	log.Infof("Build info: time[%s] git_hash[%s]", BuildTime, BuildGitCommit)

	s1 := investments.Scenario{
		AnnualYield:    7,
		InitialBalance: 20000.00,
		MonthApply:     1000.00,
		Years:          20,
	}
	log.Infoln(s1.String())
	// log.Infoln(s1.GetMonthYield())

	records := s1.Run()
	for i, r := range records {
		log.WithFields(logrus.Fields{
			"balance":               r,
			"month income (.6 a.a)": (r * .06) / 12,
		}).Infoln("Investment simulation for year", i+1)
	}

	// TODO generate chart https://github.com/gonum/plot
}
