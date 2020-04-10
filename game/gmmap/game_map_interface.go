package gmmap

type GameMapInterface interface {
	WriteValueToCell(value int, xCoordinate int, yCoordinate int) error
	GetMap() [][]int
	NewMap(verticalSize int, horizontalSize int) error
	LoadGameMapFromSlice(gameMap [][]int)
	GetCellValue(verticalCoordinate int, horizontalCoordinate int) (int, error)
	CompareCellValueToGivenValue(verticalCoordinate int, horizontalCoordinate int, valueCompareTo int) (bool, error)
	GetMapSize() [2]int
	GetMapVerticalSize() int
	GetMapHorizontalSize() int
}
