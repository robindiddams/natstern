package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
)

var nonZero *bool

func printConsumers(streamName string, includeNonZero bool) error {
	url, opts, err := getContext()
	if err != nil {
		return err
	}

	nc, err := nats.Connect(url, opts...)
	if err != nil {
		return err
	}
	js, err := nc.JetStream()
	if err != nil {
		return err
	}
	var count int
	for data := range js.Consumers(streamName) {
		count++
		if !includeNonZero && data.NumPending == 0 {
			continue
		}
		fmt.Printf("%s pending: %d, ackPending: %d\n", data.Name, data.NumPending, data.NumAckPending)
	}
	fmt.Println()
	fmt.Printf("%d consumers found for stream '%s'\n", count, streamName)
	if !includeNonZero {
		fmt.Println("consumers with zero pending messages omitted, use -z to view them")
	}
	return nil
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "natstern <stream-name>",
	Short: "Tool for scanning nats consumers",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		streamName := args[0]
		printConsumers(streamName, *nonZero)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	log.SetFlags(0)
	nonZero = rootCmd.Flags().BoolP("include-zero", "z", false, "include consumers with zero pending messages")
}
