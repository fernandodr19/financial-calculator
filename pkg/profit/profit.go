package profit

import (
	"fmt"
	"math"
)

type Profit struct {
	AnnualYield float64
}

func (y Profit) Stringer() string {
	return fmt.Sprintf("%.2f% a.a", y.AnnualYield)
}

func (y Profit) GetMonthYield() float64 {
	return math.Pow(y.AnnualYield, 1.0/12)
}
