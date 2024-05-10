package config

type serverConfig struct {
	SERVER_PORT              string `json:"SERVER_PORT" validate:"required"`
	WEBSOCKET_SERVER_PORT    string `json:"WEBSOCKET_SERVER_PORT" validate:"required"`
	APP_ENV                  string `json:"APP_ENV" validate:"required"`
	JWT_ACCESS_TOKEN_SECRETE string `json:"JWT_ACCESS_TOKEN_SECRETE" validate:"required"`

	MONGO_DB_URL  string `json:"MONGO_DB_URL" validate:"required"`
	MONGO_DB_NAME string `json:"MONGO_DB_NAME" validate:"required"`
}
