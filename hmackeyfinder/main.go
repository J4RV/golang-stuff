package main

import (
	"encoding/hex"
	"flag"
	"log"
	"math/rand"
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
