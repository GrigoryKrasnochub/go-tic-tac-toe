package gmmap

import (
	"errors"
	"fmt"
)

type InMemoryGameMap struct {
	gameMap           [][]CellValue
	verticalMapSize   int
	horizontalMapSize int
}

func (gm *InMemoryGameMap) GetMapSize() (int, int) {
	return gm.verticalMapSize, gm.horizontalMapSize
}

func (gm *InMemoryGameMap) GetMapHorizontalSize() int {
	return gm.horizontalMapSize
}

func (gm *InMemoryGameMap) GetMapVerticalSize() int {
	return gm.verticalMapSize
}

func (gm *InMemoryGameMap) LoadGameMapFromSlice(gameMap [][]CellValue) {
	gm.gameMap = gameMap
	gm.horizontalMapSize = len(gameMap[0])
	gm.verticalMapSize = len(gameMap)
}

func (gm *InMemoryGameMap) NewMap(verticalSize int, horizontalSize int) error {
	if verticalSize <= 0 || horizontalSize <= 0 {
		return errors.New("size matters, the sizeNumbers should be bigger")
	}
	gameMap := make([][]CellValue, 0)

	for i := 0; i < verticalSize; i++ {
		//создаем пустой срез
		tmp := make([]CellValue, horizontalSize)
		//допихиваем в 2d срез обычный срез
		gameMap = append(gameMap, tmp)
	}
	gm.gameMap = gameMap
	gm.verticalMapSize = verticalSize
	gm.horizontalMapSize = horizontalSize
	return nil
}

func (gm *InMemoryGameMap) WriteValueToCell(value CellValue, verticalCoordinate int, horizontalCoordinate int) error {

	cellVal, err := gm.GetCellValue(verticalCoordinate, horizontalCoordinate)
	if err != nil {
		return errors.New(fmt.Sprintf("cell doesn't exist! vertical coordinate: %d, horizontal coordinate: %d", verticalCoordinate, horizontalCoordinate))
	}

	if err := cellVal.CheckCellAvailableForWritingValue(value); err != nil {
		return err
	}

	gm.gameMap[verticalCoordinate][horizontalCoordinate] = value

	return nil
}

func (gm *InMemoryGameMap) GetMap() [][]CellValue {
	return gm.gameMap
}

func (gm *InMemoryGameMap) GetCellValue(verticalCoordinate int, horizontalCoordinate int) (CellValue, error) {
	if !gm.checkIsCellAvailable(verticalCoordinate, horizontalCoordinate) {
		return 0, errors.New(fmt.Sprintf("cell doesn't exist. vertical coordinate: %d, horizontal coordinate: %d", verticalCoordinate, horizontalCoordinate))
	}
	return gm.gameMap[verticalCoordinate][horizontalCoordinate], nil
}

func (gm *InMemoryGameMap) CompareCellValueToGivenValue(verticalCoordinate int, horizontalCoordinate int, valueCompareTo CellValue) (bool, error) {
	cellValue, getCellValueError := gm.GetCellValue(verticalCoordinate, horizontalCoordinate)
	if getCellValueError != nil {
		return false, getCellValueError
	}

	if cellValue == valueCompareTo {
		return true, nil
	}

	return false, nil
}

func (gm *InMemoryGameMap) checkIsCellAvailable(verticalCoordinate int, horizontalCoordinate int) bool {
	return verticalCoordinate >= 0 && horizontalCoordinate >= 0 && verticalCoordinate < gm.verticalMapSize && horizontalCoordinate < gm.horizontalMapSize
}
