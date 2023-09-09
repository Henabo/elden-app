package request

type SatelliteRegister struct {
	ID        string `json:"ID"`
	PublicKey string `json:"publicKey"`
}

type UserRegister struct {
	ID        string `json:"ID"`
	MacAddr   string `json:"macAddr"`
	PublicKey string `json:"publicKey"`
}
