package main

import (
	"log"

	"github.com/na4ma4/rpi-iscsi-boot/namegen"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use: "namegen",
	Run: mainCommand,
}

func init() {
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Debug output")
	_ = viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
	_ = viper.BindEnv("debug", "DEBUG")
}

func main() {
	_ = rootCmd.Execute()
}

const maxSerialLength = 8

func mainCommand(_ *cobra.Command, args []string) {
	if len(args) > 0 {
		name := args[0]
		if len(name) > maxSerialLength {
			name = name[len(name)-8:]
		}

		log.Printf("Deterministic Name[%s]: %s", args[0], namegen.HexName(name))
		log.Printf("Random Name[%s]: %s", args[0], namegen.LeftRightName(name))
	}
}
