package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error)
	StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error)
	StoreManyCategory(ctx context.Context, categories []entity.Category) error
	GetCategoryByID(ctx context.Context, id int) (entity.Category, error)
	UpdateCategory(ctx context.Context, category *entity.Category) error
	DeleteCategory(ctx context.Context, id int) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error) {
	category := []entity.Category{}

	ctg := r.db.WithContext(ctx).Where("user_id = ?", id).Find(&category)
	if ctg.Error != nil {
		return []entity.Category{}, ctg.Error
	}
	return category, nil // TODO: replace this
}

func (r *categoryRepository) StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error) {
	ctg := entity.Category{}

	categorys := r.db.WithContext(ctx).Create(&category)
	if categorys.Error != nil {
		return 0, categorys.Error
	}
	categorys = r.db.WithContext(ctx).First(&ctg).Order("id DESC")
	if categorys.Error != nil {
		return 0, categorys.Error
	}
	return ctg.ID, nil // TODO: replace this
}

func (r *categoryRepository) StoreManyCategory(ctx context.Context, categories []entity.Category) error {
	category := r.db.WithContext(ctx).Create(&categories)
	if category.Error != nil {
		return category.Error
	}
	return nil // TODO: replace this
}

func (r *categoryRepository) GetCategoryByID(ctx context.Context, id int) (entity.Category, error) {
	category := entity.Category{}

	ctg := r.db.WithContext(ctx).Where("id = ?", id).First(&category)
	if ctg.Error != nil {
		return entity.Category{}, ctg.Error
	}
	return category, nil // TODO: replace this
}

func (r *categoryRepository) UpdateCategory(ctx context.Context, category *entity.Category) error {
	ctg := r.db.WithContext(ctx).Model(&category).Where("id = ?", category.ID).Updates(category)
	if ctg.Error != nil {
		return ctg.Error
	}
	return nil // TODO: replace this
}

func (r *categoryRepository) DeleteCategory(ctx context.Context, id int) error {
	category := r.db.WithContext(ctx).Delete(&entity.Category{}, id)
	if category.Error != nil {
		return category.Error
	}
	return nil // TODO: replace this
}
