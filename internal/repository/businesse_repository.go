package repository

import (
	"be/internal/entity"
	"be/internal/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BusinesseRepository struct {
	Repository[entity.Businesses]
	Log *logrus.Logger
}

func NewBusinesseRepository(log *logrus.Logger) *BusinesseRepository {
	return &BusinesseRepository{
		Log: log,
	}
}

func (r *BusinesseRepository) FindById(db *gorm.DB, business *entity.Businesses, id string) error {
	return db.Where("id = ?", id).Take(business).Error
}

func (r *BusinesseRepository) Search(db *gorm.DB, request *model.SearchBusinesseRequest) ([]entity.Businesses, int64, error) {
	var business []entity.Businesses
	if err := db.Scopes(r.FilterBusinesse(request)).Offset((request.Page - 1) * request.Size).Limit(request.Size).Find(&business).Error; err != nil {
		return nil, 0, err
	}

	var total int64 = 0
	if err := db.Model(&entity.Businesses{}).Scopes(r.FilterBusinesse(request)).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return business, total, nil
}

func (r *BusinesseRepository) FilterBusinesse(request *model.SearchBusinesseRequest) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if term := request.Term; term != "" {
			term = "%" + term + "%"
			tx = tx.Where("term LIKE ? ", term)
		}
		if location := request.Location; location != "" {
			location = "%" + location + "%"
			tx = tx.Where("location LIKE ? ", location)
		}
		if latitude := request.Latitude; latitude != "" {
			latitude = "%" + latitude + "%"
			tx = tx.Where("latitude LIKE ? ", latitude)
		}
		if longitude := request.Longitude; longitude != "" {
			longitude = "%" + longitude + "%"
			tx = tx.Where("longitude LIKE ? ", longitude)
		}
		if radius := request.Radius; radius != "" {
			radius = "%" + radius + "%"
			tx = tx.Where("radius LIKE ? ", radius)
		}
		return tx
	}
}
