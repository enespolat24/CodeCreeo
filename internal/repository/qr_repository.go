package repository

import (
	"codecreeo/internal/model"

	"gorm.io/gorm"
)

type QrCodeRepository struct {
	db *gorm.DB
}

func NewQrCodeRepository(db *gorm.DB) *QrCodeRepository {
	return &QrCodeRepository{db}
}

func (qr *QrCodeRepository) Create(userID uint, url string) (*model.QRCode, error) {
	newQrCode := &model.QRCode{
		UserID: userID,
		Url:    url,
	}
	if err := qr.db.Create(newQrCode).Error; err != nil {
		return nil, err
	}
	return newQrCode, nil
}

func (qr *QrCodeRepository) GetByID(id uint) (*model.QRCode, error) {
	qrCode := &model.QRCode{}
	if err := qr.db.First(qrCode, id).Error; err != nil {
		return nil, err
	}
	return qrCode, nil
}

func (qr *QrCodeRepository) GetByUserID(userID uint) ([]model.QRCode, error) {
	qrCodes := []model.QRCode{}
	if err := qr.db.Where("user_id = ?", userID).Find(&qrCodes).Error; err != nil {
		return nil, err
	}
	return qrCodes, nil
}

func (qr *QrCodeRepository) Update(qrCode *model.QRCode) error {
	if err := qr.db.Save(qrCode).Error; err != nil {
		return err
	}
	return nil
}

func (qr *QrCodeRepository) Delete(qrCode *model.QRCode) error {
	if err := qr.db.Delete(qrCode).Error; err != nil {
		return err
	}
	return nil
}
