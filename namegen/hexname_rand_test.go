package namegen_test

import (
	"testing"

	"github.com/na4ma4/rpi-iscsi-boot/namegen"
)

func TestHexName_Deterministic(t *testing.T) {
	t.Parallel()

	tests := []struct {
		serial       string
		expectedName string
	}{
		{"100000002fdf68ca", "cool-pike-wing-pike-nice-zen-borg-benz"},
		{"1000000000000000", "bold-bold-bold-bold-bold-bold-bold-bold"},
		{"1000000011111111", "busy-busy-busy-busy-busy-busy-busy-busy"},
		{"1000000022222222", "cool-cool-cool-cool-cool-cool-cool-cool"},
		{"1000000033333333", "epic-epic-epic-epic-epic-epic-epic-epic"},
		{"1000000044444444", "keen-keen-keen-keen-keen-keen-keen-keen"},
		{"1000000055555555", "kind-kind-kind-kind-kind-kind-kind-kind"},
		{"1000000066666666", "nice-nice-nice-nice-nice-nice-nice-nice"},
		{"1000000077777777", "sad-sad-sad-sad-sad-sad-sad-sad"},
		{"1000000088888888", "zen-zen-zen-zen-zen-zen-zen-zen"},
		{"1000000099999999", "bell-bell-bell-bell-bell-bell-bell-bell"},
		{"10000000aaaaaaaa", "benz-benz-benz-benz-benz-benz-benz-benz"},
		{"10000000bbbbbbbb", "ride-ride-ride-ride-ride-ride-ride-ride"},
		{"10000000cccccccc", "borg-borg-borg-borg-borg-borg-borg-borg"},
		{"10000000dddddddd", "wing-wing-wing-wing-wing-wing-wing-wing"},
		{"10000000eeeeeeee", "ride-ride-ride-ride-ride-ride-ride-ride"},
		{"10000000ffffffff", "pike-pike-pike-pike-pike-pike-pike-pike"},
		{"10000000AAaaAAaa", "benz-benz-benz-benz-benz-benz-benz-benz"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.serial, func(t *testing.T) {
			t.Parallel()

			serial := namegen.Truncate(tt.serial)

			hostname := namegen.HexName(serial)

			if hostname != tt.expectedName {
				t.Errorf("namegen.LeftRightName(%s): got '%s', want '%s'", serial, hostname, tt.expectedName)
			}
		})
	}
}
