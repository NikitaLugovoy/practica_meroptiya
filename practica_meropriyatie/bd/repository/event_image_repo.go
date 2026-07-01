// repository/event_image_repository.go
package repository

import (
	"context"
	"fmt"
	"practica_meropriyatie/bd/model"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InsertNewEventImageInBD(pool *pgxpool.Pool, image *model.EventImage) (int, error) {
	rowInsert := pool.QueryRow(context.Background(),
		"INSERT INTO event_images(event_id, url_images) VALUES ($1,$2) RETURNING id",
		image.GetEvent_id(),
		image.GetUrl_image(),
	)
	var idInsert int
	if err := rowInsert.Scan(&idInsert); err != nil {
		return 0, err
	}
	return idInsert, nil
}

func SelectFullEventImages(pool *pgxpool.Pool) ([]model.EventImage, error) {

	var url_image string
	var id, event_id int
	var created_at time.Time

	rows, err := pool.Query(context.Background(),
		"SELECT id, event_id, url_images, created_at FROM event_images")

	var images []model.EventImage
	for rows.Next() {
		if err := rows.Scan(&id, &event_id, &url_image, &created_at); err != nil {
			return nil, err
		}
		image := model.NewEventImage(uint(id), event_id, url_image, created_at)
		images = append(images, *image)
	}
	rows.Close()
	fmt.Println(err)
	return images, nil
}

func SelectEventImagesByEventId(pool *pgxpool.Pool, event_id int) ([]model.EventImage, error) {

	var url_image string
	var id int
	var created_at time.Time

	rows, err := pool.Query(context.Background(),
		"SELECT id, event_id, url_images, created_at FROM event_images WHERE event_id = $1", event_id)

	var images []model.EventImage
	for rows.Next() {
		if err := rows.Scan(&id, &event_id, &url_image, &created_at); err != nil {
			return nil, err
		}
		image := model.NewEventImage(uint(id), event_id, url_image, created_at)
		images = append(images, *image)
	}
	rows.Close()
	fmt.Println(err)
	return images, nil
}

func SelectEventImagesByImageId(pool *pgxpool.Pool, image_id int) ([]model.EventImage, error) {

	var url_image string
	var id, event_id int
	var created_at time.Time

	rows, err := pool.Query(context.Background(),
		"SELECT id, event_id, url_images, created_at FROM event_images WHERE id = $1", image_id)

	var images []model.EventImage
	for rows.Next() {
		if err := rows.Scan(&id, &event_id, &url_image, &created_at); err != nil {
			return nil, err
		}
		image := model.NewEventImage(uint(id), event_id, url_image, created_at)
		images = append(images, *image)
	}
	rows.Close()
	fmt.Println(err)
	return images, nil
}

func DeleteOneEventImage(pool *pgxpool.Pool, ids int) (int64, error) {

	result, err := pool.Exec(context.Background(), "DELETE FROM event_images WHERE id = $1", ids)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), err
}

func UpdateOneEventImage(pool *pgxpool.Pool, ids int, image *model.EventImage) int64 {

	result, err := pool.Exec(context.Background(),
		"UPDATE event_images SET event_id=$1, url_images=$2 WHERE id = $3",
		image.GetEvent_id(),
		image.GetUrl_image(),
		ids)
	if err != nil {
		return 0
	}
	return result.RowsAffected()
}
