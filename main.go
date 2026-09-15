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
	_, _ = os.Stdout.WriteString(value_ext)
	_, _ = os.Stdout.WriteString(strings.ReplaceAll(stringData, "o", "o World") + "\n")
	cryptoKey := makeCrypto(stringData)

Loop:
	for value <= 5 {
		switch value {
		case -1.0:
			_, _ = os.Stdout.WriteString("You went to the wrong place\n")
		case 0.0:
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

	_, _ = os.Stdout.WriteString("Gotcha!!!\nMONTNAHP PROTOCOL:\nsudo getRoot -- true\n  ROOTING.....\n  ROOT ACCESS BY 'SKYNET'\n")
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
