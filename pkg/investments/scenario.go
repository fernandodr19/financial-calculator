package investments

import (
	"fmt"
	"math"
)

// Scenario is a simulation of a hypothetical investment scenario
type Scenario struct {
	AnnualYield    float64
	InitialBalance float64
	MonthApply     float64
	Years          int
}

func (s Scenario) Run() []Result {
	months := s.Years * 12

	records := make([]Result, 0)

	currentBalance := s.InitialBalance
	for i := 0; i < months; i++ {
		currentBalance = (currentBalance * (1 + s.GetMonthYield())) + s.MonthApply
		if (i+1)%12 == 0 {
			records = append(records, Result(currentBalance))
		}
	}

	return records
}

func (s Scenario) String() string {
	return fmt.Sprintf("%.2f%% a.a (%.2f%% a.m)", s.AnnualYield, s.GetMonthYield()*100)
}

func (s Scenario) GetMonthYield() float64 {
	return math.Pow(1+s.AnnualYield/100, 1.0/12) - 1
}
