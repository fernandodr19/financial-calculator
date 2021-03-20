package investments

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc               string
		Scenario           Scenario
		ExpectedLen        int
		ExpectedMonthYield float64
		ExpectedFinalValue Result
	}{
		{
			desc: "10%",
			Scenario: Scenario{
				AnnualYield:    10,
				InitialBalance: 0,
				MonthApply:     0,
				Years:          1,
			},
			ExpectedLen:        1,
			ExpectedMonthYield: 0.7974140428903764,
			ExpectedFinalValue: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// test
			results := tC.Scenario.Run()

			// assert
			assert.Equal(t, tC.ExpectedLen, len(results))
			assert.Equal(t, tC.ExpectedMonthYield, tC.Scenario.GetMonthYield()*100)

			if tC.ExpectedLen == 0 {
				return
			}

			finalBalance := results[len(results)-1]

			assert.Equal(t, tC.ExpectedFinalValue, finalBalance)
		})
	}
}
