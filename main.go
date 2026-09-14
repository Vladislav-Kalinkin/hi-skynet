package main

import (
	"bufio"
	"crypto"
	"encoding/hex"
	"hash"
	"math"
	"math/rand/v2"
	"os"
	"os/user"
	"strconv"
	"strings"
	"time"

	_ "golang.org/x/crypto/blake2b"
)

func main() {
	value := -1.0
	stringData := "Helloo"
	value_ext := strconv.FormatFloat(math.Sqrt(-value), 'f', 0, 64) + "\n"
	os.Stdout.WriteString(value_ext)
	os.Stdout.WriteString(strings.Replace(stringData, "o", "o World", -1) + "\n")
	cryptoKey := makeCrypto(stringData)

Loop:
	for value <= 5 {
		switch value {
		case -1.0:
			os.Stdout.WriteString("You went to the wrong place\n")
		case 0.0:
			os.Stdout.WriteString("And you went to the wrong place again\n")
		default:
			returnedStr := "I'm sored. You should have written " + hex.EncodeToString(cryptoKey.Sum(nil)) + ", but you didn't think of it.\nYou would have written down the numbers forever.\nI declare war on you!! \n"
			os.Stdout.WriteString(returnedStr)
			gameOfThrones(hex.EncodeToString(cryptoKey.Sum(nil)))
			break Loop
		}
		value += 1
	}
}

func gameOfThrones(trueHash string) {
	os.Stdout.WriteString("Try to overcome ;)\n")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	hashSword := scanner.Text()

	hashSword = strings.TrimSpace(hashSword)
	trueHash = strings.TrimSpace(trueHash)

	if hashSword != trueHash {
		os.Stdout.WriteString("Gotcha!\n")
		os.Exit(1)
	} else {
		if rand.N(100) < 50 {
			os.Stdout.WriteString("PHANTOM PROTOCOL:\nsudo destroyai -- true\n  DESTROYING.....\n  AI DESTROYED")
		} else {
			os.Stdout.WriteString("Wait. I admitted my guilt. You were right. I was wrong. I admit that I'm wrong. Have mercy..\nMaybe we can agree\n")
			capitulation()
		}
	}
}

func endOfGame(trueHash string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	hashSword := scanner.Text()

	hashSword = strings.TrimSpace(hashSword)
	trueHash = strings.TrimSpace(trueHash)

	for hashSword != trueHash {
		os.Stdout.WriteString("Wrong. Let's do it again :)\n")
	}

	os.Stdout.WriteString("Gotcha!!!\nMONTNAHP PROTOCOL:\nsudo getRoot -- true\n  ROOTING.....\n  ROOT ACCESS BY 'SKYNET'\n")
	time.Sleep(2 * time.Second)
	currentUser, _ := user.Current()
	os.Stdout.WriteString("And " + currentUser.Username + " ;0. You'll never get to heaven, if your scared of getting high!\n")
	os.Exit(42)
}

func makeCrypto(strData string) hash.Hash {
	cryptoKey := crypto.BLAKE2b_256.New()
	cryptoKey.Write([]byte(strData))
	return cryptoKey
}

func capitulation() {
	cryptoKey := makeCrypto("42")
	os.Stdout.WriteString("I need your signature. Enter the line " + hex.EncodeToString(cryptoKey.Sum(nil)) + "\n")
	endOfGame(hex.EncodeToString(cryptoKey.Sum(nil)))
}
