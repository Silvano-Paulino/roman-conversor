package converter_test

import (
	"testing"

	"github.com/silvano-paulino/internal/converter"
)

func TestIntToRoman(t *testing.T) {
	t.Run("should convert int to roman", func(t *testing.T) {
		value := 2025
		expected := "MMXXV"
		service := converter.NewDecimalToRomanService()

		result, _ := service.DecimalToRoman(value)

		if result != expected {
			t.Errorf("Expected %v, got %v", expected, value)
		}
	})

	t.Run("should return an error if value is zero", func(t *testing.T) {
		value := 0
		service := converter.NewDecimalToRomanService()

		_, err := service.DecimalToRoman(value)

		if err != converter.ErrValueMustBeGreaterThanZero {
			t.Errorf("Expected %v, got %v", converter.ErrValueMustBeGreaterThanZero, err)
		}
	})
}
