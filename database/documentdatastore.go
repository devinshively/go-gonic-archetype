package database

import "github.com/devinshively/go-gonic-archetype/model"

type DocumentDataStore interface {
	CreateOrUpdate(doc *model.Document) error
	Get(docId string) (*model.Document, error)
	Delete(docId string) error
	GetAll() ([]model.Document, error)
}
