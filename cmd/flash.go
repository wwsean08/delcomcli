package cmd

import (
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wwsean08/golight"
)

var color *string

// flashCmd represents the flash command
var flashCmd = &cobra.Command{
	Use:   "flash",
	Short: "Begins flashing a light",
	Long:  "Begins flashing a delcom light a given color",
	Run: func(cmd *cobra.Command, args []string) {
		if color == nil || *color == "" {
			log.Panic("color must be set.  The available colors are, green, red, yellow, blue, bluegreen, purple, or white")
		}
		colorCode, err := colorToByte(*color)
		if err != nil {
			log.Panic(err)
		}
		golight.Flash(colorCode)
	},
}

func init() {
	RootCmd.AddCommand(flashCmd)
	color = flashCmd.PersistentFlags().StringP("color", "c", "", "set which color to flash.  The available colors are, green, red, yellow, blue, bluegreen, purple, or white")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// flashCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// flashCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
