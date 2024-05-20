package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
)

type RiotAcc struct {
	GameName  string
	TagLine   string
	Puuid     string
	Id        string
	AccountId string
}

func (Acc RiotAcc) Print() {
	v := reflect.ValueOf(Acc)
	typeOfThis := v.Type()
	fmt.Println()
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("%s: %v\n", typeOfThis.Field(i).Name, v.Field(i).Interface())
	}
}

type Server string

const (
	Brazil              Server = "br1"
	Eu_northeast        Server = "eun1"
	Eu_west             Server = "euw1"
	Japan               Server = "jp1"
	Korea               Server = "kr"
	Latin_america_north Server = "la1"
	Latin_america_south Server = "la2"
	North_america       Server = "na1"
	Oceania             Server = "oc1"
	Phillippines        Server = "ph2"
	Russia              Server = "ru"
	Singapore           Server = "sg2"
	Thailand            Server = "th2"
	Turkey              Server = "tr1"
	Taiwan              Server = "tw2"
	Vietnam             Server = "vn2"
)

// Doesn't modify RiotAcc if Status != 200 OK
func GetRiotAccByGameNameTagLine(riotAcc *RiotAcc, gameName string, tagLine string, server Server) {
	var api_token string = os.Getenv("RIOT_TOKEN")
	//println(api_token)
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
	} else {
		fmt.Println(resp.Status) //! TEMP
	}
	link = fmt.Sprintf("https://%s.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/%s?api_key=%s", server, riotAcc.Puuid, api_token)
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
	} else {
		fmt.Println(resp.Status) //! TEMP
	}
}
