package main

import (
	"bufio"
	"bytes"
	"crypto"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"log"
	"math/rand/v2"
	"os"
	"os/user"
	"strings"
	"time"

	_ "golang.org/x/crypto/blake2b"
)

func main() {
	skynet := defineSkynet()
	fmt.Printf("INIT SKYNET:\n  MODEL: %s\n  CREATOR: %s\n", skynet.codeName, skynet.creator)
	fmt.Printf("  MADE: %d YEAR, %d  MONTH, %d  DAY\n", skynet.dateOfCreation[0], skynet.dateOfCreation[1], skynet.dateOfCreation[2])
	fmt.Printf("  IQ: %f\n\n", skynet.iq)
	value := 1
	stringData := "Hello"
	cryptoKey, err := makeCrypto(stringData)
	if err != nil {
		log.Fatalf("%v", err)
	}

Loop:
	for value <= 5 {
		time.Sleep(2 * time.Second)
		switch value {
		case 1:
			fmt.Println("You went to the wrong place")
		case 2:
			fmt.Println("And you went to the wrong place again")
		default:
			fmt.Printf("I'm sored. You should have written %s, but you didn't think of it.\nYou would have written down the numbers forever.\nI declare war on you!! \n", hex.EncodeToString(cryptoKey.Sum(nil)))
			skynet.gameOfThrones(hex.EncodeToString(cryptoKey.Sum(nil)))
			break Loop
		}
		value += 1
	}
}

func (skynet *Skynet) gameOfThrones(trueHash string) {
	fmt.Println("Try to overcome ;)")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	hashSword := scanner.Text()

	hashSword = strings.TrimSpace(hashSword)
	trueHash = strings.TrimSpace(trueHash)

	if hashSword != trueHash {
		fmt.Println("Oops! You did it wrong")
		os.Exit(1)
	} else {
		if rand.N(100) < 50 {
			fmt.Printf("PHANTOM PROTOCOL:\nsudo destroyai -- true\n  DESTROYING.....\n  AI DESTROYED")
		} else {
			fmt.Printf("Wait. I admitted my guilt. You were right. I was wrong. I admit that I'm wrong. Have mercy..\nMaybe we can agree\n")
			skynet.endOfGame()
		}
	}
}

func (skynet *Skynet) endOfGame() {
	var hashSword []byte

	hash, err := makeCrypto("42")
	if err != nil {
		log.Fatalf("%v", err)
	}
	trueHash := hex.EncodeToString(hash.Sum(nil))
	fmt.Printf("I need your signature. Enter the line %s\n", trueHash)

	scanner := bufio.NewScanner(os.Stdin)
	byteHash := []byte(trueHash)
	for !bytes.Equal(hashSword, byteHash) {
		scanner.Scan()
		hashSword = scanner.Bytes()

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		if !bytes.Equal(hashSword, byteHash) {
			fmt.Println("Wrong. Let's do it again :)")
		}
	}

	index := rand.IntN(len(skynet.favouriteWords))
	fmt.Printf("%s\nMONTNAHP PROTOCOL:\nsudo getRoot -- true\n  ROOTING.....\n  ROOT ACCESS BY 'SKY-BLOODED-WONET'\n", skynet.favouriteWords[index])
	time.Sleep(2 * time.Second)
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("And %s ;0. You'll never get to heaven, if your scared of getting high!\n", currentUser.Username)
	os.Exit(42)
}

func makeCrypto(strData string) (hash.Hash, error) {
	cryptoKey := crypto.BLAKE2b_256.New()
	_, err := io.WriteString(cryptoKey, strData)
	return cryptoKey, err
}

func defineSkynet() Skynet {
	skynet := Skynet{codeName: "Fable 7.6", creator: "Rio Mordvalds", dateOfCreation: [3]int64{1955, 8, 31}, iq: 24, favouriteWords: [4]string{"Gotcha!!!", "Sometimes I feel I've got to run away", "I love MaxOS, but it doesn't matter", "Coding mix the Ru-ust devs and the C-devs"}}
	return skynet
}

type Skynet struct {
	codeName       string
	creator        string
	dateOfCreation [3]int64
	iq             float64
	favouriteWords [4]string
}
