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
	router.PUT("/albums/:id", updateAlbum)
	router.DELETE("/albums/:id", deleteAlbum)
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

//TODO: deleteAlbum by ID
func deleteAlbum(c *gin.Context) {
	id := c.Param("id")

	//find the index(i) of the album with the matching ID.
	index := -1
	for i, a := range albums {
		if a.ID == id {
			index = i
			break
		}
	}
	//if no album is found
	if index == -1 {
		c.AbortWithStatus(http.StatusNotFound)
	}
	/*delete the album from the slice at the index i.
	albums[:index]: This is a slice that includes all the elements of albums before the album to be deleted.
	It starts at the beginning of the albums slice and goes up to (but not including) the album to be deleted, which has index index.
	albums[index+1:]: This is a slice that includes all the elements of albums after the album to be deleted.
	It starts at the element immediately after the album to be deleted and goes up to the end of the albums slice.
	...: This operator unpacks the slices from steps 1 and 2 into separate arguments for the append() function.
	append(albums[:index], albums[index+1:]...): This concatenates the two slices from steps 1 and 2 
	and returns a new slice that excludes the album to be deleted.
	albums = append(albums[:index], albums[index+1:]...): This assigns the new slice back to the albums variable, 
	effectively removing the album at the specified index from the albums slice.*/
	albums = append(albums[:index], albums[index+1:]...)

	//return a status code of 204 No Content.
	c.Status(http.StatusNoContent)
}

//updateAlbum by id
func updateAlbum(c *gin.Context) {
	id := c.Param("id")

	//find the index(i) of the album with the matching ID.
	index := -1
	for i, a := range albums {
		if a.ID == id {
			index = i
			break
		}
	}

	//if no album is found.
	if index == -1 {
		//AbortWithStatus() is used to immediately abort processing the request and return a specified HTTP status code to the client.
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	//BindJSON to bind the received JSON to newAlbum.
	var updatedAlbum album
	if err := c.BindJSON(&updatedAlbum); err != nil {
		return
	}
	//update the album in the slice at the index i.
	albums[index] = updatedAlbum

	/*IndentedJSON serializes the given struct as pretty JSON (indented + endlines) into the response body.*/
	c.IndentedJSON(http.StatusOK, updatedAlbum)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	//initialize a newAlbum of type album.
	var newAlbum album

	//BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	//add the new album to the slice.
	albums = append(albums, newAlbum)

	/*IndentedJSON serializes the given struct as pretty JSON (indented + endlines) into the response body.
	It also sets the Content-Type as "application/json".*/
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
