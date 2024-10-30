package models

// Secret - модель секрета
type Secret struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}
