package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"log"
	"runtime"
)

func bruteforce(messageBytes, hashcodeBytes []byte) {
	for i := 0; i < runtime.NumCPU(); i++ {
		log.Println("Creating worker nº", i)
		go func() {
			for {
				block := <-blocks
				tryBlock(messageBytes, hashcodeBytes, block)
			}
		}()
	}
}

func tryBlock(messageBytes, hashcodeBytes []byte, block uint8) {
	log.Println("Brute-forcing block:", block)
	defer log.Println("The key is not in block:", block)

	for x := 0; x <= 255; x++ {
		for y := 0; y <= 255; y++ {
			for z := 0; z <= 255; z++ {
				key := []byte{block, uint8(x), uint8(y), uint8(z)}
				isValid := checkMAC(messageBytes, hashcodeBytes, key)
				if isValid {
					log.Println("Found valid key: ", key, hex.EncodeToString(key))
					done <- struct{}{}
					return
				}
			}
		}
	}
}

func checkMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha1.New, key)
	mac.Write(message)
	expected := mac.Sum(nil)
	return hmac.Equal(messageMAC, expected)
}
