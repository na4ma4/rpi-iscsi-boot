package namegen_test

import (
	"testing"

	"github.com/na4ma4/rpi-iscsi-boot/namegen"
)

func TestLeftRight_Deterministic(t *testing.T) {
	t.Parallel()

	tests := []struct {
		serial       string
		expectedName string
	}{
		{"100000002fdf68ca", "fervent-dijkstra"},
		{"1000000000000000", "condescending-chatterjee"},
		{"1000000011111111", "naughty-heyrovsky"},
		{"1000000022222222", "busy-feistel"},
		{"1000000033333333", "wizardly-moore"},
		{"1000000044444444", "serene-robinson"},
		{"1000000055555555", "jolly-mcnulty"},
		{"1000000066666666", "affectionate-cori"},
		{"1000000077777777", "tender-chandrasekhar"},
		{"1000000088888888", "friendly-swartz"},
		{"1000000099999999", "nervous-pike"},
		{"10000000aaaaaaaa", "beautiful-matsumoto"},
		{"10000000bbbbbbbb", "funny-ptolemy"},
		{"10000000cccccccc", "mystifying-hodgkin"},
		{"10000000dddddddd", "xenodochial-fermi"},
		{"10000000eeeeeeee", "eloquent-tesla"},
		{"10000000ffffffff", "gallant-mahavira"},
		{"10000000AAaaAAaa", "beautiful-matsumoto"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.serial, func(t *testing.T) {
			t.Parallel()

			serial := namegen.Truncate(tt.serial)

			hostname := namegen.LeftRightName(serial)

			if hostname != tt.expectedName {
				t.Errorf("namegen.LeftRightName(%s): got '%s', want '%s'", serial, hostname, tt.expectedName)
			}
		})
	}
}

func TestLeftRight_NonDeterministic(t *testing.T) {
	t.Parallel()

	tests := []string{
		"100000002fdf68cz",
		"100000000000000z",
		"100000001111111z",
		"100000002222222z",
		"100000003333333z",
		"100000004444444z",
		"100000005555555z",
		"100000006666666z",
		"100000007777777z",
		"100000008888888z",
		"100000009999999z",
		"10000000aaaaaaaz",
		"10000000bbbbbbbz",
		"10000000cccccccz",
		"10000000dddddddz",
		"10000000eeeeeeez",
		"10000000fffffffz",
		"10000000AAaaAAaz",
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt, func(t *testing.T) {
			t.Parallel()

			serial := namegen.Truncate(tt)

			hostnames := []string{}
			for i := 0; i < 10; i++ {
				hostnames = append(hostnames, namegen.LeftRightName(serial))
			}

			for i := 0; i < len(hostnames)-1; i++ {
				for j := i + 1; j < len(hostnames); j++ {
					if hostnames[i] == hostnames[j] {
						t.Errorf("namegen.LeftRightName(): expected random names not to re-appear, got duplicate '%s'", hostnames[0])
					}
				}
			}
		})
	}
}
