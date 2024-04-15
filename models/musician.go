package models

type Musician struct {
	ID   string `json:"id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name" binding:"required,min=3"`
	Type string `json:"type" binding:"required"`
}
