package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/crossworth/nextjs-graphql-gqlgen-ent/ent"
	"github.com/crossworth/nextjs-graphql-gqlgen-ent/ent/post"
	"github.com/crossworth/nextjs-graphql-gqlgen-ent/ent/user"
	"github.com/crossworth/nextjs-graphql-gqlgen-ent/graph/generated"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input ent.CreateUserInput) (*ent.User, error) {
	return ent.FromContext(ctx).User.Create().SetInput(input).Save(ctx)
}

func (r *mutationResolver) CreatePost(ctx context.Context, input ent.CreatePostInput) (*ent.Post, error) {
	return ent.FromContext(ctx).Post.Create().SetInput(input).Save(ctx)
}

func (r *postResolver) User(ctx context.Context, obj *ent.Post) (*ent.User, error) {
	return r.client.User.Query().Where(user.IDEQ(obj.Edges.Author.ID)).First(ctx)
}

func (r *queryResolver) User(ctx context.Context, id string) (*ent.User, error) {
	return r.client.User.Query().Where(user.IDEQ(stringIDToInt(id))).First(ctx)
}

func (r *queryResolver) Post(ctx context.Context, id string) (*ent.Post, error) {
	return r.client.Post.Query().Where(post.IDEQ(stringIDToInt(id))).First(ctx)
}

func (r *queryResolver) Posts(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.PostOrder, where *ent.PostWhereInput) (*ent.PostConnection, error) {
	return r.client.Post.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithPostOrder(orderBy),
			ent.WithPostFilter(where.Filter),
		)
}

func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.UserOrder, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	return r.client.User.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithUserOrder(orderBy),
			ent.WithUserFilter(where.Filter),
		)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
