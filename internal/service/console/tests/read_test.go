package tests

import (
	"bufio"
	"bytes"
	"client/internal/console"
	cs "client/internal/service/console"
	"errors"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestReadLine(t *testing.T) {
	t.Helper()
	testCases := []struct {
		testName string
		str      string
		ans      string
		error
	}{
		{"success",
			"haha\n",
			"haha",
			nil,
		},

		{"readerError",
			"haha",
			"",
			errors.New("EOF"),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			strReader := strings.NewReader(tc.str)
			reader := bufio.NewReader(strReader)
			cons := console.NewConsole(reader, nil)
			cService := cs.New(cons)
			line, err := cService.ReadLine()
			require.Equal(t, tc.ans, line)
			require.Equal(t, tc.error, err)
		})
	}
}

func TestReadData(t *testing.T) {
	t.Helper()
	testCases := []struct {
		testName string
		str      string
		ans      string
		error
	}{
		{"success",
			"haha\n",
			"haha",
			nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			strReader := strings.NewReader(tc.str)
			reader := bufio.NewReader(strReader)
			writer := new(bytes.Buffer)
			write := bufio.NewWriter(writer)
			cons := console.NewConsole(reader, write)
			cService := cs.New(cons)
			line, err := cService.ReadData("test")
			require.Equal(t, tc.ans, line)
			require.Equal(t, tc.error, err)
		})
	}
}
