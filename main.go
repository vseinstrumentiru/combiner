package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/vseinstrumentiru/combiner/cmd"
)

func main() {
	root := &cobra.Command{
		RunE: func(_ *cobra.Command, _ []string) error {
			v := viper.New()
			v.AddConfigPath(".")
			v.SetConfigName("combine")
			v.SetDefault("path", "./values")
			v.SetDefault("out", "values.yaml")
			v.SetDefault("baseName", "default")
			v.SetDefault("withoutBase", false)
			if err := v.MergeInConfig(); err != nil {
				return err
			}

			var cfg cmd.CombineArgs
			if err := v.Unmarshal(&cfg); err != nil {
				return err
			}

			return cmd.Combine(cfg)
		},
	}

	root.AddCommand(
		cmd.CombineCmd(),
	)

	if err := root.Execute(); err != nil {
		println(err.Error())
	}
}
