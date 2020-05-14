package main

import (
	"flag"
	"fmt"
	"github.com/shammishailaj/gotp"
	"log"
)

func main() {
	totpSecretkey := flag.String("k", "", "Secret key for TOTP")
	flag.Parse()

	if *totpSecretkey == "" {
		log.Printf("Must provide secret key to generate TOTP\n\n")
		flag.Usage()
		fmt.Println("")
	} else {
		fmt.Printf("\nCurrent OTP is %s\n", gotp.NewDefaultTOTP(*totpSecretkey).Now())
	}
}
