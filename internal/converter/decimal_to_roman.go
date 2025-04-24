package converter

import "errors"

type decimalToRomanService struct{}

func NewDecimalToRomanService() *decimalToRomanService {
	return &decimalToRomanService{}
}

var ErrValueMustBeGreaterThanZero = errors.New("value must be greater than zero")

func (d decimalToRomanService) DecimalToRoman(value int) (string, error) {
	if value == 0 {
		return "", ErrValueMustBeGreaterThanZero
	}

	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syms := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	roman := ""

	for i := range vals {
		for value >= vals[i] {
			value -= vals[i]
			roman += syms[i]
		}
	}

	return roman, nil
}
