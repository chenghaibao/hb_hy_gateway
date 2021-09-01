package gateway

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewHttpGatewayCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "http",
		Short: "http gateway",
		Long:  "check and distribution",
		Run: func(cmd *cobra.Command, args []string) {
			check()
		},
	}
	return cmd
}

func check() {
	fmt.Println(111)
}
