package init

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "init",
	Short: "init vblog service",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Herea
		fmt.Println("init....")
	},
}

func init() {
	// Root --> init
}
