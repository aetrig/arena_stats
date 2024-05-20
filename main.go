package main

import (
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
	aetrig.Print()
	sunny.Print()
	ragna.Print()
	var matches []string = GetMatchesByPUUID(aetrig.Puuid)
	fmt.Println(matches)
	WriteToFile(matches, "test.txt")
}
