package cmd

import (
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the DNS record to use the current public IP",
	Long:  "This command will update the DNS record for the provided FQDN using the public IP address for this host.",
	Args:  cobra.ExactArgs(1),
	RunE:  executeUpdate,
}

func init() {
	usage := `Usage:
  canary update <FQDN>

Flags:
  -h, --help   help for check
`

	queryCmd.SetUsageTemplate(usage)
	rootCmd.AddCommand(updateCmd)
}

// executeUpdate will
func executeUpdate(cmd *cobra.Command, args []string) error {
	return nil
}
