package cmd

import (
	"dna-ethnicity-compare/pkg/compare"
	"fmt"
	"github.com/spf13/cobra"
)

var compareCmd = &cobra.Command{
	Use:   "compare [csv file]",
	Short: "Compare ethnicity data using different methods",
	Long: `Compare ethnicity data from a CSV file using different statistical methods:
	    - Bayesian method
	    - Simple average
	    - Weighted average`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := compare.LoadCSV(args[0])
		if err != nil {
			return fmt.Errorf("error loading CSV file: %v", err)
		}

		bayesianResults := compare.BayesianMethod(data)
		averageResults := compare.AverageMethod(data)
		weightedAverageResults := compare.WeightedAverageMethod(data)

		compare.DisplayResults(bayesianResults, averageResults, weightedAverageResults)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(compareCmd)
}
