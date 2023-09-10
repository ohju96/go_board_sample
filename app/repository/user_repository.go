package repository

import "ginSample/ent"

type UserRepository interface {
}

type userRepository struct {
	client ent.Client
}

func NewUserRepository(client ent.Client) UserRepository {
	return &userRepository{
		client: client,
	}
}
