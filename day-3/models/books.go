package models

type Books struct {
	Data []BookData `json:"data"`
}

type BookData struct {
	Id        int    `json:"id"`
	Author    string `json:"author"`
	Country   string `json:"country"`
	ImageLink string `json:"imageLink"`
	Language  string `json:"language"`
	Link      string `json:"link"`
	Pages     int    `json:"pages"`
	Title     string `json:"title"`
	Year      int    `json:"year"`
}
