package gmmap

import (
	"errors"
)

const (
	EmptyCell  CellValue = 0
	CrossCell  CellValue = 1
	CircleCell CellValue = 2
)

type CellValue int

func (val CellValue) String() string {
	switch val {
	case EmptyCell:
		return "_"
	case CrossCell:
		return "X"
	case CircleCell:
		return "O"
	default:
		return "E"
	}
}

func (val CellValue) IsEmpty() bool {
	return val == EmptyCell
}

func (val CellValue) IsValueAllowed(value CellValue) bool {
	return value == EmptyCell || value == CircleCell || value == CrossCell
}

func (val CellValue) CheckCellAvailableForWritingValue(applicantForWriting CellValue) error {
	if !val.IsEmpty() {
		return errors.New("cell is not empty")
	}

	if !val.IsValueAllowed(applicantForWriting) {
		return errors.New("unknown value")
	}

	return nil
}
