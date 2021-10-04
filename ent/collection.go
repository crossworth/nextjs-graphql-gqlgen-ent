// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (po *PostQuery) CollectFields(ctx context.Context, satisfies ...string) *PostQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		po = po.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return po
}

func (po *PostQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *PostQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "author":
			po = po.WithAuthor(func(query *UserQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return po
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) *UserQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		u = u.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return u
}

func (u *UserQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *UserQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "posts":
			u = u.WithPosts(func(query *PostQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return u
}
