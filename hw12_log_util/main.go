package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/pflag"
)

func main() {
	const (
		fileLogDefault   = "LOG_ANALYZER_FILE"
		levelLogDefault  = "LOG_ANALYZER_LEVEL"
		outputLogDefault = "LOG_ANALYZER_OUTPUT"
	)

	requireEnvVars := map[string]string{
		fileLogDefault:   "",
		levelLogDefault:  "",
		outputLogDefault: "",
	}

	if err := loadAndValidateConfig(requireEnvVars); err != nil {
		log.Fatal(err)
	}

	var (
		fileNameLog string
		fileOutput  string
		levelLog    string
	)

	pflag.StringVarP(&fileNameLog, "file", "f", "", "path to file with logs")
	pflag.Lookup("file").NoOptDefVal = requireEnvVars[fileLogDefault]
	pflag.StringVarP(&levelLog, "level", "l", requireEnvVars[levelLogDefault], "level logs by analyzed")
	pflag.Lookup("level").NoOptDefVal = requireEnvVars[levelLogDefault]
	pflag.StringVarP(&fileOutput, "output", "o", "", "path to output analyzed logs")
	pflag.Lookup("output").NoOptDefVal = requireEnvVars[outputLogDefault]
	pflag.Parse()

	if fileNameLog == "" {
		fmt.Println("flag --file (-f) is required")
		os.Exit(1)
	}

	result, err := analyzeStatistic(fileNameLog, levelLog)
	if err != nil {
		log.Fatal(err)
	}

	err = writeOutput(fileOutput, result)
	if err != nil {
		log.Fatal(err)
	}
}

func loadAndValidateConfig(rArgs map[string]string) error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("loading and parse config file: %w", err)
	}

	for rArg := range rArgs {
		value, ok := os.LookupEnv(rArg)

		if !ok {
			return fmt.Errorf("env var not set: %s", rArg)
		}

		rArgs[rArg] = value
	}

	return nil
}

func analyzeStatistic(fileNameLog string, levelLog string) (map[string]int, error) {
	fileLog, err := os.Open(fileNameLog)
	if err != nil {
		return nil, fmt.Errorf("unable to open src file: %w", err)
	}

	defer fileLog.Close()

	result := map[string]int{levelLog: 0}
	br := bufio.NewReader(fileLog)

	for {
		line, rErr := br.ReadString('\n')

		if rErr != nil {
			if rErr == io.EOF {
				break
			}

			return nil, fmt.Errorf("unable to read line from file: %w", rErr)
		}

		if !strings.HasPrefix(line, levelLog) {
			continue
		}

		result[levelLog]++
	}

	return result, nil
}

func writeOutput(fileOutput string, result map[string]int) error {
	var outputStr string
	for levelLog, count := range result {
		outputStr += fmt.Sprintf("Log level: %v | Count: %d\n", levelLog, count)
	}

	if fileOutput == "" {
		for keyLog, value := range result {
			fmt.Printf("Log level: %v | Count: %d\n", keyLog, value)
		}

		return nil
	}

	file, err := os.Create(fileOutput)
	if err != nil {
		return fmt.Errorf("unable to open output file: %w", err)
	}

	defer file.Close()

	_, err = file.WriteString(outputStr)
	if err != nil {
		return fmt.Errorf("unable to write output file: %w", err)
	}

	return nil
}
