package repository

import (
	"errors"
	"log"
	"math/rand"
	"psideris/albums-api/domain"
	"strconv"
	"time"
)

var albums []domain.Album

func init() {
	log.Println("Initialising Albums Repo")
	albums = []domain.Album{
		{ID: "1", Title: "Nevermind", Artist: "Nirvana", Price: 20.99},
		{ID: "2", Title: "Smash", Artist: "The Offspring", Price: 19.99},
		{ID: "3", Title: "QOTSA", Artist: "Queens of the Stone Age", Price: 39.99},
	}
}

func FindAllAlbums() *[]domain.Album {
	return &albums
}

func SaveAlbum(baseAlbum domain.BaseAlbum) *domain.Album {
	rand.Seed(time.Now().UnixNano())
	randomId := strconv.Itoa(rand.Intn(10000))
	album := domain.Album{
		ID:     randomId,
		Title:  baseAlbum.Title,
		Artist: baseAlbum.Artist,
		Price:  baseAlbum.Price,
	}

	albums = append(albums, album)
	return &album
}

func UpdateAlbum(id string, baseAlbum domain.BaseAlbum) (*domain.Album, error) {
	idx := findAlbumIndex(id)
	if idx < 0 {
		return nil, errors.New("album id not found")
	}

	album := domain.Album{
		ID:     id,
		Title:  baseAlbum.Title,
		Artist: baseAlbum.Artist,
		Price:  baseAlbum.Price,
	}

	albums[idx] = album
	return &album, nil
}

func DeleteAlbum(id string) {
	idx := findAlbumIndex(id)
	if idx < 0 {
		return
	}
	albums = append(albums[:idx], albums[idx+1:]...)
}

func findAlbumIndex(albumId string) int {
	albumIndex := -1
	for i := range albums {
		if albums[i].ID == albumId {
			albumIndex = i
		}
	}

	return albumIndex
}
