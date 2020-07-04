package console

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/GrigoryKrasnochub/go-tic-tac-toe/game"
)

var userInputRegexp = regexp.MustCompile(`(\d+)\s+(\d+)`)

func PrintMap(userGame game.Game) {
	gameMap := userGame.GetMap()
	verticalSize := userGame.GetMapVerticalSize()
	horizontalSize := userGame.GetMapHorizontalSize()
	var horizontallyCharsIndentCount int = (horizontalSize / 10) + 2
	var verticallyCharsIndentCount int = (verticalSize / 10) + 1

	//print first line
	fmt.Printf("%*s", verticallyCharsIndentCount, "")
	for y := 0; y < horizontalSize; y++ {
		fmt.Printf("%*d", horizontallyCharsIndentCount, y)
	}
	fmt.Print("\n")
	//end print first line

	for i, gameMapRow := range gameMap {
		fmt.Printf("%*d", verticallyCharsIndentCount, i)
		for _, gameMapCell := range gameMapRow {
			fmt.Printf("%*s", horizontallyCharsIndentCount, gameMapCell.String())
		}
		fmt.Print("\n")
	}
}

func DoConsoleGame() {
	reader := bufio.NewReader(os.Stdin)
	errorMessage := ""
	var userGame *game.Game
	var verticalMapSize int
	var horizontalMapSize int
	for {
		if errorMessage != "" {
			fmt.Println("Ooops! Something went wrong. Error text: ", errorMessage)
			errorMessage = ""
		}

		fmt.Println("Please write game map size in format \"VerticalMapSize HorizontalMapSize\" (without semicolons)")
		fmt.Print(">")
		text, _ := reader.ReadString('\n')
		userMapSize := userInputRegexp.FindStringSubmatch(text)
		if userMapSize == nil || userMapSize[1] == "" || userMapSize[2] == "" {
			errorMessage = "incorrect format"
			continue
		}

		verticalMapSize, _ = strconv.Atoi(userMapSize[1])
		horizontalMapSize, _ = strconv.Atoi(userMapSize[2])
		createdGame, createdGameError := game.New(verticalMapSize, horizontalMapSize)

		if createdGameError != nil {
			errorMessage = createdGameError.Error()
			continue
		}

		userGame = createdGame
		if consoleGame(*userGame) {
			fmt.Println("Restart game? y/n")
			fmt.Print(">")
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
		fmt.Print(">")
		text, _ := reader.ReadString('\n')
		userCoordinates := userInputRegexp.FindStringSubmatch(text)
		if userCoordinates == nil || userCoordinates[1] == "" || userCoordinates[2] == "" {
			errorMessage = "incorrect format"
			continue
		}

		verticalCoordinate, _ := strconv.Atoi(userCoordinates[1])
		horizontalCoordinate, _ := strconv.Atoi(userCoordinates[2])
		gameTurnError := userGame.MakeTurn(verticalCoordinate, horizontalCoordinate)
		if gameTurnError != nil {
			errorMessage = gameTurnError.Error()
			continue
		}

		if userGame.IsWinningCombinationExistForCell(verticalCoordinate, horizontalCoordinate) {
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
