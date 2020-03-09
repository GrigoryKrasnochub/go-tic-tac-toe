package console

import (
	"awesomeProject/game"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func PrintMap(userGame game.Game){

	verticallyLen := len(userGame.GameMap)
	horizontallyLen := len(userGame.GameMap[0])
	var horizontallyCharsIndentCount int = (horizontallyLen / 10) + 2
	var verticallyCharsIndentCount int = (verticallyLen / 10) + 1

	//print first line
	fmt.Printf("%*s",verticallyCharsIndentCount,"");
	for y := 0; y < horizontallyLen; y++{
		fmt.Printf("%*d",horizontallyCharsIndentCount,y)
	}
	fmt.Print("\n");
	//end print first line

	for i, gameMapRow := range userGame.GameMap{
		fmt.Printf("%*d",verticallyCharsIndentCount,i)
		for _, gameMapCell := range gameMapRow{
			switch gameMapCell {
			case 0:
				fmt.Printf("%*s",horizontallyCharsIndentCount,"_")
			case 1:
				fmt.Printf("%*s",horizontallyCharsIndentCount,"X")
			case 2:
				fmt.Printf("%*s",horizontallyCharsIndentCount,"O")
			}
		}
		fmt.Print("\n")
	}
}

func DoConsoleGame(){
	regex, _ := regexp.Compile(`(\d+)\s+(\d+)`)
	reader := bufio.NewReader(os.Stdin)
	for {
		var xMapSize int;
		var yMapSize int;
		for {
			fmt.Println("Please write game map size in format \"xMapSize yMapSize\" (without semicolons)")
			text, _ := reader.ReadString('\n')
			userMapSize := regex.FindStringSubmatch(text)
			if userMapSize != nil {
				if userMapSize[1] != "" && userMapSize[2] != "" {
					xMapSize, _ = strconv.Atoi(userMapSize[1])
					yMapSize, _ = strconv.Atoi(userMapSize[2])
					break
				}
			}
			fmt.Println("Ooops! Something went wrong")
		}
		userGame := game.New(xMapSize, yMapSize)
		if consoleGame(userGame) {
			fmt.Println("Restart game? y/n")
			text, _ := reader.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)
			if strings.Compare("y", text) == 0 {
				fmt.Println("Game RESTARTED!")
			}else{
				break
			}
		}
	}
}

func consoleGame(userGame game.Game) bool{
	reader := bufio.NewReader(os.Stdin)
	regex, _ := regexp.Compile(`(\d+)\s+(\d+)`)
	players := [2]string {"Cross","Circle"}
	for !game.IsGameEnded(userGame) {
		PrintMap(userGame)
		for {
			fmt.Println("Make your turn, write coordinates in format \"xCoordinate yCoordinate\" (without semicolons)")
			text, _ := reader.ReadString('\n')
			userCoordinates := regex.FindStringSubmatch(text)
			if userCoordinates != nil{
				if userCoordinates[1] != "" && userCoordinates[2] != "" {
					xCoordinate, _ := strconv.Atoi(userCoordinates[1])
					yCoordinate, _ := strconv.Atoi(userCoordinates[2])
					if game.MakeTurn(&userGame,xCoordinate, yCoordinate){
						if game.IsWinningCombinationExistForCell(&userGame, xCoordinate, yCoordinate) {
							var player string
							//game turn was reverted
							if userGame.UserTurn {
								player = players[1]
							}else {
								player = players[0]
							}
							fmt.Printf("%s win\n", player)
						}
						break;
					}
				}
			}
			fmt.Println("Ooops! Something went wrong")
		}
	}
	fmt.Println("Game Over")
	return true
}
