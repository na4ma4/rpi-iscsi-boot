package namegen

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// LeftRightName returns a "left right" name based on the docker name dictionary
// using a random seed from the supplied serial.
//
//nolint:gosec // I know it's weak rand, it's being used deterministically.
func LeftRightName(serial string) string {
	var rs *rand.Rand
	if hex, err := strconv.ParseInt(serial, 16, 64); err == nil {
		rs = rand.New(rand.NewSource(hex))
	}
	if rs == nil {
		rs = rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	left := Left[rs.Intn(len(Left))]
	right := Right[rs.Intn(len(Right))]
	// log.Printf("Random Name[%s]: %s-%s", args[0], left, right)
	return fmt.Sprintf("%s-%s", left, right)
}
