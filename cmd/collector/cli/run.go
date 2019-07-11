package cli

import (
	"io/ioutil"

	"github.com/replicatedhq/troubleshoot/pkg/collect"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Run() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "run a single collector",
		Long:  `...`,
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlag("collector", cmd.Flags().Lookup("collector"))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			v := viper.GetViper()

			specContents, err := ioutil.ReadFile(v.GetString("collector"))
			if err != nil {
				return err
			}

			collector := collect.Collector{
				Spec: string(specContents),
			}
			if err := collector.RunCollectorSync(); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().String("collector", "", "path to a single collector spec to collect")

	cmd.MarkFlagRequired("collector")

	viper.BindPFlags(cmd.Flags())

	return cmd
}