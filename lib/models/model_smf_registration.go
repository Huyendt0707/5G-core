package models

type SmfRegistraion struct {
	SmfInstanceId              string  `json:"smfInstanceId"`
	SupportedFeatures          string  `json:"supportedFeatures"`
	PduSessionId               uint8   `json:"pduSessionId"`
	SingleNssai                *Snssai `json:"singleNssai"`
	Dnn                        string  `json:"dnn"`
	EmergencyServices          bool    `json:"dnn"`
	PcscRestorationCallbackUri string  `json:"pcscRestorationCallbackUri"`
	PlmnId                     *PlmnId `json:"plmnId"`
	PgwFqdn                    string  `json:"pgwFqdn"`
}
