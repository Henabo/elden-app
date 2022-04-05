package request

type SatelliteRegister struct {
	Id        string `json:"id"`
	PublicKey string `json:"publicKey"`
}

type UserRegister struct {
	Id        string `json:"id"`
	MacAddr   string `json:"macAddr"`
	PublicKey string `json:"publicKey"`
}
