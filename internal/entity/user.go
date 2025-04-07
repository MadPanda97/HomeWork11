package entity

type User struct {
	ID       int
	Name     string
	Email    string
	Phone    string
	Login    string
	Password string
	Balance  int64
}

type Address struct {
	ID        int
	UserID    int
	Country   string
	City      string
	Street    string
	Number    string
	Floor     string
	Apartment string
	Long      int64
	Lat       int64
}

type UpdateUserRequest struct {
	ID    int
	Name  string
	Email string
	Phone string
}
