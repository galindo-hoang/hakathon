package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DetectImage(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	files := form.File["updaload[]"]
	for _, f := range files {
		log.Printf("name: %v || size: %v\n", f.Filename, f.Size)
	}
	ctx.String(http.StatusAccepted, fmt.Sprintf("%d files uploaded", len(files)))
}
