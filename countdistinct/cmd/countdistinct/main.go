package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/banaio/golang/countdistinct/lib"
	"github.com/banaio/golang/countdistinct/pcsa"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	algorithm string
	file      string
)

var (
	errUnsupportedAlgorithm = errors.New("algorithm is unsupported")
)

var cmd = &cobra.Command{
	Use:   "countdistinct",
	Short: "count-distinct problem aka cardinality estimation problem",
	Long: `In computer science, the count-distinct problem (also
known in applied mathematics as the cardinality estimation
problem) is the problem of finding the number of distinct
elements in a data stream with repeated elements. This is
a well-known problem with numerous applications. The elements
might represent IP addresses of packets passing through a router,
unique visitors to a web site, elements in a large database,
motifs in a DNA sequence, or elements of RFID/sensor networks.

See: https://en.wikipedia.org/wiki/Count-distinct_problem`,
	RunE: func(cmd *cobra.Command, args []string) error {
		file, err := os.Open(file)
		if err != nil {
			return err
		}
		defer file.Close()

		_, found := lib.Algorithms()[algorithm]
		if !found {
			return errors.Wrap(errUnsupportedAlgorithm, algorithm)
		}

		scanner := bufio.NewScanner(file)

		set, err := makeSet(algorithm)
		if err != nil {
			return err
		}

		for scanner.Scan() {
			// set.Add(scanner.Text())
			if err := set.Add(scanner.Bytes()); err != nil {
				return err
			}
		}

		if err := scanner.Err(); err != nil {
			return err
		}

		// fmt.Printf("pcsa: count=%v\n", set.Count())
		// fmt.Printf("pcsa: %s", pcsa.String())
		fmt.Printf("%+v\n", set.String())

		return nil
	},
}

func makeSet(algorithm string) (lib.Set, error) {
	if algorithm == "pcsa" {
		set, err := pcsa.NewPCSA(16)
		return set, err
	}
	set, err := pcsa.NewPCSA(16)
	return set, err
}

func init() {
	cw, err := os.Getwd()
	if err != nil {
		exitWithError(err)
	}
	defaultFile := fmt.Sprintf("%s/countdistinct_elements.txt", cw)

	cmd.PersistentFlags().StringVarP(&file, "file", "f", defaultFile, "File containing the elements to add")
	cmd.PersistentFlags().StringVarP(&algorithm, "algorithm", "a", "pcsa", "The algorithm to use")
}

func main() {
	if err := cmd.Execute(); err != nil {
		exitWithError(err)
	}
}

func exitWithError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
