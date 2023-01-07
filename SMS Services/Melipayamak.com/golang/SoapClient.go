package main

import (
    "bytes"
    "strings"
    "fmt"
    "strconv"
    "io/ioutil"
    "net/http"
)

func main() {
    SendSimpleSMS2("9120000000", "****", "09120000000", "5000***", "go soap test", false)
}

//request methods
func makeRequestToSend(Data string) {

    dataAsByte := []byte(Data)
    response, err := http.Post("https://api.payamak-panel.com/post/send.asmx", "text/xml; charset=utf-8", bytes.NewBuffer(dataAsByte))
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }
}
func makeRequestToReceive(Data string) {

    dataAsByte := []byte(Data)
    response, err := http.Post("https://api.payamak-panel.com/post/receive.asmx", "text/xml; charset=utf-8", bytes.NewBuffer(dataAsByte))
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }
}
func makeRequestToContacts(Data string) {

    dataAsByte := []byte(Data)
    response, err := http.Post("https://api.payamak-panel.com/post/contacts.asmx", "text/xml; charset=utf-8", bytes.NewBuffer(dataAsByte))
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }
}
func makeRequestToActions(Data string) {

    dataAsByte := []byte(Data)
    response, err := http.Post("https://api.payamak-panel.com/post/contacts.asmx", "text/xml; charset=utf-8", bytes.NewBuffer(dataAsByte))
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }
}
func makeRequestToSchedule(Data string) {

    dataAsByte := []byte(Data)
    response, err := http.Post("https://api.payamak-panel.com/post/schedule.asmx", "text/xml; charset=utf-8", bytes.NewBuffer(dataAsByte))
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }
}
func makeRequestToTickets(Data string) {

    dataAsByte := []byte(Data)
    response, err := http.Post("https://api.payamak-panel.com/post/tickets.asmx", "text/xml; charset=utf-8", bytes.NewBuffer(dataAsByte))
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }
}
func makeRequestToUsers(Data string) {

    dataAsByte := []byte(Data)
    response, err := http.Post("https://api.payamak-panel.com/post/users.asmx", "text/xml; charset=utf-8", bytes.NewBuffer(dataAsByte))
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }
}
func makeRequestToVoice(Data string) {

    dataAsByte := []byte(Data)
    response, err := http.Post("https://api.payamak-panel.com/post/voice.asmx", "text/xml; charset=utf-8", bytes.NewBuffer(dataAsByte))
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }
}





func SendSimpleSMS2(username string, password string, to string, from string, msg string, flash bool) {
    _func := "SendSimpleSMS2";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><to>" + to + "</to><from>" + from + "</from><text>" + msg + "</text><isflash>" + strconv.FormatBool(flash) + "</isflash></"+ _func +"></soap:Body></soap:Envelope>";

    makeRequestToSend(wsReq)
}
func SendSimpleSMS(username string, password string, to []string, from string, msg string, flash bool) {

    _to := "<string>" + strings.Join(to, "</string><string>") + "</string>";

    _func := "SendSimpleSMS";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><to>" + _to + "</to><from>" + from + "</from><text>" + msg + "</text><isflash>" + strconv.FormatBool(flash) + "</isflash></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToSend(wsReq)
}
func SendSms(username string, password string, to []string, from string, msg string, flash bool, udh string, recid []int64) {
    
    _to := "<string>" + strings.Join(to, "</string><string>") + "</string>";
    _recid := "<long>" + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(recid)), "</long><long>"), "[]") + "</long>";

    _func := "SendSms";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><to>" + _to + "</to><from>" + from + "</from><text>" + msg + "</text><isflash>" + strconv.FormatBool(flash) + "</isflash><udh>" + udh + "</udh><recId>" + _recid + "</recId></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToSend(wsReq);
}
func SendWithDomain(username string, password string, to string, from string, msg string, flash bool, domain string) {
   
    _func := "SendWithDomain";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><to>" + to + "</to><from>" + from + "</from><text>" + msg + "</text><isflash>" + strconv.FormatBool(flash) + "</isflash><domainName>" + domain + "</domainName></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToSend(wsReq);
}
func SendByBaseNumber(username string, password string, text []string, to string, bodyId int) {

    _text := "<string>" + strings.Join(text, "</string><string>") + "</string>";

    _func := "SendByBaseNumber";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><text>" + _text + "</text><to>" + to + "</to><bodyId>" + strconv.Itoa(bodyId) + "</bodyId></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToSend(wsReq)
}
func SendByBaseNumber2(username string, password string, text string, to string, bodyId int) {

    _func := "SendByBaseNumber2";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><text>" + text + "</text><to>" + to + "</to><bodyId>" + strconv.Itoa(bodyId) + "</bodyId></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToSend(wsReq)
}
func getMessages(username string, password string, location int, from string, index int, count int) {

    _func := "getMessages";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><location>" + strconv.Itoa(location) + "</location><from>" + from + "</from><index>" + strconv.Itoa(index) + "</index><count>" + strconv.Itoa(count) + "</count></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToSend(wsReq);
}
func GetSmsPrice(username string, password string, irancellCount int, mtnCount int, from string, text string) {

    _func := "GetSmsPrice";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><irancellCount>" + strconv.Itoa(irancellCount) + "</irancellCount><mtnCount>" + strconv.Itoa(mtnCount) + "</mtnCount><from>"+ from +"</from><text>"+ text +"</text></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToSend(wsReq);
}
func GetMultiDelivery2(username string, password string, recId string) {

    _func := "GetMultiDelivery2";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><recId>" + recId + "</recId></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToSend(wsReq);
}
func GetMultiDelivery(username string, password string, recId string) {

    _func := "GetMultiDelivery";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><recId>" + recId + "</recId></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToSend(wsReq);
}
func GetInboxCount(username string, password string, isRead bool) {

    _func := "GetInboxCount";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><isRead>" + strconv.FormatBool(isRead) + "</isRead></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToSend(wsReq);
}
func GetDelivery2(username string, password string, recId string) {

    _func := "GetDelivery2";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><recId>" + recId + "</recId></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToSend(wsReq);
}
func GetDelivery(username string, password string, recId string) {

    _func := "GetDelivery";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><recId>" + recId + "</recId></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToSend(wsReq);
}
func GetDeliveries3(username string, password string, recIds []string) {

    _recids := "<string>" + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(recIds)), "</string><string>"), "[]") + "</string>";

    _func := "GetDeliveries3";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><recId>" + _recids + "</recId></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToSend(wsReq);
}
func GetDeliveries2(username string, password string, recId string) {

    _func := "GetDeliveries2";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><recId>" + recId + "</recId></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToSend(wsReq);
}
func GetDeliveries(username string, password string, recIds []int64) {

    _recids := "<long>" + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(recIds)), "</long><long>"), "[]") + "</long>";

    _func := "GetDeliveries";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><recId>" + _recids + "</recId></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToSend(wsReq);
}
func GetCredit(username string, password string) {

    _func := "GetCredit";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToSend(wsReq);
}



// //Receive API Operations
func RemoveMessages2(username string, password string, location int, msgIds string) {

    _func := "RemoveMessages2";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><location>" + strconv.Itoa(location) + "</location><msgIds>" + msgIds + "</msgIds></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToReceive(wsReq);
}
//use Received or Sent or Removed or Deleted for location in the next method
func RemoveMessages(username string, password string, location string, msgIds string) {

    _func := "RemoveMessages";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><location>" + location + "</location><msgIds>" + msgIds + "</msgIds></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToReceive(wsReq);
}
func GetUsersMessagesByDate(username string, password string, location int, from string, index int, count int, dateFrom string, dateTo string) {

    _func := "GetUsersMessagesByDate";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><location>" + strconv.Itoa(location) + "</location><from>" + from + "</from><index>" + strconv.Itoa(index) + "</index><count>" + strconv.Itoa(count) +"</count><dateFrom>" + dateFrom + "</dateFrom><dateTo>" + dateTo + "</dateTo></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToReceive(wsReq);
}
func GetOutBoxCount(username string, password string) {

    _func := "GetOutBoxCount";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToReceive(wsReq);
}
func GetMessagesWithChangeIsRead(username string, password string, location int, from string, index int, count int, isRead bool, changeIsRead bool) {

   _func := "GetMessagesWithChangeIsRead";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><location>" + strconv.Itoa(location) + "</location><from>" + from + "</from><index>" + strconv.Itoa(index) + "</index><count>" + strconv.Itoa(count) +"</count><isRead>" + strconv.FormatBool(isRead) + "</isRead><changeIsRead>" + strconv.FormatBool(changeIsRead) + "</changeIsRead></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToReceive(wsReq);
}
func GetMessagesReceptions(username string, password string, msgId int, fromRows int) {

    _func := "GetMessagesReceptions";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><msgId>" + strconv.Itoa(msgId) + "</msgId><fromRows>" + strconv.Itoa(fromRows) + "</fromRows></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToReceive(wsReq);
}
func GetMessagesFilterByDate(username string, password string, location int, from string, index int, count int, dateFrom string, dateTo string, isRead bool) {

    _func := "GetMessagesFilterByDate";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><location>" + strconv.Itoa(location) + "</location><from>" + from + "</from><index>" + strconv.Itoa(index) + "</index><count>" + strconv.Itoa(count) + "</count><dateFrom>" + dateFrom + "</dateFrom><dateTo>" + dateTo + "</dateTo><isRead>" + strconv.FormatBool(isRead) + "</isRead></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToReceive(wsReq);
}
func GetMessagesByDate(username string, password string, location int, from string, index int, count int, dateFrom string, dateTo string) {

    _func := "GetMessagesByDate";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><location>" + strconv.Itoa(location) + "</location><from>" + from + "</from><index>" + strconv.Itoa(index) + "</index><count>" + strconv.Itoa(count) + "</count><dateFrom>" + dateFrom + "</dateFrom><dateTo>" + dateTo + "</dateTo></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToReceive(wsReq);
}
func GetMessagesAfterIDJson(username string, password string, location int, from string, count int, msgId int) {

    _func := "GetMessagesAfterIDJson";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><location>" + strconv.Itoa(location) + "</location><from>" + from + "</from><count>" + strconv.Itoa(count) + "</count><msgId>" + strconv.Itoa(msgId) + "</msgId></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToReceive(wsReq);
}
func GetMessagesAfterID(username string, password string, location int, from string, count int, msgId int) {

    _func := "GetMessagesAfterID";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><location>" + strconv.Itoa(location) + "</location><from>" + from + "</from><count>" + strconv.Itoa(count) + "</count><msgId>" + strconv.Itoa(msgId) + "</msgId></"+ _func +"></soap:Body></soap:Envelope>";
    makeRequestToReceive(wsReq);
}
func GetMessages(username string, password string, location int, from string, index int, count int) {

    _func := "GetMessages";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><location>" + strconv.Itoa(location) +"</location><from>" + from + "</from><index>" + strconv.Itoa(index) + "</index><count>" + strconv.Itoa(count) + "</count></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToReceive(wsReq);
}
func GetMessageStr(username string, password string, location int, from string, index int, count int) {

    _func := "GetMessageStr";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><location>" + strconv.Itoa(location) +"</location><from>" + from + "</from><index>" + strconv.Itoa(index) + "</index><count>" + strconv.Itoa(count) + "</count></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToReceive(wsReq);
}
func ChangeMessageIsRead(username string, password string, msgIds string) {

    _func := "ChangeMessageIsRead";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><msgIds>" + msgIds + "</msgIds></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToReceive(wsReq);
}

// //Users API Operations
func AddPayment(username string, password string, name string, family string, bankName string, code string, amount int,  cardNumber string) {

    _func := "AddPayment";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><name>"+name+"</name><family>"+family+"</family><bankName>"+bankName+"</bankName><code>"+code+"</code><amount>" + strconv.Itoa(amount)+"</amount><cardNumber>" + cardNumber + "</cardNumber></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToUsers(wsReq);
}
func AddUser(username string, password string, productId int, descriptions string, mobileNumber string, emailAddress string, nationalCode string, name string, family string, corporation string, phone string, fax string, address string, postalCode string, certificateNumber string) {

    _func := "AddUser";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><productId>" + strconv.Itoa(productId) + "</productId><descriptions>"+descriptions+"</descriptions><mobileNumber>"+mobileNumber+"</mobileNumber><emailAddress>"+emailAddress+"</emailAddress><nationalCode>"+nationalCode+"</nationalCode><name>"+name+"</name><family>"+family+"</family><corporation>"+corporation+"</corporation><phone>"+phone+"</phone><fax>"+fax+"</fax><address>"+address+"</address><postalCode>"+postalCode+"</postalCode><certificateNumber>"+certificateNumber+"</certificateNumber></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToUsers(wsReq);
}
func AddUserComplete(username string, password string, productId int, descriptions string, mobileNumber string, emailAddress string, nationalCode string, name string, family string, corporation string, phone string, fax string, address string, postalCode string, certificateNumber string, country int, province int, city int, howFindUs string, commercialCode string, saleId string, recommanderId string) {

    _func := "AddUserComplete";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><productId>" + strconv.Itoa(productId) + "</productId><descriptions>"+descriptions+"</descriptions><mobileNumber>"+mobileNumber+"</mobileNumber><emailAddress>"+emailAddress+"</emailAddress><nationalCode>"+nationalCode+"</nationalCode><name>"+name+"</name><family>"+family+"</family><corporation>"+corporation+"</corporation><phone>"+phone+"</phone><fax>"+fax+"</fax><address>"+address+"</address><postalCode>"+postalCode+"</postalCode><certificateNumber>"+certificateNumber+"</certificateNumber><country>"+strconv.Itoa(country)+"</country><province>"+ strconv.Itoa(province)+"</province><city>"+strconv.Itoa(city)+"</city><howFindUs>"+howFindUs+"</howFindUs><commercialCode>"+commercialCode+"</commercialCode><saleId>"+saleId+"</saleId><recommanderId>"+recommanderId+"</recommanderId></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToUsers(wsReq);
}
func AddUserWithLocation(username string, password string, productId int, descriptions string, mobileNumber string, emailAddress string, nationalCode string, name string, family string, corporation string, phone string, fax string, address string, postalCode string, certificateNumber string, country int, province int, city int) {

    _func := "AddUserWithLocation";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><productId>" + strconv.Itoa(productId) + "</productId><descriptions>"+descriptions+"</descriptions><mobileNumber>"+mobileNumber+"</mobileNumber><emailAddress>"+emailAddress+"</emailAddress><nationalCode>"+nationalCode+"</nationalCode><name>"+name+"</name><family>"+family+"</family><corporation>"+corporation+"</corporation><phone>"+phone+"</phone><fax>"+fax+"</fax><address>"+address+"</address><postalCode>"+postalCode+"</postalCode><certificateNumber>"+certificateNumber+"</certificateNumber><country>"+strconv.Itoa(country)+"</country><province>"+strconv.Itoa(province)+"</province><city>"+strconv.Itoa(city)+"</city></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToUsers(wsReq);
}
func AddUserWithMobileNumber(username string, password string, productId int, mobileNumber string, firstName string, lastName string, email string) {

    _func := "AddUserWithMobileNumber";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><productId>"+strconv.Itoa(productId)+"</productId><mobileNumber>"+mobileNumber+"</mobileNumber><firstName>"+firstName+"</firstName><lastName>"+lastName+"</lastName><email>"+email+"</email></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToUsers(wsReq);
}
func AddUserWithUserNameAndPass(username string, password string, targetUserName string, targetUserPassword string, productId int, descriptions string, mobileNumber string, emailAddress string, nationalCode string, name string, family string, corporation string, phone string, fax string, address string, postalCode string, certificateNumber string) {

    _func := "AddUserWithUserNameAndPass";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><targetUserName>"+targetUserName+"</targetUserName><targetUserPassword>"+targetUserPassword+"</targetUserPassword><productId>"+strconv.Itoa(productId)+"</productId><descriptions>"+descriptions+"</descriptions><mobileNumber>"+mobileNumber+"</mobileNumber><emailAddress>"+emailAddress+"</emailAddress><nationalCode>"+nationalCode+"</nationalCode><name>"+name+"</name><family>"+family+"</family><corporation>"+corporation+"</corporation><phone>"+phone+"</phone><fax>"+fax+"</fax><address>"+address+"</address><postalCode>"+postalCode+"</postalCode><certificateNumber>"+certificateNumber+"</certificateNumber></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToUsers(wsReq);
}
func AuthenticateUser(username string, password string) {

    _func := "AuthenticateUser";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToUsers(wsReq);
}
func ChangeUserCredit(username string, password string, amount int, description string, targetUsername string, GetTax bool) {

    _func := "ChangeUserCredit";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><amount>" + strconv.Itoa(amount) + "</amount><description>"+description+"</description><targetUsername>"+targetUsername+"</targetUsername><GetTax>"+ strconv.FormatBool(GetTax)+"</GetTax></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToUsers(wsReq);
}
func ChangeUserCredit2(username string, password string, amount string, description string, targetUsername string, GetTax bool) {

    _func := "ChangeUserCredit2";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><amount>" + amount + "</amount><description>"+description+"</description><targetUsername>"+targetUsername+"</targetUsername><GetTax>"+ strconv.FormatBool(GetTax)+"</GetTax></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToUsers(wsReq);
}
func ChangeUserCreditBySeccondPass(username string, password string, ausername string, seccondPassword string, amount int, description string, targetUsername string, GetTax bool) {

    _func := "ChangeUserCreditBySeccondPass";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>"+ausername+"</username><seccondPassword>"+seccondPassword+"</seccondPassword><amount>"+strconv.Itoa(amount)+"</amount><description>"+description+"</description><targetUsername>"+targetUsername+"</targetUsername><GetTax>"+strconv.FormatBool(GetTax)+"</GetTax></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToUsers(wsReq);
}

func DeductUserCredit(username string, password string, ausername string, apassword string, amount int, description string) {

    _func := "DeductUserCredit";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>"+ausername+"</username><password>"+apassword+"</password><amount>"+strconv.Itoa(amount)+"</amount><description>"+description+"</description></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToUsers(wsReq);
}

func ForgotPassword(username string, password string, mobileNumber string, emailAddress string, targetUsername string) {

    _func := "ForgotPassword";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><mobileNumber>"+mobileNumber+"</mobileNumber><emailAddress>"+emailAddress+"</emailAddress><targetUsername>"+targetUsername+"</targetUsername></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToUsers(wsReq);
}
func GetCities(username string, password string, provinceId int) {

    _func := "GetCities";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><provinceId>" + strconv.Itoa(provinceId) + "</provinceId></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToUsers(wsReq);
}
func GetEnExpireDate(username string, password string) {

   _func := "GetEnExpireDate";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToUsers(wsReq);
}
func GetExpireDate(username string, password string) {

    _func := "GetExpireDate";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToUsers(wsReq);
}
func GetProvinces(username string, password string) {

    _func := "GetProvinces";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToUsers(wsReq);
}
func GetUserBasePrice(username string, password string, targetUsername string) {

    _func := "GetUserBasePrice";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><targetUsername>" + targetUsername + "</targetUsername></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToUsers(wsReq);
}
func GetUserByUserID(username string, password string, pass string, userId int) {

    _func := "GetUserByUserID";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><pass>"+pass+"</pass><userId>"+strconv.Itoa(userId)+"</userId></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToUsers(wsReq);
}
func GetUserCredit(username string, password string, targetUsername string) {

    _func := "GetUserCredit";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><targetUsername>"+targetUsername+"</targetUsername></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToUsers(wsReq);
}
func GetUserCreditBySecondPass(username string, password string, secondPassword string, targetUsername string) {

    _func := "GetUserCreditBySecondPass";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><secondPassword>"+secondPassword+"</secondPassword><targetUsername>"+targetUsername+"</targetUsername></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToUsers(wsReq);
}
func GetUserDetails(username string, password string, targetUsername string) {
    _func := "GetUserDetails";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><targetUsername>"+targetUsername+"</targetUsername></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToUsers(wsReq);
}
func GetUserIsExist(username string, password string, targetUsername string) {

    _func := "GetUserIsExist";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><targetUsername>"+targetUsername+"</targetUsername></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToUsers(wsReq);
}
func GetUserNumbers(username string, password string) {

    _func := "GetUserNumbers";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToUsers(wsReq);
}
func GetUserTransactions(username string, password string, targetUsername string, creditType string, dateFrom string,  dateTo string, keyword string) {

    _func := "GetUserTransactions";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><targetUsername>"+targetUsername+"</targetUsername><creditType>"+creditType+"</creditType><dateFrom>"+dateFrom+"</dateFrom><dateTo>"+dateTo+"</dateTo><keyword>"+keyword+"</keyword></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToUsers(wsReq);
}
func GetUserWallet(username string, password string) {

    _func := "GetUserWallet";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToUsers(wsReq);
}
func GetUserWalletTransaction(username string, password string, dateFrom string, dateTo string, count int, startIndex int, payType string, PayLoc string) {

    _func := "GetUserWalletTransaction";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><dateFrom>" + dateFrom + "</dateFrom><dateTo>" + dateTo + "</dateTo><count>" + strconv.Itoa(count) + "</count><startIndex>" + strconv.Itoa(startIndex) + "</startIndex><payType>" + payType + "</payType><payLoc>" + PayLoc + "</payLoc></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToUsers(wsReq);
}
func GetUsers(username string, password string) {

    _func := "GetUsers";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToUsers(wsReq);
}
func HasFilter(username string, password string, text string) {

    _func := "HasFilter";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><text>" + text + "</text></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToUsers(wsReq);
}
func RemoveUser(username string, password string, targetUsername string) {

    _func := "RemoveUser";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><targetUsername>" +targetUsername + "</targetUsername></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToUsers(wsReq);
}
func RevivalUser(username string, password string, targetUsername string) {

    _func := "RevivalUser";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><targetUserName>" + targetUsername + "</targetUserName></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToUsers(wsReq);
}

// //Contact API Operations
func AddContact(username string, password string, groupIds string, firstname string, lastname string, nickname string, corporation string, mobilenumber string, phone string, fax string, birthdate string, email string, gender int, province int, city int, address string, postalCode string, additionaldate string, additionaltext string, descriptions string) {

    _func := "AddContact";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><groupIds>" + groupIds + "</groupIds><firstname>" + firstname + "</firstname><lastname>" + lastname + "</lastname><nickname>" + nickname + "</nickname><corporation>" + corporation + "</corporation><mobilenumber>" + mobilenumber + "</mobilenumber><phone>" + phone + "</phone><fax>" + fax + "</fax><birthdate>" + birthdate + "</birthdate><email>" + email + "</email><gender>" + strconv.Itoa(gender) + "</gender><province>" + strconv.Itoa(province) + "</province><city>" + strconv.Itoa(city) + "</city><address>" + address + "</address><postalCode>" + postalCode + "</postalCode><additionaldate>" + additionaldate + "</additionaldate><additionaltext>" + additionaltext + "</additionaltext><descriptions>" + descriptions + "</descriptions></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToContacts(wsReq);
}
func AddContactEvents(username string, password string, contactId int, eventName string, eventType int, eventDate string) {

    _func := "AddContactEvents";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><contactId>" + strconv.Itoa(contactId) + "</contactId><eventName>" + eventName + "</eventName><eventDate>" + eventDate + "</eventDate><eventType>" + strconv.Itoa(eventType) + "</eventType></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToContacts(wsReq);
}
func AddGroup(username string, password string, groupName string, Descriptions string, showToChilds bool) {

    _func := "AddGroup";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><groupName>" + groupName + "</groupName><Descriptions>" + Descriptions + "</Descriptions><showToChilds>" + strconv.FormatBool(showToChilds) + "</showToChilds></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToContacts(wsReq);
}
func ChangeContact(username string, password string, contactId int, mobilenumber string, firstname string, lastname string, nickname string, corporation string, phone string, fax string, email string, gender int, province int, city int, address string, postalCode string, additionaltext string, descriptions string, contactStatus int) {

    _func := "ChangeContact";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><contactId>" + strconv.Itoa(contactId) + "</contactId><mobilenumber>" + mobilenumber + "</mobilenumber><firstname>" + firstname + "</firstname><lastName>" + lastname + "</lastname><nickname>" + nickname + "</nickname><corporation>" + corporation + "</corporation><phone>" + phone + "</phone><fax>" + fax + "</fax><email>" + email + "</email><gender>" + strconv.Itoa(gender) + "</gender><province>" + strconv.Itoa(province) + "</province><city>" + strconv.Itoa(city) + "</city><address>" + address + "</address><postalCode>" + postalCode + "</postalCode><additionaltext>" + additionaltext + "</additionaltext><descriptions>" + descriptions + "</descriptions><contactStatus>" + strconv.Itoa(contactStatus) + "</contactStatus></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToContacts(wsReq);
}
func ChangeGroup(username string, password string, groupId int, groupName string, Descriptions string, showToChilds bool, groupStatus int) {

    _func := "ChangeGroup";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><groupId>" + strconv.Itoa(groupId) + "</groupId><groupName>" + groupName + "</groupName><Descriptions>" + Descriptions + "</Descriptions><showToChilds>" + strconv.FormatBool(showToChilds) + "</showToChilds><groupStatus>" + strconv.Itoa(groupStatus) + "</groupStatus></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToContacts(wsReq);
}
func CheckMobileExistInContact(username string, password string, mobileNumber string) {

    _func := "CheckMobileExistInContact";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><mobileNumber>" + mobileNumber + "</mobileNumber></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToContacts(wsReq);
}
func GetContactEvents(username string, password string, contactId int) {

    _func := "GetContactEvents";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><contactId>" + strconv.Itoa(contactId) + "</contactId></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToContacts(wsReq);
}
func GetContacts(username string, password string, groupId int, keyword string, from int, count int) {

    _func := "GetContacts";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><groupId>" + strconv.Itoa(groupId) + "</groupId><keyword>" + keyword + "</keyword><from>" + strconv.Itoa(from) + "</from><count>" + strconv.Itoa(count) + "</count></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToContacts(wsReq);
}
func GetContactsByID(username string, password string, contactId int, status int) {

    _func := "GetContactsByID";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><contactId>" + strconv.Itoa(contactId) + "</contactId><status>" + strconv.Itoa(status) + "</status></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToContacts(wsReq);
}
func GetGroups(username string, password string) {

    _func := "GetGroups";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToContacts(wsReq);
}
func MergeGroups(username string, password string, originGroupId int, destinationGroupId int, deleteOriginGroup bool) {

    _func := "MergeGroups";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><originGroupId>" + strconv.Itoa(originGroupId) + "</originGroupId><destinationGroupId>" + strconv.Itoa(destinationGroupId) + "</destinationGroupId><deleteOriginGroup>" + strconv.FormatBool(deleteOriginGroup) + "</deleteOriginGroup></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToContacts(wsReq);
}
func RemoveContact(username string, password string, mobilenumber string) {

    _func := "RemoveContact";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><mobilenumber>" + mobilenumber + "</mobilenumber></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToContacts(wsReq);
}
func RemoveContactByContactID(username string, password string, contactId int) {

    _func := "RemoveContactByContactID";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><contactId>" + strconv.Itoa(contactId) + "</contactId></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToContacts(wsReq);
}
func RemoveGroup(username string, password string, groupId int) {

    _func := "RemoveGroup";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><groupId>" + strconv.Itoa(groupId) + "</groupId></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToContacts(wsReq);
}



// //ACtions API Operations
func AddBranch(username string, password string, branchName string, owner int) {

    _func := "AddBranch";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><branchName>" + branchName + "</branchName><owner>" + strconv.Itoa(owner) + "</owner></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToActions(wsReq);
}
func AddBulk(username string, password string, from string, branch int, bulkType int, title string, message string, rangeFrom string, rangeTo string, DateToSend string, requestCount int, rowFrom int) {

    _func := "AddBulk";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><from>" + from + "</from><branch>" +strconv.Itoa(branch) + "</branch><bulkType>" + strconv.Itoa(bulkType) + "</bulkType><title>" + title + "</title><message>" + message + "</message><rangeFrom>" + rangeFrom + "</rangeFrom><rangeTo>" + rangeTo + "</rangeTo><DateToSend>" + DateToSend + "</DateToSend><requestCount>" + strconv.Itoa(requestCount) + "</requestCount><rowFrom>" + strconv.Itoa(rowFrom) + "</rowFrom></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToActions(wsReq);
}
func AddBulk2(username string, password string, from string, branch int, bulkType int, title string, message string, rangeFrom string, rangeTo string, DateToSend string, requestCount int, rowFrom int) {

    _func := "AddBulk2";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><from>" + from + "</from><branch>" +strconv.Itoa(branch) + "</branch><bulkType>" + strconv.Itoa(bulkType) + "</bulkType><title>" + title + "</title><message>" + message + "</message><rangeFrom>" + rangeFrom + "</rangeFrom><rangeTo>" + rangeTo + "</rangeTo><DateToSend>" + DateToSend + "</DateToSend><requestCount>" + strconv.Itoa(requestCount) + "</requestCount><rowFrom>" + strconv.Itoa(rowFrom) + "</rowFrom></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToActions(wsReq);
}
func AddNewBulk(username string, password string, from string, branch int, bulkType int, title string, message string, rangeFrom string, rangeTo string, DateToSend string, requestCount int, rowFrom int) {

    _func := "AddNewBulk";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><from>" + from + "</from><branch>" +strconv.Itoa(branch) + "</branch><bulkType>" + strconv.Itoa(bulkType) + "</bulkType><title>" + title + "</title><message>" + message + "</message><rangeFrom>" + rangeFrom + "</rangeFrom><rangeTo>" + rangeTo + "</rangeTo><DateToSend>" + DateToSend + "</DateToSend><requestCount>" + strconv.Itoa(requestCount) + "</requestCount><rowFrom>" + strconv.Itoa(rowFrom) + "</rowFrom></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToActions(wsReq);
}
func AddNumber(username string, password string, branchId int, mobileNumbers []string) {

    _mobileNumbers := "<string>" + strings.Join(mobileNumbers, "</string><string>") + "</string>";

    _func := "AddNumber";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><branchId>" + strconv.Itoa(branchId) + "</branchId><mobileNumbers>" + _mobileNumbers + "</mobileNumbers></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToActions(wsReq);
}
func GetBranchs(username string, password string, owner int) {

    _func := "GetBranchs";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><owner>" + strconv.Itoa(owner) + "</owner></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToActions(wsReq);
}
func GetBulk(username string, password string) {

    _func := "GetBulk";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToActions(wsReq);
}
func GetBulkCount(username string, password string, branch int, rangeFrom string, rangeTo string) {

    _func := "GetBulkCount";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><branch>" + strconv.Itoa(branch) + "</branch><rangeFrom>" + rangeFrom + "</rangeFrom><rangeTo>" + rangeTo + "</rangeTo></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToActions(wsReq);
}
func GetBulkReceptions(username string, password string, bulkId int, fromRows int) {

    _func := "GetBulkReceptions";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><bulkId>" + strconv.Itoa(bulkId) + "</bulkId><fromRows>" + strconv.Itoa(fromRows) + "</fromRows></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToActions(wsReq);
}
func GetBulkStatus(username string, password string, bulkId int, sent int, failed int, status int) {

    _func := "GetBulkStatus";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><bulkId>" + strconv.Itoa(bulkId) + "</bulkId><sent>" + strconv.Itoa(sent) + "</sent><failed>" + strconv.Itoa(failed) + "</failed><status>" + strconv.Itoa(status) + "</status></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToActions(wsReq);
}
//duplicate method
// func GetMessagesReceptions(username string, password string, msgId int64, fromRows int) {

//     _func := "GetMessagesReceptions";
//  wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><msgId>" + strconv.Itoa(msgId) + "</msgId><fromRows>" + strconv.Itoa(fromRows) + "</fromRows></"+ _func +"></soap:Body></soap:Envelope>";
 
//     makeRequestToActions(wsReq);
// }
func GetMobileCount(username string, password string, branch int, rangeFrom string, rangeTo string) {

    _func := "GetMobileCount";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><branch>"+ strconv.Itoa(branch) + "</branch><rangeFrom>" + rangeFrom + "</rangeFrom><rangeTo>" + rangeTo + "</rangeTo></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToActions(wsReq);
}
func GetSendBulk(username string, password string) {

    _func := "GetSendBulk";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToActions(wsReq);
}
func GetTodaySent(username string, password string) {

    _func := "GetTodaySent";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToActions(wsReq);
}
func GetTotalSent(username string, password string) {

    _func := "GetTotalSent";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToActions(wsReq);
}
func RemoveBranch(username string, password string, branchId int) {

    _func := "RemoveBranch";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><branchId>" + strconv.Itoa(branchId) + "</branchId></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToActions(wsReq);
}

func RemoveBulk(username string, password string, bulkId int) {

    _func := "RemoveBulk";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><bulkId>" + strconv.Itoa(bulkId) + "</bulkId></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToActions(wsReq);
}
func SendMultipleSMS(username string, password string, to []string, from string, text []string, isflash bool, udh string, recId []int64, status string) {

    _to := "<string>" + strings.Join(to, "</string><string>") + "</string>";
    _text := "<string>" + strings.Join(text, "</string><string>") + "</string>";
    _recId := "<long>" + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(recId)), "</long><long>"), "[]") + "</long>";

    _func := "SendMultipleSMS";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><to>" + _to + "</to><from>" + from + "</from><text>" + _text + "</text><isflash>" + strconv.FormatBool(isflash) + "</isflash><udh>" + udh + "</udh><recId>" + _recId + "</recId><status>" + status + "</status></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToActions(wsReq);
}
func SendMultipleSMS2(username string, password string, to []string, from []string, text []string, isflash bool, udh string, recId []int64, status string) {

    _to := "<string>" + strings.Join(to, "</string><string>") + "</string>";
    _text := "<string>" + strings.Join(text, "</string><string>") + "</string>";
    _from := "<string>" + strings.Join(from, "</string><string>") + "</string>";
    _recId := "<long>" + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(recId)), "</long><long>"), "[]") + "</long>";


    _func := "SendMultipleSMS2";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><to>" + _to + "</to><from>" + _from + "</from><text>" + _text + "</text><isflash>" + strconv.FormatBool(isflash) + "</isflash><udh>" + udh + "</udh><recId>" + _recId + "</recId><status>" + status + "</status></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToActions(wsReq);
}
func UpdateBulkDelivery(username string, password string, bulkId int) {

    _func := "UpdateBulkDelivery";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><bulkId>" + strconv.Itoa(bulkId) + "</bulkId></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToActions(wsReq);
}



// //Schedule API Operations
func AddMultipleSchedule(username string, password string, to []string, from string, text []string, isflash bool, scheduleDateTime []string, period string) {

    _to := "<string>" + strings.Join(to, "</string><string>") + "</string>";
    _text := "<string>" + strings.Join(text, "</string><string>") + "</string>";
    _schDates := "<string>" + strings.Join(scheduleDateTime, "</string><string>") + "</string>";

    _func := "AddMultipleSchedule";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><to>"+_to+"</to><from>" + from + "</from><text>" + _text + "</text><isflash>" + strconv.FormatBool(isflash) + "</isflash><scheduleDateTime>" + _schDates + "</scheduleDateTime><period>" + period + "</period></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToSchedule(wsReq);
}
func AddNewUsance(username string, password string, to string, from string, text string, isflash bool, scheduleStartDateTime string, countrepeat int, scheduleEndDateTime string, periodType string) {

    _func := "AddNewUsance";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><to>" + to + "</to><from>" + from + "</from><text>" + text + "</text><isflash>" + strconv.FormatBool(isflash) + "</isflash><scheduleStartDateTime>" +scheduleStartDateTime + "</scheduleStartDateTime><countrepeat>" + strconv.Itoa(countrepeat) + "</countrepeat><periodType>" + periodType + "</periodType></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToSchedule(wsReq);
}
func AddSchedule(username string, password string, to string, from string, text string, isflash bool, scheduleDateTime string, period string) {

    _func := "AddSchedule";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><to>" + to + "</to><from>" + from + "</from><text>" + text + "</text><isflash>" + strconv.FormatBool(isflash) + "</isflash><scheduleDateTime>" + scheduleDateTime + "</scheduleDateTime><period>" + period + "</period></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToSchedule(wsReq);
}
func AddUsance(username string, password string, to string, from string, text string, isflash bool, scheduleStartDateTime string, repeatAfterDays int, scheduleEndDateTime string) {

    _func := "AddUsance";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><to>" + to + "</to><from>" + from + "</from><text>" + text + "</text><isflash>" + strconv.FormatBool(isflash) + "</isflash><scheduleStartDateTime>" + scheduleStartDateTime + "</scheduleStartDateTime><repeatAfterDays>" + strconv.Itoa(repeatAfterDays) + "</repeatAfterDays><scheduleEndDateTime>" + scheduleEndDateTime + "</scheduleEndDateTime></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToSchedule(wsReq);
}
func GetScheduleStatus(username string, password string, scheduleId int) {

    _func := "GetScheduleStatus";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><scheduleId>" + strconv.Itoa(scheduleId) + "</scheduleId></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToSchedule(wsReq);
}
func RemoveSchedule(username string, password string, scheduleId int) {

    _func := "RemoveSchedule";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><scheduleId>" + strconv.Itoa(scheduleId) + "</scheduleId></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToSchedule(wsReq);
}
func RemoveScheduleList(username string, password string, scheduleIdList []int) {

    _list := "<int>" + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(scheduleIdList)), "</int><int>"), "[]") + "</int>";
    _func := "RemoveScheduleList";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><scheduleIdList>" + _list + "</scheduleIdList></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToSchedule(wsReq);
}


// Voice API Operations
func GetSendSMSWithSpeechTextStatus(username string, password string, recId int) {

    _func := "GetSendSMSWithSpeechTextStatus";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><recId>" + strconv.Itoa(recId) + "</recId></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToVoice(wsReq);
}
func SendBulkSpeechText(username string, password string, title string, body string, receivers string, DateToSend string, repeatCount int) {

    _func := "SendBulkSpeechText";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><title>" + title + "</title><body>" + body + "</body><receivers>" + receivers + "</receivers><DateToSend>" + DateToSend + "</DateToSend><repeatCount>" + strconv.Itoa(repeatCount) + "</repeatCount></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToVoice(wsReq);
}
func SendBulkVoiceSMS(username string, password string, title string, voiceFileId int, receivers string, DateToSend string, repeatCount int) {

    _func := "SendBulkVoiceSMS";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><title>" + title + "</title><voiceFileId>" + strconv.Itoa(voiceFileId) + "</voiceFileId><receivers>" + receivers + "</receivers><DateToSend>" + DateToSend + "</DateToSend><repeatCount>" + strconv.Itoa(repeatCount) + "</repeatCount></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToVoice(wsReq);
}
func SendSMSWithSpeechText(username string, password string, smsBody string, speechBody string, from string, to string) {

    _func := "SendSMSWithSpeechText";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><smsBody>" + smsBody + "</smsBody><speechBody>" + speechBody + "</speechBody><from>" + from + "</from><to>" + to + "</to></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToVoice(wsReq);
}
func SendSMSWithSpeechTextBySchduleDate(username string, password string, smsBody string, speechBody string, from string, to string, scheduleDate string) {

    _func := "SendSMSWithSpeechTextBySchduleDate";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><smsBody>" + smsBody + "</smsBody><speechBody>" + speechBody + "</speechBody><from>" + from + "</from><to>" + to + "</to><scheduleDate>" + scheduleDate + "</scheduleDate></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToVoice(wsReq);
}
func UploadVoiceFile(username string, password string, title string, base64StringFile string) {

    _func := "UploadVoiceFile";
    wsReq := "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><" + _func + " xmlns=\"http://tempuri.org/\"><username>" + username + "</username><password>" + password + "</password><title>" + title + "</title><base64StringFile>" + base64StringFile + "</base64StringFile></"+ _func +"></soap:Body></soap:Envelope>";
 
    makeRequestToVoice(wsReq);
}
