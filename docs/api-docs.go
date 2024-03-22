package docs

import (
	"github.com/Kotodian/go-redoc"
	ginredoc "github.com/Kotodian/go-redoc/gin"
	"github.com/gin-gonic/gin"
)

func InitApiDocs() gin.HandlerFunc {
	doc := redoc.Redoc{
		Title:       "Tile",
		Description: "Description",
		SpecFile:    "./docs/swagger.yml",
		SpecPath:    "/docs/swagger.yml",
		DocsPath:    "/docs",
	}

	return ginredoc.New(doc)
}
