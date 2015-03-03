package postgres

import (
	"github.com/devinshively/go-gonic-archetype/model"
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
	"log"
)

var db *sqlx.DB

func init() {
	newDb, err := sqlx.Connect("postgres", "dbname=test host=localhost port=5432 user=test password=password sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	db = newDb
}

type PostgresDocumentDB struct{}

func (ds PostgresDocumentDB) CreateOrUpdate(doc *model.Document) error {
	if doc.Id == 0 {
		rows, err := db.NamedQuery(`INSERT INTO document(title, text) VALUES(:title, :text) RETURNING id`, &doc)
		if err != nil {
			return err
		}
		if rows.Next() {
			rows.Scan(&doc.Id)
		}
		return err
	} else {
		_, err := db.NamedExec(`UPDATE document SET title=:title, text=:text WHERE id=:id`, &doc)
		return err
	}
}

func (ds PostgresDocumentDB) Get(docId string) (*model.Document, error) {
	doc := model.Document{}
	err := db.Get(&doc, `SELECT * FROM document WHERE id=$1`, docId)
	return &doc, err
}

func (ds PostgresDocumentDB) Delete(docId string) error {
	_, err := db.NamedExec(`DELETE FROM document WHERE id=$1`, docId)
	return err
}

func (ds PostgresDocumentDB) GetAll() ([]model.Document, error) {
	docs := []model.Document{}
	err := db.Select(&docs, `SELECT * FROM document`)
	if err != nil {
		log.Fatal(err)
	}
	return docs, err
}
