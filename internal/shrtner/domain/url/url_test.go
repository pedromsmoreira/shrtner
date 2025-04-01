package url

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecimalToBase62(t *testing.T) {
	type test struct {
		decimalValue int64
		shorUrl      string
	}

	tests := []test{
		{
			decimalValue: 1,
			shorUrl:      "1",
		},
		{
			decimalValue: 100,
			shorUrl:      "1C",
		},
		{
			decimalValue: 1000,
			shorUrl:      "g8",
		},
		{
			decimalValue: 123456789089898,
			shorUrl:      "z3wBXxG2",
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("when is %d should return %s", tc.decimalValue, tc.shorUrl), func(t *testing.T) {
			s := decimalToBase62(tc.decimalValue)
			require.Equal(t, tc.shorUrl, s)
		})
	}
}

func TestShorten(t *testing.T) {
	type test struct {
		url     string
		shorUrl string
	}

	tests := []test{
		{
			url:     "https://www.google.com",
			shorUrl: "d2kID5O",
		},
		{
			url:     "https://www.youtube.com",
			shorUrl: "gqgnHTy",
		},
		{
			url:     "https://www.gmail.com",
			shorUrl: "g7DgV6c",
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("when is %s should return %s", tc.url, tc.shorUrl), func(t *testing.T) {
			result := shorten(tc.url)
			require.Equal(t, tc.shorUrl, result)
		})
	}
}
