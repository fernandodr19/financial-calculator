package investments

import (
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

type Result float64

func (r Result) String() string {
	// []string{"en_US", "pt_BR", "de", "ja", "hi"}
	lang := language.MustParse("pt_BR")
	cur, _ := currency.FromTag(lang)
	scale, _ := currency.Cash.Rounding(cur) // fractional digits
	dec := number.Decimal(float64(r), number.Scale(scale))
	p := message.NewPrinter(lang)
	return p.Sprintf("%v%v", currency.Symbol(cur), dec)
}
