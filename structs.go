package main

// Response struct
type Response struct {
	Page []Page
}

//Page struct
type Page []struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  Address `json:"address"`
	Phone    string  `json:"phone"`
	Website  string  `json:"website"`
	Company  Company `json:"company"`
}

//Geo struct
type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

//Address struct
type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     Geo    `json:"geo"`
}

//Company struct
type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

//ResponseDatabase struct
type ResponseDatabase struct {
	Database []Database
}

//Database struct
type Database struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
