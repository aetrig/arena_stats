package main

import (
	//"fmt"

	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	var aetrig RiotAcc
	var sunny RiotAcc
	var ragna RiotAcc
	GetRiotAccByGameNameTagLine(&aetrig, "aetrig", "uwu", Eu_west)
	GetRiotAccByGameNameTagLine(&sunny, "SunnyAsh", "AETI", Eu_west)
	GetRiotAccByGameNameTagLine(&ragna, "abs woman", "1312", Eu_west)
	// aetrig.Print()
	// sunny.Print()
	// ragna.Print()
	var matches []string = GetMatchesByRiotAcc(aetrig)
	WriteToFile(matches, "matches.txt")
	match := GetMatchByID(matches[0], aetrig)
	fmt.Printf("ID: %s\nChampion: %s\nPlacement: %d\n", match.MatchID, match.Champion, match.Placement)
}
