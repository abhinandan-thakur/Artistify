package handlers

import (
	"github.com/abhinandan-thakur/Artistify/music-service/internal/models"
	"github.com/abhinandan-thakur/Artistify/music-service/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"net/http"
	"strconv"
)

type GetAlbumsResponse struct {
	Source string `json:"source"`
	Data []models.Albums `json:"data"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

// GET REQUESTS FUNCTION

// GetAlbums godoc
//
// @Summary Get all albums
// @Description Get All albums inside the database
// @Tags Music
// 
// @Produce json
//
// @Success 200 {object} GetAlbumsResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
//
// @Router /albums [get]


func GetAlbums(pool *pgxpool.Pool, rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		albums, source, err := service.GetAlbums(pool, rdb)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"source": source, "data": albums})
	}
}

// GetAlbumByID godoc
//
// @Summary Get an album by its ID
// @Description Get an album from the database based on their id
// @Tags Music
// 
// @Produce json
//
// @Param id path int true "Album ID"
//
// @Success 200 {object} models.Albums
// @Failure 404 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
//
// @Router /albums/{id} [get]
func GetAlbumByID(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {

		Stringid := c.Param("id")
		id, err := strconv.Atoi(Stringid)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}

		album, err := service.GetAlbumByID(pool, id)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, album)
	}
}

// GetAlbumsByArtist godoc
//
// @Summary get albums by artist
// @Description get albums of an artist in the database
// @Tags Music
// 
// @Produce json
//
// @Param artist path string true "Artist name"
//
// @Success 200 {array} models.Albums
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
//
// @Router /albums/{artist} [get]
func GetAlbumsByArtist(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {

		artist := c.Param("artist")

		albums, err := service.GetAlbumsByArtist(pool, artist)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, albums)
	}
}

// POST REQUESTS FUNCTION HERE

// PostAlbum godoc
//
// @Summary Add Album
// @Description Add an album inside the database
// @Tags Music
// 
// @Accept json
// @Produce json
// 
// @Param album body models.Albums true "Album"
//
// @Success 201 {object} models.Albums
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
//
// @Router /albums [post]
func PostAlbum(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {

		var newAlbum models.Albums
		err := c.BindJSON(&newAlbum)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newAlbum, err = service.PostAlbum(pool, newAlbum)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, newAlbum)
	}
}

// PUT REQUESTS FUNCTION HERE

// UpdateAlbumByID godoc
//
// @Summary Update Album by ID
// @Description Update Album by ID
// @Tags Music
// 
// @Accept json
// @Produce json
//
// @Param id path int true "Album id"
// @Param album body models.Albums true "Updated Album"
//
// @Success 200 {object} models.Albums
// @Failure 404 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
//
// @Router /albums/{id} [put]
func UpdateAlbumByID(pool *pgxpool.Pool) gin.HandlerFunc {
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

		updatedAlbum, err = service.UpdateAlbumByID(pool, id, updatedAlbum)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, updatedAlbum)

	}
}

// DELETE REQUESTS FUNCTION HERE

// DeleteAlbumByID godoc
//
// @Summary Delete an Album
// @Description Delete an Album from the database
// @Tags Music
// 
// @Produce json
//
// @Param id path int true "Album id"
// 
// @Success 200 {object} MessageResponse 
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
//
// @Router /albums/{id} [delete]
func DeleteAlbumByID(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {

		Stringid := c.Param("id")
		id, err := strconv.Atoi(Stringid)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}

		err = service.DeleteAlbumByID(pool, id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Album delete Successfully"})
	}
}
