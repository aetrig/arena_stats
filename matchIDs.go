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

type _Match struct {
	MatchID   string
	Champion  string
	Placement int
	//PlayedWithSunny bool
}

type _Metadata struct {
	Participants []string `json:"participants"`
}

type _Participant struct {
	ChampionName string `json:"championName"`
	Placement    int    `json:"placement"`
}

type _Info struct {
	Participants []_Participant `json:"participants"`
}

type _MatchJSON struct {
	Metadata _Metadata `json:"metadata"`
	Info     _Info     `json:"info"`
}

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

func WriteToFile(data []string, fileName string) {
	fileContentBytes, _ := os.ReadFile(fileName)
	content := strings.Fields(string(fileContentBytes))
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

func GetMatchByID(matchID string, acc RiotAcc) _Match {
	var match _Match
	var matchJSON _MatchJSON
	var api_token string = os.Getenv("RIOT_TOKEN")
	link := fmt.Sprintf(
		"https://europe.api.riotgames.com/lol/match/v5/matches/%s?api_key=%s",
		matchID,
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
		err = json.Unmarshal(JSONBody, &matchJSON)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println(resp.Status)
	}
	var accIndex int = -1
	for i := 0; i < len(matchJSON.Metadata.Participants); i++ {
		if acc.Puuid == matchJSON.Metadata.Participants[i] {
			accIndex = i
		}
	}
	match.MatchID = matchID
	match.Champion = matchJSON.Info.Participants[accIndex].ChampionName
	match.Placement = matchJSON.Info.Participants[accIndex].Placement
	return match
}
