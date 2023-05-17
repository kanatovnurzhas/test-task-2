package models

type Request struct {
	Method string `json:"method"`
	Url    string `json:"url"`
}

type Response struct {
	Course []Course `json:"courses"`
}

type Course struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Teacher  string   `json:"teacher"`
	Students []string `json:"students"`
}
