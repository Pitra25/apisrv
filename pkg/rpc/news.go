package rpc

import (
	"context"
)

/*** News ***/

// News return list news.
//
//zenrpc:200 ok
//zenrpc:404 not found
func (ns *NewsService) News(ctx context.Context, params *queryParams) ([]NewsSummary, error) {
	list, err := ns.m.GetNewsByFilters(ctx, params.NewFilter())
	if err != nil {
		return nil, newInternalError(err)
	} else if list == nil {
		return nil, ErrInternal
	}

	return NewNewsSummaries(list), nil
}

// GetById returns news by ID.
//
//zenrpc:id news id
//zenrpc:200 ok
//zenrpc:404 not found
func (ns *NewsService) GetById(ctx context.Context, id int) (*News, error) {
	news, err := ns.m.GetNewsByID(ctx, id)
	if err != nil {
		return nil, newInternalError(err)
	} else if news == nil {
		return nil, ErrInternal
	}

	return NewNews(news), nil
}

// CountNews returns count news by filters.
//
//zenrpc:200 ok
//zenrpc:404 not found
func (ns *NewsService) CountNews(ctx context.Context, params *queryParams) (int, error) {
	count, err := ns.m.GetNewsCount(ctx, params.NewFilter())
	if err != nil {
		return 0, newInternalError(err)
	}

	return count, nil
}

/*** Category ***/

// Categories return list category
//
//zenrpc:200 ok
//zenrpc:404 not found
func (ns *NewsService) Categories(ctx context.Context) ([]Category, error) {
	list, err := ns.m.GetAllCategory(ctx)
	if err != nil {
		ns.l.Error(ctx, "error getting all categories", "err", err)
		return nil, newInternalError(err)
	} else if list == nil {
		return nil, ErrInternal
	}

	return NewCategories(list), nil
}

/*** Tag ***/

// Tags return list tag.
//
//zenrpc:200 ok
//zenrpc:404 not found
func (ns *NewsService) Tags(ctx context.Context) ([]Tag, error) {
	list, err := ns.m.GetAllTag(ctx)
	if err != nil {
		return nil, newInternalError(err)
	} else if list == nil {
		return nil, ErrInternal
	}

	return NewTags(list), nil
}

/*** User ***/

// User return user by id.
//
//zenrpc:200 ok
//zenrpc:404 not found
func (ns *NewsService) UserByID(ctx context.Context, id int) (*User, error) {
	user, err := ns.m.GetUserByID(ctx, id)
	if err != nil {
		return nil, newInternalError(err)
	} else if user == nil {
		return nil, ErrInternal
	}

	return NewUser(user), nil
}
