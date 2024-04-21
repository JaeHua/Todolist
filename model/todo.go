package model

import (
	"JTools/dao"
)

// Todo module
type Todo struct {
	ID     int    `gorm:"primaryKey"`
	Title  string `json:"title"`
	Status bool   `json:"status" gorm:"default:true"`
}

// Todo 增删改查

func CreateATodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

func GetAllTodo() (todoList []*Todo, err error) {
	if err = dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

func GetATodo(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err = dao.DB.Debug().Where("id=?", id).First(&todo).Error; err != nil {

		return nil, err
	}
	return
}

func UpdateTodo(todo *Todo) (err error) {
	err = dao.DB.Save(&todo).Error
	return
}

func DeleteTodo(id string) (err error) {

	err = dao.DB.Debug().Where("id=?", id).Delete(Todo{}).Error
	return

}
