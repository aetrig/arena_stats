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
	Puuid    string
	GameName string
	TagLine  string
}

// Return empty struct if Status != 200 OK
func getRiotAcc(gameName string, tagLine string) RiotAcc {
	var Acc RiotAcc
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
		err = json.Unmarshal(JSONBody, &Acc)
		if err != nil {
			log.Fatal(err)
		}
		return Acc
	}
	fmt.Println(resp.Status)
	return RiotAcc{}
}

func main() {
	godotenv.Load()
	api_token = os.Getenv("RIOT_TOKEN")
	aetrig := getRiotAcc("aetrig", "uwu")
	fmt.Printf("puuid: %s\ngameName: %s\ntagLine: %s\n", aetrig.Puuid, aetrig.GameName, aetrig.TagLine)
}
