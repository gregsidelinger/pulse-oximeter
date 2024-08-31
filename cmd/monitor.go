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
	"net/http"

	"github.com/gregsidelinger/pulse-oximeter/pkg/serial"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// monitorCmd represents the monitor command
var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Read pulse oximeter data from serial device",
	Long: `Read pulse oximeter data.
Push to promometheous
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("monitor called")

		serial.Config.Name = viper.GetString("monitor.device")
		serial.Config.Baud = viper.GetInt("monitor.baud-rate")

		fmt.Println("baud-rate: ", serial.Config.Baud)
		fmt.Println("device: ", serial.Config.Name)

		go serial.Read()

		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":9100", nil)
	},
}

func init() {
	rootCmd.AddCommand(monitorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// monitorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// monitorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	monitorCmd.Flags().StringP("device", "d", "/dev/ttyUSB0", "Serial Device")
	viper.BindPFlag("monitor.device", monitorCmd.Flags().Lookup("device"))
	monitorCmd.Flags().Int("baud-rate", 19200, "Serial Baud Rate")
	viper.BindPFlag("monitor.baud-rate", monitorCmd.Flags().Lookup("baud-rate"))

	monitorCmd.Flags().StringP("push-gateway", "p", "", "Push gateway URL")

}
