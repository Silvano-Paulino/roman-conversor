package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/silvano-paulino/internal/converter"
)

func Cli() {
	IntToRoman := flag.NewFlagSet("int", flag.ExitOnError)
	intValue := IntToRoman.Int("value", 0, "Decimal value to convert to Roman numeral")

	RomanToInt := flag.NewFlagSet("roman", flag.ExitOnError)
	romanValue := RomanToInt.String("value", "", "Roman numeral to convert to decimal")

	if len(os.Args) < 3 {
		println("expected 'int' or 'roman' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "int":
		IntToRoman.Parse(os.Args[2:])

		service := converter.NewDecimalToRomanService()

		v, err := service.DecimalToRoman(*intValue)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}

		fmt.Println("Decimal to Roman:", *intValue, "->", v)
	case "roman":
		RomanToInt.Parse(os.Args[2:])

		service := converter.NewRomanToDecimalService()

		v, err := service.RomanToDecimal(*romanValue)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}

		fmt.Println("Roman to Decimal:", *romanValue, "->", v)
	default:
		println("expected 'int' or 'roman' subcommands")
		os.Exit(1)
	}
}
