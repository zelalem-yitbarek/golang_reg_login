package handlers

import ("context" "encoding/json" "net/http" "strings" "time"  "golang_reg_login/models" "go.mongodb.org/mongo-driver/bson" "go.mongodb.org/mongo-driver/bson/primitive" "go.mongodb.org/mongo-driver/mongo" "golang.org/x/crypto/bcrypt")

type RegisterRequest struct{
	Name string `json:"name"`
	UserName string `json:"username"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Password string `json:"password"`
	Role string `json:"role"`
}

func RegisterHandler(userCollection *mongo.Collection)http.HandlerFunc{
	return func(w http.ResponseWriter,r*http.Request){
		if r.Method !=http.MethodPost{
			http.Error(w,"Method Not Allowed",http.StatusMethodNotAllowed)
			return
		}
	}
}