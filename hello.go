package hello

import (
	"fmt"

	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)

func Hello() string {
	str := quote.Hello()
	fmt.Printf("quote.Hello() returned %s\n", str)
	return str
}

func Proverb() string {
	str := quoteV3.Concurrency()
	fmt.Printf("quoteV3.Concurrency() returned %s\n", str)
	return str
}
