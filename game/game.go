package game

const emptyCellValue  = 0

type Game struct {
	GameMap        [][]int
	UserTurn       bool
	GameFinished   bool
}

func New(mapSizeX int, mapSizeY int) Game {

	//TODO добавить проверку на всякое
	//Создаем пустой 2d срез
	gameMap := make([][]int,0);

	for i := 0; i < mapSizeX; i ++{
		//создаем пустой срез
		tmp := make([]int, mapSizeY)
		//допихиваем в 2d срез обычный срез
		gameMap = append(gameMap, tmp)
	}
	userGame := Game{GameMap : gameMap,UserTurn : true, GameFinished : false}
	return userGame;
}

func MakeTurn(game *Game,x int,y int) bool{
	if x > len(game.GameMap) || x < 0 || y > len(game.GameMap[0]) || y < 0 {
		return false
	}

	if game.GameMap[x][y] != emptyCellValue {
		return false
	}

	if game.UserTurn{
		game.GameMap[x][y] = 1;
	}else{
		game.GameMap[x][y] = 2;
	}
	game.UserTurn = !game.UserTurn
	return true
}

func IsGameEnded(game Game) bool{
	return !isNextTurnAvailable(game) || game.GameFinished
}

func IsWinningCombinationExist(game Game,x int ,y int)  bool{
	//TODO метод проверяет всю карту на предмет выигрышных комбинаций
	return true
}

func IsWinningCombinationExistForCell(game *Game,x int ,y int) bool{
	cellValue := game.GameMap[x][y]
	gameMapXLen := len(game.GameMap)
	gameMapYLen := len(game.GameMap[0])

	if x >= 0 && x < gameMapXLen && y >= 0 && y < gameMapYLen && cellValue != emptyCellValue {
		//TODO написать более изысканную проверку
		verticalCellValueStrike := getVerticallyValueStrikeForCell(*game, x, y)

		horizontalCellValueStrike := getHorizontalValueStrikeForCell(*game, x, y)

		backSlashValueStrike := getBackSlashValueStrikeForCell(*game, x, y)

		slashValueStrike := getSlashValueStrikeForCell(*game, x, y)

		cellValueStrike := getMaxInt(verticalCellValueStrike, horizontalCellValueStrike, backSlashValueStrike,slashValueStrike)

		if getGameCellsForWinCount(*game) <= cellValueStrike {
			game.GameFinished = true
			return true
		}
	}
	return false
}

func getVerticallyValueStrikeForCell(game Game,x int ,y int) int{
	cellValue := game.GameMap[x][y]
	cellValueStrike := 0

	holdValue := false
	for i := 0; i < len(game.GameMap); i++{
		if i == x {
			holdValue = true
		}
		if game.GameMap[i][y] == cellValue{
			cellValueStrike++
		}else{
			if !holdValue{
				//если мы еще не прошли через ячейку, то сбрасываем счетчик
				cellValueStrike = 0;
			}else {
				//если уже прошли, то выходим
				break;
			}
		}
	}
	return cellValueStrike
}

func getHorizontalValueStrikeForCell(game Game,x int ,y int) int{
	cellValue := game.GameMap[x][y]
	cellValueStrike := 0

	holdValue := false
	for i := 0; i < len(game.GameMap[0]); i++{
		if i == y {
			holdValue = true
		}
		if game.GameMap[x][i] == cellValue{
			cellValueStrike++
		}else{
			if !holdValue{
				//если мы еще не прошли через ячейку, то сбрасываем счетчик
				cellValueStrike = 0;
			}else {
				//если уже прошли, то выходим
				break;
			}
		}
	}
	return cellValueStrike
}


/*
	\
*/
func getBackSlashValueStrikeForCell(game Game,x int ,y int) int{
	cellValue := game.GameMap[x][y]
	cellValueStrike := 0

	holdValue := false

	var startCell [2]int;

	if x>=y {
		startCell[0] = x-y
		startCell[1] = 0
	}else{
		startCell[0] = 0
		startCell[1] = y-x
	}

	horizontalCellCoordinate := startCell[1]
	mapHorizontalLen := len(game.GameMap[0])
	for i := startCell[0]; i < len(game.GameMap); i++{
		if i == x && horizontalCellCoordinate == y {
			holdValue = true
		}
		if game.GameMap[i][horizontalCellCoordinate] == cellValue{
			cellValueStrike++
		}else{
			if !holdValue{
				//если мы еще не прошли через ячейку, то сбрасываем счетчик
				cellValueStrike = 0;
			}else {
				//если уже прошли, то выходим
				break;
			}
		}

		horizontalCellCoordinate++

		if horizontalCellCoordinate >= mapHorizontalLen {
			break
		}
	}
	return cellValueStrike
}


/*
	/
*/
func getSlashValueStrikeForCell(game Game,x int ,y int) int{
	cellValue := game.GameMap[x][y]
	cellValueStrike := 0

	holdValue := false

	var startCell [2]int;
	startCell[0] = 0
	startCell[1] = y+x
	mapHorizontalLen := len(game.GameMap[0])
	if startCell[1] >= mapHorizontalLen {
		startCell[0] = startCell[1] - mapHorizontalLen + 1
		startCell[1] = mapHorizontalLen - 1
	}

	horizontalCellCoordinate := startCell[1]


	for i := startCell[0]; i < len(game.GameMap); i++{
		if i == x && horizontalCellCoordinate == y {
			holdValue = true
		}
		if game.GameMap[i][horizontalCellCoordinate] == cellValue{
			cellValueStrike++
		}else{
			if !holdValue{
				//если мы еще не прошли через ячейку, то сбрасываем счетчик
				cellValueStrike = 0;
			}else {
				//если уже прошли, то выходим
				break;
			}
		}

		horizontalCellCoordinate--

		if horizontalCellCoordinate < 0 {
			break
		}
	}
	return cellValueStrike
}

func getGameCellsForWinCount(game Game) int{
	verticalLen := len(game.GameMap)
	horizontalLen := len(game.GameMap[0])

	if verticalLen <= horizontalLen{
		return verticalLen
	}else{
		return horizontalLen
	}
}

func isNextTurnAvailable(game Game) bool{
	for _, gameMapRow := range game.GameMap{
		for _, gameMapCell := range gameMapRow{
			if gameMapCell == emptyCellValue {
				return true
			}
		}
	}
	return false
}

func getMaxInt(nums ... int) int{
	//TODO add nil check
	maxNum := nums[0]
	for _, num := range nums {
		if maxNum < num {
			maxNum = num
		}
	}
	return maxNum
}

