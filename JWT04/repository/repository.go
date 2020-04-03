package repository

import (
	"GOLANG101/JWT04/config"
	"GOLANG101/JWT04/models"
	"fmt"
)

//fetch all todos at once
func GetAllTodos(todo *[]models.Todo) (err error) {
	if err = config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

//insert a todo
func CreateATodo(todo *models.Todo) (err error) {
	if err = config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

//fetch one todo
func GetATodo(todo *models.Todo, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

//update a todo
func UpdateATodo(todo *models.Todo, id string) (err error) {
	fmt.Println(todo)
	config.DB.Save(todo)
	return nil
}

//delete a todo
func DeleteATodo(todo *models.Todo, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(todo)
	return nil
}
