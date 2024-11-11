package main

import (
	"crypto/md5"
	"fmt"
	"math/rand"
)

// hashToSeed converts a string to a numerical seed based on its hash.
func hashToSeed(input string) int64 {
	hash := md5.Sum([]byte(input))
	var seed int64
	for _, b := range hash[:8] { // Using first 8 bytes for seed generation
		seed = (seed << 8) | int64(b)
	}
	return seed
}

// generatePastelColor generates a lighter pastel color hex code based on a given seed.
func generatePastelColor(seed int64) string {
	rnd := rand.New(rand.NewSource(seed))

	lightness := 220
	randPart := 255 - lightness
	// Generate pastel color by ensuring each RGB component is in a lighter range.
	r := (rnd.Intn(randPart) + lightness) // Range: 191-255 for lighter colors
	g := (rnd.Intn(randPart) + lightness) // Range: 191-255 for lighter colors
	b := (rnd.Intn(randPart) + lightness) // Range: 191-255 for lighter colors

	// Format as hex color code.
	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}

func GetStringColor(input string) string {

	seed := hashToSeed(input)
	pastelColor := generatePastelColor(seed)

	return pastelColor
}
