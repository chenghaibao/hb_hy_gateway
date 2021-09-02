package gateway

import (
	"context"
	"github.com/spf13/cobra"
	"hb_hy_gateway/logic/http"
)

func NewHttpGatewayCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "http",
		Short: "http gateway",
		Long:  "check and distribution",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runHttpGateway()
		},
	}
	return cmd
}

func runHttpGateway() error {
	g := http.NewGateway(context.TODO())
	return g.ListenAndServe()
}
