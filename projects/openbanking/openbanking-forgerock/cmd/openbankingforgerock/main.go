package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/banaio/openbankingforgerock/config"
	"github.com/banaio/openbankingforgerock/lib"
)

// flags
var (
	verbose              bool   // nolint:gochecknoglobals
	registerResponseFile string // nolint:gochecknoglobals
)

func createRootCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "openbankingforgerock",
		Short: "openbankingforgerock",
		Long: `openbankingforgerock
Connect to ForgeRock's directory:

1. Registers on ForgeRock's directory.
2. Tests the MATLS setup.
3. Get an access token to represent you as a TPP using the Client credential flow.
4. Create an account request.
5. Consume the accounts API.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	return rootCmd
}

func createVersionCmd() *cobra.Command {
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of openbankingforgerock",
		Long:  `All software has versions. This is openbankingforgerock's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("https://github.com/banaio/openbankingforgerock v0.0.1 -- HEAD\n")
		},
	}
	return versionCmd
}

func createValidateCmd() *cobra.Command {
	var validateCmd = &cobra.Command{
		Use:   "validate",
		Short: "validate",
		Long:  `validate`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	var validateConfigCmd = &cobra.Command{
		Use:   "config",
		Short: "config",
		Long:  `config`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if _, err := config.NewConfig("./config.yml"); err != nil {
				return err
			}
			return nil
		},
	}
	validateCmd.AddCommand(validateConfigCmd)

	return validateCmd
}

func createMTLSCmd() *cobra.Command {
	var mtlsCmd = &cobra.Command{
		Use:   "mtls",
		Short: "mtls",
		Long:  `mtls`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("mtls:\n")

			config, err := config.NewConfig("./.ignore/config.yml")
			if err != nil {
				return err
			}

			client, err := lib.NewClient(config)
			if err != nil {
				return err
			}

			res, err := client.MTLSTest()
			if err != nil {
				return err
			}
			fmt.Printf("%+v\n", res)

			return nil
		},
	}
	return mtlsCmd
}

func createRegisterCmd() *cobra.Command {
	var registerCmd = &cobra.Command{
		Use:   "register",
		Short: "register",
		Long:  `register`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("register:\n")
			fmt.Printf("register: register-response=%+v\n", registerResponseFile)
			return nil
		},
	}
	return registerCmd
}

func createAccountsCmd() *cobra.Command {
	var accountsCmd = &cobra.Command{
		Use:   "accounts",
		Short: "accounts",
		Long:  `accounts`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("accounts:\n")
			fmt.Printf("accounts: register-response=%+v\n", registerResponseFile)
			return nil
		},
	}
	return accountsCmd
}

func initCommands() (*cobra.Command, error) {
	// wd, err := os.Getwd()
	// if err != nil {
	// 	return err
	// }
	// registerResponseFileDefault := fmt.Sprintf("%s/config/register-response.json", wd)

	rootCmd := createRootCmd()
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	rootCmd.AddCommand(createVersionCmd())

	rootCmd.AddCommand(createValidateCmd())
	rootCmd.AddCommand(createMTLSCmd())

	registerCmd := createRegisterCmd()
	registerCmd.Flags().StringVarP(&registerResponseFile, "register-response", "r", "", "File to store OpenID Connect Dynamic Client Registration response")
	if err := registerCmd.MarkFlagRequired("register-response"); err != nil {
		return nil, err
	}
	rootCmd.AddCommand(registerCmd)

	accountsCmd := createAccountsCmd()
	accountsCmd.Flags().StringVarP(&registerResponseFile, "register-response", "r", "", "File containing OpenID Connect Dynamic Client Registration response")
	if err := accountsCmd.MarkFlagRequired("register-response"); err != nil {
		return nil, err
	}
	rootCmd.AddCommand(accountsCmd)

	return rootCmd, nil
}

func main() {
	exitWithError := func(err error) {
		red := color.New(color.FgRed)
		red.Fprint(os.Stderr, err, "\n")
		os.Exit(1)
	}

	rootCmd, err := initCommands()
	if err != nil {
		exitWithError(err)
	}

	if err := rootCmd.Execute(); err != nil {
		exitWithError(err)
	}
}
