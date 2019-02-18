package output

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func Table(header []string) *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)

	return table
}
