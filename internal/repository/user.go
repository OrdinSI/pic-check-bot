package repository

import (
	"context"
	"fmt"
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

func (r *user) TopReposts(ctx context.Context) ([]*model.TopRepost, error) {
	var topReposts []*model.TopRepost

	err := r.TrOrDB(ctx).
		Model(&model.User{}).
		Select(`
            users.id AS user_id,
            users.username AS username,
            COUNT(reposts.id) AS count
        `).
		Joins("LEFT JOIN reposts ON users.id = reposts.user_id").
		Group("users.id, users.username").
		Order("count DESC").
		Limit(5).
		Scan(&topReposts).Error

	if err != nil {
		return nil, fmt.Errorf("failed to get top reposts: %w", err)
	}

	return topReposts, nil
}
