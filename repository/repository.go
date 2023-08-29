package repository

import (
	"errors"
	"url-shortener/model/entity"
)

type KVStore interface {
	Set(key string, value string) (string, error)
	Get(key string) (string, error)
	Delete(key string) (int, error)
}

type AppRepository struct {
	Db KVStore
}

func NewAppRepository(db KVStore) *AppRepository {
	return &AppRepository{Db: db}
}

type IAppRepository interface {
	Save(role *entity.Url) (*entity.Url, error)
	Find(key string) (*entity.Url, error)
	Delete(key string) (int, error)
}

func (r *AppRepository) Save(uri *entity.Url) (*entity.Url, error) {
	_, err := r.Db.Set(uri.Hash, uri.LongUrl)
	if err != nil {
		return nil, err
	}
	return uri, nil
}

func (r *AppRepository) Find(hash string) (*entity.Url, error) {
	value, err := r.Db.Get(hash)
	if err != nil {
		return nil, err
	}
	if value == "" {
		return nil, errors.New("error key not found")
	}
	uri := entity.Url{LongUrl: value, Hash: value}
	return &uri, nil
}

func (r *AppRepository) Delete(hash string) (int, error) {
	value, err := r.Db.Delete(hash)
	if err != nil {
		return 0, err
	}
	return value, nil
}
