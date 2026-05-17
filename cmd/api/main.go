package main

import (
	"fmt"
	"net/http"
	"rsc.io/quote"
	// "strconv"
	// "html/template"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"context"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"artistify/internal/middleware"
	"artistify/internal/models"
	"artistify/internal/database"
)



func register(c *gin.Context) {
	type Register struct {
		Username string `json:"name"`
		Email string `json:"email"`
		Password string `json:"password"`
	}

	var registerInput Register

	err := c.BindJSON(&registerInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var id int
	var createdAt time.Time

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(registerInput.Password), bcrypt.DefaultCost)

	err = conn.QueryRow(context.Background(),
						"INSERT INTO users(username, email, password, type) VALUES($1, $2, $3, $4) RETURNING id, created_at",
						registerInput.Username, 
						registerInput.Email,
						string(hashedPassword),
						"user",
						).Scan(
							&id,
							&createdAt,
						)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":"User Registered",
		"User id":id,
		"Created at":createdAt,
	})

}

func registerWithRole(c *gin.Context) {
	type Register struct {
		Username string `json:"name"`
		Email string `json:"email"`
		Password string `json:"password"`
		Type int `json:"type"`
	}

	var registerInput Register

	err := c.BindJSON(&registerInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var id int
	var createdAt time.Time

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(registerInput.Password), bcrypt.DefaultCost)

	var Role string

	if registerInput.Type == 0 {
		Role = "admin"
	} else if registerInput.Type == 1 {
		Role = "artist"
	} else {
		Role = "user"
	}

	err = conn.QueryRow(context.Background(),
						"INSERT INTO users(username, email, password, type) VALUES($1, $2, $3, $4) RETURNING id, created_at",
						registerInput.Username, 
						registerInput.Email,
						string(hashedPassword),
						Role,
						).Scan(
							&id,
							&createdAt,
						)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":"User Registered",
		"User id":id,
		"Created at":createdAt,
		"Acces:": Role,
	})

}

func login(c *gin.Context) {
	type loginInput struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

	var input loginInput

	err := c.BindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Invalid Input"})
		return
	}

	var id int
	var password string
	var name string
	var role string

	err = conn.QueryRow(context.Background(),
						"SELECT id, username, password, type FROM users WHERE email = $1",
						input.Email,
					).Scan(
						&id,
						&name,
						&password,
						&role,
					)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(password), 
		[]byte(input.Password),
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Wrong Password"})
		return
	}

	expires := time.Now().Add(time.Hour)
	claims := jwt.MapClaims{
		"id": id,
		"role": role,
		"exp": expires.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("super-secret-key"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"could not create token"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"token": tokenString})
}

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"quote": quote.Go(),
	})
}

func getAlbums(c *gin.Context) {
	rows, err := conn.Query(context.Background(), "SELECT * FROM albums")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(),})
		return
	}
	
	defer rows.Close()

	var albums []models.Albums
	for rows.Next() {
		var album models.Albums
		err := rows.Scan(&album.ID, &album.AlbumName, &album.Artist, &album.Sales, &album.Rating, &album.CreatedAt)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(),})
			return
		}

		albums = append(albums, album)
	}

	c.JSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var album models.Albums
	err := conn.QueryRow(
		context.Background(),
		"SELECT * FROM albums WHERE id =$1", 
		id,
		).Scan(
			&album.ID, 
			&album.AlbumName, 
			&album.Artist, 
			&album.Sales, 
			&album.Rating,
			&album.CreatedAt,
		)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Album not found" })
		return
	}
	c.JSON(http.StatusOK, album)
}

func getAlbumsByArtist(c *gin.Context) {
	artist := c.Param("artist")
	var albums []models.Albums
	rows, err := conn.Query(context.Background(),
						"SELECT id, album_name, artist, sales, rating, created_at FROM albums WHERE artist = $1", 
						artist)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// log.Fatal(err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var album models.Albums
		err := rows.Scan(&album.ID, &album.AlbumName, &album.Artist, &album.Sales, &album.Rating, &album.CreatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			// log.Fatal(err)
			return
		}
		albums = append(albums, album)
	}

	c.JSON(http.StatusOK, albums)
}

func updateAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var updatedAlbum models.Albums
	err := c.BindJSON(&updatedAlbum) 

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Album not found"})
		return
	}

	err = conn.QueryRow(context.Background(),
						"UPDATE albums SET album_name = $2, artist = $3, sales = $4, rating = $5 WHERE id = $1 RETURNING id, created_at", 
						id, 
						updatedAlbum.AlbumName, 
						updatedAlbum.Artist, 
						updatedAlbum.Sales, 
						updatedAlbum.Rating,
						).Scan(
							&updatedAlbum.ID,
							&updatedAlbum.CreatedAt,
						)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can't update album right now... Try again later."})
		return
	}
	c.JSON(http.StatusCreated, updatedAlbum)
	
}

func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	commandTag, err := conn.Exec(context.Background(),
					"DELETE FROM albums WHERE id = $1",
					id,
				)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if commandTag.RowsAffected() == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Album not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"Album delete Successfully"})
	return
}

func addAlbum(c *gin.Context) {
	var newAlbum models.Albums
	err := c.BindJSON(&newAlbum)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	err = conn.QueryRow(context.Background(), 
						"INSERT INTO albums (album_name, artist, sales, rating) VALUES($1, $2, $3, $4) RETURNING id, created_at", 
						newAlbum.AlbumName, 
						newAlbum.Artist, 
						newAlbum.Sales,
						newAlbum.Rating,
						).Scan(
							&newAlbum.ID,
							&newAlbum.CreatedAt,
						)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Not Able to add album"})
		return
	}
	c.JSON(http.StatusCreated, newAlbum)
}


var conn *pgx.Conn
func main() {

	var err error

	conn, err = database.ConnectDB()

	if err != nil {
		panic(err)
	}

	defer conn.Close(context.Background())

	fmt.Println("DATABASE Successfully connected!!!")

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	
	fmt.Println("Running server at 8080...", quote.Go())

	ipRateLimiter := middleware.NewIPRateLimiter()
	
	router.GET("/", home)
	router.POST("/auth/login", middleware.RateLimitMiddleware(ipRateLimiter), login)
	router.POST("/auth/register", register)
	router.POST("/auth/registerWithRole", registerWithRole)
	router.GET("/albums", getAlbums)
	router.GET("/albums/artist/:artist", getAlbumsByArtist)
	router.GET("/albums/:id", getAlbumByID)
	

	artist := router.Group("/artist")
	artist.Use(middleware.AuthMiddleware())
	artist.Use(middleware.RequireRole("artist")) 
	{
		artist.POST("/albums", addAlbum)
		artist.PUT("/albums/:id", updateAlbumByID)
		artist.DELETE("/albums/:id", deleteAlbumByID)
	}

	admin := router.Group("/admin")
	admin.Use(middleware.AuthMiddleware())
	admin.Use(middleware.RequireRole("admin")) 
	{
		admin.POST("/albums", addAlbum)
		admin.PUT("/albums/:id", updateAlbumByID)
		admin.DELETE("/albums/:id", deleteAlbumByID)
	}


	// artist.GET("/albums/stats/:id", albumStatsByID)



	router.Run(":8080")
}