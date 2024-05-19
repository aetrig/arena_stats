package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
)

var api_token string

type RiotAcc struct {
	GameName  string
	TagLine   string
	Puuid     string
	Id        string
	AccountId string
}

func (Acc RiotAcc) print() {
	v := reflect.ValueOf(Acc)
	typeOfThis := v.Type()
	fmt.Println()
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("%s: %v\n", typeOfThis.Field(i).Name, v.Field(i).Interface())
	}
}

type Server string

const (
	brazil              Server = "br1"
	eu_northeast        Server = "eun1"
	eu_west             Server = "euw1"
	japan               Server = "jp1"
	korea               Server = "kr"
	latin_america_north Server = "la1"
	latin_america_south Server = "la2"
	north_america       Server = "na1"
	oceania             Server = "oc1"
	phillippines        Server = "ph2"
	russia              Server = "ru"
	singapore           Server = "sg2"
	thailand            Server = "th2"
	turkey              Server = "tr1"
	taiwan              Server = "tw2"
	vietnam             Server = "vn2"
)

// Doesn't modify RiotAcc if Status != 200 OK
func getRiotAccByGameNameTagLine(riotAcc *RiotAcc, gameName string, tagLine string, server Server) {
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

func main() {
	godotenv.Load()
	api_token = os.Getenv("RIOT_TOKEN")
	var aetrig RiotAcc
	var sunny RiotAcc
	var ragna RiotAcc
	getRiotAccByGameNameTagLine(&aetrig, "aetrig", "uwu", eu_west)
	getRiotAccByGameNameTagLine(&sunny, "SunnyAsh", "AETI", eu_west)
	getRiotAccByGameNameTagLine(&ragna, "abs woman", "1312", eu_west)
	aetrig.print()
	sunny.print()
	ragna.print()
}
