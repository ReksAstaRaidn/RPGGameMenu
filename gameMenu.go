package main

import (
	"fmt"
	"math/rand"
	"time"
)

const defaultJitter = 10 * time.Millisecond

func init() {
	rand.Seed(time.Now().UnixNano())
}

func typewriter(text string, delay time.Duration, jitter time.Duration) {
	for _, r := range text {
		fmt.Printf("%c", r)
		j := time.Duration(rand.Int63n(int64(jitter)*2)) - jitter
		d := delay + j
		if d < 0 {
			d = 0
		}
		time.Sleep(d)
	}
	fmt.Println()
}

func menu1() {
	fmt.Println("╔═══════════════════════════════╗")
	fmt.Println("║         MONSTER HUNTER        ║")
	fmt.Println("╠═══════════════════════════════╣")
	fmt.Println("║         1. Start Game         ║")
	fmt.Println("║         2. Load Game          ║")
	fmt.Println("║         3. Exit Game (˶°ㅁ°)  ║")
	fmt.Println("╚═══════════════════════════════╝")
}

func SC() {
	fmt.Println("1. Sora")
	fmt.Println("2. Riku")
	fmt.Println("3. Kairi")
	fmt.Println("0. Back to Main Menu")
	fmt.Println("Select a character to start the game.")

	var pilihChar int
	fmt.Scan(&pilihChar)

	switch pilihChar {
	case 1:
		soraStat()
	case 2:
		rikuStat()
	case 3:
		kairiStat()
	case 0:
		return
	default:
		fmt.Println("Invalid choice!")
	}
}

func soraStat() {
	typewriter("You selected Sora!", 40*time.Millisecond, defaultJitter)
	typewriter("Starting game with Sora...", 200*time.Millisecond, defaultJitter)
	typewriter("Game Started!", 20*time.Millisecond, defaultJitter)
	typewriter("Sora Stats: Level 1, HP 100, MP 50", 30*time.Millisecond, defaultJitter)
	menu1()
	fmt.Println("Choose an option:")
	// return to previous menu
	return
}

func rikuStat() {
	typewriter("You selected Riku!", 40*time.Millisecond, defaultJitter)
	typewriter("Starting game with Riku...", 50*time.Millisecond, defaultJitter)
	typewriter("Game Started!", 20*time.Millisecond, defaultJitter)
	typewriter("Riku Stats: Level 1, HP 120, MP 40", 30*time.Millisecond, defaultJitter)
	menu1()
	fmt.Println("Choose an option:")
	// return to previous menu
	return
}

func kairiStat() {
	typewriter("You selected Kairi!", 20*time.Millisecond, defaultJitter)
	typewriter("Starting game with Kairi...", 60*time.Millisecond, defaultJitter)
	typewriter("Game Started!", 20*time.Millisecond, defaultJitter)
	typewriter("Kairi Stats: Level 1, HP 90, MP 60", 30*time.Millisecond, defaultJitter)
	menu1()
	fmt.Println("Choose an option:")
	// return to previous menu
	return
}

func saved1() {

	fmt.Println("1. Save File 1")
	fmt.Println("2. Save File 2")
	fmt.Println("3. Save File 3")
	fmt.Println("0. Back to Main Menu")
	fmt.Println("Select Saved File to load.")

	var pilihSave int
	fmt.Scan(&pilihSave)
	switch pilihSave {
	case 1:
		SF1()
	case 2:
		SF2()
	case 3:
		SF3()
	case 0:
		return
	default:
		fmt.Println("Invalid choice!")
	}
}

func SF1() {
	fmt.Println("Loading Save File 1...")
	fmt.Println("Save File 1 Loaded!")
}

func SF2() {
	fmt.Println("Loading Save File 2...")
	fmt.Println("Save File 2 Loaded!")
}

func SF3() {
	fmt.Println("Loading Save File 3...")
	fmt.Println("Save File 3 Loaded!")
}

func main() {
	menu1()
	q := 0
	fmt.Println("Choose an option:")
	for q != 3 {
		var choice int
		fmt.Scan(&choice)
		fmt.Print("\n")
		switch choice {
		case 1:
			SC()
		case 2:
			saved1()
		case 3:
			fmt.Println("Exiting the game. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice!")
		}
	}
}
