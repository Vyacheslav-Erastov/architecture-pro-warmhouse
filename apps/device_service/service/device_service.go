package service

import (
	"context"
	"device_service/models"
	"device_service/repository"
)

type DeviceService struct {
	repo *repository.DeviceRepository
}

func NewDeviceService(repo *repository.DeviceRepository) *DeviceService {
	return &DeviceService{repo: repo}
}

func (s *DeviceService) GetDevices(ctx context.Context) ([]models.Device, error) {
	return s.repo.GetAll(ctx)
}

func (s *DeviceService) CreateDevice(ctx context.Context, d models.DeviceCreate) (int, error) {
	return s.repo.Create(ctx, d)
}

func (s *DeviceService) GetDevice(ctx context.Context, id string) (*models.Device, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *DeviceService) UpdateStatus(ctx context.Context, id string, status string) error {
	return s.repo.UpdateStatus(ctx, id, status)
}
