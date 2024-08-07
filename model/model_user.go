package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	User_id   primitive.ObjectID `bson:"publication_id,omitempty"`
	Name      string             `bson:"name" validate:"required"`
	Surname   string             `bson:"surname,omitempty"`
	Bio       string             `bson:"bio,omitempty" `
	Nick      string             `bson:"nick" validate:"required"`
	Email     string             `bson:"email" validate:"required"`
	Password  string             `bson:"password" validate:"required"`
	Role      string             `bson:"role,omitempty" default:"role_user"`
	Image     string             `bson:"image,omitempty" default:"default.png"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
}
