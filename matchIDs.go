package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	//"reflect"
)

const April20 string = "1713618542"
const ArenaID string = "1700"
const n_matches string = "100"

func GetMatchesByPUUID(puuid string) []string {
	var matches []string
	var api_token string = os.Getenv("RIOT_TOKEN")
	link := fmt.Sprintf("https://europe.api.riotgames.com/lol/match/v5/matches/by-puuid/%s/ids?startTime=%s&queue=%s&start=0&count=%s&api_key=%s", puuid, April20, ArenaID, n_matches, api_token)
	resp, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		JSONBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(JSONBody, &matches)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println(resp.Status)
	}
	return matches
}

func WriteToFile(arrayToWrite []string, fileName string) {
	for i, j := 0, len(arrayToWrite)-1; i < j; i, j = i+1, j-1 {
		arrayToWrite[i], arrayToWrite[j] = arrayToWrite[j], arrayToWrite[i]
	}
	var writeToFile string = ""
	for i := 0; i < len(arrayToWrite); i++ {
		writeToFile += arrayToWrite[i] + "\n"
	}
	err := os.WriteFile(fileName, []byte(writeToFile), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
