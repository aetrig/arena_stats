package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	var aetrig RiotAcc
	// var sunny RiotAcc
	// var ragna RiotAcc
	GetRiotAccByGameNameTagLine(&aetrig, "aetrig", "uwu", Eu_west)
	//GetRiotAccByGameNameTagLine(&sunny, "SunnyAsh", "AETI", Eu_west)
	//GetRiotAccByGameNameTagLine(&ragna, "abs woman", "1312", Eu_west)
	// aetrig.Print()
	// sunny.Print()
	// ragna.Print()
	var last100Matches []string = GetMatchesByRiotAcc(aetrig)
	WriteToMatchesFile(last100Matches)
	// start := time.Now()
	// match := GetMatchByID(last100Matches[2], aetrig)
	// end := time.Since(start)
	// fmt.Println(end)
	// fmt.Printf(
	// 	"Newest match:\nID: %s\nChampion: %s\nPlacement: %d\n",
	// 	match.MatchID,
	// 	match.Champion,
	// 	match.Placement,
	// )
	// WriteToStatsFile(match)
	StatsMatches(aetrig)
}
