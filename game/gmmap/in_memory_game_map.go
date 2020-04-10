package gmmap

import "errors"

type InMemoryGameMap struct {
	gameMap           [][]int
	verticalMapSize   int
	horizontalMapSize int
}

func (inMemoryGameMap *InMemoryGameMap) GetMapSize() [2]int {
	result := [2]int{inMemoryGameMap.verticalMapSize, inMemoryGameMap.horizontalMapSize}
	return result
}

func (inMemoryGameMap *InMemoryGameMap) GetMapHorizontalSize() int {
	return inMemoryGameMap.horizontalMapSize
}

func (inMemoryGameMap *InMemoryGameMap) GetMapVerticalSize() int {
	return inMemoryGameMap.verticalMapSize
}

func (inMemoryGameMap *InMemoryGameMap) LoadGameMapFromSlice(gameMap [][]int) {
	inMemoryGameMap.gameMap = gameMap
	inMemoryGameMap.horizontalMapSize = len(gameMap[0])
	inMemoryGameMap.verticalMapSize = len(gameMap)
}

func (inMemoryGameMap *InMemoryGameMap) NewMap(verticalSize int, horizontalSize int) error {
	if verticalSize <= 0 || horizontalSize <= 0 {
		return errors.New("size matters, the sizeNumbers should be bigger")
	}
	gameMap := make([][]int, 0)

	for i := 0; i < verticalSize; i++ {
		//создаем пустой срез
		tmp := make([]int, horizontalSize)
		//допихиваем в 2d срез обычный срез
		gameMap = append(gameMap, tmp)
	}
	inMemoryGameMap.gameMap = gameMap
	inMemoryGameMap.verticalMapSize = verticalSize
	inMemoryGameMap.horizontalMapSize = horizontalSize
	return nil
}

func (inMemoryGameMap *InMemoryGameMap) WriteValueToCell(value int, verticalCoordinate int, horizontalCoordinate int) error {
	//TODO add check inArray()
	inMemoryGameMap.gameMap[verticalCoordinate][horizontalCoordinate] = value
	return nil
}

func (inMemoryGameMap *InMemoryGameMap) GetMap() [][]int {
	return inMemoryGameMap.gameMap
}

func (inMemoryGameMap *InMemoryGameMap) GetCellValue(verticalCoordinate int, horizontalCoordinate int) (int, error) {
	if !inMemoryGameMap.checkIsCellAvailable(verticalCoordinate, horizontalCoordinate) {
		return 0, errors.New("cell is unreachable")
	}
	return inMemoryGameMap.gameMap[verticalCoordinate][horizontalCoordinate], nil
}

func (inMemoryGameMap *InMemoryGameMap) CompareCellValueToGivenValue(verticalCoordinate int, horizontalCoordinate int, valueCompareTo int) (bool, error) {
	cellValue, getCellValueError := inMemoryGameMap.GetCellValue(verticalCoordinate, horizontalCoordinate)
	if getCellValueError != nil {
		return false, getCellValueError
	}

	if cellValue == valueCompareTo {
		return true, nil
	}

	return false, nil
}

func (inMemoryGameMap *InMemoryGameMap) checkIsCellAvailable(verticalCoordinate int, horizontalCoordinate int) bool {
	if verticalCoordinate >= 0 && horizontalCoordinate >= 0 && verticalCoordinate < inMemoryGameMap.verticalMapSize && horizontalCoordinate < inMemoryGameMap.horizontalMapSize {
		return true
	}
	return false
}
