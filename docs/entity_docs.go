package docs

import (
	"example.com/go-swagger-test/entity"
)

// Create body for a new entity
//
// swagger:parameters createEntity
type CreateBody struct {
	// in: body
	Body entity.EntityRequest
}

// Id of the entity
//
// swagger:parameters getEntityById updateEntity
type IDParam struct {
	// Id of the entity 2
	//
	// in: path
	// required: true
	Id string `json:"id"`
}

// Search value
//
// swagger:parameters getEntityById
type SearchParam struct {
	// the value to search
	//
	// in: query
	// required: false
	SearchVal string `json:"searchVal"`
}
