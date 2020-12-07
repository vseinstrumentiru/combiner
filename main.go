package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/vseinstrumentiru/combiner/cmd"
)

type config struct {
	Files []cmd.CombineArgs
}

func main() {
	root := &cobra.Command{
		RunE: func(_ *cobra.Command, _ []string) error {
			v := viper.New()
			v.AddConfigPath(".")
			v.SetConfigName("combine")

			if err := v.MergeInConfig(); err != nil {
				return err
			}

			var cfg config
			if err := v.Unmarshal(&cfg); err != nil {
				return err
			}

			for _, file := range cfg.Files {
				if file.Path == "" {
					file.Path = "./values"
				}

				if file.Out == "" {
					file.Out = "values.yaml"
				}

				if err := cmd.Combine(file); err != nil {
					return err
				}
			}

			return nil
		},
	}

	root.AddCommand(
		cmd.CombineCmd(),
	)

	if err := root.Execute(); err != nil {
		println(err.Error())
	}
}
