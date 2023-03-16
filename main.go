package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*album represents data about a record album.
`json:"id"` specifies how the data will look in json*/
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Taylor Allderdice", Artist: "Wiz Khalifa", Price: 10.99},
	{ID: "2", Title: "The College Dropout", Artist: "Kanye West", Price: 12.99},
	{ID: "3", Title: "Blueprint", Artist: "Jay-Z", Price: 11.99},
}

/*assign the handler function to an endpoint path.
This sets up an association in which getAlbums handles requests to the /albums endpoint path.*/
func main() {
	//Initialize a Gin router using Default.
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	//Use the Run function to attach the router to an http.Server and start the server.
	router.Run("localhost:8080")
}

/*getAlbums returns a list of albums as json.
gin.Context is the most important part of Gin.
It carries request details, validates and serializes JSON, and more.*/
func getAlbums(c *gin.Context) {
	/*Call Context.IndentedJSON to serialize the struct into JSON and add it to the response.
	The function’s first argument is the HTTP status code you want to send to the client.
	Here, you’re passing the StatusOK constant from the net/http package to indicate 200 OK.
	Note that you can replace Context.IndentedJSON with a call to Context.JSON to send more compact JSON.
	In practice, the indented form is much easier to work with when debugging and the size difference is usually small.*/
	c.IndentedJSON(http.StatusOK, albums)
}

//get a single album by ID
func getAlbumByID(c *gin.Context) {
	//extract the id from the request path.
	id := c.Param("id")

	//loop through the list of albums, looking for an album whose ID value matches the parameter value.
	for _, a := range albums {
		if a.ID == id {
			//if album is found, call IndentedJSON to return the album(a) as a response.
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	//if no album is found.
	c.IndentedJSON(http.StatusNotFound, gin.H{"msg": "Album not found"})
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	//initialize a newAlbum of type album.
	var newAlbum album

	//call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	//add the new album to the slice.
	albums = append(albums, newAlbum)

	/*IndentedJSON serializes the given struct as pretty JSON (indented + endlines) into the response body.
	It also sets the Content-Type as "application/json".*/
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
