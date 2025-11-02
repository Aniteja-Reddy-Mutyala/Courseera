/*
Birthday Paradox Simulator
This program calculates the probability that at least two people 
in a room share the same birthday using Monte Carlo simulation.
It runs simulations for groups of 10 to 100 people (in increments of 10)
and performs 10,000 trials for each group size.
Author: Aniteja Reddy Mutyala
Date: November 2, 2025
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// simulateBirthdays checks if any two people in a group share a birthday
func simulateBirthdays(numPeople int) bool {
	birthdays := make(map[int]bool)
	
	for i := 0; i < numPeople; i++ {
		birthday := rand.Intn(365) + 1  // Random day from 1 to 365
		
		if birthdays[birthday] {
			return true  // Found a matching birthday
		}
		birthdays[birthday] = true
	}
	
	return false  // No matching birthdays
}

// calculateProbability runs multiple trials and calculates probability
func calculateProbability(numPeople, numTrials int) float64 {
	matches := 0
	
	for i := 0; i < numTrials; i++ {
		if simulateBirthdays(numPeople) {
			matches++
		}
	}
	
	return float64(matches) / float64(numTrials)
}

func main() {
	rand.Seed(time.Now().UnixNano())  // Seed for different results each run
	const numTrials = 10000
	
	fmt.Println("Birthday Paradox Probability Simulation")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Printf("Number of trials per group size: %d\n", numTrials)
	fmt.Println()
	fmt.Println("People in Room    Probability of 2+ Same Birthday")
	fmt.Println("------------------------------------------------")
	
	for people := 10; people <= 100; people += 10 {
		probability := calculateProbability(people, numTrials)
		fmt.Printf("%-17d %.6f\n", people, probability)
	}
	
	fmt.Println()
	fmt.Println("Note: Probability > 0.5 typically occurs around 23 people!")
}