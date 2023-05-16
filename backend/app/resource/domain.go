package resource

type SendCurrentGamePOSTAPI struct {
	AreaCode int    `json:"area_code" binding:"required"`
	GameMode string `json:"game_mode"  binding:"required"`
	Playing  int    `json:"playing"`
}

type GetMostPlayedGameGETAPI struct {
	AreaCode int `json:"area_code" binding:"required"`
}

var GAME_MODES = []string{"battle_royale", "team_deathmatch", "capture_the_flag"}
