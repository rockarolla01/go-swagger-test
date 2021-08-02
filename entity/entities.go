package entity


// TestRequest is the create/update body for the
//
// swagger:model testEntityRequest
type EntityRequest struct {
	Value string `json:"value"`
	Num   int    `json:"num"`
}

// TestResponse is the main model class used in this app
//
// swagger:model testEntity
type Entity struct {
	EchoValue string `json:"echoValue"`
	EchoNum   string `json:"echoNum"`
	QParam    string `json:"qParam"`
	PParam    string `json:"pParam"`
}
