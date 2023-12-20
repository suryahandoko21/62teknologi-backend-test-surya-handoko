package http

import (
	"be/internal/model"
	"be/internal/usecase"
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type BusinessController struct {
	UseCase *usecase.BusinesseUseCase
	Log     *logrus.Logger
}

func NewBusinessController(useCase *usecase.BusinesseUseCase, log *logrus.Logger) *BusinessController {
	return &BusinessController{
		UseCase: useCase,
		Log:     log,
	}
}

func (c *BusinessController) Create(ctx *fiber.Ctx) error {
	request := new(model.BusinesseCreateRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}
	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error creating Businesse")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.BusinesseResponse]{Data: response})
}

func (c *BusinessController) Update(ctx *fiber.Ctx) error {

	request := new(model.BusinesseUpdateRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}
	request.ID = ctx.Params("uuid")
	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error updating business")
		return err
	}
	return ctx.JSON(model.WebResponse[*model.BusinesseResponse]{Data: response})

}

func (c *BusinessController) List(ctx *fiber.Ctx) error {
	request := &model.SearchBusinesseRequest{
		Term:      ctx.Query("term", ""),
		Location:  ctx.Query("location", ""),
		Latitude:  ctx.Query("latitude", ""),
		Longitude: ctx.Query("longitude", ""),
		Radius:    ctx.Query("radius", ""),
		Page:      ctx.QueryInt("page", 1),
		Size:      ctx.QueryInt("size", 10),
	}
	responses, total, err := c.UseCase.Search(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error searching business")
		return err
	}
	paging := &model.PageMetadata{
		Page:      request.Page,
		Size:      request.Size,
		TotalItem: total,
		TotalPage: int64(math.Ceil(float64(total) / float64(request.Size))),
	}

	return ctx.JSON(model.WebResponse[[]model.BusinesseResponse]{
		Data:   responses,
		Paging: paging,
	})

}
func (c *BusinessController) Delete(ctx *fiber.Ctx) error {
	uuid := ctx.Params("uuid")
	request := &model.DeleteBusinesseRequest{
		ID: uuid,
	}
	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("error deleting businesse")
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: true})
}
