package seven

import (
	"aoc24/internal/common"
	"bufio"
	"fmt"
	"io"
	"os"
)

type calibration struct {
	result   int
	equation []int
}

func parseCalibration(r *bufio.Reader) (*calibration, error) {
	result, c, err := common.ReadInt(r)
	if err != nil {
		return nil, err
	}
	if c != ':' {
		return nil, fmt.Errorf("expected ':', got %c", c)
	}

	c, _, err = r.ReadRune()
	if err != nil {
		return nil, err
	}
	if c != ' ' {
		return nil, fmt.Errorf("expected ' ', got %c", c)
	}

	equation := []int{}
	for {
		n, c, err := common.ReadInt(r)
		equation = append(equation, n)
		if c == '\n' || err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if c != ' ' {
			return nil, fmt.Errorf("expected ' ', got %c", c)
		}
	}

	return &calibration{result, equation}, nil
}

func parseCalibrations(r *bufio.Reader) ([]*calibration, error) {
	var calibrations []*calibration
	for {
		calibration, err := parseCalibration(r)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		calibrations = append(calibrations, calibration)
	}
	return calibrations, nil
}

func readInput() ([]*calibration, error) {
	f, err := os.Open("input/7")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return parseCalibrations(bufio.NewReader(f))
}
