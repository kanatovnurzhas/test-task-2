package models

type Request struct {
	Method string `json:"method"`
	Url    string `json:"url"`
}

type Response struct {
	Students []Student `json:"students"`
}

type Student struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Email   string   `json:"email"`
	Grade   string   `json:"grade"`
	Courses []string `json:"courses"`
}
