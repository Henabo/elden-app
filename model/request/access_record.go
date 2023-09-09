package request

import "github.com/hiro942/elden-app/model"

type CreateAccessRecord struct {
	ID           string                 `json:"ID"`
	MacAddr      string                 `json:"macAddr"`
	AccessRecord model.UserAccessRecord `json:"accessRecord"`
}
