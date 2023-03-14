package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	user := entity.User{}

	users := r.db.WithContext(ctx).Where("id = ?", id).Find(&user)
	if users.Error != nil {
		return entity.User{}, users.Error
	}
	return user, nil // TODO: replace this
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	user := entity.User{}

	users := r.db.WithContext(ctx).Where("email = ?", email).Find(&user)
	if users.Error != nil {
		return entity.User{}, users.Error
	}
	return user, nil // TODO: replace this
}

func (r *userRepository) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	userr := entity.User{}

	users := r.db.WithContext(ctx).Create(&user)
	if users.Error != nil {
		return entity.User{}, users.Error
	}

	// users = r.db.WithContext(ctx).First(&userr).Order("id DESC")
	usr := r.db.WithContext(ctx).Where("fullname = ? and email = ?", user.Fullname, user.Email).First(&userr)
	if usr.Error != nil {
		return entity.User{}, usr.Error
	}
	return user, nil // TODO: replace this
}

func (r *userRepository) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
	userr := entity.User{}
	users := r.db.WithContext(ctx).Model(&user).Where("id = ?", user.ID).Updates(user)
	if users.Error != nil {
		return entity.User{}, users.Error
	}

	users = r.db.WithContext(ctx).First(&userr).Order("updated_at DESC")
	if users.Error != nil {
		return entity.User{}, users.Error
	}
	return user, nil // TODO: replace this
}

func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	user := r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.User{})
	if user.Error != nil {
		return user.Error
	}
	return nil // TODO: replace this
}
