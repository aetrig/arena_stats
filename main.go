package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	//"reflect"
)

var api_token string

type RiotAcc struct {
	GameName  string
	TagLine   string
	Puuid     string
	Id        string
	AccountId string
}

// Doesn't modify RiotAcc if Status != 200 OK
func getRiotAccByGameNameTagLine(riotAcc *RiotAcc, gameName string, tagLine string) {
	//var Acc RiotAcc
	link := fmt.Sprintf("https://europe.api.riotgames.com/riot/account/v1/accounts/by-riot-id/%s/%s?api_key=%s", gameName, tagLine, api_token)
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
		err = json.Unmarshal(JSONBody, &riotAcc)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(resp.Status)
	link = fmt.Sprintf("https://euw1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/%s?api_key=%s", riotAcc.Puuid, api_token)
	resp, err = http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		JSONBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(JSONBody, &riotAcc)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(resp.Status)
}

func main() {
	godotenv.Load()
	api_token = os.Getenv("RIOT_TOKEN")
	var aetrig RiotAcc
	getRiotAccByGameNameTagLine(&aetrig, "aetrig", "uwu")
	fmt.Println(aetrig)
	//fmt.Printf("puuid: %s\ngameName: %s\ntagLine: %s\n", aetrig.Puuid, aetrig.GameName, aetrig.TagLine)
}
