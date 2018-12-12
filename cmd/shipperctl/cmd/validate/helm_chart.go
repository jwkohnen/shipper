package validate

import (
	"fmt"

	shipperchart "github.com/bookingcom/shipper/pkg/chart"
	"github.com/spf13/cobra"
	"k8s.io/helm/pkg/chartutil"
)

var helmChartCmd = &cobra.Command{
	Use:   "chart",
	Short: "Validate Helm chart",
	RunE:  runValidateHelmChartCommand,
}

func HelmChartCmd() *cobra.Command {
	return helmChartCmd
}

func runValidateHelmChartCommand(cmd *cobra.Command, args []string) error {
	for _, chartPath := range args {
		// TODO: make it understand remote chart URLs
		// use chart.downloadChart/3
		chart, loadErr := chartutil.Load(chartPath)
		if loadErr != nil {
			return fmt.Errorf("Failed to load chart under path %q: %s", chartPath, loadErr.Error())
		}
		if validateErr := shipperchart.Validate(chart); validateErr != nil {
			return fmt.Errorf("Chart validation failed: %s\n", validateErr.Error())
		} else {
			cmd.Printf("Chart %s successfully passed all validation checks\n", chartPath)
		}
	}

	return nil
}
