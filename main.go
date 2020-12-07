package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/vseinstrumentiru/combiner/cmd"
)

type fileConfig struct {
	Path     string
	Out      string
	BaseName string
	Groups   map[string]map[string]interface{}
}

func (f fileConfig) toCombineArgs() cmd.CombineArgs {
	arg := cmd.CombineArgs{
		Path:     f.Path,
		Out:      f.Out,
		BaseName: f.BaseName,
		Groups:   make(map[string][]string, len(f.Groups)),
	}

	for s, m := range f.Groups {
		for n, _ := range m {
			arg.Groups[s] = append(arg.Groups[s], n)
		}
	}

	return arg
}

type config struct {
	Files []fileConfig
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

				if err := cmd.Combine(file.toCombineArgs()); err != nil {
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
