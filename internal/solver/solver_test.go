package solver_test

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"test_assigment/internal/solver"
)

func serialize(in [][]any) string {
	var sb strings.Builder

	for _, row := range in {
		for _, value := range row {
			sb.WriteString(fmt.Sprintf("%v ", value))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

type input struct {
	n     int
	array [][]any
}

type output struct {
	ok  bool
	err error
}

func TestSolve(t *testing.T) {
	for _, tt := range []struct {
		name   string
		input  input
		output output
	}{
		{
			name: "solution exist",
			input: input{
				n: 2,
				array: [][]any{
					{1, 2},
					{2, 1},
				},
			},
			output: output{
				ok:  true,
				err: nil,
			},
		},
		{
			name: "no solution",
			input: input{
				n: 3,
				array: [][]any{
					{10, 20, 30},
					{1, 1, 1},
					{0, 0, 1},
				},
			},
			output: output{
				ok:  false,
				err: nil,
			},
		},
		{
			name: "valid input",
			input: input{
				n: 3,
				array: [][]any{
					{1, 1, 1},
					{1, 1, 1},
					{1, 1, 1},
				},
			},
			output: output{
				ok:  true,
				err: nil,
			},
		},
		{
			name: "invalid input len",
			input: input{
				n: 3,
				array: [][]any{
					{1, 1, 1},
					{1, 1, 1},
					{1, 1},
				},
			},
			output: output{
				ok:  false,
				err: io.EOF,
			},
		},
		{
			name: "negative number of balls",
			input: input{
				n: 3,
				array: [][]any{
					{1, 1, 1},
					{1, 1, 1},
					{1, -1, 1},
				},
			},
			output: output{
				ok:  false,
				err: solver.ErrNegativeBallNumber,
			},
		},
		{
			name: "non-integer input",
			input: input{
				n: 3,
				array: [][]any{
					{1, 1, 1},
					{1, 1, 1},
					{1, "a", 1},
				},
			},
			output: output{
				ok:  false,
				err: fmt.Errorf("error reading input: invalid syntax"),
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(serialize(tt.input.array))

			ok, err := solver.Solve(tt.input.n, io.Reader(r))
			require.Equal(t, tt.output.ok, ok)

			if tt.output.err != nil {
				require.Error(t, tt.output.err, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
