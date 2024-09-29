package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"kredit-plus/service/model/request"
	"kredit-plus/service/usecase"
	"kredit-plus/service/utils"
	"net/http"
)

type Handlers struct {
	CustomerService    usecase.CustomerUsecase
	TransactionService usecase.TransactionService
}

func NewHandler(customer usecase.CustomerUsecase, transaction usecase.TransactionService) Handlers {
	return Handlers{customer, transaction}
}

func (h *Handlers) Route(apps *fiber.App) {
	apps.Post("/api/v1/create-transaction", h.CreateTransaction)
	apps.Post("/api/v1/create-customer", h.CreateCustomer)
	apps.Get("/api/v1/verif", h.VerifyTransaction)

}

func (h *Handlers) VerifyTransaction(c *fiber.Ctx) error {
	var input request.RequestTransaction
	input.Nik = c.Query("nik")
	input.Otp = c.Query("otp")

	err := h.TransactionService.Transaction(&input)
	if err != nil {
		response := fiber.Map{
			"statusMessage": err.Error(),
			"statusCode":    404,
		}
		return c.Status(http.StatusNotFound).JSON(response)
	}

	response := fiber.Map{
		"statusMessage": "success",
		"statusCode":    200,
	}
	return c.Status(http.StatusOK).JSON(response)

}

func (h *Handlers) CreateTransaction(c *fiber.Ctx) error {
	var input request.TransactionRequest

	c.Params("")

	err := c.BodyParser(&input)
	if err != nil {
		response := fiber.Map{
			"statusMessage": err.Error(),
			"statusCode":    400,
		}
		return c.Status(400).JSON(response)
	}

	otp, err := h.TransactionService.CreateTransaction(&input)
	if err != nil {
		response := fiber.Map{
			"statusMessage": err.Error(),
			"statusCode":    502,
		}
		return c.Status(http.StatusBadGateway).JSON(response)
	}

	response := fiber.Map{
		"statusMessage": "success",
		"statusCode":    200,
		"otp":           otp,
	}
	return c.Status(http.StatusOK).JSON(response)

}

func (h *Handlers) CreateCustomer(c *fiber.Ctx) error {
	var input request.InputCustomer

	err := c.BodyParser(&input)
	if err != nil {
		response := fiber.Map{
			"statusMessage": err.Error(),
			"statusCode":    400,
		}
		return c.Status(400).JSON(response)
	}
	path := fmt.Sprintf("./assets/images/ktp/%s-%s", input.Nik, input.FullName)
	selfieImages := fmt.Sprintf("./assets/images/selfie/%s-%s", input.Nik, input.FullName)
	imagesPath, TypePath, err := utils.ConvertImages(input.KtpImage)
	if err != nil {
		response := fiber.Map{
			"statusMessage": err.Error(),
			"statusCode":    400,
		}
		return c.Status(400).JSON(response)
	}

	imagesSelfie, TypeSelfie, err := utils.ConvertImages(input.KtpImage)
	if err != nil {
		response := fiber.Map{
			"statusMessage": err.Error(),
			"statusCode":    400,
		}
		return c.Status(400).JSON(response)
	}

	path += TypePath
	selfieImages += TypeSelfie
	input.KtpImage = path

	err = utils.SaveImage(imagesPath, TypePath, path)
	if err != nil {
		response := fiber.Map{
			"statusMessage": err.Error(),
			"statusCode":    400,
		}
		return c.Status(400).JSON(response)
	}

	err = utils.SaveImage(imagesSelfie, TypeSelfie, selfieImages)
	if err != nil {
		response := fiber.Map{
			"statusMessage": err.Error(),
			"statusCode":    400,
		}
		return c.Status(400).JSON(response)
	}

	err = h.CustomerService.CreateCustomer(&input)
	if err != nil {
		response := fiber.Map{
			"statusMessage": err.Error(),
			"statusCode":    502,
		}
		return c.Status(http.StatusBadGateway).JSON(response)
	}

	response := fiber.Map{
		"statusMessage": "succes",
		"statusCode":    200,
	}
	return c.Status(http.StatusOK).JSON(response)

}
