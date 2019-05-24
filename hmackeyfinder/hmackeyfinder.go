package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"flag"
	"log"
	"math/rand"
	"runtime"
	"time"
)

var done = make(chan struct{}, 1)
var blocks = make(chan uint8, 256)

var messageBytes []byte
var hashcodeBytes []byte

func main() {
	parseFlagsAndInitGlobals()

	start := time.Now().Unix()
	bruteforce(messageBytes, hashcodeBytes)
	<-done

	log.Println("Total time:", time.Now().Unix()-start, "seconds")
}

func parseFlagsAndInitGlobals() {
	var messageInput, hashcodeInput string
	var randomOrder bool
	var err error

	flag.BoolVar(&randomOrder, "randomOrder", false, "Set to true if you want to brute force search 'in order'.")
	flag.StringVar(&messageInput, "message", "341567891 487654 500", "The message fits the provided hashcode with a particular key. Default: 341567891 487654 500")
	flag.StringVar(&hashcodeInput, "hashcode", "f3c2ae334dc98a387601c85ef83c77360943023a", "The hashcode of the message calculated with an unknown key. Default: f3c2ae334dc98a387601c85ef83c77360943023a")
	flag.Parse()

	fillBlocks(randomOrder)
	messageBytes = []byte(messageInput)
	hashcodeBytes, err = hex.DecodeString(hashcodeInput)
	if err != nil {
		panic(err)
	}

	log.Println("Value of messageBytes:", messageBytes)
	log.Println("Value of hashcodeBytes:", hashcodeBytes)
}

func fillBlocks(randomOrder bool) {
	if randomOrder {
		log.Println("Filling blocks channel in a random order")
		for _, randI := range shuffled() {
			blocks <- randI
		}
	} else {
		log.Println("Filling blocks channel with [0,255]")
		for i := 0; i < 256; i++ {
			blocks <- uint8(i)
		}
	}
}

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

func shuffled() [256]uint8 {
	var slice [256]uint8
	for i := 0; i < 256; i++ {
		slice[i] = uint8(i)
	}
	seed := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(seed)
	for i := 256; i > 0; i-- {
		randi := rng.Intn(i)
		slice[i-1], slice[randi] = slice[randi], slice[i-1]
	}
	return slice
}
