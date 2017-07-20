package cmd

import (
	log "github.com/Sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/wwsean08/golight"
)

var onColor *string

// onCmd represents the on command
var onCmd = &cobra.Command{
	Use:   "on",
	Short: "Turns on a light",
	Long:  "Turns on a delcom USB light.  If you already have one light on that light will not be turned off",
	Run: func(cmd *cobra.Command, args []string) {
		if onColor == nil || *onColor == "" {
			log.Panic("color must be set.  The available colors are, green, red, yellow, blue, bluegreen, purple, or white")
		}
		colorCode, err := colorToByte(*onColor)
		if err != nil {
			log.Panic(err)
		}
		golight.On(colorCode)
	},
}

func init() {
	RootCmd.AddCommand(onCmd)
	onColor = onCmd.PersistentFlags().StringP("color", "c", "", "set which color to flash.  The available colors are, green, red, yellow, blue, bluegreen, purple, or white")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// onCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// onCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
