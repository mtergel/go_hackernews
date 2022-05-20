package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/tergelm/go_hackernews/graph/generated"
	"github.com/tergelm/go_hackernews/graph/model"
	"github.com/tergelm/go_hackernews/internal/auth"
	"github.com/tergelm/go_hackernews/internal/links"
	"github.com/tergelm/go_hackernews/internal/pkg/jwt"
	"github.com/tergelm/go_hackernews/internal/users"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return &model.Link{}, fmt.Errorf("access denied")
	}

	var link links.Link
	link.User = user
	link.Title = input.Title
	link.Address = input.Address
	createdLinkId := link.Save()
	gqlUser := &model.User{
		ID:   user.Id,
		Name: user.Username,
	}
	return &model.Link{
		ID:      strconv.FormatInt(createdLinkId, 10),
		Title:   link.Title,
		Address: link.Address,
		User:    gqlUser,
	}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	user.Create()

	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	valid := user.Authenticate()

	if !valid {
		return "", &users.WrongUsernameOrPasswordError{}
	}

	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	username, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}

	token, err := jwt.GenerateToken(username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	var resultsLinks []*model.Link
	var dbLinks []links.Link
	dbLinks = links.GetAll()
	for _, link := range dbLinks {
		gqlUser := &model.User{
			ID:   link.User.Id,
			Name: link.User.Username,
		}
		resultsLinks = append(resultsLinks, &model.Link{
			ID:      link.Id,
			Title:   link.Title,
			Address: link.Address,
			User:    gqlUser,
		})
	}

	return resultsLinks, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type (
	mutationResolver struct{ *Resolver }
	queryResolver    struct{ *Resolver }
)
