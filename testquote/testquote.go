package testquote

import "rsc.io/quote"


func Random() string {
	message := quote.Go()
	return message
}