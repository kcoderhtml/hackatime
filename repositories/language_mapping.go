package repositories

import (
	"errors"

	"github.com/hackclub/hackatime/config"
	"github.com/hackclub/hackatime/models"
	"gorm.io/gorm"
)

type LanguageMappingRepository struct {
	config *config.Config
	db     *gorm.DB
}

func NewLanguageMappingRepository(db *gorm.DB) *LanguageMappingRepository {
	return &LanguageMappingRepository{config: config.Get(), db: db}
}

func (r *LanguageMappingRepository) GetAll() ([]*models.LanguageMapping, error) {
	var mappings []*models.LanguageMapping
	if err := r.db.Find(&mappings).Error; err != nil {
		return nil, err
	}
	return mappings, nil
}

func (r *LanguageMappingRepository) GetById(id uint) (*models.LanguageMapping, error) {
	mapping := &models.LanguageMapping{}
	if err := r.db.Where(&models.LanguageMapping{ID: id}).First(mapping).Error; err != nil {
		return mapping, err
	}
	return mapping, nil
}

func (r *LanguageMappingRepository) GetByUser(userId string) ([]*models.LanguageMapping, error) {
	var mappings []*models.LanguageMapping
	if userId == "" {
		return mappings, nil
	}
	if err := r.db.
		Where(&models.LanguageMapping{UserID: userId}).
		Find(&mappings).Error; err != nil {
		return mappings, err
	}
	return mappings, nil
}

func (r *LanguageMappingRepository) Insert(mapping *models.LanguageMapping) (*models.LanguageMapping, error) {
	if !mapping.IsValid() {
		return nil, errors.New("invalid mapping")
	}
	result := r.db.Create(mapping)
	if err := result.Error; err != nil {
		return nil, err
	}
	return mapping, nil
}

func (r *LanguageMappingRepository) Delete(id uint) error {
	return r.db.
		Where("id = ?", id).
		Delete(models.LanguageMapping{}).Error
}
