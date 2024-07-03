package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"slices"
	"sort"
	"strings"
	//"reflect"
)

const April20 string = "1713618542"
const ArenaID string = "1700"
const n_matches string = "100"

func GetMatchesByRiotAcc(riotAcc RiotAcc) []string {
	puuid := riotAcc.Puuid
	var matches []string
	var api_token string = os.Getenv("RIOT_TOKEN")
	link := fmt.Sprintf(
		"https://europe.api.riotgames.com/lol/match/v5/matches/by-puuid/%s/ids?startTime=%s&queue=%s&start=0&count=%s&api_key=%s",
		puuid,
		April20,
		ArenaID,
		n_matches,
		api_token,
	)
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

func WriteToMatchesFile(data []string, acc RiotAcc) {
	fileName := "matchIDs_" + acc.GameName + ".txt"
	fileContentBytes, _ := os.ReadFile(fileName)
	content := strings.Split(string(fileContentBytes), "\n")
	//contentString := string(fileContentBytes)
	//fmt.Print(content)
	//fmt.Print(contentString)

	for _, matchID := range data {
		if !slices.Contains(content, matchID) {
			content = append(content, matchID)
		}
	}

	sort.Strings(content)
	slices.Reverse(content)
	contentString := strings.Join(content, "\n")
	_ = os.WriteFile(fileName, []byte(contentString), 0666)
}
