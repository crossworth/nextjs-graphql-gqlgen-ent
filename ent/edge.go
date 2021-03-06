// Code generated by entc, DO NOT EDIT.

package ent

import "context"

func (po *Post) Author(ctx context.Context) (*User, error) {
	result, err := po.Edges.AuthorOrErr()
	if IsNotLoaded(err) {
		result, err = po.QueryAuthor().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (u *User) Posts(ctx context.Context) ([]*Post, error) {
	result, err := u.Edges.PostsOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryPosts().All(ctx)
	}
	return result, err
}
