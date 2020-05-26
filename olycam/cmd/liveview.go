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
	"time"

	"github.com/spf13/cobra"

	"github.com/mangelajo/go-olympuscam/camera"
)

// poweroffCmd represents the poweroff command
var livepreviewCmd = &cobra.Command{
	Use:   "liveview",
	Short: "start the livepreview mode",

	Run: func(cmd *cobra.Command, args []string) {
		cam := camera.NewClient()
		err := cam.StartLivePreview()
		exitOnError(err, "setup live preview")
		time.Sleep(5*time.Second)
		cam.StopLivePreview()
	},
}

func init() {
	rootCmd.AddCommand(livepreviewCmd)
}
