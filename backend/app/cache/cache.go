package cache

import (
	"LilaGames/services"
	"context"
	"errors"
	"log"
	"sync"
)

var logger = log.Default()

type LocalCache struct {
	mu                      sync.RWMutex
	areaCodeGameModeCounter map[int]map[string]int
}

func NewCache() *LocalCache {
	lc := &LocalCache{
		areaCodeGameModeCounter: make(map[int]map[string]int),
	}
	return lc
}

func (lc *LocalCache) UpdateCounter(areaCode int, gameMode string) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	_, ok := lc.areaCodeGameModeCounter[areaCode]
	if !ok {
		lc.areaCodeGameModeCounter[areaCode] = make(map[string]int)
	}

	_, ok = lc.areaCodeGameModeCounter[areaCode][gameMode]
	if !ok {
		lc.areaCodeGameModeCounter[areaCode][gameMode] = 0
	}

	lc.areaCodeGameModeCounter[areaCode][gameMode] += 1
}

func (lc *LocalCache) GetCounter(areaCode int) (map[string]int, error) {
	lc.mu.RLock()
	defer lc.mu.RUnlock()

	_, ok := lc.areaCodeGameModeCounter[areaCode]
	if !ok {
		return nil, errors.New("no active games in area")
	}

	return lc.areaCodeGameModeCounter[areaCode], nil
}

func (lc *LocalCache) WriteToDB() {
	gameCollection := services.GetCollection(services.DB, "lilaGames")

	lc.mu.RLock()
	defer lc.mu.RUnlock()

	logger.Println(lc.areaCodeGameModeCounter)
	_, err := gameCollection.InsertOne(context.Background(), lc.areaCodeGameModeCounter)
	if err != nil {
		log.Fatal("failed writing to db")
	}
}
