package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/vseinstrumentiru/combiner/cmd"
)

type envFile struct {
	Name  string
	Range *[2]int
	Files []string
}

type fileConfig struct {
	Path     string
	Out      string
	BaseName string
	Envs     []envFile `mapstrusture:"groups"`
}

func (f fileConfig) toCombineArgs() cmd.CombineArgs {
	arg := cmd.CombineArgs{
		Path:     f.Path,
		Out:      f.Out,
		BaseName: f.BaseName,
		Groups:   make(map[string][]string),
	}

	if arg.Path == "" {
		arg.Path = "./values"
	}

	if arg.Out == "" {
		arg.Out = "values.yaml"
	}

	for _, env := range f.Envs {
		if env.Range != nil {
			for i := (*env.Range)[0]; i <= (*env.Range)[1]; i++ {
				name := fmt.Sprintf(env.Name, i)
				arg.Groups[name] = append(arg.Groups[name], env.Files...)
			}
		} else {
			name := env.Name
			arg.Groups[name] = append(arg.Groups[name], env.Files...)
		}
	}

	return arg
}

type config struct {
	Items []fileConfig `mapstructure:"combine"`
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

			for _, file := range cfg.Items {
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
