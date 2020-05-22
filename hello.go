 package main
 import(
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)


type Joke struct {
	ID     int     `json:"id" binding:"required"`
	Likes  int     `json:"likes"`
	Joke   string  `json:"joke" binding:"required"`
  }


  var jokes = []Joke{
	Joke{1, 0, "Did you hear about the restaurant on the moon? Great food, no atmosphere."},
	Joke{2, 0, "What do you call a fake noodle? An Impasta."},
	Joke{3, 0, "How many apples grow on a tree? All of them."},
	Joke{4, 0, "Want to hear a joke about paper? Nevermind it's tearable."},
	Joke{5, 0, "I just watched a program about beavers. It was the best dam program I've ever seen."},
	Joke{6, 0, "Why did the coffee file a police report? It got mugged."},
	Joke{7, 0, "How does a penguin build it's house? Igloos it together."},
  }
  

var router *gin.Engine

 func main() {
	router := gin.Default()
	router.Use(cors.Default())
	// v1 := r.Group("api/v1")
	// {
	// 	// v1.POST("/users",PostUser)
	// 	v1.GET("/house",GetHouse)
	// }
	router.GET("/",GetMessgae )
	router.GET("/test",GetHouse)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}


func GetMessgae (c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}


func GetHouse(c *gin.Context) {
	jsonFile, err := os.Open("./page1.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)

    var result map[string]interface{}
    json.Unmarshal([]byte(byteValue), &result)
	
	c.Header("Content-Type", "application/json")
    c.JSON(200, result)

    // curl -i http://localhost:8080/api/v1/users
}