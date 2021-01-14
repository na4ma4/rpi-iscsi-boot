package namegen

import (
	"fmt"
	"math/rand"
	"strconv"
)

// LeftRightName returns a "left right" name based on the docker name dictionary
// using a random seed from the supplied serial.
//nolint:gosec // I know it's weak rand, it's being used deterministically.
func LeftRightName(serial string) string {
	if hex, err := strconv.ParseInt(serial, 16, 64); err == nil {
		rand.Seed(hex)
	}

	left := Left[rand.Intn(len(Left))]
	right := Right[rand.Intn(len(Right))]
	// log.Printf("Random Name[%s]: %s-%s", args[0], left, right)
	return fmt.Sprintf("%s-%s", left, right)
}
