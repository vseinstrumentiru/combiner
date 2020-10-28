package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Combine() *cobra.Command {
	cfg := combineArgs{groups: map[string][]string{}}
	cmd := &cobra.Command{
		Use:   "combine group1:file1,file2 file3 group3:file4,file5 ...",
		Short: "Combines several config files into one.",
		Long: `All files in groups will be merged and overwrite the default config file.
All merged config writes by sections, where section name is a group name.`,
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			for i, arg := range args {
				parsed := strings.Split(arg, ":")

				switch len(parsed) {
				case 1:
					cfg.groups[parsed[0]] = []string{parsed[0]}
				case 2:
					cfg.groups[parsed[0]] = strings.Split(parsed[1], ",")
				default:
					return fmt.Errorf("wrong argument %d", i+1)
				}
			}

			return combine(cfg)
		},
	}

	cmd.Flags().StringVarP(&cfg.path, "path", "p", "./values", "Folder with yaml configs to merge")
	cmd.Flags().StringVarP(&cfg.out, "out", "o", "values.yaml", "Output file name")
	cmd.Flags().StringVarP(
		&cfg.baseName,
		"default",
		"d",
		"default",
		"Default config file name (without extension). This is base config file for other groups",
	)
	cmd.Flags().BoolVarP(&cfg.withoutBase, "no-default", "n", false, "Without default file config")

	return cmd
}

type combineArgs struct {
	path        string
	out         string
	withoutBase bool
	baseName    string
	groups      map[string][]string
}

func combine(args combineArgs) error {
	writer := viper.New()

	for key, group := range args.groups {
		reader := viper.New()
		reader.AddConfigPath(args.path)

		if !args.withoutBase {
			reader.SetConfigName(args.baseName)
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

	err := writer.WriteConfigAs(args.out)

	if err == nil {
		fmt.Printf("successful combined to %s\n", args.out)
	}

	return err
}
