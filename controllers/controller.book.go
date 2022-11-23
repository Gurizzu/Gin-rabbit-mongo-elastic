package controllers

import (
	"gin_es-rabbit/models"
	"gin_es-rabbit/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

type BookController struct {
	router  *gin.RouterGroup
	service *services.BookServices
}

func NewBookController(router *gin.RouterGroup) *BookController {
	this := &BookController{router: router, service: services.NewBookServices()}

	book := this.router.Group("/book")
	book.POST("/create", this.Create)
	book.GET("/getOne/:book_id", this.GetOne)

	return this
}

// Create godoc
// @Summary     Add Book
// @Description  add book
// @Tags         Book
// @Accept       json
// @Param request body models.BookInput true "query params"
// @Produce      json
// @Success      200  {object}  models.Book
// @Router       /api/v1/book/create [post]
func (this *BookController) Create(c *gin.Context) {
	resp := models.Response{}
	var book models.Book
	if err := c.BindJSON(&book); err != nil {
		log.Fatal(err)
		return
	}

	book.ID = primitive.NewObjectID()
	book.Created_At = time.Now().Local().Unix()
	book.Updated_At = time.Now().Local().Unix()
	resp = this.service.InsertOne(book)
	c.JSON(200, resp.Data)

}

// GetOne godoc
// @Summary     get Book
// @Description  get book
// @Tags         Book
// @Accept       json
// @Param        id   path      string  true  "Note ID"
// @Produce      json
// @Success      200  {object}  models.Book
// @Router       /api/v1/book/getOne/{id} [get]
func (this *BookController) GetOne(c *gin.Context) {
	resp := models.Response{}
	resp = this.service.GetOne("_id", c.Param("book_id"))
	c.JSON(200, resp.Data)

}
