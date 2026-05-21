package service

import (
	"artistify/internal/repository"
	"artistify/internal/database"
	"artistify/internal/models"
	"encoding/json"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
	"time"
)


func GetAlbums(conn *pgx.Conn, rdb *redis.Client) ([]models.Albums, string, error) {
	cacheKey := "albums:all"
		
	cachedAlbums, err := rdb.Get(database.Ctx, cacheKey).Result()

	var albums []models.Albums
		
	if err == nil {
		json.Unmarshal([]byte(cachedAlbums), &albums)
		return albums, "Redis", err
	}

	albums, err = repository.GetAlbums(conn)

	if err != nil {
		return nil, "Postgres", err
	}

	jsonData, _ := json.Marshal(albums)

	err = rdb.Set(database.Ctx, cacheKey, jsonData, 30*time.Minute).Err()

	return albums, "Postgres", nil

}

func GetAlbumByID(conn *pgx.Conn, id int) (models.Albums, error) {
	album, err := repository.GetAlbumByID(conn, id)

	if err != nil {
		return models.Albums{}, err
	}

	return album, nil
}

func GetAlbumsByArtist(conn *pgx.Conn, artist string) ([]models.Albums, error) {
		albums, err := repository.GetAlbumsByArtist(conn, artist)

		if err != nil {
			return nil, err
		}
		
		return albums, nil
}

func PostAlbum(conn *pgx.Conn, newAlbum models.Albums) (models.Albums, error) {

		newAlbum, err := repository.PostAlbum(conn, newAlbum)
		
		if err != nil {
			return models.Albums{}, err
		}

		return newAlbum, nil
}

func DeleteAlbumByID(conn *pgx.Conn, id int) (error) {

		err := repository.DeleteAlbumByID(conn, id)
		
		if err != nil {
			return err
		}
		
		return nil
}

func UpdateAlbumByID(conn *pgx.Conn, id int, updatedAlbum models.Albums) (models.Albums, error) {

	updatedAlbum, err := repository.UpdateAlbumByID(conn, id, updatedAlbum)

	if err != nil {
		return models.Albums{}, err
	}

	return updatedAlbum, nil
}