package game

import (
	"github.com/GrigoryKrasnochub/go-tic-tac-toe/game/gmmap"
)

type Game struct {
	UserTurn     bool
	GameFinished bool
	gmmap.GameMapInterface
}

func New(mapSizeVertical int, mapSizeHorizontal int) (*Game, error) {
	userGame := &Game{UserTurn: true, GameFinished: false, GameMapInterface: &gmmap.InMemoryGameMap{}}
	newMapError := userGame.NewMap(mapSizeVertical, mapSizeHorizontal)
	if newMapError != nil {
		return nil, newMapError
	}

	return userGame, nil
}

func (game Game) IsGameEnded() bool {
	return !game.isNextTurnAvailable() || game.GameFinished
}

func (game *Game) MakeTurn(verticalValue int, horizontalValue int) error {
	var writeValueToCellError error
	if game.UserTurn {
		writeValueToCellError = game.WriteValueToCell(gmmap.CrossCell, verticalValue, horizontalValue)
	} else {
		writeValueToCellError = game.WriteValueToCell(gmmap.CircleCell, verticalValue, horizontalValue)
	}

	if writeValueToCellError != nil {
		return writeValueToCellError
	}

	game.UserTurn = !game.UserTurn
	return nil
}

func (game *Game) IsWinningCombinationExistForCell(verticalCoordinate int, horizontalCoordinate int) bool {
	cellValue, getCellValueError := game.GetCellValue(verticalCoordinate, horizontalCoordinate)

	if getCellValueError != nil {
		return false
	}

	if cellValue.IsEmpty() {
		return false
	}

	//TODO Do it in better way
	verticalCellValueStrike := game.getVerticallyValueStrikeForCell(verticalCoordinate, horizontalCoordinate)

	horizontalCellValueStrike := game.getHorizontalValueStrikeForCell(verticalCoordinate, horizontalCoordinate)

	backSlashValueStrike := game.getBackSlashValueStrikeForCell(verticalCoordinate, horizontalCoordinate)

	slashValueStrike := game.getSlashValueStrikeForCell(verticalCoordinate, horizontalCoordinate)

	cellValueStrike := getMaxInt(verticalCellValueStrike, horizontalCellValueStrike, backSlashValueStrike, slashValueStrike)

	if game.getGameCellsForWinCount() <= cellValueStrike {
		game.GameFinished = true
		return true
	}

	return false
}

func (game Game) getVerticallyValueStrikeForCell(verticalCoordinate int, horizontalCoordinate int) int {
	cellValue, _ := game.GetCellValue(verticalCoordinate, horizontalCoordinate)
	cellValueStrike := 0

	holdValue := false
	for i := 0; i < game.GetMapVerticalSize(); i++ {
		if i == verticalCoordinate {
			holdValue = true
		}
		compareResult, _ := game.CompareCellValueToGivenValue(i, horizontalCoordinate, cellValue)
		if compareResult {
			cellValueStrike++
		} else {
			if !holdValue {
				//если мы еще не прошли через ячейку, то сбрасываем счетчик
				cellValueStrike = 0
			} else {
				//если уже прошли, то выходим
				break
			}
		}
	}
	return cellValueStrike
}

func (game Game) getHorizontalValueStrikeForCell(verticalCoordinate int, horizontalCoordinate int) int {
	cellValue, _ := game.GetCellValue(verticalCoordinate, horizontalCoordinate)
	cellValueStrike := 0

	holdValue := false
	for i := 0; i < game.GetMapHorizontalSize(); i++ {
		if i == horizontalCoordinate {
			holdValue = true
		}
		compareResult, _ := game.CompareCellValueToGivenValue(verticalCoordinate, i, cellValue)
		if compareResult {
			cellValueStrike++
		} else {
			if !holdValue {
				//если мы еще не прошли через ячейку, то сбрасываем счетчик
				cellValueStrike = 0
			} else {
				//если уже прошли, то выходим
				break
			}
		}
	}
	return cellValueStrike
}

/*
	\
*/
func (game Game) getBackSlashValueStrikeForCell(verticalCoordinate int, horizontalCoordinate int) int {
	cellValue, _ := game.GetCellValue(verticalCoordinate, horizontalCoordinate)
	cellValueStrike := 0

	holdValue := false

	var startCell [2]int

	if verticalCoordinate >= horizontalCoordinate {
		startCell[0] = verticalCoordinate - horizontalCoordinate
		startCell[1] = 0
	} else {
		startCell[0] = 0
		startCell[1] = horizontalCoordinate - verticalCoordinate
	}

	horizontalCellCoordinate := startCell[1]
	mapHorizontalSize := game.GetMapHorizontalSize()
	for i := startCell[0]; i < game.GetMapVerticalSize(); i++ {
		if i == verticalCoordinate && horizontalCellCoordinate == horizontalCoordinate {
			holdValue = true
		}
		compareResult, _ := game.CompareCellValueToGivenValue(i, horizontalCellCoordinate, cellValue)
		if compareResult {
			cellValueStrike++
		} else {
			if !holdValue {
				//если мы еще не прошли через ячейку, то сбрасываем счетчик
				cellValueStrike = 0
			} else {
				//если уже прошли, то выходим
				break
			}
		}

		horizontalCellCoordinate++

		if horizontalCellCoordinate >= mapHorizontalSize {
			break
		}
	}
	return cellValueStrike
}

/*
	/
*/
func (game Game) getSlashValueStrikeForCell(verticalCoordinate int, horizontalCoordinate int) int {
	cellValue, _ := game.GetCellValue(verticalCoordinate, horizontalCoordinate)
	cellValueStrike := 0

	holdValue := false

	var startCell [2]int
	startCell[0] = 0
	startCell[1] = horizontalCoordinate + verticalCoordinate
	mapHorizontalSize := game.GetMapHorizontalSize()
	if startCell[1] >= mapHorizontalSize {
		startCell[0] = startCell[1] - mapHorizontalSize + 1
		startCell[1] = mapHorizontalSize - 1
	}

	horizontalCellCoordinate := startCell[1]

	for i := startCell[0]; i < game.GetMapVerticalSize(); i++ {
		if i == verticalCoordinate && horizontalCellCoordinate == horizontalCoordinate {
			holdValue = true
		}
		compareResult, _ := game.CompareCellValueToGivenValue(i, horizontalCellCoordinate, cellValue)
		if compareResult {
			cellValueStrike++
		} else {
			if !holdValue {
				//если мы еще не прошли через ячейку, то сбрасываем счетчик
				cellValueStrike = 0
			} else {
				//если уже прошли, то выходим
				break
			}
		}

		horizontalCellCoordinate--

		if horizontalCellCoordinate < 0 {
			break
		}
	}
	return cellValueStrike
}

func (game Game) getGameCellsForWinCount() int {
	verticalMapSize := game.GetMapVerticalSize()
	horizontalMapSize := game.GetMapHorizontalSize()

	if verticalMapSize <= horizontalMapSize {
		return verticalMapSize
	} else {
		return horizontalMapSize
	}
}

func (game Game) isNextTurnAvailable() bool {
	gameMap := game.GetMap()
	for _, gameMapRow := range gameMap {
		for _, gameMapCell := range gameMapRow {
			if gameMapCell == gmmap.EmptyCell {
				return true
			}
		}
	}
	return false
}

func getMaxInt(nums ...int) int {
	//TODO add nil check
	maxNum := nums[0]
	for _, num := range nums {
		if maxNum < num {
			maxNum = num
		}
	}
	return maxNum
}
