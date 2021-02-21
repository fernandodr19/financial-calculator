package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	anualYield := 10
	log.Infoln(anualYield)
}
