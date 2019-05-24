package handler

import (
	"strconv"

	"../model"
	"../storage"

	"github.com/gin-gonic/gin"
)

// NoteCreate create a new note
func NoteCreate(c *gin.Context, noteDBLayer storage.NoteDBLayer) (*model.Note, error) {
	note := model.Note{}
	if err := c.ShouldBind(&note); err != nil {
		return nil, err
	}
	return noteDBLayer.Create(note)
}

//NoteList list all note using userid
func NoteList(c *gin.Context, noteDBLayer storage.NoteDBLayer) ([]model.Note, error) {
	id, _ := strconv.Atoi(c.Param("id"))
	return noteDBLayer.List(id)
}

//NoteUpdate update a note
func NoteUpdate(c *gin.Context, noteDBLayer storage.NoteDBLayer) error {
	id, _ := strconv.Atoi(c.Param("id"))
	note := model.Note{}
	if err := c.ShouldBind(&note); err != nil {
		return err
	}
	return noteDBLayer.Update(id, note)
}

//NoteDelete delete a note
func NoteDelete(c *gin.Context, noteDBLayer storage.NoteDBLayer) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return noteDBLayer.Delete(id)
}
