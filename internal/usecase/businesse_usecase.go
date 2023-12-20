package usecase

import (
	"be/internal/entity"
	"be/internal/model"
	converter "be/internal/model/convertor"
	"be/internal/repository"
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BusinesseUseCase struct {
	DB                  *gorm.DB
	Log                 *logrus.Logger
	Validate            *validator.Validate
	BusinesseRepository *repository.BusinesseRepository
}

func NewBusinesseUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	businesseRepository *repository.BusinesseRepository) *BusinesseUseCase {
	return &BusinesseUseCase{
		DB:                  db,
		Log:                 logger,
		Validate:            validate,
		BusinesseRepository: businesseRepository,
	}
}

func (c *BusinesseUseCase) Create(ctx context.Context, request *model.BusinesseCreateRequest) (*model.BusinesseResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, fiber.ErrBadRequest
	}
	businesse := &entity.Businesses{
		ID:                  uuid.New().String(),
		Location:            request.Location,
		Latitude:            request.Latitude,
		Longitude:           request.Longitude,
		Term:                request.Term,
		Radius:              request.Radius,
		Locale:              request.Locale,
		OpenNow:             request.OpenNow,
		OpenAt:              request.OpenAt,
		SortBy:              request.SortBy,
		DevicePlatform:      request.DevicePlatform,
		ReservationDate:     request.ReservationDate,
		ReservationTime:     request.ReservationTime,
		ReservationCover:    request.ReservationCover,
		MatchPartySizeParam: request.MatchPartySizeParam,
		Limit:               request.Limit,
		Offset:              request.Offset,
	}

	categories, err := json.Marshal(request.Categories)
	if err != nil {
		fmt.Println("Error mengonversi array ke JSON:", err)
		return nil, fiber.ErrInternalServerError
	}
	attributes, err := json.Marshal(request.Attributes)
	if err != nil {
		fmt.Println("Error mengonversi array ke JSON:", err)
		return nil, fiber.ErrInternalServerError
	}
	prices, err := json.Marshal(request.Price)
	if err != nil {
		fmt.Println("Error mengonversi array ke JSON:", err)
		return nil, fiber.ErrInternalServerError
	}
	businesse.Categories = string(categories)
	businesse.Attributes = string(attributes)
	businesse.Price = string(prices)
	if err := c.BusinesseRepository.Create(tx, businesse); err != nil {
		c.Log.WithError(err).Error("error creating businesse", err)
		return nil, fiber.ErrInternalServerError
	}
	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error creating businesse")
		return nil, fiber.ErrInternalServerError
	}
	return converter.BusinesseToResponse(businesse), nil
}

func (c *BusinesseUseCase) Update(ctx context.Context, request *model.BusinesseUpdateRequest) (*model.BusinesseResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
	busines := new(entity.Businesses)
	if err := c.BusinesseRepository.FindById(tx, busines, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting businesse")
		return nil, fiber.ErrNotFound
	}
	fmt.Println("datanua", busines)
	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, fiber.ErrBadRequest
	}
	categories, err := json.Marshal(request.Categories)
	if err != nil {
		fmt.Println("Error mengonversi array ke JSON:", err)
		return nil, fiber.ErrInternalServerError
	}
	attributes, err := json.Marshal(request.Attributes)
	if err != nil {
		fmt.Println("Error mengonversi array ke JSON:", err)
		return nil, fiber.ErrInternalServerError
	}
	prices, err := json.Marshal(request.Price)
	if err != nil {
		fmt.Println("Error mengonversi array ke JSON:", err)
		return nil, fiber.ErrInternalServerError
	}
	busines.Location = request.Location
	busines.Latitude = request.Latitude
	busines.Longitude = request.Longitude
	busines.Term = request.Term
	busines.Radius = request.Radius
	busines.Categories = string(categories)
	busines.Locale = request.Locale
	busines.Price = string(prices)
	busines.OpenNow = request.OpenNow
	busines.OpenAt = request.OpenAt
	busines.Attributes = string(attributes)
	busines.SortBy = request.SortBy
	busines.DevicePlatform = request.DevicePlatform
	busines.ReservationDate = request.ReservationDate
	busines.ReservationTime = request.ReservationTime
	busines.ReservationCover = request.ReservationCover
	busines.MatchPartySizeParam = request.MatchPartySizeParam
	busines.Limit = request.Limit
	busines.Offset = request.Offset
	if err := c.BusinesseRepository.Update(tx, busines); err != nil {
		c.Log.WithError(err).Error("error updating business")
		return nil, fiber.ErrInternalServerError
	}
	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error updating business")
		return nil, fiber.ErrInternalServerError
	}
	return converter.BusinesseToResponse(busines), nil
}

func (c *BusinesseUseCase) Delete(ctx context.Context, request *model.DeleteBusinesseRequest) error {

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return fiber.ErrBadRequest
	}
	businesse := new(entity.Businesses)
	if err := c.BusinesseRepository.FindById(tx, businesse, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting businesse")
		return fiber.ErrNotFound
	}
	if err := c.BusinesseRepository.Delete(tx, businesse); err != nil {
		c.Log.WithError(err).Error("error deleting businesse")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting businesse")
		return fiber.ErrInternalServerError
	}
	return nil
}

func (c *BusinesseUseCase) Search(ctx context.Context, request *model.SearchBusinesseRequest) ([]model.BusinesseResponse, int64, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, 0, fiber.ErrBadRequest
	}

	businesse, total, err := c.BusinesseRepository.Search(tx, request)
	if err != nil {
		c.Log.WithError(err).Error("error getting businesse")
		return nil, 0, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting businesse")
		return nil, 0, fiber.ErrInternalServerError
	}
	responses := make([]model.BusinesseResponse, len(businesse))
	for i, business := range businesse {
		responses[i] = *converter.BusinesseToResponse(&business)
	}
	return responses, total, nil
}
