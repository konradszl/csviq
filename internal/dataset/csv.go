package dataset

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
)

// FromCSV reads CSV records from r and returns them as a Dataset. The first
// record becomes the Dataset's columns, and each record after it becomes a row.
//
// It returns an error if r cannot be parsed as CSV, or if r holds no records
// at all.
func FromCSV(r io.Reader) (*Dataset, error) {
	csvReader := csv.NewReader(r)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("parsing csv: %w", err)
	}

	if len(records) == 0 {
		return nil, errors.New("csv has no records")
	}

	return &Dataset{
		Columns: records[0],
		Rows:    records[1:],
	}, nil
}
