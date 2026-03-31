package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var jsonOutput bool

func main() {
	rand.Seed(time.Now().UnixNano())

	// Check for --json flag
	args := []string{}
	for _, arg := range os.Args[1:] {
		if arg == "--json" {
			jsonOutput = true
		} else {
			args = append(args, arg)
		}
	}

	if len(args) < 1 {
		printUsage()
		return
	}

	switch args[0] {
	case "draw", "d":
		n := 1
		if len(args) > 1 {
			fmt.Sscanf(args[1], "%d", &n)
		}
		drawCards(n)
	case "spread", "s":
		spread := "three"
		if len(args) > 1 {
			spread = args[1]
		}
		doSpread(spread)
	case "list", "l":
		listCards()
	case "card", "c":
		if len(args) < 2 {
			fmt.Println("Usage: tarot card <id>")
			return
		}
		var id int
		fmt.Sscanf(args[1], "%d", &id)
		showCard(id)
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println(`ASCII Tarot

Usage:
  tarot draw [n] [--json]   Draw n random cards (default: 1)
  tarot spread [type] [--json]  Do a spread (three, celtic)
  tarot list [--json]       List all cards
  tarot card <id> [--json]  Show specific card

Flags:
  --json    Output in JSON format for robot/agent consumption`)
}

type CardOutput struct {
	ID        int      `json:"id"`
	Numeral   string   `json:"numeral"`
	Name      string   `json:"name"`
	Reversed  bool     `json:"reversed"`
	Meanings  []string `json:"meanings"`
	Position  string   `json:"position,omitempty"`
}

func drawCards(n int) {
	drawn := pickRandom(n)
	var cards []CardOutput
	
	for i, idx := range drawn {
		reversed := rand.Intn(2) == 1
		c := Deck[idx]
		
		meanings := c.Desc
		if reversed {
			meanings = c.RDesc
		}
		
		if jsonOutput {
			cards = append(cards, CardOutput{
				ID:       c.ID,
				Numeral:  c.Numeral,
				Name:     c.Name,
				Reversed: reversed,
				Meanings: meanings,
			})
		} else {
			printCard(c, reversed)
			if i < len(drawn)-1 {
				fmt.Println()
			}
		}
	}
	
	if jsonOutput {
		out, _ := json.MarshalIndent(cards, "", "  ")
		fmt.Println(string(out))
	}
}

func doSpread(spread string) {
	positions := []string{"Past", "Present", "Future"}
	n := 3
	if spread == "celtic" {
		positions = []string{"Present", "Challenge", "Past", "Future", "Above", "Below", "Advice", "External", "Hopes", "Outcome"}
		n = 10
	}
	
	drawn := pickRandom(n)
	var cards []CardOutput
	
	for i, idx := range drawn {
		reversed := rand.Intn(2) == 1
		c := Deck[idx]
		
		meanings := c.Desc
		if reversed {
			meanings = c.RDesc
		}
		
		if jsonOutput {
			cards = append(cards, CardOutput{
				ID:       c.ID,
				Numeral:  c.Numeral,
				Name:     c.Name,
				Reversed: reversed,
				Meanings: meanings,
				Position: positions[i],
			})
		} else {
			fmt.Printf("=== %s ===\n", positions[i])
			printCard(c, reversed)
			fmt.Println()
		}
	}
	
	if jsonOutput {
		out, _ := json.MarshalIndent(cards, "", "  ")
		fmt.Println(string(out))
	}
}

func pickRandom(n int) []int {
	if n > len(Deck) {
		n = len(Deck)
	}
	perm := rand.Perm(len(Deck))
	return perm[:n]
}

func printCard(c Card, reversed bool) {
	orient := ""
	if reversed {
		orient = "(Reversed)"
	}
	
	fmt.Println("+---------------------+")
	fmt.Printf("| %-19s |\n", c.Numeral)
	
	if c.Art != "" {
		lines := strings.Split(c.Art, "\n")
		for _, line := range lines {
			if line != "" {
				fmt.Printf("|%-21s|\n", line)
			}
		}
	} else {
		fmt.Println("|                     |")
		fmt.Println("|                     |")
	}
	
	name := c.Name
	if len(name) > 19 {
		name = name[:19]
	}
	fmt.Printf("| %-19s |\n", name)
	fmt.Printf("| %-19s |\n", orient)
	fmt.Println("+---------------------+")

	meanings := c.Desc
	if reversed {
		meanings = c.RDesc
	}
	fmt.Printf("\nMeanings: %s\n", strings.Join(meanings, ", "))
}

func listCards() {
	if jsonOutput {
		type CardInfo struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}
		var cards []CardInfo
		for _, c := range Deck {
			cards = append(cards, CardInfo{ID: c.ID, Name: c.Name})
		}
		out, _ := json.MarshalIndent(cards, "", "  ")
		fmt.Println(string(out))
	} else {
		for _, c := range Deck {
			fmt.Printf("%2d: %s\n", c.ID, c.Name)
		}
	}
}

func showCard(id int) {
	if id < 0 || id >= len(Deck) {
		if jsonOutput {
			fmt.Println(`{"error": "Invalid card ID"}`)
		} else {
			fmt.Println("Invalid card ID")
		}
		return
	}
	
	c := Deck[id]
	if jsonOutput {
		out, _ := json.MarshalIndent(CardOutput{
			ID:       c.ID,
			Numeral:  c.Numeral,
			Name:     c.Name,
			Reversed: false,
			Meanings: c.Desc,
		}, "", "  ")
		fmt.Println(string(out))
	} else {
		printCard(c, false)
		fmt.Println("\nReversed meanings: " + strings.Join(c.RDesc, ", "))
	}
}
