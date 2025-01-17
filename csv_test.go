package csvalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadCSV(t *testing.T) {
	expected := [][]string{
		{"id", "comment"},
		{"1", "comment"},
		{"2", "comment"},
		{"3", "comment"},
		{"4", "comment"},
		{"5", "comment"},
	}

	cases := []struct {
		filePath string
		expected [][]string
	}{
		{
			filePath: "./samples/file.csv",
			expected: expected,
		},
		{
			filePath: "./samples/fileWithBOM.csv",
			expected: expected,
		},
	}

	for _, tc := range cases {
		records, err := readCSV(tc.filePath)
		assert.NoError(t, err)
		assert.Equal(t, tc.expected, records)
	}
}
