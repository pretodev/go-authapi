package data

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"math/rand"
	"time"
)

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"nome"`
	Email   string `json:"email"`
	Balance string `json:"saldo"`
}

func (u *User) GenerateBalance() {
	// random balance
	rand.Seed(time.Now().UnixNano())
	balance := float64(rand.Intn(756)+100) + rand.Float64()
	u.Balance = fmt.Sprintf("%.2f", balance)
}

func (u *User) Password() string {
	return "12345"
}

func (u *User) CreateToken() (string, error) {
	claims := jwt.MapClaims{}
	claims["sub"] = u.ID
	claims["email"] = u.Email
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	fmt.Println(claims)

	// Gerar o token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetUserFromToken(tokenString string) (User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inválido: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return User{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return User{}, err
	}
	id := claims["sub"].(float64)
	email := claims["email"].(string)

	fmt.Println("id: ", id)
	fmt.Println("email: ", email)

	var user User
	for _, u := range Users {
		if int(id) == u.ID {
			user = u
		}
	}

	return user, nil
}

var Users = []User{
	{
		ID:    1,
		Name:  "Maria Viana",
		Email: "maria@gmail.com",
	},
	{
		ID:    2,
		Name:  "João Rodrigues",
		Email: "joao@gmail.com",
	},
}
