package utils

import (
	"github.com/run-bigpig/conduit/internal/storage"
	"time"
)

// SetData 设置缓存
func SetData(key string, data []byte, expire int) error {
	return storage.GetStorage().Set(key, data, time.Duration(expire)*time.Second)
}

// GetData 获取缓存
func GetData(key string) ([]byte, error) {
	dataByte, err := storage.GetStorage().Get(key)
	if err != nil {
		return nil, err
	}
	return dataByte, nil
}

func DelData(key string) error {
	return storage.GetStorage().Delete(key)
}
