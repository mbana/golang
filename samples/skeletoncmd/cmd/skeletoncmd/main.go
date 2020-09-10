package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"

	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	choices = []string{
		"stat",
		"contents",
	}
	choice   string
	filename string
)

var (
	errUnsupportedChoice = errors.New("unsupported choice")
)

var cmd = &cobra.Command{
	Use:   "skeletoncmd",
	Short: "skeletoncmd",
	Long:  `skeletoncmd`,
	RunE: func(cmd *cobra.Command, args []string) error {
		file, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer file.Close()

		sort.StringSlice(choices).Sort()
		choiceIndex := sort.SearchStrings(choices, choice)
		if !(choiceIndex < len(choices) && choices[choiceIndex] == choice) {
			return errors.Wrap(errUnsupportedChoice, choice)
		}

		if choice == "stat" {
			stat, err := file.Stat()
			if err != nil {
				return err
			}

			fmt.Printf("filename=%#v\n", filename)
			fmt.Printf("stat=%#v\n", stat)
			return nil
		}

		if choice == "contents" {
			fileContents, err := ioutil.ReadFile(filename)
			if err != nil {
				return err
			}
			fmt.Printf("filename=%#v\n", filename)
			fmt.Printf("contents=%#v\n", string(fileContents))
		}

		return nil
	},
}

func init() {
	cw, err := os.Getwd()
	if err != nil {
		exitWithError(err)
	}
	defaultFile := fmt.Sprintf("%s/skeletoncmd.txt", cw)

	cmd.PersistentFlags().StringVarP(&filename, "filename", "f", defaultFile, "file")
	cmd.PersistentFlags().StringVarP(&choice, "choice", "c", "", fmt.Sprintf("Available choices: %+v", choices))
}

func main() {
	if err := cmd.Execute(); err != nil {
		exitWithError(err)
	}
}

func exitWithError(err error) {
	red := color.New(color.FgRed)
	red.Fprint(os.Stderr, err, "\n")
	os.Exit(1)
}
