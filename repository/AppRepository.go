package repository

import (
	"errors"
	"url-shortener/model/entity"
	redis "url-shortener/redis/client"
)

type KVStore interface {
	Set(key string, value string) (string, error)
	Get(key string) (string, error)
}

type AppRepository struct {
	Db KVStore
}

func NewAppRepository() *AppRepository {
	var redisClient = redis.NewClient(&redis.RedisOptions{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &AppRepository{Db: redisClient}
}

type IAppRepository interface {
	Save(role *entity.Url) (*entity.Url, error)
	Find(key string) (*entity.Url, error)
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
