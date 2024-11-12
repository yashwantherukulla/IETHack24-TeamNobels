package helpers

import (
    "os"
    "github.com/olekukonko/tablewriter"
)

// InitTable initializes a tablewriter table with given headers and a customizable style.
func InitTable(headers []string, alignment int, border bool) *tablewriter.Table {
    table := tablewriter.NewWriter(os.Stdout)
    table.SetHeader(headers)
    table.SetAlignment(alignment) // Center alignment: tablewriter.ALIGN_CENTER
    table.SetBorder(border)       // Border style
    table.SetRowSeparator("-")    // Set row separator if needed
    table.SetColumnSeparator("|") // Set column separator if needed
    table.SetHeaderAlignment(tablewriter.ALIGN_CENTER)
    return table
}

func DisplayMetricsTable(metrics map[string]string) {
    headers := []string{"Metric", "Score"}
    table := InitTable(headers, tablewriter.ALIGN_LEFT, true)

    for metric, score := range metrics {
        table.Append([]string{metric, score})
    }

    table.Render()
}