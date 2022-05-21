package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"psideris/albums-api/domain"
	"psideris/albums-api/repository"
)

const (
	AlbumsRoute       = "/albums"
	AlbumsWithIdRoute = "/albums/:id"
)

type apiError struct {
	Code    int    `json:"code"`
	Details string `json:"details"`
}

func ConfigureRoutes(router gin.IRouter) {
	router.GET(AlbumsRoute, getAlbums)
	router.POST(AlbumsRoute, addAlbum)
	router.PUT(AlbumsWithIdRoute, updateAlbum)
	router.DELETE(AlbumsWithIdRoute, deleteAlbum)
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, repository.FindAllAlbums())
}

func addAlbum(c *gin.Context) {
	var album domain.BaseAlbum
	if err := c.ShouldBindJSON(&album); err != nil {
		c.IndentedJSON(http.StatusBadRequest, apiError{Code: 400, Details: "Invalid body"})
		return
	}
	c.IndentedJSON(http.StatusCreated, repository.SaveAlbum(album))
}

func updateAlbum(c *gin.Context) {
	var album domain.BaseAlbum
	if err := c.ShouldBindJSON(&album); err != nil {
		c.IndentedJSON(http.StatusBadRequest, apiError{Code: 400, Details: "Invalid body"})
		return
	}
	id := c.Param("id")
	savedAlbum, err := repository.UpdateAlbum(id, album)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, apiError{Code: 400, Details: "Invalid album id"})
		return
	}
	c.IndentedJSON(http.StatusCreated, savedAlbum)
}

func deleteAlbum(c *gin.Context) {
	id := c.Param("id")
	repository.DeleteAlbum(id)
	c.Status(http.StatusNoContent)
}
