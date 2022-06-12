package dao

import (
	"github.com/go-kit/log"
	"github.com/google/uuid"
	"github.com/stkr89/modelsvc/common"
	"github.com/stkr89/modelsvc/config"
	"github.com/stkr89/modelsvc/models"
	"gorm.io/gorm"
)

//go:generate mocker --dst ../mock/model_dao_mock.go --pkg mock model_dao.go ModelDao
type ModelDao interface {
	Create(obj *models.Model) (*models.Model, error)
	Get(id uuid.UUID) (*models.Model, error)
	List() ([]*models.Model, error)
	Update(obj *models.Model) (*models.Model, error)
	Delete(id uuid.UUID) error
}

type ModelDaoImpl struct {
	logger log.Logger
	db     gorm.DB
}

func NewModelDaoImpl() *ModelDaoImpl {
	return &ModelDaoImpl{
		logger: common.NewLogger(),
		db:     config.NewDB(),
	}
}

func (m ModelDaoImpl) Delete(id uuid.UUID) error {
	result := m.db.Delete(&models.Model{ID: id})
	if result.Error != nil {
		m.logger.Log("message", "failed to delete", "error", result.Error)
		return common.NewError(common.SomethingWentWrong, "failed to delete")
	}

	m.logger.Log("message", "deleted successfully", "return", id)

	return nil
}

func (m ModelDaoImpl) Update(obj *models.Model) (*models.Model, error) {
	existing, err := m.Get(obj.ID)
	if err != nil {
		return nil, err
	}

	obj.CreatedAt = existing.CreatedAt

	result := m.db.Save(obj)
	if result.Error != nil {
		m.logger.Log("message", "failed to update", "error", result.Error)
		return nil, common.NewError(common.SomethingWentWrong, "failed to update")
	}

	m.logger.Log("message", "updated successfully", "return", obj.ID)

	return obj, nil
}

func (m ModelDaoImpl) List() ([]*models.Model, error) {
	var objs []*models.Model
	result := m.db.Find(&objs)
	if result.Error != nil {
		m.logger.Log("message", "failed to list", "error", result.Error)
		return nil, common.NewError(common.SomethingWentWrong, "failed to list")
	}

	return objs, nil
}

func (m ModelDaoImpl) Get(id uuid.UUID) (*models.Model, error) {
	obj := models.Model{ID: id}

	result := m.db.First(&obj)
	if result.Error != nil {
		m.logger.Log("message", "failed to get", "error", result.Error)
		return nil, common.NewError(common.SomethingWentWrong, "failed to get")
	}

	return &obj, nil
}

func (m ModelDaoImpl) Create(obj *models.Model) (*models.Model, error) {
	result := m.db.Create(&obj)
	if result.Error != nil {
		m.logger.Log("message", "failed to create", "error", result.Error)
		return nil, common.NewError(common.SomethingWentWrong, "failed to create")
	}

	m.logger.Log("message", "created successfully", "return", obj.ID)

	return obj, nil
}
