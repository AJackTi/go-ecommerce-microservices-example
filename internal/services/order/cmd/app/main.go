package main

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:              "orders-microservice",
	Short:            "orders-microservice based on vertical slice architecture",
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func main() {

}
