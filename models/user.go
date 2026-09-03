packages models

import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct(
	ID   primitive.ObjectId  `bson:"_id,omitempty" json:"id,omitempty"`
	Name string              `bson:"name" json:"name"`
	UserName string          `bson:"username" json:"username"`
	Phone string             `bson:"phone" json:"phone"`
	Email string             `bson:"email" json:"email"`
	Password string          `bson:"password" json:"password"`
	Role string              `bson:"role" json:"role"`
	CreatedAt time.Time      `bson:"created_at" json:"created_at"`

)