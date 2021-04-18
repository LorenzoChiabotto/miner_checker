package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func workersStatusChecker(wallet string) bool {
	response, err := http.Get(fmt.Sprintf("https://api.ethermine.org/miner/%s/dashboard", wallet))

	if err != nil {
		log.Fatalln(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	var responseObject Dashboard
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		log.Println(err)
	}

	return len(responseObject.Data.Workers) == responseObject.Data.CurrentStatistics.ActiveWorkers
}

func sendSMSNotification(cellphone string) bool {
	body := SMSRequestBody{
		APIKey:    os.Getenv("NEXMO_API_KEY"),
		APISecret: os.Getenv("NEXMO_API_SECRET"),
		To:        cellphone,
		From:      "miner_checker",
		Text:      "Hey, some of your pool workers are down. Check it now!",
	}

	smsBody, err := json.Marshal(body)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("https://rest.nexmo.com/sms/json", "application/json", bytes.NewBuffer(smsBody))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK

}

func main() {
	log.Println("--- miner_checker started---")
	log.Println("--- loading environment variables ---")
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	wallet := os.Getenv("WALLET_ID")
	cellphone := os.Getenv("CELLPHONE_NUMBER")

	checkTicker := time.NewTicker(30 * time.Minute)
	isMinerActive := true
	inactiveRounds := 0
	messageWasSent := false
	for {
		select {
		case <-checkTicker.C:
			log.Println(" Checking workers status ")
			isMinerActive = workersStatusChecker(wallet)
			if isMinerActive {
				inactiveRounds = 0
				messageWasSent = false
			}else{
				inactiveRounds += 1
			}
		}
		if !isMinerActive && (inactiveRounds == 1 || !messageWasSent) {
			log.Println(" Sending Message ")
			messageWasSent = sendSMSNotification(cellphone)
		}
	}
}