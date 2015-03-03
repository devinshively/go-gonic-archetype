package api

import (
	"github.com/devinshively/gogonic-archetype/database"
	"github.com/devinshively/gogonic-archetype/database/postgres"
	"github.com/devinshively/gogonic-archetype/model"
    "github.com/gin-gonic/gin"
)

var DocumentRoutes []model.Route
var db database.DocumentDataStore

func init() {

	DocumentRoutes = []model.Route{
		{
			Method:  "GET",
			Path:    "/documents",
			Handler: getAllDocumentsHandler,
		}, {
			Method:  "POST",
			Path:    "/documents",
			Handler: createOrUpdateDocumentHandler,
		}, {
			Method:  "GET",
			Path:    "/documents/:id",
			Handler: getDocumentHandler,
		}, {
			Method:  "DELETE",
			Path:    "/documents/:id",
			Handler: deleteDocumentHandler,
		},
	}

	// Set datastore
	db = new(postgres.PostgresDocumentDB)
}

func getAllDocumentsHandler(c *gin.Context) {
	docs, err := db.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error":err.Error()})
	} else {
		c.JSON(200, docs)
	}
}

func createOrUpdateDocumentHandler(c *gin.Context) {
	var doc model.Document
	c.Bind(&doc)

	if err := db.CreateOrUpdate(&doc); err != nil {
		c.JSON(500, gin.H{"error":err.Error()})
	} else {
		c.JSON(200, doc)
	}
}

func getDocumentHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	doc, err := db.Get(id)
	if err != nil {
		c.JSON(500, gin.H{"error":err.Error()})
	} else {
		c.JSON(200, doc)
	}
}

func deleteDocumentHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	if err := db.Delete(id); err != nil {
		c.JSON(500, gin.H{"error":err.Error()})
	} else {
		c.JSON(204,nil)
	}
}
