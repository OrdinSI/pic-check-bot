package repository

import (
	"context"
	"time"

	"github.com/OrdinSI/pic-check-bot/internal/database/txgorm"
	"github.com/OrdinSI/pic-check-bot/internal/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type user struct {
	*txgorm.TxGetter
}

func NewUserRepository(db *gorm.DB) User {
	return &user{
		TxGetter: txgorm.NewTxGetter(db),
	}
}

func (r *user) GetUser(ctx context.Context, id int64) (*model.User, error) {
	u := &model.User{}
	err := r.TrOrDB(ctx).WithContext(ctx).Where("id = ?", id).Take(u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to get user")
	}
	return u, nil
}

func (r *user) CreateUser(ctx context.Context, user *model.User) (int64, error) {
	user.CreatedAt = time.Time{}
	err := r.TrOrDB(ctx).WithContext(ctx).
		Session(&gorm.Session{FullSaveAssociations: true}).
		Omit().
		Create(user).Error
	return user.ID, err
}

func (r *user) UpdateUser(ctx context.Context, id int64, user model.User) error {
	err := r.TrOrDB(ctx).WithContext(ctx).
		Select("*").
		Omit("CreatedAt").
		Session(&gorm.Session{FullSaveAssociations: true}).
		Where("id = ?", id).
		Updates(&user).Error
	return err
}

func (r *user) GetGroup(ctx context.Context, id int64) (*model.Group, error) {
	g := &model.Group{}
	err := r.TrOrDB(ctx).WithContext(ctx).Where("id = ?", id).Take(g).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to get group")
	}
	return g, nil
}

func (r *user) CreateGroup(ctx context.Context, group *model.Group) (int64, error) {
	group.CreatedAt = time.Time{}
	err := r.TrOrDB(ctx).WithContext(ctx).
		Session(&gorm.Session{FullSaveAssociations: true}).
		Omit().
		Create(group).Error
	return group.ID, err
}

func (r *user) UpdateGroup(ctx context.Context, id int64, group model.Group) error {
	err := r.TrOrDB(ctx).WithContext(ctx).
		Select("*").
		Omit("CreatedAt").
		Session(&gorm.Session{FullSaveAssociations: true}).
		Where("id = ?", id).
		Updates(&group).Error
	return err
}
