package models

type Food struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Price uint   `json:"price"`
}

type EditFood struct {
	Name  string `json:"name"`
	Price uint   `json:"price"`
}

var Foods = []Food{
	{ID: 1, Name: "rice", Price: 50},
	{ID: 2, Name: "noodle", Price: 100},
}
