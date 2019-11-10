package hello

import (
	"fmt"

	"rsc.io/quote"
)

func Hello() string {
	str := quote.Hello()
	fmt.Printf("quote.Hello() returned %s\n", str)
	return str
}
