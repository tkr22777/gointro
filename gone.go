package gone

import (
	 "rsc.io/quote/v3"
)

func Gone() string {
	return "Hello, world!"
}

func Quote() string  {
	return quote.HelloV3()
}

func Proverb() string {
	return quote.Concurrency()
}


