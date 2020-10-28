package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func CombineCmd() *cobra.Command {
	cfg := CombineArgs{Groups: map[string][]string{}}
	cmd := &cobra.Command{
		Use:   "combine group1:file1,file2 file3 group3:file4,file5 ...",
		Short: "Combines several config files into one.",
		Long: `All files in groups will be merged and overwrite the default config file.
All merged config writes by sections, where section name is a group name.`,
		Args: cobra.MinimumNArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			for i, arg := range args {
				parsed := strings.Split(arg, ":")

				switch len(parsed) {
				case 1:
					cfg.Groups[parsed[0]] = []string{parsed[0]}
				case 2:
					cfg.Groups[parsed[0]] = strings.Split(parsed[1], ",")
				default:
					return fmt.Errorf("wrong argument %d", i+1)
				}
			}

			return Combine(cfg)
		},
	}

	cmd.Flags().StringVarP(&cfg.Path, "path", "p", "./values", "Folder with yaml configs to merge")
	cmd.Flags().StringVarP(&cfg.Out, "out", "o", "values.yaml", "Output file name")
	cmd.Flags().StringVarP(
		&cfg.BaseName,
		"default",
		"d",
		"default",
		"Default config file name (without extension). This is base config file for other Groups",
	)
	cmd.Flags().BoolVarP(&cfg.WithoutBase, "no-default", "n", false, "Without default file config")

	return cmd
}

type CombineArgs struct {
	Path        string
	Out         string
	WithoutBase bool
	BaseName    string
	Groups      map[string][]string
}

func Combine(args CombineArgs) error {
	writer := viper.New()

	for key, group := range args.Groups {
		reader := viper.New()
		reader.AddConfigPath(args.Path)

		if !args.WithoutBase {
			reader.SetConfigName(args.BaseName)
			if err := reader.MergeInConfig(); err != nil {
				return err
			}
		}

		for _, name := range group {
			reader.SetConfigName(name)

			if err := reader.MergeInConfig(); err != nil {
				return err
			}

			writer.Set(key, reader.AllSettings())
		}
	}

	err := writer.WriteConfigAs(args.Out)

	if err == nil {
		fmt.Printf("successful combined to %s\n", args.Out)
	}

	return err
}
