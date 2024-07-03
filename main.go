package main

import (
	"log"

	"github.com/joho/godotenv"
)

func prepareMatchHistory(name string, tag string, server Server) {
	var acc RiotAcc
	GetRiotAccByGameNameTagLine(&acc, name, tag, server)
	WriteToMatchesFile(GetMatchesByRiotAcc(acc), acc)
	CreateMatchHistoryFile(acc)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	prepareMatchHistory("aetrig", "uwu", Eu_west)
}
