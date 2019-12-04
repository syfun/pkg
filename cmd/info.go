/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	_package "github.com/syfun/package/pkg/package"
	"log"

	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		getPackage(args[0])
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getPackage(name string) {
	url := fmt.Sprintf("%v/api/v1/packages/%v/", viper.GetString("server"), name)
	resp, err := Get(url)
	if err != nil {
		log.Fatal(err)
	}
	var p _package.Package
	if err := resp.Decode(&p); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Package:\n\nID: %v\nName: %v\n\n", p.ID, p.Name)
	if p.Versions == nil || len(p.Versions) == 0 {
		return
	}
	var data [][]interface{}
	for _, v := range p.Versions {
		data = append(data, []interface{}{v.ID, v.Name, v.Size, v.FileName})
	}
	printTable([]string{"ID", "Name", "Size", "FileName"}, data)
}