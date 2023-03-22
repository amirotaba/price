package cmd

import (
	"github.com/spf13/cobra"
	"price/internal/app"
	"price/internal/app/template"
)

var startCmd = &cobra.Command{
	Use:     "start",
	Aliases: []string{"s"},
	Short:   "starts the service",
	Run: func(cmd *cobra.Command, args []string) {
		app.StartServer()
	},
}

var createCmd = &cobra.Command{
	Use:   "migrate",
	Short: "create the data models on database",
	Run: func(cmd *cobra.Command, args []string) {
		template.CreateItem()
		template.CreateMaterial()
		template.CreateRecipe()
		template.CreateItemRecipe()
	},
}

var dropCmd = &cobra.Command{
	Use:   "drop",
	Short: "drop the data models of database",
	Run: func(cmd *cobra.Command, args []string) {
		template.DropItem()
		template.DropMaterial()
		template.DropRecipe()
		template.DropItemRecipe()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(dropCmd)
}
