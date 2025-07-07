package authz

import (
	"context"
	"errors"
)

const UserCtxKey string = "user"

func GetUserFromContext(ctx context.Context) (User, error) {
	v := ctx.Value(UserCtxKey)
	if v == nil {
		return User{}, errors.New("user not found in context")
	}

	user, ok := v.(User)
	if !ok {
		return User{}, errors.New("user not found in context")
	}

	if user.Username == "" {
		return User{}, errors.New("user not found in context")
	}

	return user, nil
}
