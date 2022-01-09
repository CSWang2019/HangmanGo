package main

import (
	"fmt"
	"unicode"
)

type Hangman struct {
	secret  []byte
	hidden  []byte
	history []byte
	life    int
}

func NewHangman() *Hangman {
	secret := []byte("hello world!")
	hidden := make([]byte, len(secret))
	history := make([]byte, 0)
	for i := 0; i < len(secret); i++ {
		if unicode.IsLetter(rune(secret[i])) {
			hidden[i] = '*'
		} else {
			hidden[i] = secret[i]
		}
	}
	return &Hangman{secret, hidden, history, 5}
}

func (hm *Hangman) GetGuess() byte {

	// prompt
	var input string
	fmt.Print("please guess a letter: ")
	fmt.Scanf("%s", &input)

	if len(input) > 1 {
		fmt.Println("please only guess one letter")
		return hm.GetGuess()
	}

	if !unicode.IsLetter(rune(input[0])) {
		fmt.Println("please guess a letter (a-z or A-Z)")
		return hm.GetGuess()
	}

	lower := unicode.ToLower(rune(input[0]))
	for _, c := range hm.history {
		if lower == rune(c) {
			fmt.Println("you've guessed this letter")
			return hm.GetGuess()
		}
	}

	hm.history = append(hm.history, byte(lower))

	return input[0]
}

func (hm *Hangman) ProcessGuess(guess byte) {
	hit := false
	for i := 0; i < len(hm.secret); i++ {
		if unicode.ToLower(rune(guess)) == unicode.ToLower(rune(hm.secret[i])) {
			hit = true
			hm.hidden[i] = hm.secret[i]
		}
	}

	if !hit {
		hm.life--
	}
}

func (hm *Hangman) IsWin() bool {
	return string(hm.secret) == string(hm.hidden)
}

func (hm *Hangman) Alive() bool {
	return hm.life > 0
}

func (hm *Hangman) Print() {
	fmt.Println(string(hm.secret))
	fmt.Println(string(hm.hidden))
	fmt.Println("guess history:", string(hm.history))
	fmt.Println("life:", hm.life)
}

func main() {
	fmt.Println("welcome to hangman")
	hm := NewHangman()

	for !hm.IsWin() && hm.Alive() {
		hm.Print()
		c := hm.GetGuess()
		hm.ProcessGuess(c)
	}

	if hm.Alive() {
		fmt.Println("you win!")
	} else {
		fmt.Println("you lose!")
	}
}
