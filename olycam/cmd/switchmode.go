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
		mode := validateModeName(args[0])
		qty := validateLiveQuality(lvqty)
		cam := camera.NewClient()
		err := cam.SwitchMode(mode, qty)
		exitOnError(err, "switching camera mode")
	},
}

var lvqty string

func init() {
	rootCmd.AddCommand(switchModeCmd)
	switchModeCmd.Flags().StringVarP(&lvqty, "live-quality", "l", "480p", "Live view quality, 240p or 480p")
}

func validateModeName(mode string) camera.CameraMode{
	switch mode {
	case "play": return camera.ModePlay
	case "shutter": return camera.ModeShutter
	case "rec": return camera.ModeRec
	default:
		exitError("Unknown camera mode, please use play, shutter or rec")
	}
	return ""
}

func validateLiveQuality(lvqty string) camera.LiveQuality {
	switch lvqty {
	case "240p": return camera.Live240p
	case "480p": return camera.Live480p
	default:
		exitError("Unsupported preview size: use 240p or 480p")
	}
	return ""
}
