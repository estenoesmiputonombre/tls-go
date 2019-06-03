package main

//Employer ... Struct that handles the data of the entity
type Employer struct {
	ID int
	Name,
	Surname string
	language Language
}

//Language ... Struct that handles programming language
type Language struct {
	ID int
	Name,
	Description string
}
