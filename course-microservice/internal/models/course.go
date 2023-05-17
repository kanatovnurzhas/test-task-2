package models

type Course struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Teacher  string   `json:"teacher"`
	Students []string `json:"students"`
}
