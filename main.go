package main

/*album represents data about a record album.
`json:"id"` specifies how the data will look in json*/
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album {
	{ID: "1", Title: "Taylor Allderdice", Artist: "Wiz Khalifa", Price: 10.99},
	{ID: "2", Title: "The College Dropout", Artist: "Kanye West", Price: 12.99},
	{ID: "3", Title: "Blueprint", Artist: "Jay-Z", Price: 11.99},
}