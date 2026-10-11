package cmd

import (
	"fmt"
	"os"

	"github.com/Davidvu1/my-go-cli/internal/installer"
	"github.com/Davidvu1/my-go-cli/internal/templates"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{

	Use:   "apply",
	Short: "apply a resource",
	RunE: func(cmd *cobra.Command, args []string) error {
		embedSkills, err := templates.FS.ReadDir("skills")

		if err != nil {
			return fmt.Errorf("there was an error trying to apply a resource: %w", err)
		}

		// get the names of the embedded skills
		var skillNames []string
		for _, skill := range embedSkills {
			skillNames = append(skillNames, skill.Name())
		}

		var selected []string

		form := huh.NewForm(
			huh.NewGroup(
				huh.NewMultiSelect[string]().
					Title("Pick skills to apply").
					Options(huh.NewOptions(skillNames...)...).
					Value(&selected),
			),
		)

		if err := form.Run(); err != nil {
			return fmt.Errorf("run picker: %w", err)
		}
		fmt.Println("you picked:", selected)

		targetDir, err := os.Getwd()

		if err != nil {
			return fmt.Errorf("get current directory: %w", err)
		}

		for _, name := range selected {
			err := installer.CopySkill(templates.FS, name, targetDir)

			if err != nil {
				return fmt.Errorf("copy skill %s: %w", name, err)
			}

		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)

}
