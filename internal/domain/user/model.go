package user

type User struct {
	UUID    string `json:"uuid,omitempty"`
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
	Email   string `json:"email,omitempty"`
	Age     int    `json:"age,omitempty"`
}
