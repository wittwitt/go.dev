package door3

import (
	"math/rand"
	"testing"
	"time"
)

// simulateCoinToss simulates a single coin toss and returns true for heads, false for tails.
func simulateCoinToss() bool {
	// Generate a random number (0 or 1) to simulate the coin toss
	return rand.Intn(2) == 0
}

func TestCoinTossProbability(t *testing.T) {
	const numTosses = 1000000 // Number of coin tosses to perform

	rand.Seed(time.Now().UnixNano())

	var headsCount, tailsCount int

	for i := 0; i < numTosses; i++ {
		if simulateCoinToss() {
			headsCount++
		} else {
			tailsCount++
		}
	}

	// Calculate the probability of heads and tails
	headsProbability := float64(headsCount) / float64(numTosses)
	tailsProbability := float64(tailsCount) / float64(numTosses)

	// You can check the results by printing the probabilities:
	t.Logf("Probability of heads: %.2f%%", headsProbability*100)
	t.Logf("Probability of tails: %.2f%%", tailsProbability*100)
}

// simulateMontyHall simulates one round of the Monty Hall problem.
// Returns true if switching doors leads to a win, false otherwise.
func simulateMontyHall() bool {

	doorsToChooseFrom := []int{0, 1, 2}
	prizeDoor := rand.Intn(3)
	initialChoice := rand.Intn(3)

	doorsToChooseFrom = removeElement(doorsToChooseFrom, initialChoice)
	other := prizeDoor
	if initialChoice == prizeDoor {
		other = doorsToChooseFrom[rand.Intn(len(doorsToChooseFrom))]
	}

	return other == prizeDoor
}

// removeElement removes the specified element from a slice and returns the new slice.
func removeElement(slice []int, elem int) []int {
	for i := range slice {
		if slice[i] == elem {
			slice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	return slice
}

func TestMontyHall(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	const numSimulations = 100000 // Number of Monty Hall simulations

	var winsSwitch int
	var winsStick int

	for i := 0; i < numSimulations; i++ {
		if simulateMontyHall() {
			winsSwitch++
		} else {
			winsStick++
		}
	}

	t.Logf("Wins after switching: %d (%.2f%%)", winsSwitch, float64(winsSwitch)/float64(numSimulations)*100)
	t.Logf("Wins without switching: %d (%.2f%%)", winsStick, float64(winsStick)/float64(numSimulations)*100)
}
