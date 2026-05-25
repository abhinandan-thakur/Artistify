package service

import (
	"github.com/abhinandan-thakur/Artistify/music-service/internal/database"
	"github.com/abhinandan-thakur/Artistify/music-service/internal/models"
	"github.com/abhinandan-thakur/Artistify/music-service/internal/repository"
	"encoding/json"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

func GetAlbums(pool *pgxpool.Pool, rdb *redis.Client) ([]models.Albums, string, error) {
	log.Println("Entered Sevice/GetAlbums()")
	cacheKey := "albums:all"
	var albums []models.Albums

	// Try Redis only if client exists
	if rdb != nil {

		cachedAlbums, err := rdb.Get(database.Ctx, cacheKey).Result()

		if err == nil {

			err = json.Unmarshal([]byte(cachedAlbums), &albums)

			if err == nil {
				return albums, "Redis", nil
			}
		}
	}

	log.Println("going to repository/GetAlbums()")
	albums, err := repository.GetAlbums(pool)

	if err != nil {
		return nil, "Postgres", err
	}

	jsonData, _ := json.Marshal(albums)

	err = rdb.Set(database.Ctx, cacheKey, jsonData, 30*time.Minute).Err()

	if err != nil {
		log.Println("failed to cache data:", err)
	}
	log.Println("function ended")
	return albums, "Postgres", nil

}

func GetAlbumByID(pool *pgxpool.Pool, id int) (models.Albums, error) {
	album, err := repository.GetAlbumByID(pool, id)

	if err != nil {
		return models.Albums{}, err
	}

	return album, nil
}

func GetAlbumsByArtist(pool *pgxpool.Pool, artist string) ([]models.Albums, error) {
	albums, err := repository.GetAlbumsByArtist(pool, artist)

	if err != nil {
		return nil, err
	}

	return albums, nil
}

func PostAlbum(pool *pgxpool.Pool, newAlbum models.Albums) (models.Albums, error) {

	newAlbum, err := repository.PostAlbum(pool, newAlbum)

	if err != nil {
		return models.Albums{}, err
	}

	return newAlbum, nil
}

func DeleteAlbumByID(pool *pgxpool.Pool, id int) error {

	err := repository.DeleteAlbumByID(pool, id)

	if err != nil {
		return err
	}

	return nil
}

func UpdateAlbumByID(pool *pgxpool.Pool, id int, updatedAlbum models.Albums) (models.Albums, error) {

	updatedAlbum, err := repository.UpdateAlbumByID(pool, id, updatedAlbum)

	if err != nil {
		return models.Albums{}, err
	}

	return updatedAlbum, nil
}
