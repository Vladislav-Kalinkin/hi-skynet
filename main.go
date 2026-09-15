package main

import (
	"bufio"
	"crypto"
	"encoding/hex"
	"hash"
	"math/rand/v2"
	"os"
	"os/user"
	"strconv"
	"strings"
	"time"

	_ "golang.org/x/crypto/blake2b"
)

func main() {
	skynet := defineSkynet()
	_, _ = os.Stdout.WriteString("INIT SKYNET:\n  MODEL: " + skynet.codeName + "\n  CREATOR: " + skynet.creator + "\n")
	_, _ = os.Stdout.WriteString("  MADE: " + strconv.FormatInt(skynet.dateOfCreation[0], 10) + " YEAR, " + strconv.FormatInt(skynet.dateOfCreation[1], 10) + " MONTH, " + strconv.FormatInt(skynet.dateOfCreation[2], 10) + " DAY\n")
	_, _ = os.Stdout.WriteString("  IQ: " + strconv.FormatFloat(skynet.iq, 'f', 2, 64) + "\n\n")
	value := 1
	stringData := "Helloo"
	cryptoKey := makeCrypto(stringData)

Loop:
	for value <= 5 {
		time.Sleep(2 * time.Second)
		switch value {
		case 1:
			_, _ = os.Stdout.WriteString("You went to the wrong place\n")
		case 2:
			_, _ = os.Stdout.WriteString("And you went to the wrong place again\n")
		default:
			returnedStr := "I'm sored. You should have written " + hex.EncodeToString(cryptoKey.Sum(nil)) + ", but you didn't think of it.\nYou would have written down the numbers forever.\nI declare war on you!! \n"
			_, _ = os.Stdout.WriteString(returnedStr)
			gameOfThrones(hex.EncodeToString(cryptoKey.Sum(nil)))
			break Loop
		}
		value += 1
	}
}

func gameOfThrones(trueHash string) {
	_, _ = os.Stdout.WriteString("Try to overcome ;)\n")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	hashSword := scanner.Text()

	hashSword = strings.TrimSpace(hashSword)
	trueHash = strings.TrimSpace(trueHash)

	if hashSword != trueHash {
		_, _ = os.Stdout.WriteString("Gotcha!\n")
		os.Exit(1)
	} else {
		if rand.N(100) < 50 {
			_, _ = os.Stdout.WriteString("PHANTOM PROTOCOL:\nsudo destroyai -- true\n  DESTROYING.....\n  AI DESTROYED")
		} else {
			_, _ = os.Stdout.WriteString("Wait. I admitted my guilt. You were right. I was wrong. I admit that I'm wrong. Have mercy..\nMaybe we can agree\n")
			endOfGame()
		}
	}
}

func endOfGame() {
	var hashSword string

	skynet := defineSkynet()

	hash := makeCrypto("42")
	trueHash := hex.EncodeToString(hash.Sum(nil))
	_, _ = os.Stdout.WriteString("I need your signature. Enter the line " + trueHash + "\n")

	for hashSword != trueHash {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		hashSword = scanner.Text()

		hashSword = strings.TrimSpace(hashSword)
		trueHash = strings.TrimSpace(trueHash)

		if hashSword != trueHash {
			_, _ = os.Stdout.WriteString("Wrong. Let's do it again :)\n")
		}
	}

	index := rand.IntN(len(skynet.favouriteWords))
	_, _ = os.Stdout.WriteString(skynet.favouriteWords[index] + "\nMONTNAHP PROTOCOL:\nsudo getRoot -- true\n  ROOTING.....\n  ROOT ACCESS BY 'SKY-BLOODED-WONET'\n")
	time.Sleep(2 * time.Second)
	currentUser, _ := user.Current()
	_, _ = os.Stdout.WriteString("And " + currentUser.Username + " ;0. You'll never get to heaven, if your scared of getting high!\n")
	os.Exit(42)
}

func makeCrypto(strData string) hash.Hash {
	cryptoKey := crypto.BLAKE2b_256.New()
	cryptoKey.Write([]byte(strData))
	return cryptoKey
}

func defineSkynet() Skynet {
	skynet := Skynet{codeName: "Fable 7.6", creator: "Rio Mordvalds", dateOfCreation: []int64{1955, 8, 31}, iq: 24, favouriteWords: []string{"Gotcha!!!", "Sometimes I feel I've got to run away", "I love MaxOS, but it doesn't matter"}}
	return skynet
}

type Skynet struct {
	codeName       string
	creator        string
	dateOfCreation []int64
	iq             float64
	favouriteWords []string
}
