package resource

import (
	"LilaGames/cache"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

var logger = log.Default()

type GameResource struct {
	localCache *cache.LocalCache
}

func NewGameHandler() GameResource {
	gameCache := cache.NewCache()
	handler := GameResource{localCache: gameCache}

	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for ; ; <-ticker.C {
			logger.Println("Writing to DB")
			handler.localCache.WriteToDB()
			logger.Println("Completed writing to DB")
		}
	}()
	return handler
}

func (gc *GameResource) GetPopularGameMode(ctx *gin.Context) {
	var data GetMostPlayedGameGETAPI

	err := ctx.BindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "required data not passed",
			"error":   err.Error(),
		})
		return
	}

	if ok := ValidateAreaCode(data.AreaCode); !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid area code, should be of three digits",
		})
		return
	}

	areaCodeData, err := gc.localCache.GetCounter(data.AreaCode)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	maxPlayers := 0
	mostPlayedMode := ""

	for _, gameMode := range GAME_MODES {
		_, ok := areaCodeData[gameMode]
		if ok {
			if areaCodeData[gameMode] > maxPlayers {
				mostPlayedMode = gameMode
				maxPlayers = areaCodeData[gameMode]
			}
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": mostPlayedMode,
	})
}

func (gc *GameResource) WriteGameMode(ctx *gin.Context) {
	var data SendCurrentGamePOSTAPI

	err := ctx.BindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "required data not passed",
			"error":   err.Error(),
		})
		return
	}

	if ok := ValidateAreaCode(data.AreaCode); !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid area code, should be of three digits",
		})
		return
	}

	ok := ValidateGameMode(data.GameMode)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid game mode",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
	go gc.localCache.UpdateCounter(data.AreaCode, data.GameMode)
}

func ValidateGameMode(key string) bool {
	for _, val := range GAME_MODES {
		if val == key {
			return true
		}
	}
	return false
}

func ValidateAreaCode(code int) bool {
	if code/1000 != 0 || code < 100 {
		return false
	}
	return true
}
