package category

import (
	"blog/blog/storage"
	bgvc "blog/gunk/v1/category"
	"context"
)

type PostCoreLink interface{
	CreateCat(context.Context,storage.Category)(int64,error)
	ListCat(context.Context)([]storage.Category, error)
	GetCat(context.Context, int64)(storage.Category, error)
	UpdateCat(context.Context, storage.Category) error
	DeleteCat(context.Context,int64)error
}

type PostSvc struct {
	bgvc.UnimplementedCategoryServiceServer
	store PostCoreLink
}

func NewPostSvc(s PostCoreLink) *PostSvc{
	return &PostSvc{
		store: s,
	}
}