package delivery

import (
	"backend/domain"
	"backend/features/common"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type postHandler struct {
	postUseCase domain.PostUseCase
}

func New(puc domain.PostUseCase) domain.PostHandler {
	return &postHandler{
		postUseCase: puc,
	}
}

// Create implements domain.PostHandler
func (ph *postHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newpost PostFormat
		id := common.ExtractData(c)
		bind := c.Bind(&newpost)

		if bind != nil {
			log.Println("cant bind")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "There is an error in internal server",
			})
		}

		status := ph.postUseCase.CreatePost(newpost.ToModel(), id)

		if status == 400 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    status,
				"message": "Wrong input",
			})
		}

		if status == 404 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    status,
				"message": "Data not found",
			})
		}

		if status == 500 {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    status,
				"message": "There is an error in internal server",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    status,
			"message": "Register success",
		})
	}
}

// Update implements domain.PostHandler
func (ph *postHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newpost PostFormat
		postid := c.Param("postid")
		cnv, err := strconv.Atoi(postid)

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "Wrong input",
			})
		}

		userid := common.ExtractData(c)
		bind := c.Bind(&newpost)

		if bind != nil {
			log.Println("cant bind")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "There is an error in internal server",
			})
		}

		status := ph.postUseCase.UpdatePost(newpost.ToModel(), cnv, userid)

		if status == 400 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    status,
				"message": "Wrong input",
			})
		}

		if status == 404 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    status,
				"message": "Data not found",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    status,
			"message": "Update success",
		})
	}
}