package models

type ClientPushingRequest struct {
	ServerClientAPIKey    string `json:"SC_API_Key"`
	OragnizationAPIKey    string `json:"O_API_Key"`
	ClientOrganizationKey string `json:"CO_API_Key"`
	CrocCodePhrase        string `json:"code-phrase"`
}

type ClientPullingRequest struct {
	ServerClientAPIKey    string `json:"SC_API_Key"`
	OragnizationAPIKey    string `json:"O_API_Key"`
	ClientOrganizationKey string `json:"CO_API_Key"`
}

type ClientPushingResponse struct {
	Status string `json:"status"`
}

type ClientPullingResponse struct {
	CrocCodePhrase string `json:"code-phrase"`
}
