//go:generate go run github.com/99designs/gqlgen

package accounts

import (
	"context"
)

var _ ResolverRoot = (*resolver)(nil)

func NewResolver() ResolverRoot {
	ptrString := func(s string) *string {
		return &s
	}

	return &resolver{
		users: []*User{
			{
				ID:       "1",
				Name:     ptrString("Jad"),
				Username: ptrString("@jad"),
				Email:    ptrString("jad@email.com"),
			},
			{
				ID:       "2",
				Name:     ptrString("Marouan"),
				Username: ptrString("@marouan"),
				Email:    ptrString("marouan@email.com"),
			},
			{
				ID:       "3",
				Name:     ptrString("Anas"),
				Username: ptrString("@anas"),
				Email:    ptrString("anas@email.com"),
			},
			{
				ID:       "4",
				Name:     ptrString("Chaymae"),
				Username: ptrString("@chaymae"),
				Email:    ptrString("chaymae@email.com"),
			},
			{
				ID:       "5",
				Name:     ptrString("Adam"),
				Username: ptrString("@adam"),
				Email:    ptrString("adam@email.com"),
			},
		},
	}
}

type resolver struct {
	users []*User
}

func (r *resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

func (r *resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func (r *resolver) Entity() EntityResolver {
	return &entityResolver{r}
}

type queryResolver struct{ *resolver }

func (r *queryResolver) Me(ctx context.Context) (*User, error) {
	return r.users[0], nil
}

func (r *queryResolver) AllUsers(ctx context.Context) ([]*User, error) {
	return r.users, nil
}

type mutationResolver struct{ *resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input *NewUser) (*User, error) {
	user := &User{
		ID:       input.ID,
		Name:     &input.Name,
		Username: &input.Username,
		Email:    &input.Email,
	}
	r.users = append(r.users, user)
	return user, nil
}

type entityResolver struct{ *resolver }

func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, nil
}
