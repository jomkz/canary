package cmd

import (
	"fmt"

	"github.com/jomkz/canary/engine"
	"github.com/spf13/cobra"
)

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query the public and FQDN IP addresses",
	Long:  "This command will query the public IP address for this host as well as the IP address for the provided FQDN.",
	Args:  cobra.ExactArgs(1),
	RunE:  executeQuery,
}

func init() {
	usage := `Usage:
  canary query <FQDN>

Flags:
  -h, --help   help for check
`

	queryCmd.SetUsageTemplate(usage)
	rootCmd.AddCommand(queryCmd)
}

// executeQuery will run query the public facing IP address for the host and
// the IP address for the given FQDN.
func executeQuery(cmd *cobra.Command, args []string) error {
	pip, err := engine.PublicAddress()
	if err != nil {
		return err
	}

	fip, err := engine.FQDNAddress(args[0])
	if err != nil {
		return err
	}

	fmt.Println("Public IP:", pip)
	fmt.Println("  FQDN IP:", fip)
	return nil
}
