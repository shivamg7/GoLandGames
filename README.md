## Get started
#### Pre-requisites
1. Containerization software like podman/docker installed

#### Run the application
1. Set working directory to LilaGames
2. Source .env file `source .env`
3Run `docker-compose up -d --build` or `podman-compose up -d --build`

##### Assumptions & Improvements
1. The application doesn't take into account that a game could end.
2. Application doesn't have persistence storage being loaded back when application restarts.

#### How does the application scale
1. The API is written in GO, which in itself is every fast.
2. In-memory caching helps fetch records faster.
3. The POST API makes use of go coroutines to update the counter async & not let the client block.
4. To scale this further a distributed cache service (like Redis) could have been used
   The backend service could be horizontally scaled by adding multiple containers in the same cluster.
   APIs could be routed to specific part of the cluster depending upon their areacode.


#### Documentation of the webservice APIs
1. POST /games
Payload
```json
{
  "area_code": 123,
  "game_mode": "battle_royale"
}
```

Expected Response: Status code 200
```azure
{
"message": "OK"
}
```

2. GET /games
Payload
```json
{
  "area_code": 123
}
```
Expected Payload
```json
{
  "message": "battle_royale"
}
```