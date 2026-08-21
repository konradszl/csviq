package dataset

// Dataset represents a table with column names and rows. Each row is expected
// to hold one value per column, in the same order as Columns.
type Dataset struct {
	Columns []string
	Rows    [][]string
}

// RowCount returns the number of data rows in the Dataset, not counting the
// column header.
func (d *Dataset) RowCount() int {
	return len(d.Rows)
}
