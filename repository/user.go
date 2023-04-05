package repository

import (
	"context"
	"errors"

	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/model"
	"gorm.io/gorm"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository interface {
	FindByID(ctx context.Context, id int64) (*model.User, error)
	FindByUuid(ctx context.Context, uuid string) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id int64) error
	GetPageLinkedList(ctx context.Context, uuid string, t model.PageLinkedListType) (*model.PageLinkedList, error)
}

type userRepository struct {
	*BaseRepo
}

func NewUserRepository(orm *gorm.DB) UserRepository {
	return &userRepository{NewRepo(orm.Model(model.User{}))}
}

func (repo *userRepository) FindByID(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	err := repo.orm.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (repo *userRepository) FindByUuid(ctx context.Context, uuid string) (*model.User, error) {
	var user model.User
	err := repo.orm.WithContext(ctx).Where("uuid = ?", uuid).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (repo *userRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := repo.orm.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (repo *userRepository) Create(ctx context.Context, user *model.User) error {
	err := repo.orm.WithContext(ctx).Create(user).Error
	return err
}

func (repo *userRepository) Update(ctx context.Context, user *model.User) error {
	err := repo.orm.WithContext(ctx).Model(user).Updates(model.User{Name: user.Name, Email: user.Email}).Error
	return err
}

func (repo *userRepository) Delete(ctx context.Context, id int64) error {
	err := repo.orm.WithContext(ctx).Where("id = ?", id).Delete(&model.User{}).Error
	return err
}

func (r *userRepository) GetPageLinkedList(ctx context.Context, uuid string, t model.PageLinkedListType) (*model.PageLinkedList, error) {
	var pageLinkedList *model.PageLinkedList

	user, err := r.FindByUuid(ctx, uuid)
	if err != nil {
		return nil, err
	}

	associations := r.orm.Model(&user).Where("type", t).Association("PageLinkedList")
	if err := associations.Find(&pageLinkedList); err != nil {
		return nil, err
	}

	return pageLinkedList, nil
}
