package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/banaio/golang/openbanking/tools/swagger"
)

var (
	logger  = logrus.StandardLogger()
	rootCmd = &cobra.Command{
		Use:   "openbanking_tools",
		Short: "TODO",
		Long:  `TODO`,
	}
	conditionalPropertiesCmd = &cobra.Command{
		Use:   "conditional_properties",
		Short: "TODO",
		Long:  `TODO`,
		// Neither of the two lines below seem to work ... :(
		// Args:  cobra.ExactArgs(2),
		// Args:  cobra.ExactValidArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := logger.WithFields(logrus.Fields{
				// "app": "fcs_discovery_scripts",
			})

			swaggerFile := viper.GetString("swagger_file")
			outputFile := viper.GetString("output_file")
			logger.WithFields(logrus.Fields{
				"swagger_file": swaggerFile,
				"output_file":  outputFile,
			}).Infof("Parsing started ...")

			nonManadatoryFields, err := swagger.ParseSchema(swaggerFile, logger.Logger)
			if err != nil {
				return err
			}

			logger.WithFields(logrus.Fields{
				"total": len(nonManadatoryFields.Endpoints),
			}).Infof("Finished generating conditional properties")

			data, err := json.MarshalIndent(nonManadatoryFields, "", "  ")
			if err != nil {
				return err
			}

			dirName := filepath.Dir(outputFile)
			// Make intermediate directories, e.g., if given the path `generated/v3.1/account-info-swagger.json` create
			// `v3.1` if necessary.
			if _, err := os.Stat(dirName); err != nil && !os.IsExist(err) {
				if err := os.MkdirAll(dirName, os.ModePerm); err != nil && !os.IsExist(err) {
					return err
				}
			}

			err = ioutil.WriteFile(outputFile, data, 0644)
			if err != nil {
				return err
			}

			logger.WithFields(logrus.Fields{
				"output_file": outputFile,
			}).Infof("Finished writing to file ...")

			return nil
		},
	}
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}

func init() {
	viper.SetEnvPrefix("")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	rootCmd.AddCommand(conditionalPropertiesCmd)
	rootCmd.PersistentFlags().String("log_level", "INFO", "Log level")
	if err := viper.BindPFlags(rootCmd.PersistentFlags()); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}

	conditionalPropertiesCmd.PersistentFlags().StringP("swagger_file", "s", "", "Swagger file path, e.g., 'specifications/read-write/v3.1/account-info-swagger.yaml'")
	conditionalPropertiesCmd.PersistentFlags().StringP("output_file", "o", "", "Swagger file path, e.g., 'generated/v3.1/account-info-swagger.json'")
	if err := viper.BindPFlags(conditionalPropertiesCmd.PersistentFlags()); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
	if err := cobra.MarkFlagRequired(conditionalPropertiesCmd.PersistentFlags(), "swagger_file"); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
	if err := cobra.MarkFlagRequired(conditionalPropertiesCmd.PersistentFlags(), "output_file"); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}

	cobra.OnInitialize(onInitialize)
}

func onInitialize() {
	logger.SetNoLock()
	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors:    false,
		ForceColors:      true,
		FullTimestamp:    false,
		DisableTimestamp: true,
	})
	level, err := logrus.ParseLevel(viper.GetString("log_level"))
	if err != nil {
		printCommandFlags()
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
	logger.SetLevel(level)

	printCommandFlags()
}

func printCommandFlags() {
	rootCmd.PersistentFlags().PrintDefaults()

	logger.WithFields(logrus.Fields{
		"log_level": viper.GetString("log_level"),
	}).Info("flags")
}
