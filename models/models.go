package models

import (
	"fmt"
	"go-gin/config"
)

//fetch all todos at once
func GetTodos(todo *[]Todo) error {
	if err := config.DB.Find(todo).Error; err != nil {
		return err
	}

	return nil
}

//insert a todo
func CreateTodo(todo *Todo) error {
	if err := config.DB.Create(todo).Error; err != nil {
		return err
	}

	return nil
}

//fetch one todo
func GetTodo(todo *Todo, id string) error {
	if err := config.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}

	return nil
}

//update a todo
func UpdateTodo(todo *Todo, id string) error {
	fmt.Println(todo)
	if err := config.DB.Save(todo).Error; err != nil {
		return err
	}

	return nil
}

//delete a todo
func DeleteTodo(todo *Todo, id string) error {
	if err := config.DB.Where("id = ?", id).Delete(todo).Error; err != nil {
		return err
	}

	return nil
}
