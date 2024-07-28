package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	printLine      bool
	numberNonBlank bool
	squeezeBlank   bool
	showEnds       bool
	showTabs       bool
)

var rootCmd = &cobra.Command{
	Use:     "cat [FILE]...",
	Short:   "Concatenate file(s) and print the result to terminal",
	Long:    "A go port of the cat command used to concatenate file(s) and print the result to terminal",
	Example: "cat file1.txt file2.txt",
	Version: "0.0.1",
	Args:    cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		var files []string
		if len(args) == 0 {
			files = append(files, "-")
		} else {
			files = args
		}

		content := getFilesContent(files)
		if squeezeBlank {
			content = squeezeBlankLines(content)
		}
		if printLine {
			content = numberLines(content, numberNonBlank)
		}
		if showEnds {
			content = displayEndOfLineChars(content)
		}
		if showTabs {
			content = displayTabChars(content)
		}
		fmt.Print(content)
	},
}

func getFilesContent(files []string) string {
	var result strings.Builder
	for _, file := range files {
		var content string
		if file == "-" {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				content += scanner.Text() + "\n"
			}
		} else {
			data, err := os.ReadFile(file)
			if err != nil {
				content = err.Error() + "\n"
			} else {
				content = string(data)
			}
		}
		result.WriteString(content)
	}
	return result.String()
}

func numberLines(content string, numberNonBlank bool) string {
	var result strings.Builder
	lines := strings.Split(content, "\n")
	lineNumber := 1
	for _, line := range lines {
		if numberNonBlank && strings.TrimSpace(line) == "" {
			result.WriteString("\n")
		} else {
			result.WriteString(fmt.Sprintf("%4d %s\n", lineNumber, line))
			lineNumber++
		}
	}
	return result.String()
}

func squeezeBlankLines(content string) string {
	var result strings.Builder
	lines := strings.Split(content, "\n")
	prevBlank := false
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			if prevBlank {
				continue
			}
			prevBlank = true
		} else {
			prevBlank = false
		}
		result.WriteString(line + "\n")
	}
	return result.String()
}

func displayEndOfLineChars(content string) string {
	var result strings.Builder
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		result.WriteString(line + "$\n")
	}
	return result.String()
}

func displayTabChars(content string) string {
	return strings.ReplaceAll(content, "\t", "^I")
}

func Execute() {
	rootCmd.PersistentFlags().BoolVarP(&printLine, "number", "n", false, "number all output lines")
	rootCmd.PersistentFlags().BoolVarP(&numberNonBlank, "number-nonblank", "b", false, "number non-blank output lines")
	rootCmd.PersistentFlags().BoolVarP(&squeezeBlank, "squeeze-blank", "s", false, "suppress repeated empty output lines")
	rootCmd.PersistentFlags().BoolVarP(&showEnds, "show-ends", "E", false, "display $ at end of each line")
	rootCmd.PersistentFlags().BoolVarP(&showTabs, "show-tabs", "T", false, "display TAB characters as ^I")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
