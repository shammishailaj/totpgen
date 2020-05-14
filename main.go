package main

import (
	"fmt"
	"github.com/xlzd/gotp"
)

func main() {
	fmt.Println("Current OTP is", gotp.NewDefaultTOTP("4S62BZNFXXSZLCRO").Now())
}
