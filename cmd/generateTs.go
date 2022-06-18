package cmd

import (
	"fmt"

	"github.com/maxifom/eos-abigen/pkg/commands/generate-ts"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var generateTsCmd = &cobra.Command{
	Use:   "generate-ts [flags] [abi_file]",
	Short: "Generates client and table structures from ABI contract file for Typescript",
	Long:  "Generates client and table structures from ABI contract file for Typescript. \nYou can also provide .eos-abigen.yaml file to generate multiple contracts with one command",
	Args: func(cmd *cobra.Command, args []string) error {
		var contracts []ContractConfig
		viper.UnmarshalKey("generate.contracts", &contracts)

		if len(contracts) > 0 {
			for _, c := range contracts {
				exists, err := fileExists(c.File)
				if err != nil {
					return err
				}
				if !exists {
					return fmt.Errorf("file %s does not exist", c.File)
				}
			}

			return nil
		}

		if len(args) != 1 {
			return fmt.Errorf("1 file is required. provided: %d", len(args))
		}

		exists, err := fileExists(args[0])
		if err != nil {
			return err
		}
		if !exists {
			return fmt.Errorf("file %s does not exist. \nYou can use eos-abigen get-contract %s command to download it", args[0], args[0])
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		contractNameOverride, err := cmd.Flags().GetString("contract_name_override")
		if err != nil {
			return err
		}

		if len(args) > 0 {
			opts := generate.Opts{
				ContractFilePath:     args[0],
				ContractNameOverride: contractNameOverride,
				GeneratedFolder:      viper.GetString("generate.folder"),
				Version:              VERSION,
				FS:                   afero.NewOsFs(),
			}

			err = generate.Run(opts)
			if err != nil {
				return fmt.Errorf("failed to generate for %s: %w", args[0], err)
			}

			return nil
		}

		var contracts []ContractConfig
		err = viper.UnmarshalKey("generate.contracts", &contracts)
		if err != nil {
			return fmt.Errorf("failed to get contracts from config file: %w", err)
		}

		for _, c := range contracts {
			opts := generate.Opts{
				ContractFilePath:     c.File,
				ContractNameOverride: c.NameOverride,
				GeneratedFolder:      viper.GetString("generate.folder"),
				Version:              VERSION,
				FS:                   afero.NewOsFs(),
			}

			err = generate.Run(opts)
			if err != nil {
				return fmt.Errorf("failed to generate for %s: %w", c.File, err)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateTsCmd)
	generateTsCmd.Flags().StringP("contract_name_override", "c", "", "contract name to use in calls to RPC. (default abi filename without extension)")
	generateTsCmd.Flags().StringP("folder", "f", "generated", "folder for generated files output")
	err := viper.BindPFlag("generate.folder", generateTsCmd.Flags().Lookup("folder"))
	if err != nil {
		panic(err)
	}
}
