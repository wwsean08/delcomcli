package cmd

import (
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wwsean08/golight"
)

var offColor *string

// offCmd represents the off command
var offCmd = &cobra.Command{
	Use:   "off",
	Short: "Turns off a light",
	Long:  "Turns off a delcom USB light whether it's flashing or not.",
	Run: func(cmd *cobra.Command, args []string) {
		if offColor == nil || *offColor == "" {
			log.Panic("color must be set.  The available colors are, green, red, yellow, blue, bluegreen, purple, or white")
		}
		colorCode, err := colorToByte(*offColor)
		if err != nil {
			log.Panic(err)
		}
		golight.Off(colorCode)
	},
}

func init() {
	RootCmd.AddCommand(offCmd)
	offColor = offCmd.PersistentFlags().StringP("color", "c", "", "set which color to flash.  The available colors are, green, red, yellow, blue, bluegreen, purple, or white")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// offCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// offCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
