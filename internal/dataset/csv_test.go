package dataset_test

import (
	"slices"
	"strings"
	"testing"

	"github.com/konradszl/csviq/internal/dataset"
)

func TestFromCSV(t *testing.T) {
	t.Run("reads columns and rows", func(t *testing.T) {
		in := strings.NewReader("name,age\nkonrad,22\nklaudia,23\n")

		got, err := dataset.FromCSV(in)
		if err != nil {
			t.Fatalf("FromCSV() returned unexpected error: %v", err)
		}

		want := []string{"name", "age"}
		if !slices.Equal(got.Columns, want) {
			t.Errorf("Columns = %v, want %v", got.Columns, want)
		}

		if got.RowCount() != 2 {
			t.Errorf("RowCount() = %d, want 2", got.RowCount())
		}
	})

	t.Run("returns error for empty input", func(t *testing.T) {
		in := strings.NewReader("")

		_, err := dataset.FromCSV(in)
		if err == nil {
			t.Fatal("FromCSV() expected an error for empty input, got nil")
		}
	})
}
