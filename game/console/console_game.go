package console

import (
	"bufio"
	"fmt"
	"github.com/GrigoryKrasnochub/go-tic-tac-toe/game"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func PrintMap(userGame game.Game) {
	gameMap := userGame.GetMap()
	verticallyLen := len(gameMap)
	horizontallyLen := len(gameMap[0])
	var horizontallyCharsIndentCount int = (horizontallyLen / 10) + 2
	var verticallyCharsIndentCount int = (verticallyLen / 10) + 1

	//print first line
	fmt.Printf("%*s", verticallyCharsIndentCount, "")
	for y := 0; y < horizontallyLen; y++ {
		fmt.Printf("%*d", horizontallyCharsIndentCount, y)
	}
	fmt.Print("\n")
	//end print first line

	for i, gameMapRow := range gameMap {
		fmt.Printf("%*d", verticallyCharsIndentCount, i)
		for _, gameMapCell := range gameMapRow {
			switch gameMapCell {
			case 0:
				fmt.Printf("%*s", horizontallyCharsIndentCount, "_")
			case 1:
				fmt.Printf("%*s", horizontallyCharsIndentCount, "X")
			case 2:
				fmt.Printf("%*s", horizontallyCharsIndentCount, "O")
			}
		}
		fmt.Print("\n")
	}
}

func DoConsoleGame() {
	regex, _ := regexp.Compile(`(\d+)\s+(\d+)`)
	reader := bufio.NewReader(os.Stdin)
	errorMessage := ""
	var userGame *game.Game
	var xMapSize int
	var yMapSize int
	for {
		if errorMessage != "" {
			fmt.Println("Ooops! Something went wrong. Error text: ", errorMessage)
			errorMessage = ""
		}

		fmt.Println("Please write game map size in format \"VerticalMapSize HorizontalMapSize\" (without semicolons)")
		text, _ := reader.ReadString('\n')
		userMapSize := regex.FindStringSubmatch(text)
		if userMapSize == nil || userMapSize[1] == "" || userMapSize[2] == "" {
			errorMessage = "incorrect format"
			continue
		}

		xMapSize, _ = strconv.Atoi(userMapSize[1])
		yMapSize, _ = strconv.Atoi(userMapSize[2])
		createdGame, createdGameError := game.New(xMapSize, yMapSize)

		if createdGameError != nil {
			errorMessage = createdGameError.Error()
			continue
		}

		userGame = createdGame
		if consoleGame(*userGame) {
			fmt.Println("Restart game? y/n")
			text, _ := reader.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)
			if strings.Compare("y", text) == 0 {
				fmt.Println("Game RESTARTED!")
			} else {
				break
			}
		}
	}
}

func consoleGame(userGame game.Game) bool {
	reader := bufio.NewReader(os.Stdin)
	regex, _ := regexp.Compile(`(\d+)\s+(\d+)`)
	players := [2]string{"Cross", "Circle"}
	errorMessage := ""
	for !userGame.IsGameEnded() {

		if errorMessage != "" {
			fmt.Println("Ooops! Something went wrong. Error text: ", errorMessage)
			errorMessage = ""
		} else {
			PrintMap(userGame)
		}

		fmt.Println("Make your turn, write coordinates in format \"VerticalCoordinate HorizontalCoordinate\" (without semicolons)")
		text, _ := reader.ReadString('\n')
		userCoordinates := regex.FindStringSubmatch(text)
		if userCoordinates == nil || userCoordinates[1] == "" || userCoordinates[2] == "" {
			errorMessage = "incorrect format"
			continue
		}

		xCoordinate, _ := strconv.Atoi(userCoordinates[1])
		yCoordinate, _ := strconv.Atoi(userCoordinates[2])
		_, gameTurnError := userGame.MakeTurn(xCoordinate, yCoordinate)
		if gameTurnError != nil {
			errorMessage = gameTurnError.Error()
			continue
		}

		if userGame.IsWinningCombinationExistForCell(xCoordinate, yCoordinate) {
			var player string
			//game turn was reverted
			if userGame.UserTurn {
				player = players[1]
			} else {
				player = players[0]
			}
			PrintMap(userGame)
			fmt.Printf("%s win\n", player)
		}
	}
	fmt.Println("Game Over")
	return true
}
