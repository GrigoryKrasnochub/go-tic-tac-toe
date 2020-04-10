package gmmap

type GameMapInterface interface {
	WriteValueToCell(value CellValue, verticalCoordinate int, horizontalCoordinate int) error
	GetMap() [][]CellValue
	NewMap(verticalSize int, horizontalSize int) error
	LoadGameMapFromSlice(gameMap [][]CellValue)
	GetCellValue(verticalCoordinate int, horizontalCoordinate int) (CellValue, error)
	GetCellValuePointer(verticalCoordinate int, horizontalCoordinate int) (*CellValue, error)
	CompareCellValueToGivenValue(verticalCoordinate int, horizontalCoordinate int, valueCompareTo CellValue) (bool, error)
	GetMapSize() [2]int
	GetMapVerticalSize() int
	GetMapHorizontalSize() int
}
