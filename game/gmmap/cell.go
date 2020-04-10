package gmmap

import (
	"errors"
)

const (
	EmptyCell 	CellValue = 0
	CrossCell 	CellValue = 1
	CircleCell 	CellValue = 2
)

type CellValue int

func (cellValue CellValue) String () string {
	switch cellValue {
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

func (cellValue CellValue) IsEmpty () bool  {
	return cellValue == EmptyCell
}

func (cellValue *CellValue) SetValue (value CellValue) error {
	if !cellValue.IsEmpty() {
		return errors.New("cell is not empty")
	}

	if value != EmptyCell && value != CircleCell && value != CrossCell {
		return errors.New("unknown value")
	}

	*cellValue = value
	return nil
}