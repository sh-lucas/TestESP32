package main

import (
	"fmt"
	"log"

	"github.com/SherClockHolmes/webpush-go"
)

func main() {
	privateKey, publicKey, err := webpush.GenerateVAPIDKeys()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("VAPID Public Key: %s\n", publicKey)
	fmt.Printf("VAPID Private Key: %s\n", privateKey)
}
