package hmackeyfinder

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"log"
	"os"
	"time"
)

var done chan bool
var guard chan bool

var messageBytes []byte
var hashcodeBytes []byte

var rangeOfUint8 = [256]uint8{}

func init() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(file)
	}
}

func init() {
	for i := range rangeOfUint8 {
		rangeOfUint8[i] = uint8(i)
	}
	if !Config.SearchInOrder {
		rangeOfUint8 = Shuffle(rangeOfUint8)
	}
	done = make(chan bool)
	guard = make(chan bool, Config.MaxGoroutines)
	messageBytes = []byte(Config.Message)
	hashcodeBytes, _ = hex.DecodeString(Config.Hashcode)
}

func main() {
	log.Println("Starting with", Config.MaxGoroutines, "max goroutines")
	log.Println("messageBytes:", messageBytes)
	log.Println("hashcodeBytes:", hashcodeBytes)

	start := time.Now().Unix()
	go tryAllPossibilities(messageBytes, hashcodeBytes)
	<-done
	log.Println("Total time:", (time.Now().Unix() - start), "seconds")
}

func tryAllPossibilities(messageBytes, hashcodeBytes []byte) {
	var x uint8
	for _, x = range rangeOfUint8 {
		go func(x uint8) {
			tryBlock(messageBytes, hashcodeBytes, x)
			<-guard
		}(x)
		guard <- true
	}
}

func tryBlock(messageBytes, hashcodeBytes []byte, block uint8) {
	log.Println("Started block:", block)
	defer log.Println("Finished block:", block)

	// Probably not the cleanest way to do it...
	var y, z, w uint8
	for _, y = range rangeOfUint8 {
		for _, z = range rangeOfUint8 {
			for _, w = range rangeOfUint8 {
				key := []byte{block, y, z, w}
				isValid := CheckMAC(
					messageBytes,
					hashcodeBytes,
					key)
				if isValid {
					log.Println("Valid key: ", key, hex.EncodeToString(key))
					done <- true
					return
				}
			}
		}
	}
}

/*CheckMAC checks that the messageMAC is correct*/
func CheckMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha1.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}
