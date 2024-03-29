package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/HuanLiu-hotstar/search-graphql/graph/generated"
	"github.com/HuanLiu-hotstar/search-graphql/graph/model"
)

func (r *queryResolver) Search(ctx context.Context, searchWord string) ([]model.SearchTamplate, error) {
	if searchWord == "IPL" {
		return convertScoreCardData(), nil
	}
	return convertListData(), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
