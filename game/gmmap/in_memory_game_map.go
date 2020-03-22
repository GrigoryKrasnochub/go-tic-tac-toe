package gmmap

import "errors"

type InMemoryGameMap struct {
	gameMap [][]int
}

func (inMemoryGameMap *InMemoryGameMap) LoadGameMapFromSlice(gameMap [][]int) {
	inMemoryGameMap.gameMap = gameMap
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
	inMemoryGameMap.gameMap = gameMap;
	return nil
}

func (inMemoryGameMap *InMemoryGameMap) WriteValueToCell(value int, xCoordinate int, yCoordinate int) error {
	//TODO add check inArray()
	inMemoryGameMap.gameMap[xCoordinate][yCoordinate] = value
	return nil
}

func (inMemoryGameMap *InMemoryGameMap) GetMap() [][]int {
	return inMemoryGameMap.gameMap
}
