package model

// PublicKeys indicates the structure how public keys are saved
type PublicKeys map[string]string

// UserAccessRecord is single access log
type UserAccessRecord struct {
	AccessType          string `json:"accessType"`          // "first" || "normal" || "handover"
	SatelliteID         string `json:"satelliteID"`         // current satellite
	PreviousSatelliteID string `json:"previousSatelliteID"` // previous satellite in handover
	StartAt             string `json:"startAt"`             // when to start
	EndAt               string `json:"endAt"`               // when to end
}

// UserAccessRecords indicates access records for a specific device
// macAddr -> []UserAccessRecord
type UserAccessRecords map[string][]UserAccessRecord

type Node struct {
	ID           string            `json:"ID"`
	NodeType     string            `json:"nodeType"`
	PublicKey    PublicKeys        `json:"publicKey"`
	AccessRecord UserAccessRecords `json:"accessRecord"`
	Status       string            `json:"status"`
	CreatedAt    string            `json:"createdAt"`
	UpdatedAt    string            `json:"updatedAt"`
}
