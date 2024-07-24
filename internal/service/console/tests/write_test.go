package tests

import (
	"bufio"
	"bytes"
	"client/internal/console"
	cs "client/internal/service/console"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestWrite(t *testing.T) {
	t.Helper()
	testCases := []struct {
		testName string
		str      string
		ans      string
		error
	}{
		{"success",
			"haha",
			"haha",
			nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			writer := new(bytes.Buffer)
			write := bufio.NewWriter(writer)
			cons := console.NewConsole(nil, write)
			cService := cs.New(cons)
			cService.Write(tc.str)
			require.Equal(t, tc.ans, writer.String())
		})
	}
}

func TestWriteLine(t *testing.T) {
	t.Helper()
	testCases := []struct {
		testName string
		str      string
		ans      string
		error
	}{
		{"success",
			"haha",
			"haha\n",
			nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			writer := new(bytes.Buffer)
			write := bufio.NewWriter(writer)
			cons := console.NewConsole(nil, write)
			cService := cs.New(cons)
			cService.WriteLine(tc.str)
			require.Equal(t, tc.ans, writer.String())
		})
	}
}
