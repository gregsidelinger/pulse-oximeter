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
	"log"
	"net/http"

	"github.com/gregsidelinger/pulse-oximeter/pkg/simulate"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/spf13/cobra"
)

// simulateCmd represents the simulate command
var simulateCmd = &cobra.Command{
	Use:   "simulate",
	Short: "Generate fake oximeter data",
	Long: `Generate fake oximeter data.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("simulate called")
		spo2Min, err := cmd.Flags().GetInt("spo2-min")
		if err != nil {
			log.Fatal(err)
		}
		spo2Max, err := cmd.Flags().GetInt("spo2-max")
		if err != nil {
			log.Fatal(err)
		}
		bpaMin, err := cmd.Flags().GetInt("bpa-min")
		if err != nil {
			log.Fatal(err)
		}
		bpaMax, err := cmd.Flags().GetInt("bpa-max")
		if err != nil {
			log.Fatal(err)
		}
		bpMin, err := cmd.Flags().GetInt("bp-min")
		if err != nil {
			log.Fatal(err)
		}
		bpMax, err := cmd.Flags().GetInt("bp-max")
		if err != nil {
			log.Fatal(err)
		}
		go simulate.Read(spo2Min, spo2Max, bpaMin, bpaMax, bpMin, bpMax)

		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":9100", nil)
	},
}

func init() {
	rootCmd.AddCommand(simulateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// simulateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// simulateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	simulateCmd.Flags().Int("spo2-min", 80, "Spo2 min simulated value")
	simulateCmd.Flags().Int("spo2-max", 100, "Spo2 max simulated value")

	simulateCmd.Flags().Int("bpa-min", 60, "bpa min simulated value")
	simulateCmd.Flags().Int("bpa-max", 160, "bpa max simulated value")

	simulateCmd.Flags().Int("bp-min", 10, "bp min simulated value")
	simulateCmd.Flags().Int("bp-max", 40, "bp max simulated value")

}
