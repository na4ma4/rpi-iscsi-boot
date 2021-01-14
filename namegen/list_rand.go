package namegen

import (
	"strconv"
	"strings"
)

// HexName takes a Raspberry Pi serial number (truncated to 8 bytes) and
// returns a name based on each 4 bit segment.
func HexName(serial string) string {
	out := []string{}

	for i := 0; i < len(serial); i++ {
		if hex, err := strconv.ParseInt(string(serial[i]), 16, 64); err == nil {
			out = append(out, List[hex])
		}
	}

	return strings.Join(out, "-")
}
