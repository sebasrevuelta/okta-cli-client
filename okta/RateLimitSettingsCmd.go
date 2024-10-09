package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var RateLimitSettingsCmd = &cobra.Command{
	Use:  "rateLimitSettings",
	Long: "Manage RateLimitSettingsAPI",
}

func init() {
	rootCmd.AddCommand(RateLimitSettingsCmd)
}

func NewGetRateLimitSettingsAdminNotificationsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "getAdminNotifications",

		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RateLimitSettingsAPI.GetRateLimitSettingsAdminNotifications(apiClient.GetConfig().Context)

			resp, err := req.Execute()
			if err != nil {
				return err
			}
			d, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			utils.PrettyPrintByte(d)
			// cmd.Println(string(d))
			return nil
		},
	}

	return cmd
}

func init() {
	GetRateLimitSettingsAdminNotificationsCmd := NewGetRateLimitSettingsAdminNotificationsCmd()
	RateLimitSettingsCmd.AddCommand(GetRateLimitSettingsAdminNotificationsCmd)
}

var ReplaceRateLimitSettingsAdminNotificationsdata string

func NewReplaceRateLimitSettingsAdminNotificationsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "replaceAdminNotifications",

		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RateLimitSettingsAPI.ReplaceRateLimitSettingsAdminNotifications(apiClient.GetConfig().Context)

			if ReplaceRateLimitSettingsAdminNotificationsdata != "" {
				req = req.Data(ReplaceRateLimitSettingsAdminNotificationsdata)
			}

			resp, err := req.Execute()
			if err != nil {
				return err
			}
			d, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			utils.PrettyPrintByte(d)
			// cmd.Println(string(d))
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplaceRateLimitSettingsAdminNotificationsdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	return cmd
}

func init() {
	ReplaceRateLimitSettingsAdminNotificationsCmd := NewReplaceRateLimitSettingsAdminNotificationsCmd()
	RateLimitSettingsCmd.AddCommand(ReplaceRateLimitSettingsAdminNotificationsCmd)
}

func NewGetRateLimitSettingsPerClientCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "getPerClient",

		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RateLimitSettingsAPI.GetRateLimitSettingsPerClient(apiClient.GetConfig().Context)

			resp, err := req.Execute()
			if err != nil {
				return err
			}
			d, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			utils.PrettyPrintByte(d)
			// cmd.Println(string(d))
			return nil
		},
	}

	return cmd
}

func init() {
	GetRateLimitSettingsPerClientCmd := NewGetRateLimitSettingsPerClientCmd()
	RateLimitSettingsCmd.AddCommand(GetRateLimitSettingsPerClientCmd)
}

var ReplaceRateLimitSettingsPerClientdata string

func NewReplaceRateLimitSettingsPerClientCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "replacePerClient",

		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RateLimitSettingsAPI.ReplaceRateLimitSettingsPerClient(apiClient.GetConfig().Context)

			if ReplaceRateLimitSettingsPerClientdata != "" {
				req = req.Data(ReplaceRateLimitSettingsPerClientdata)
			}

			resp, err := req.Execute()
			if err != nil {
				return err
			}
			d, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			utils.PrettyPrintByte(d)
			// cmd.Println(string(d))
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplaceRateLimitSettingsPerClientdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	return cmd
}

func init() {
	ReplaceRateLimitSettingsPerClientCmd := NewReplaceRateLimitSettingsPerClientCmd()
	RateLimitSettingsCmd.AddCommand(ReplaceRateLimitSettingsPerClientCmd)
}

func NewGetRateLimitSettingsWarningThresholdCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "getWarningThreshold",

		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RateLimitSettingsAPI.GetRateLimitSettingsWarningThreshold(apiClient.GetConfig().Context)

			resp, err := req.Execute()
			if err != nil {
				return err
			}
			d, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			utils.PrettyPrintByte(d)
			// cmd.Println(string(d))
			return nil
		},
	}

	return cmd
}

func init() {
	GetRateLimitSettingsWarningThresholdCmd := NewGetRateLimitSettingsWarningThresholdCmd()
	RateLimitSettingsCmd.AddCommand(GetRateLimitSettingsWarningThresholdCmd)
}

var ReplaceRateLimitSettingsWarningThresholddata string

func NewReplaceRateLimitSettingsWarningThresholdCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "replaceWarningThreshold",

		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RateLimitSettingsAPI.ReplaceRateLimitSettingsWarningThreshold(apiClient.GetConfig().Context)

			if ReplaceRateLimitSettingsWarningThresholddata != "" {
				req = req.Data(ReplaceRateLimitSettingsWarningThresholddata)
			}

			resp, err := req.Execute()
			if err != nil {
				return err
			}
			d, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			utils.PrettyPrintByte(d)
			// cmd.Println(string(d))
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplaceRateLimitSettingsWarningThresholddata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	return cmd
}

func init() {
	ReplaceRateLimitSettingsWarningThresholdCmd := NewReplaceRateLimitSettingsWarningThresholdCmd()
	RateLimitSettingsCmd.AddCommand(ReplaceRateLimitSettingsWarningThresholdCmd)
}