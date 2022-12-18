package main

import (
	"github.com/Depado/ginprom"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	testdata "go-api-gin-swagger/data"
	docs "go-api-gin-swagger/docs"
	model "go-api-gin-swagger/model"
	"net/http"
	"os"
)

var albums = testdata.Testalbums
var serviceCodes = testdata.TestserviceCodes

//// albums slice to seed record album data.
//var albums = []model.Album{
//	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
//	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
//	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
//}
//
//var serviceCodes = []model.ServiceCode{
//	{ID: 1, Name: "SVC01", Description: "super svc01", BusinessUnit: "BFI"},
//	{ID: 2, Name: "SVC02", Description: "super svc02", BusinessUnit: "BSX"},
//	{ID: 3, Name: "SVC03", Description: "super svc03", BusinessUnit: "BSS"},
//}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	// TextFormatter JSONFormatter
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

// @title           Trinity Metadata API
// @version         1.0
// @description     This provides properties access for cloud automation
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {

	r := gin.Default()
	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOrigins = []string{"*"}
	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true

	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS")

	p := ginprom.New(
		ginprom.Engine(r),
		ginprom.Subsystem("gin"),
		ginprom.Path("/metrics"),
	)
	r.Use(p.Instrument())
	// Register the middleware
	r.Use(cors.New(corsConfig))
	// prometheus metrics

	// register swagger
	docs.SwaggerInfo.BasePath = "/api/v1"
	log.WithFields(log.Fields{
		"swagger": "docs",
	}).Info("http://localhost:8080/swagger/index.html")
	log.WithFields(log.Fields{
		"prometheus": "metrics",
	}).Info("http://localhost:8080/metrics")

	r.GET("/hello-world", Helloworld)

	v1 := r.Group("/api/v1")
	{
		group := v1.Group("/alb")
		{
			group.GET("/albums", getAlbums)
			group.GET("/albums/:id", getAlbumByID)
			group.POST("/albums", postAlbums)
		}
		example := v1.Group("/example")
		{
			example.GET("/hello-world", Helloworld)
		}
		svc := v1.Group("/generic")
		{
			svc.GET("/svc", getServiceCodes)
			svc.POST("/svc", postServiceCode)
		}
		//...
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}

// Get all Albums godoc
// @Summary ping example
// @Schemes
// @Description return a list of Albums
// @Tags album
// @Accept json
// @Produce json
// @Success 200 {array} model.Album
// @Router /alb/albums [get]
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// Post a Album godoc
// @Summary ping example
// @Schemes
// @Description return a list of Albums
// @Tags album
// @Accept json
// @Produce json
// @Success 201 {array} model.Album
// @Router /alb/albums [post]
// @Param   collection  query     []string   false  "string collection"  collectionFormat(multi)
func postAlbums(c *gin.Context) {
	var newAlbum model.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// ShowAccount godoc
// @Summary      Show an album
// @Description  get string by ID
// @Tags         albums
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Album ID"
// @Success      200  {object}  model.Album
// @Router       /alb/albums/{id} [get]
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// Get all ServiceCodes godoc
// @Summary ping example
// @Schemes
// @Description return a list of Albums
// @Tags svc
// @Accept json
// @Produce json
// @Success 200 {array} model.ServiceCode
// @Router /generic/svc [get]
func getServiceCodes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, serviceCodes)
}

// Get all ServiceCode godoc
// @Summary ping example
// @Schemes
// @Description return a list of Albums
// @Tags svc
// @Accept json
// @Produce json
// @Success 200 {array} model.ServiceCode
// @Router /generic/svc [post]
func postServiceCode(c *gin.Context) {
	var newServiceCode model.ServiceCode

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newServiceCode); err != nil {
		return
	}

	// Add the new album to the slice.
	serviceCodes = append(serviceCodes, newServiceCode)
	c.IndentedJSON(http.StatusCreated, newServiceCode)
}

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}
