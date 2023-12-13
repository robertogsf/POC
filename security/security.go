package security

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/robertogsf/POC/models"
	"golang.org/x/crypto/bcrypt"
)

func WithRol(role string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		value := os.Getenv("secret")
		secretKey := []byte(value)
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		claims, _ := token.Claims.(jwt.MapClaims)
		userRole := claims["rol"].(string)

		if userRole != role {
			http.Error(w, "Acceso denegado", http.StatusForbidden)
			return
		}

		next(w, r)
	}
}

func CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
func HashedPassword(user models.User) models.User {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	if err != nil {
		log.Fatal(err)
	}
	return user
}

func GenerateAccessToken(u models.User) string {
	claims := jwt.MapClaims{
		"email": u.Email,
		"rol":   u.Rol,
		"iat":   time.Now().Unix(),                // Issued At
		"exp":   time.Now().Add(time.Hour).Unix(), // Expiration Time
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	value := os.Getenv("secret")
	secretKey := []byte(value)
	accessToken, err := token.SignedString(secretKey)
	if err != nil {
		return ""
	}

	return accessToken
}
func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/*requiresAuthorization := checkIfAuthorizationIsRequired(r.URL.Path)

		if !requiresAuthorization {
			next.ServeHTTP(w, r)
			return
		}*/
		// Lista de endpoints que no requieren autenticación
		paths := []string{"/login", "/register"}

		// Comprueba si la ruta actual está en la lista de paths
		for _, path := range paths {
			if r.URL.Path == path {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenString := r.Header.Get("Authorization")
		/*tokenString = tokenString[1:]
		tokenString = strings.TrimSuffix(tokenString, string(tokenString[len(tokenString)-1]))
		*/
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing access token"))
			return
		}

		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		value := os.Getenv("secret")
		secretKey := []byte(value)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid access token"))
			return
		}

		next.ServeHTTP(w, r)
	})
}

/*
func checkIfAuthorizationIsRequired(path string) bool {
	if path == "/users/login" {
		return false
	} else if path == "/users/signup" {
		return false
	}
	return true
}
*/
