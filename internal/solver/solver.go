package solver

import (
	"errors"
	"fmt"
	"io"
	"sort"
)

var ErrNegativeBallNumber = errors.New("negative ball number")

func readRow(row []int, r io.Reader) error {
	for i := 0; i < len(row); i++ {
		_, err := fmt.Fscan(r, &row[i])
		if err != nil {
			return fmt.Errorf("error reading input: %v", err)
		}
	}
	return nil
}

func Solve(n int, r io.Reader) (bool, error) {
	containers := make([]int, n)
	colors := make([]int, n)
	row := make([]int, n)

	// read input and count stats
	for i := 0; i < n; i++ {
		err := readRow(row, r)
		if err != nil {
			return false, err
		}

		for k, elem := range row {
			if elem < 0 {
				return false, ErrNegativeBallNumber
			}

			containers[i] += elem
			colors[k] += elem
		}
	}

	// sorting arrays
	sort.Ints(containers)
	sort.Ints(colors)

	// checking is solution exist
	for i := 0; i < n; i++ {
		if containers[i] != colors[i] {
			return false, nil
		}
	}
	return true, nil
}
