package client

import (
	"fmt"
	"net"
	"url-shortener/redis/resp"
	"url-shortener/redis/server"
)

type RedisOptions struct {
	Addr     string
	Password string
	DB       int
}

type RedisClient struct {
	Options *RedisOptions
}

func NewClient(config *RedisOptions) *RedisClient {
	client := RedisClient{config}
	redisServer := server.NewServer("6379")
	go redisServer.Run()
	return &client
}

func (rc *RedisClient) Set(key, value string) (string, error) {
	result, err := rc.send(fmt.Sprintf("%s %s %s", "SET", key, value))
	if err != nil {
		return "", err
	}
	return anyToString(result), nil
}

func (rc *RedisClient) Get(key string) (string, error) {
	result, err := rc.send(fmt.Sprintf("%s %s", "GET", key))
	if err != nil {
		return "", err
	}
	return anyToString(result), nil
}

func (rc *RedisClient) Delete(key string) (int, error) {
	result, err := rc.send(fmt.Sprintf("%s %s", "DEL", key))
	if err != nil {
		return 0, err
	}
	res := result.(int)
	return res, nil
}

func (rc *RedisClient) send(cmd string) (any, error) {
	respCmd, err := resp.Encode(cmd)
	if err != nil {
		return "", err
	}
	conn, err := net.Dial("tcp", rc.Options.Addr)
	defer func(conn net.Conn) {
		_ = conn.Close()
	}(conn)
	if err != nil {
		return "", err
	}
	_, err = conn.Write([]byte(respCmd))
	if err != nil {
		return "", err
	}
	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		return "", err
	}

	response, err := resp.Decode(buf)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// helper method to convert an any value to a string
func anyToString(value any) string {
	switch v := value.(type) {
	case string:
		return v
	default:
		return ""
	}
}
