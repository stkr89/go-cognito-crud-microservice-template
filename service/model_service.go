package service

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/stkr89/modelsvc/common"
	"github.com/stkr89/modelsvc/dao"
	"github.com/stkr89/modelsvc/models"
	"github.com/stkr89/modelsvc/types"
)

// ModelService interface
type ModelService interface {
	Create(ctx context.Context, request *types.CreateRequest) (*types.CreateResponse, error)
	Get(ctx context.Context, request *types.GetRequest) (*types.GetResponse, error)
	List(ctx context.Context, request *types.ListRequest) (*types.ListResponse, error)
	Update(ctx context.Context, request *types.UpdateRequest) (*types.UpdateResponse, error)
	Delete(ctx context.Context, request *types.DeleteRequest) error
}

type ModelServiceImpl struct {
	logger   log.Logger
	modelDao dao.ModelDao
}

func NewModelServiceImpl() *ModelServiceImpl {
	return &ModelServiceImpl{
		logger:   common.NewLogger(),
		modelDao: dao.NewModelDaoImpl(),
	}
}

func (m ModelServiceImpl) Delete(ctx context.Context, request *types.DeleteRequest) error {
	err := m.modelDao.Delete(request.ID)
	if err != nil {
		return err
	}

	m.logger.Log("message", "deleted successfully", "return", request.ID)

	return nil
}

func (m ModelServiceImpl) Update(ctx context.Context, request *types.UpdateRequest) (*types.UpdateResponse, error) {
	updatedModel, err := m.modelDao.Update(&models.Model{})
	if err != nil {
		return nil, err
	}

	m.logger.Log("message", "updated successfully", "return", updatedModel.ID)

	return &types.UpdateResponse{}, nil
}

func (m ModelServiceImpl) List(ctx context.Context, request *types.ListRequest) (*types.ListResponse, error) {
	list, err := m.modelDao.List()
	if err != nil {
		return nil, err
	}

	var data []*types.ListResponseData
	for _, model := range list {
		data = append(data, &types.ListResponseData{
			ID: model.ID,
		})
	}

	return &types.ListResponse{
		Data: data,
	}, err
}

func (m ModelServiceImpl) Get(ctx context.Context, request *types.GetRequest) (*types.GetResponse, error) {
	_, err := m.modelDao.Get(request.ID)
	if err != nil {
		return nil, err
	}

	return &types.GetResponse{}, nil
}

func (m ModelServiceImpl) Create(ctx context.Context, request *types.CreateRequest) (*types.CreateResponse, error) {
	createdModel, err := m.modelDao.Create(&models.Model{})
	if err != nil {
		return nil, err
	}

	m.logger.Log("message", "created successfully", "return", createdModel.ID)

	return &types.CreateResponse{}, nil
}
