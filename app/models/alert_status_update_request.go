package models

type AlertStatusUpdateRequest struct {
	Status string `json:"status"`
	Text   string `json:"text"`
}
