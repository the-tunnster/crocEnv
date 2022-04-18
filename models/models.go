package models

type Request struct {
	ServerClientAPIKey    string `json:"SC_API_Key"`
	OragnizationAPIKey    string `json:"O_API_Key"`
	ClientOrganizationKey string `json:"CO_API_Key"`
	CrocCodePhrase        string `json:"code-phrase"`
}

type Response struct {
	Status         string `json:"status"`
	CrocCodePhrase string `json:"code-phrase"`
}

type UserInfo struct {
	ServerClientAPIKey    string `json:"SC_API_Key"`
	OragnizationAPIKey    string `json:"O_API_Key"`
	ClientOrganizationKey string `json:"CO_API_Key"`
}
