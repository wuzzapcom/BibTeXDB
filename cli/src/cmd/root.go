package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

var rootCommand = &cobra.Command{
	Use: "cli",
	Run: func (cmd *cobra.Command, args []string){
		fmt.Println("cli")
	},
}


func init(){
	rootCommand.AddCommand(searchCommand)
}

func Execute(){
	rootCommand.Execute()
}
