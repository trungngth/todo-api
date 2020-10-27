package storage

import (
	//"../helper"
	"todo/model"
	"github.com/jinzhu/gorm"
)

//NoteDBLayer is the interface for note database
type NoteDBLayer interface {
	Create(model.Note) (*model.Note, error)
	List(int) ([]model.Note, error)
	Update(int, model.Note) error
	Delete(int) error
}

//NoteDB is the layer for database
type NoteDB struct {
	DB *gorm.DB
}

//Create a new note
func (n *NoteDB) Create(note model.Note) (*model.Note, error) {
	err := n.DB.Create(&note).Error
	return &note, err
}

//List all the notes using user_id
func (n *NoteDB) List(id int) ([]model.Note, error) {
	notes := []model.Note{}
	err := n.DB.Where("user_id = ?", id).Find(&notes).Error
	return notes, err
}

//Update a note
func (n *NoteDB) Update(id int, note model.Note) error {
	err := n.DB.
		Model(&model.Note{}).
		Where("id = ?", id).
		Update(&note).Error
	return err
}

//Delete marks a note as deleted. Not delete from database
func (n *NoteDB) Delete(id int) error {
	err := n.DB.Where("id = ?", id).Delete(&model.Note{}).Error
	return err
}
