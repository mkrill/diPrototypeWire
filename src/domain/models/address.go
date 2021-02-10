// package "models" contains data structures belonging to business entities including corresponding functions
// => see also package dtos
package models

type Address struct {
	Street   string
	StreetNr string
	PostCode string
	City     string
	KlsId    string
}

// functions could also be included here, e.g. to print the address
