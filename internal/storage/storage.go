package storage

import (
	"github.com/gofiber/storage/memory/v2"
	"sync"
)

var (
	once    sync.Once
	storage *memory.Storage
)

func Init() {
	once.Do(func() {
		storage = memory.New()
	})
}

func GetStorage() *memory.Storage {
	return storage
}
