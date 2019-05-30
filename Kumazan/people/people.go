package people

// People structure in db
type People struct {
	ID        string `json:"-"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
}
