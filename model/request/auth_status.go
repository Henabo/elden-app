package request

type ChangeUserAuthStatus struct {
	ID             string `json:"ID"`
	AuthStatusCode string `json:"authStatusCode"`
}
