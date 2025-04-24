package converter_test

import (
	"testing"

	"github.com/silvano-paulino/internal/converter"
)

func TestRomanToDecimal(t *testing.T) {
	t.Run("should convert roman to decimal numeral", func(t *testing.T) {
		value := "MMXXV"
		expected := 2025
		service := converter.NewRomanToDecimalService()

		result, _ := service.RomanToDecimal(value)

		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("should return error if value is missing", func(t *testing.T) {
		value := ""
		service := converter.NewRomanToDecimalService()

		_, err := service.RomanToDecimal(value)

		if err != converter.ErrValueEmpty {
			t.Errorf("Expected %v, got %v", converter.ErrValueEmpty, err)
		}
	})
}
