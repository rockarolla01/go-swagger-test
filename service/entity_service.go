package service

import (
	"example.com/go-swagger-test/entity"
	"strconv"
)

type EntityService interface {
	GetAll() []*entity.Entity

	GetById(id string, searchVal string) *entity.Entity

	Create(request *entity.EntityRequest) *entity.Entity

	Update(id string, request *entity.EntityRequest) *entity.Entity
}

func NewEntityService() EntityService {
	return &entityServiceImpl{}
}

type entityServiceImpl struct {
}

// GetAll swagger:route GET /entities entities_tag getAllEntities
//
// This api returns all entities
//
// responses:
//   200: []testEntity
//   500:
func (es *entityServiceImpl) GetAll() []*entity.Entity {
	return []*entity.Entity{
		{"myVal 1", "1", "azaza1", "kekeke1"},
		{"myVal 2", "2", "azaza2", "kekeke2"},
	}
}

// GetById swagger:route GET /entities/{id} entities_tag getEntityById
//
// This api returns an entity found by id
//
// responses:
//   200: testEntity
//   404:
//   500:
func (es *entityServiceImpl) GetById(id string, searchVal string) *entity.Entity {
	return &entity.Entity{"get-by-id " + id, id, searchVal, "--"}
}

// Create swagger:route POST /entities entities_tag createEntity
//
// This api creates a new entity and returns it
//
// responses:
//   200: testEntity
//   500:
func (es *entityServiceImpl) Create(request *entity.EntityRequest) *entity.Entity {
	return &entity.Entity{
		request.Value,
		strconv.Itoa(request.Num),
		"no_search_param",
		"---",
	}
}

// Update swagger:route PUT /entities/{id} entities_tag updateEntity
//
// This api returns an entity found by id
//
// responses:
//   200: testEntity
//   404:
//   500:
func (es *entityServiceImpl) Update(id string, request *entity.EntityRequest) *entity.Entity {
	return &entity.Entity{
		request.Value,
		id,
		"no_search_param",
		"---",
	}
}
