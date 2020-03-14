package gone

import (
	"rsc.io/quote"
	qouteV3 "rsc.io/quote/v3"
)

func Gone() string {
	return "Hello, world!"
}

func Quote() string  {
	return quote.Hello()
}

func Proverb() string {
	return qouteV3.Concurrency()
}


