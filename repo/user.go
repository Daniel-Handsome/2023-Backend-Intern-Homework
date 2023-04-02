package repo

import (
	"context"
	"errors"

	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/orm"
	"gorm.io/gorm"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository interface {
    FindByID(ctx context.Context, id int64) (*orm.User, error)
    FindByEmail(ctx context.Context, email string) (*orm.User, error)
    Create(ctx context.Context, user *orm.User) error
    Update(ctx context.Context, user *orm.User) error
    Delete(ctx context.Context, id int64) error
}

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db: db}
}

func (r *userRepository) FindByID(ctx context.Context, id int64) (*orm.User, error) {
    var user orm.User
    err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, ErrUserNotFound
        }
        return nil, err
    }
    return &user, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*orm.User, error) {
    var user orm.User
    err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, ErrUserNotFound
        }
        return nil, err
    }
    return &user, nil
}

func (r *userRepository) Create(ctx context.Context, user *orm.User) error {
    err := r.db.WithContext(ctx).Create(user).Error
    return err
}

func (r *userRepository) Update(ctx context.Context, user *orm.User) error {
    err := r.db.WithContext(ctx).Model(user).Updates(orm.User{Name: user.Name, Email: user.Email}).Error
    return err
}

func (r *userRepository) Delete(ctx context.Context, id int64) error {
    err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&orm.User{}).Error
	return err
}