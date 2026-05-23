package repository

import (
	"artistify/internal/models"
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetAlbums(pool *pgxpool.Pool) ([]models.Albums, error) {

	var albums []models.Albums

	rows, err := pool.Query(
		context.Background(),
		"SELECT id, album_name, artist, sales, rating, created_at FROM albums",
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var album models.Albums

		err = rows.Scan(&album.ID, &album.AlbumName, &album.Artist, &album.Sales, &album.Rating, &album.CreatedAt)

		if err != nil {
			return nil, err
		}

		albums = append(albums, album)
	}

	return albums, nil

}

func GetAlbumByID(pool *pgxpool.Pool, id int) (models.Albums, error) {
	var album models.Albums

	err := pool.QueryRow(
		context.Background(),
		"SELECT id, album_name, artist, sales, rating, created_at FROM albums WHERE id =$1",
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
		return models.Albums{}, err
	}

	return album, nil

}

func GetAlbumsByArtist(pool *pgxpool.Pool, artist string) ([]models.Albums, error) {
	var albums []models.Albums

	rows, err := pool.Query(
		context.Background(),
		"SELECT id, album_name, artist, sales, rating, created_at FROM albums WHERE artist = $1",
		artist)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var album models.Albums
		err := rows.Scan(&album.ID, &album.AlbumName, &album.Artist, &album.Sales, &album.Rating, &album.CreatedAt)
		if err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}

	return albums, nil
}

func PostAlbum(pool *pgxpool.Pool, newAlbum models.Albums) (models.Albums, error) {

	err := pool.QueryRow(
		context.Background(),
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
		return models.Albums{}, err
	}

	return newAlbum, nil
}

func DeleteAlbumByID(pool *pgxpool.Pool, id int) error {
	commandTag, err := pool.Exec(context.Background(),
		"DELETE FROM albums WHERE id = $1",
		id,
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return errors.New("Album Not Found")
	}

	return nil
}

func UpdateAlbumByID(pool *pgxpool.Pool, id int, updatedAlbum models.Albums) (models.Albums, error) {

	err := pool.QueryRow(context.Background(),
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
		return models.Albums{}, err
	}

	return updatedAlbum, nil
}
