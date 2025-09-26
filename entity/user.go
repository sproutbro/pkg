package entity

type User struct {
	ID         string `json:"id"`
	ProviderID string `json:"provider_id"`
	Provider   string `json:"provider"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Picture    string `json:"picture"`
}
