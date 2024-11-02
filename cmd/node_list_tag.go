/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github/akstron/MetaManager/pkg/cmderror"
	"github/akstron/MetaManager/pkg/data"
	"github/akstron/MetaManager/pkg/utils"
	"path/filepath"
	"runtime/debug"

	"github.com/spf13/cobra"
)

func nodeListInternal(path string) ([]string, error) {
	rw, err := GetRW()
	if err != nil {
		return nil, err
	}

	tgMg := data.NewTagManager()

	err = tgMg.Load(rw)
	if err != nil {
		return nil, err
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	tags, err := tgMg.GetNodeTags(absPath)
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func nodeList(cmd *cobra.Command, args []string) {
	var err error
	var tags []string

	if len(args) != 1 {
		err = &cmderror.InvalidNumberOfArguments{}
		goto finally
	}

	_, err = utils.CommonAlreadyInitializedChecks()
	if err != nil {
		goto finally
	}

	tags, err = nodeListInternal(args[0])
	fmt.Println(tags)

finally:
	if err != nil {
		fmt.Println(err)
		// Print stack trace in case of error
		debug.PrintStack()
	}
}

// listTagCmd represents the listTag command
var nodeListTagCmd = &cobra.Command{
	Use:   "listTag",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run:     nodeList,
	Aliases: []string{"lt"},
}

func init() {
	nodeCmd.AddCommand(nodeListTagCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listTagCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listTagCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}