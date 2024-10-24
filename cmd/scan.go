/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github/akstron/MetaManager/pkg/cmderror"
	"github/akstron/MetaManager/pkg/data"
	"github/akstron/MetaManager/pkg/utils"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func scanPath(cmd *cobra.Command, args []string) {
	var err error
	var root, dataFilePath, rootDirPath string
	var mg data.TreeManager
	var serializedNode []byte
	var isDataFileEmpty bool

	isInitialized, err := utils.IsRootInitialized()
	if err != nil {
		goto finally
	}

	if !isInitialized {
		err = &cmderror.UninitializedRoot{}
		goto finally
	}

	root, err = utils.GetAbsMMDirPath()
	if err != nil {
		goto finally
	}

	/*
		Check if the data.json is already written.
		Don't override, if already written
	*/
	dataFilePath, err = filepath.Abs(root + "/data.json")
	if err != nil {
		goto finally
	}

	isDataFileEmpty, err = utils.IsFileEmpty(dataFilePath)
	if err != nil {
		goto finally
	}

	/*
		TODO: Write a more reasonable error
	*/
	if !isDataFileEmpty {
		err = &cmderror.ActionForbidden{}
		goto finally
	}

	// Get parent directory
	// TODO: Update this to extract path from config.json
	rootDirPath, err = filepath.Abs(root + "/..")
	if err != nil {
		goto finally
	}

	// Do other heavy lifting only when data file is empty
	mg = data.TreeManager{
		// Parent of root which is .mm directory path
		DirPath: rootDirPath,
	}

	err = mg.ScanDirectory()
	if err != nil {
		goto finally
	}

	// Save the tree structure in data.json
	serializedNode, err = json.Marshal(mg.Root)
	if err != nil {
		goto finally
	}

	err = os.WriteFile(dataFilePath, serializedNode, 0666)
	if err != nil {
		goto finally
	}

finally:
	if err != nil {
		fmt.Println(err)
	}
}

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: scanPath,
}

func init() {
	rootCmd.AddCommand(scanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
