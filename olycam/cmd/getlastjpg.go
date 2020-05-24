/*
Copyright © 2020 Miguel Angel Ajo <miguelangel@ajo.es>

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
	"os"

	"github.com/spf13/cobra"

	"github.com/mangelajo/go-olympuscam/camera"
)

// poweroffCmd represents the poweroff command
var getLastJpegCmd = &cobra.Command{
	Use:   "get-lastjpg",
	Short: "Get the last jpeg",

	Run: func(cmd *cobra.Command, args []string) {
		cam := camera.NewClient()
		jpegBytes, err := cam.GetLastJpeg()
		exitOnError(err, "Get last jpeg")

		file, err := os.Create(outputJpeg)
		defer file.Close()
		exitOnError(err, "Opening output file")

		_, err = file.Write(jpegBytes)
		exitOnError(err, "Writing jpeg data")
	},
}

var outputJpeg string

func init() {
	rootCmd.AddCommand(getLastJpegCmd)

	getLastJpegCmd.PersistentFlags().StringVarP(&outputJpeg, "file", "f", "lastjpg.jpg",
		"Output jpeg file")

}
