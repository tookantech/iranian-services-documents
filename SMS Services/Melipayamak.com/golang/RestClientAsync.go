package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "strconv"
    "io/ioutil"
    "net/http"
    "time"
)

func main() {
	SendSMS("username", "password", "to", "from", "go rest async test", false)
    SendSMS("username", "password", "to", "from", "go rest async test", false)
    SendSMS("username", "password", "to", "from", "go rest async test", false)

    time.Sleep(5 * time.Second)
}

func makeRequest(jsonData map[string]string, op string) {

	jsonValue, _ := json.Marshal(jsonData)
    response, err := http.Post("https://rest.payamak-panel.com/api/SendSMS/" + op, "application/json", bytes.NewBuffer(jsonValue))
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }
}

func SendSMS(username string, password string, to string, from string, text string, isFlash bool) {

    jsonData := map[string]string {
    	"username": username,
    	"password": password, 
    	"to" : to, 
    	"from" : from, 
    	"text" : text, 
    	"isFlash" : strconv.FormatBool(isFlash),
    }
    go makeRequest(jsonData, "SendSMS")
}

func SendByBaseNumber(username string, password string, text string, to string, bodyId int64) {

    jsonData := map[string]string {
    	"username": username,
    	"password": password, 
    	"text" : text, 
    	"to" : to, 
    	"bodyId" : strconv.FormatInt(bodyId, 10),
    }
    go makeRequest(jsonData, "BaseServiceNumber")
}

func GetDeliveries2(username string, password string, recID int64) {
	jsonData := map[string]string {
    	"username": username,
    	"password": password, 
    	"recID" : strconv.FormatInt(recID, 10), 
    }
    go makeRequest(jsonData, "GetDeliveries2")
}

func GetMessages(username string, password string, location int, from string, index string, count bool) {
	jsonData := map[string]string {
    	"UserName": username,
    	"PassWord": password, 
    	"Location" : strconv.Itoa(location), 
    	"From" : from,
    	"Index" : index,
    	"Count" : strconv.FormatBool(count),
    }
    go makeRequest(jsonData, "GetMessages")
}

func GetCredit(username string, password string) {
	jsonData := map[string]string {
    	"UserName": username,
    	"PassWord": password, 
    }
    go makeRequest(jsonData, "GetCredit")
}


func GetBasePrice(username string, password string) {
	jsonData := map[string]string {
    	"UserName": username,
    	"PassWord": password, 
    }
    go makeRequest(jsonData, "GetBasePrice")
}

func GetUserNumbers(username string, password string) {
	jsonData := map[string]string {
    	"UserName": username,
    	"PassWord": password, 
    }
    go makeRequest(jsonData, "GetUserNumbers")
}


