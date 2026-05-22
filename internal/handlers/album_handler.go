package handlers

import (
	"artistify/internal/models"
	"artistify/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
	"net/http"
	"strconv"
)

// GET REQUESTS FUNCTION

func GetAlbums(conn *pgx.Conn, rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		albums, source, err := service.GetAlbums(conn, rdb)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"source": source, "data": albums})
	}
}

func GetAlbumByID(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {

		Stringid := c.Param("id")
		id, err := strconv.Atoi(Stringid)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}

		album, err := service.GetAlbumByID(conn, id)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, album)
	}
}

func GetAlbumsByArtist(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {

		artist := c.Param("artist")

		albums, err := service.GetAlbumsByArtist(conn, artist)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, albums)
	}
}

// POST REQUESTS FUNCTION HERE

func PostAlbum(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {

		var newAlbum models.Albums
		err := c.BindJSON(&newAlbum)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newAlbum, err = service.PostAlbum(conn, newAlbum)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, newAlbum)
	}
}

// PUT REQUESTS FUNCTION HERE

func UpdateAlbumByID(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {

		StringID := c.Param("id")
		id, err := strconv.Atoi(StringID)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		var updatedAlbum models.Albums
		err = c.BindJSON(&updatedAlbum)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Album not found"})
			return
		}

		updatedAlbum, err = service.UpdateAlbumByID(conn, id, updatedAlbum)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		c.JSON(http.StatusCreated, updatedAlbum)

	}
}

// DELETE REQUESTS FUNCTION HERE

func DeleteAlbumByID(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {

		Stringid := c.Param("id")
		id, err := strconv.Atoi(Stringid)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}

		err = service.DeleteAlbumByID(conn, id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Album delete Successfully"})
	}
}
