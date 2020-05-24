/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"strings"

	"github.com/spf13/cobra"

	"github.com/mangelajo/go-olympuscam/camera"
)

// getCommandlistCmd represents the getCommandlist command
var getParametersCmd = &cobra.Command{
	Use:   "get-parameters",
	Short: "Get the list of parameters",
	Run: func(cmd *cobra.Command, args []string) {
		cam := camera.NewClient()
		parameterList, err := cam.GetParameters()
		exitOnError(err, "getting parameter list")
		for _, x := range parameterList.Desc {
			val := x.Value
			options := strings.Split(x.Enum, " ")
			valPrinted := false

			fmt.Printf("%s:\t", x.Propname)
			for _, o := range options {
				if o == val {
					fmt.Printf("[%s] ", o)
					valPrinted = true
				} else {
					fmt.Printf("%s ",o)
				}
			}
			if valPrinted == false {
				fmt.Print(val)
			}
			fmt.Print("\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(getParametersCmd)
}
