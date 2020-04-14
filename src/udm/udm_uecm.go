package main

import (
	"5g_core/lib/models"
	"5g_core/lib/mux"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var data []models.SmfRegistraion

func initData() {
	fileJson, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
	}
	defer fileJson.Close()
	bytedata, _ := ioutil.ReadAll(fileJson)
	json.Unmarshal(bytedata, &data)

}
func getSmfRegistration(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Come in")
	params := mux.Vars(r)
	userId := params["userId"]
	fmt.Println("userId: ", userId)
	for _, v := range data {
		if v.SmfInstanceId == userId {
			fmt.Println(v)
			fmt.Fprintln(w, v)
		}
	}
}
func putSmfRegistration(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Come in ok :)))")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	// update data of smf
	smf := models.SmfRegistraion{}
	smfSt := json.Unmarshal(body, &smf)
	smfJson, _ := json.Marshal(smfSt)
	fmt.Println(smfJson)
	fmt.Fprintln(w, smfJson)
}

// func putSmfRegistration(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	fmt.Println("Come in ok :)))")
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(body)
// 	fmt.Fprintln(w, "Send smfRegistration successful")
// }

func main() {
	r := mux.NewRouter()
	// r1 := httprouter.New()
	// r1.PUT("/{userId}/registrations/smf-registrations/{pduSessionId}", putSmfRegistration)

	r.HandleFunc("/{userId}/registrations/smf-registrations/{pduSessionId}", putSmfRegistration).Methods("PUT")
	initData()
	r.HandleFunc("/{userId}/registrations/smf-registrations", getSmfRegistration).Methods("GET")
	log.Fatal(http.ListenAndServe(":8082", r))
}
