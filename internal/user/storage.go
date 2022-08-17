package user

import "context"

type Storage interface {
	Create(ctx context.Context, user *User) (u *User, err error)
	FindAll(ctx context.Context) (u []*User, err error)
	FindOne(ctx context.Context, id string) (u *User, err error)
	Update(ctx context.Context, user *User) (u *User, err error)
	Delete(ctx context.Context, id string) error
}
