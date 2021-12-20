package category

import (
	"blog/blog/storage"
	"blog/blog/storage/postgres"
	"context"
)
type CoreSvc struct {
	core *postgres.Storage
}

func NewCoreSvc(s *postgres.Storage) *CoreSvc {
	return &CoreSvc{
		core: s,
	}
}

func (cs CoreSvc) CreateCat(ctx context.Context, t storage.Post) (int64, error) {

	return cs.core.CreateCat(ctx,t)
}
func (cs CoreSvc) ListCat(ctx context.Context) ([]storage.Post, error) {

	return cs.core.ListCat(ctx)
}
func (cs CoreSvc) GetCat(ctx context.Context, id int64) (storage.Post, error) {

	return cs.core.GetCat(ctx, id)
}
func (cs CoreSvc) UpdateCat(ctx context.Context, t storage.Post) error {

	return cs.core.UpdateCat(ctx, t)
}
func (cs CoreSvc) DeleteCat(ctx context.Context, id int64) error {

	return cs.core.DeleteCat(ctx, id)
}
