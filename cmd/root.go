package cmd

import (
	"fmt"
	"os"

	"github.com/dimiro1/banner"
	"github.com/mattn/go-colorable"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "svelte-cli",
	Short: "Simple CLI for creating svelte apps",
	Long:  "Creator: Wesley Mochiemo. This is a simple CLI for creating svelte applications. Features will include choosing the template",
	Run: func(cmd *cobra.Command, args []string) {
		BannerCreation()
	},
}

func init() {
	rootCmd.AddCommand(presetCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func BannerCreation() {
	templ := `{{ .Title "Svelte CLI" "" 4 }}
	{{ .AnsiColor.BrightCyan }}Welcome to this here awesome program{{ .AnsiColor.Default }}
	GoVersion: {{ .GoVersion }}
	GOOS: {{ .GOOS }}
	GOARCH: {{ .GOARCH }}
	NumCPU: {{ .NumCPU }}
	GOPATH: {{ .GOPATH }}
	GOROOT: {{ .GOROOT }}
	Compiler: {{ .Compiler }}
	ENV: {{ .Env "GOPATH" }}
	Now: {{ .Now "Monday, 2 Jan 2006" }}
	{{ .AnsiColor.BrightGreen }}This text will appear in Green
	{{ .AnsiColor.BrightRed }}This text will appear in Red{{ .AnsiColor.Default }}`
	banner.InitString(colorable.NewColorableStdout(), true, true, templ)
}

type presets struct {
	Name string
	Url  string
}
