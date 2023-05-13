package entities

type Movie struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Director string  `json:"director"`
	Length   float64 `json:"length"`
	Content  string
}

type Match struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	User1    string `json:"user1"`
	User2    string `json:"user2"`
	Board    string `json:"board"`
	Turn     string `json:"turn"`
}
