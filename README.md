## Overview
This is the start of a library for [Comtele](http://www.comtele.com.br/).
## License
Gocomtele is licensed under a BSD license.

## Installation
To install gocomtele, simply run `go get github.com/yuriadams/gocomtele`.

## SMS Example

	package main

	import (
		"github.com/yuriadams/gocomtele"
	)

	func main() {
    comtele := comtele.NewComteleClient(authToken)
		from := "sender"
		to := "receivers"
		message := "Hello Word!!!"
		comtele.SendSMS(from, to, message)
	}
