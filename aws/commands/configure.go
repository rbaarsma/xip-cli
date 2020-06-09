package commands

import (
	"github.com/spf13/cobra"
	"xip/aws/functions"
)

func Configure() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "configure [role]",
		Short: "Configure AWS CLI SSO by providing the role. This role will be provided to you by your administrator.",
		Args:  cobra.ExactArgs(1),
		Run:   ConfigureRun,
	}

	cmd.Flags().StringP("profile", "p", "xip", "The profile name to use")
	cmd.Flags().StringP("region", "r", "eu-west-1", "The region of your org")

	return cmd
}

func ConfigureRun(cmd *cobra.Command, args []string) {
	role := args[0]

	path, _ := cmd.Flags().GetString("config")
	profile, _ := cmd.Flags().GetString("profile")
	region, _ := cmd.Flags().GetString("region")

	functions.CreateOrUpdateSsoProfile(path, profile, role, region)
	functions.SetDefault(path, profile)
}