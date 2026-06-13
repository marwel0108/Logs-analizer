package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"logs-analizer/jsonlog"
)

var (
	inPath  string
	outPath string
)

var rootCmd = &cobra.Command{
	Use: "logs-analizer",
	Run: run,
}

func run(_ *cobra.Command, _ []string) {

	jsonData := []byte(`{"timestamp": "data","level": "ERROR", "message": "TEST", "metadata": "TEST"}`)

	l, err := jsonlog.ParseJsonToLogStruct(jsonData)

	if err != nil {
		log.Fatalf("Error while parsing the json: %v", err)
	}

	fmt.Printf("Log struct value: %v", l)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&inPath, "in", "i", "", "Path to input JSON file")
	rootCmd.Flags().StringVarP(&outPath, "out", "o", "", "Path to output CSV file")
}
