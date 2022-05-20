package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/tergelm/go_hackernews/graph/generated"
	"github.com/tergelm/go_hackernews/graph/model"
	"github.com/tergelm/go_hackernews/internal/links"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	// TODO Add user
	var link links.Link
	link.Title = input.Title
	link.Address = input.Address
	createdLinkId := link.Save()
	return &model.Link{
		ID:      strconv.FormatInt(createdLinkId, 10),
		Title:   link.Title,
		Address: link.Address,
	}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	// panic(fmt.Errorf("not implemented"))
	var resultsLinks []*model.Link
	var dbLinks []links.Link
	dbLinks = links.GetAll()
	for _, link := range dbLinks {
		resultsLinks = append(resultsLinks, &model.Link{
			ID:      link.Id,
			Title:   link.Title,
			Address: link.Address,
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
