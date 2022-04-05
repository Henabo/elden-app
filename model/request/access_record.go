package request

import "github.com/hiro942/elden-app/model"

type CreateAccessRecord struct {
	Id           string `json:"id"`
	MacAddr      string `json:"macAddr"`
	SatelliteId  string `json:"satelliteId"`
	AccessRecord model.UserAccessRecord
}