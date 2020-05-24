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
	"github.com/spf13/cobra"

	"github.com/mangelajo/go-olympuscam/camera"
)

// switchModeCmd represents the switchMode command
var switchModeCmd = &cobra.Command{
	Use:   "switch-mode [play, rec or shutter]",
	Short: "Switch camera mode between play, rec, and shutter",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			exitError("switch-mode needs the mode to set the camera in (play, rec or shutter)")
		}
		cam := camera.NewClient()
		err := cam.SwitchMode(camera.CameraMode(args[0]))
		exitOnError(err, "switching camera mode")
	},
}

func init() {
	rootCmd.AddCommand(switchModeCmd)
}
