package gmmap

type GameMapInterface interface {
	WriteValueToCell(value int, xCoordinate int, yCoordinate int) error
	GetMap() [][]int
	NewMap(verticalSize int, horizontalSize int) error
	LoadGameMapFromSlice(gameMap [][]int)
}
