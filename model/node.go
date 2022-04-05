package model

import "time"

type Node struct {
	Id            string      `json:"id"`
	NodeType      string      `json:"nodeType"`
	PublicKey    interface{} `json:"publicKey"`
	AccessRecords interface{} `json:"accessRecords"`
	CreatedAt     time.Time   `json:"createdAt"`
	UpdatedAt     time.Time   `json:"updatedAt"`
}

// UserPublicKeys is the struct of user's public key
type UserPublicKeys map[string]string

// NodePair indicates the communication parties
type NodePair struct {
	MacAddr     string `json:"macAddr"`
	SatelliteId string `json:"satelliteId"`
}

// UserAccessRecord is single access log
type UserAccessRecord struct {
	AccessType          string    `json:"accessType"`          // normal OR fast OR handover
	PreviousSatelliteId string    `json:"previousSatelliteId"` // previous satellite for handover
	StartAt             time.Time `json:"startAt"`             // when to start
	EndAt               time.Time `json:"endAt"`               // when to end
}

// UserAccessRecords indicates access records for a specific device
type UserAccessRecords map[NodePair][]UserAccessRecord
