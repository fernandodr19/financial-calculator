package investments

import (
	"fmt"
	"math"
)

type Scenario struct {
	AnnualYield    float64
	InitialBalance float64
	MonthApply     float64
	Years          int
}

func (s Scenario) Run() []Result {
	months := s.Years * 12

	records := make([]Result, months)

	currentBalance := s.InitialBalance
	for i := 0; i < months; i++ {
		currentBalance = (currentBalance * (1 + s.GetMonthYield()/100)) + s.MonthApply
		records[i] = Result(currentBalance)
	}

	return records
}

func (s Scenario) String() string {
	return fmt.Sprintf("%.2f%% a.a (%.2f%% a.m)", s.AnnualYield, s.GetMonthYield())
}

func (s Scenario) GetMonthYield() float64 {
	return math.Pow(s.AnnualYield, 1.0/12)
}
