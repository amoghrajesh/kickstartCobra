/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"strconv"

	"github.com/spf13/cobra"
)

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello called")
		flag, _ := cmd.Flags().GetBool("isfloat")
		if !flag {
			addFloats(args)
		} else {
			addInts(args)
		}
	},
}

func addInts(args []string) int {
	sum := 0
	for _, nums := range args {
		v, _ := strconv.Atoi(nums)
		sum += v
	}
	fmt.Println("Sum of numbers is", sum)
	return sum
}

func addFloats(args []string) {
	var sum float64
	for _, nums := range args {
		v, _ := strconv.ParseFloat(nums, 64)
		sum += v
	}
	fmt.Println("Sum of numbers is", sum)

}

func init() {
	rootCmd.AddCommand(helloCmd)
	helloCmd.Flags().BoolP("isfloat", "f", false, "Are the numbers float?")
}
