package handler

import (
	"codecreeo/internal/model"
	"codecreeo/internal/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type QRHandler struct {
	qrRepo repository.QrCodeRepository
}

func NewQRHandler(qrRepo repository.QrCodeRepository) *QRHandler {
	return &QRHandler{
		qrRepo: qrRepo,
	}
}

func (qh *QRHandler) ViewQRCode(c *fiber.Ctx) error {
	qrID := c.Params("qrID")
	id, err := strconv.ParseUint(qrID, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid QR code ID",
		})
	}

	qr, err := qh.qrRepo.GetByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve QR code",
		})
	}
	if qr == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "QR code not found",
		})
	}
	return c.Redirect(qr.Url, fiber.StatusPermanentRedirect)
}

func (qh *QRHandler) GetUserQRCode(c *fiber.Ctx) error {
	userID := c.Params("userID")
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	qrs, err := qh.qrRepo.GetByUserID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve user QR codes",
		})
	}
	return c.Status(fiber.StatusOK).JSON(qrs)
}

func (qh *QRHandler) CreateQRCode(c *fiber.Ctx) error {
	var qr model.QRCode
	if err := c.BodyParser(&qr); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	if qr.UserID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User ID is required",
		})
	}
	if qr.Url == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "URL is required",
		})
	}

	createdQR, err := qh.qrRepo.Create(qr.UserID, qr.Url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create QR code",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(createdQR)
}

func (qh *QRHandler) UpdateQRCode(c *fiber.Ctx) error {
	qrID := c.Params("qrID")
	var qr model.QRCode
	if err := c.BodyParser(&qr); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	qr.UDID = qrID
	if err := qh.qrRepo.Update(&qr); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update QR code",
		})
	}

	return c.Status(fiber.StatusOK).JSON(qr)
}

func (qh *QRHandler) DeleteQRCode(c *fiber.Ctx) error {
	qrID := c.Params("qrID")
	qrIDUint, err := strconv.ParseUint(qrID, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid QR ID",
		})
	}
	qrModel, err := qh.qrRepo.FindByID(uint(qrIDUint))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve QR code",
		})
	}
	if err := qh.qrRepo.Delete(qrModel); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete QR code",
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{})
}
