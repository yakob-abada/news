package domain

type Articles struct {
	Status     int
	NumResults int
	Articles   []Article
}

type Article struct {
	Url         string
	ImageUrl    string
	PubDate     string
	Title       string
	Description string
	Content     string
	Keywords    []Keyword
}

type Keyword struct {
	Name   string
	Weight float64
}
