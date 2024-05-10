package middleware

import (
	"fmt"
	"hayday/server/config"
	"hayday/server/internal/logger"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type jwtUserDetails struct {
	jwt.RegisteredClaims
	UserId interface{} `json:"id"`
}

func WsAuthValidator(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestId, _ := uuid.NewUUID()
		log := logger.Log.WithFields(map[string]interface{}{
			"EVENT":      "AUTH_VALIDATOR",
			"REQUEST_ID": requestId,
		})

		tokenString := ""
		queyrValues := r.URL.Query()
		tokenString = queyrValues.Get("token")

		if tokenString == "" {
			authHeader := r.Header.Get("Authorization")
			if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
				tokenString = authHeader[7:]
			} else {
				log.Errorf(`failed to authorize the user : TOKEN MISSING : %s`, authHeader)
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, "Unauthorized User")
				return
			}
		}

		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized User")
		}

		jwtAccessToken := config.EnvValues.JWT_ACCESS_TOKEN_SECRETE
		token, err := jwt.ParseWithClaims(tokenString, &jwtUserDetails{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Errorf(`failed to authorize the user : ERROR IN HASHING`)
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, "Unauthorized User")
			}
			return []byte(jwtAccessToken), nil
		})

		if err != nil {
			log.Errorf(`failed to parse the token claims : ERROR  : %s`, err)
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized User")
			return
		}

		userDetails, ok := token.Claims.(*jwtUserDetails)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized User")
			return
		}

		switch userDetails.UserId.(type) {
		case float64:
			userDetails.UserId = fmt.Sprintf("%d", int(userDetails.UserId.(float64)))
		case string:
			userDetails.UserId = userDetails.UserId.(string)
		}
		queyrValues.Add("userId", userDetails.UserId.(string))
		//queyrValues.Add("clubId", clubId)
		r.URL.RawQuery = queyrValues.Encode()
		if err != nil || !token.Valid {
			log.Errorf(`failed to authorize the user : TOKEN INVALID : ERROR : %s`, err)
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized User")
			return
		}

		next.ServeHTTP(w, r)
	})
}
