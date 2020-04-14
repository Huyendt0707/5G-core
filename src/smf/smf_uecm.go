package main

import (
	"5g_core/lib/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func getSmfRegistration(ueId string) {
	uri := "http://localhost:8082/" + ueId + "/registrations/smf-registrations"
	resp, err := http.Get(uri)
	if err != nil {
		fmt.Println(err)
	}
	smfRegistrationInfo, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	statusCode := resp.StatusCode
	switch statusCode {
	case 200:
		fmt.Println("smfRegistrationInfo: ", string(smfRegistrationInfo))
	case 400:
		fmt.Println("Bad Request")
	case 403:
		fmt.Println("Forbidden")
	case 404:
		fmt.Println("Not Found")
	case 500:
		fmt.Println("Internal Server Error")
	case 503:
		fmt.Println("Service Unavailable")
	default:
		fmt.Println("Unexpected error")
	}
	fmt.Println("StatusCode: ", resp.StatusCode)
	fmt.Println(resp)
	fmt.Println("smfRegistrationInfo: ", string(smfRegistrationInfo))

}

func sendSmfRegistration(smfUecm models.SmfRegistraion) {
	userId := smfUecm.SmfInstanceId
	pduSessionId := strconv.Itoa(int(smfUecm.PduSessionId))
	uri := "http://localhost:8082/" + userId + "/registrations/smf-registrations/" + pduSessionId
	fmt.Println(uri)
	bodyRequest, err := json.Marshal(smfUecm)
	if err != nil {
		fmt.Println("Not bodyRequest: ", err)
	}
	req, err := http.NewRequest(http.MethodPut, uri, bytes.NewBuffer(bodyRequest))
	if err != nil {
		fmt.Println("Not convert to byte smf_uecm: ", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	resp.Body.Close()
	smfByte, _ := ioutil.ReadAll(resp.Body)
	smfSt := models.SmfRegistraion{}
	json.Unmarshal(smfByte, &smfSt)
	// update information smfRegistration and display
	statusCode := resp.StatusCode
	fmt.Println("statusCode: ", statusCode)
	switch statusCode {
	case 201:
		fmt.Println("smfRegistrationInfo: ", smfSt)
	case 200:
		fmt.Println("smfRegistrationInfo: ", smfSt)
	case 204:
		fmt.Println("No content")
	case 400:
		fmt.Println("Bad Request")
	case 403:
		fmt.Println("Forbidden")
	case 404:
		fmt.Println("Not Found")
	case 500:
		fmt.Println("Internal Server Error")
	case 503:
		fmt.Println("Service Unavailable")
	default:
		fmt.Println("Unexpected error")
	}
}
func initSms() models.SmfRegistraion {
	smf := models.SmfRegistraion{
		SmfInstanceId: "1",
		PduSessionId:  2,
	}
	fmt.Println(smf)
	return smf
}
func main() {
	smf := initSms()
	// getSmfRegistration("1")
	sendSmfRegistration(smf)

}
