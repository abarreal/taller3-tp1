package visitcount

import (
	"fmt"
	"net/http"
	"os"

	"aba.taller3.fi.uba.ar/tp1/visitcounter/components/storage"
	"github.com/gin-gonic/gin"
)

type VisitCountController interface {
	GetTotal(*gin.Context)
}

func CreateVisitCountController(
	repo storage.VisitCounterRepository) VisitCountController {
	return &visitCountController{repo}
}

type visitCountController struct {
	repository storage.VisitCounterRepository
}

func (c *visitCountController) GetTotal(ctx *gin.Context) {
	// Get the total amount of visits from the repository.
	if count, err := c.repository.GetCount(ctx); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"Count": count})
	}
}
