package cmd

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var presetCmd = &cobra.Command{
	Use:   "create",
	Short: "Select a simple preset to install svelte template",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			if err := getPresets("svelte-template"); err != nil {
				return err
			}
		} else {
			folderName := args[0]
			if err := getPresets(folderName); err != nil {
				return err
			}
		}
		return nil
	},
}

func getPresets(folder string) error {
	items := []presets{
		{Name: "svelte-rollup", Url: "https://github.com/sveltejs/template.git"},
		{Name: "svelte-snowpack", Url: "https://svelte.dev"},
	}
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "{{ .Name | cyan}}",
		Inactive: " {{ .Name | cyan}}",
		Selected: "{{ .Name | cyan }}",
	}

	searcher := func(input string, index int) bool {
		item := items[index]
		name := strings.Replace(strings.ToLower(item.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)
		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Template type",
		Items:     items,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v \n", err)
		return err
	}
	fmt.Printf("You chose number %d:%s\n", items[i].Name)
	switch items[i].Name {
	case "svelte-rollup":
		fmt.Println("Installing the rollup template")
		clone := exec.Command("git", "clone", "--depth", "1", items[i].Url, folder)
		clone.Run()
		removeGit := exec.Command("rm", "-rf", "/svelty-trial/.git")
		removeGit.Run()
		fmt.Println("Installation done. Kindly change directory")

	case "svelte-snowpack":
		fmt.Println("Installing the Snowpack template")
		clone := exec.Command("npx", "create-snowpack-app", folder, "--template", "@snowpack/app-template-svelte")
		var stdoutBuf, stderrBuf bytes.Buffer
		clone.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
		clone.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

		err := clone.Run()
		if err != nil {
			log.Fatal(err)
		}
	}

	// clone := exec.Command("git", "clone", "git@github.com:wesleymutwiri/svelte-ecommerce-design.git")
	// clone.Run()
	return nil
}
