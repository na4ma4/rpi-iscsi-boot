package namegen_test

import (
	"testing"

	"github.com/na4ma4/rpi-iscsi-boot/namegen"
)

func TestTruncate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		serial         string
		expectedSerial string
	}{
		{"100000002fdf68ca", "2fdf68ca"},
		{"1000000000000000", "00000000"},
		{"1000000011111111", "11111111"},
		{"1000000022222222", "22222222"},
		{"1000000033333333", "33333333"},
		{"1000000044444444", "44444444"},
		{"1000000055555555", "55555555"},
		{"1000000066666666", "66666666"},
		{"1000000077777777", "77777777"},
		{"1000000088888888", "88888888"},
		{"1000000099999999", "99999999"},
		{"10000000aaaaaaaa", "aaaaaaaa"},
		{"10000000bbbbbbbb", "bbbbbbbb"},
		{"10000000cccccccc", "cccccccc"},
		{"10000000dddddddd", "dddddddd"},
		{"10000000eeeeeeee", "eeeeeeee"},
		{"10000000ffffffff", "ffffffff"},
		{"10000000AAaaAAaa", "AAaaAAaa"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.serial, func(t *testing.T) {
			t.Parallel()

			serial := namegen.Truncate(tt.serial)
			if serial != tt.expectedSerial {
				t.Errorf("namegen.Truncate(): got '%s', want '%s'", serial, tt.expectedSerial)
			}
		})
	}
}
