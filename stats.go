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
	"strconv"
	"strings"
)

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

func getMatchByID(matchID string, acc RiotAcc) _Match {
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

func writeToMatchHistoryFile(match _Match, acc RiotAcc) {
	var fileName string = "matchHistory_" + acc.GameName + ".txt"
	fileContentBytes, _ := os.ReadFile(fileName)
	fileContent := strings.Split(string(fileContentBytes), "\n")
	matchToAppend := strings.Join([]string{match.MatchID, match.Champion, strconv.Itoa(match.Placement)}, " ")
	fileContent = append(fileContent, matchToAppend)
	sort.Strings(fileContent)
	slices.Reverse(fileContent)
	contentString := strings.Join(fileContent, "\n")
	_ = os.WriteFile(fileName, []byte(contentString), 0666)
}

func CreateMatchHistoryFile(acc RiotAcc) {
	var fileNameIDs string = "matchIDs_" + acc.GameName + ".txt"
	var fileNameHistory string = "matchHistory_" + acc.GameName + ".txt"
	fileContentBytes, _ := os.ReadFile(fileNameIDs)
	matchIDs := strings.Split(string(fileContentBytes), "\n")
	fileContentBytes, _ = os.ReadFile(fileNameHistory)
	statsFileContent := string(fileContentBytes)
	for _, matchID := range matchIDs {
		if !strings.Contains(statsFileContent, matchID) {
			writeToMatchHistoryFile(getMatchByID(matchID, acc), acc)
		}
	}
}
