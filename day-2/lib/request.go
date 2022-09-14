package lib

type StoreUserRequest struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required"`
	Password    string `json:"password" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Status      int    `json:"status" validate:"required"`
}

type UpdateUserRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Status      int    `json:"status"`
}

type StoreBookRequest struct {
	Id        int    `json:"id" validate:"required"`
	Title     string `json:"title" validate:"required"`
	Author    string `json:"author" validate:"required"`
	Country   string `json:"country" validate:"required"`
	Language  string `json:"language" validate:"required"`
	Pages     int    `json:"pages" validate:"required"`
	ImageLink string `json:"imageLink"`
	Link      string `json:"link"`
	Year      int    `json:"year"`
}

type UpdatedBookRequest struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Country   string `json:"country"`
	Language  string `json:"language"`
	Pages     int    `json:"pages"`
	ImageLink string `json:"imageLink"`
	Link      string `json:"link"`
	Year      int    `json:"year"`
}
