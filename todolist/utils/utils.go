package utils

import (
	"encoding/json"
	"errors"
	"os"
	"slices"
)

type Todo struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	Completed bool   `json:"completed"`
}

const Perm os.FileMode = 0o644

var seed = 0

func SerializeJSON(data []Todo) ([]byte, error) {
	return json.MarshalIndent(data, "", "  ")
}

func DeserializeJSON(data []byte) ([]Todo, error) {
	var parsedData []Todo

	if err := json.Unmarshal(data, &parsedData); err != nil {
		return parsedData, err
	}
	return parsedData, nil
}

func ReadFile(path string) ([]byte, error) {
	if path == "" {
		return nil, errors.New("路径错误")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.New("读取文件错误")
	}

	return data, nil
}

func WriteFile(path string, data []byte) error {
	return os.WriteFile(path, data, Perm)
}

func GetTodoList(path string) ([]Todo, error) {
	data, err := ReadFile(path)
	if err != nil {
		return nil, err
	}
	return DeserializeJSON(data)
}

func InitTodoList(path string) ([]Todo, error) {
	data, err := GetTodoList(path)
	if err == nil {
		return data, err
	}

	writeErr := os.WriteFile(path, []byte("[]"), Perm)
	if writeErr != nil {
		return nil, writeErr
	}
	return []Todo{}, nil
}

func AddTodo(path string, content string) (Todo, error) {
	data, err := GetTodoList(path)
	if err != nil {
		return Todo{}, err
	}
	seed++
	todo := Todo{
		ID:        seed,
		Content:   content,
		Completed: false,
	}
	newTodoList := append(data, todo)
	todolistJson, err2 := SerializeJSON(newTodoList)
	if err2 != nil {
		return Todo{}, err2
	}
	if err3 := WriteFile(path, todolistJson); err3 != nil {
		return Todo{}, err3
	}
	return todo, nil
}

func FindTodo(todolist []Todo, id int) (Todo, int) {
	targetTodo := Todo{}
	index := -1
	for i, currentTodo := range todolist {
		if currentTodo.ID == id {
			index = i
			targetTodo = currentTodo
			break
		}
	}
	if index == -1 {
		return targetTodo, -1
	}
	return targetTodo, index
}

func RemoveTodo(path string, id int) (Todo, error) {
	data, err := GetTodoList(path)
	if err != nil {
		return Todo{}, err
	}
	todo, index := FindTodo(data, id)
	if index == -1 {
		return Todo{}, errors.New("未找到数据")
	}
	newTodoList := slices.Delete(slices.Clone(data), index, index+1)
	todolistJson, err2 := SerializeJSON(newTodoList)
	if err2 != nil {
		return Todo{}, err2
	}
	if err3 := WriteFile(path, todolistJson); err3 != nil {
		return Todo{}, err3
	}
	return todo, nil
}

func ToggleTodo(path string, id int) (Todo, error) {
	data, err := GetTodoList(path)
	if err != nil {
		return Todo{}, err
	}
	todo, index := FindTodo(data, id)
	todo.Completed = !todo.Completed

	newTodoList := slices.Replace(slices.Clone(data), index, index+1, todo)
	todolistJson, err2 := SerializeJSON(newTodoList)
	if err2 != nil {
		return Todo{}, err2
	}
	if err3 := WriteFile(path, todolistJson); err3 != nil {
		return Todo{}, err3
	}
	return todo, nil
}

func ClearCompletedTodos(path string) ([]Todo, error) {
	data, err := GetTodoList(path)
	if err != nil {
		return []Todo{}, err
	}
	newTodoList := slices.DeleteFunc(data, func(todo Todo) bool {
		return todo.Completed == true
	})
	todolistJson, err2 := SerializeJSON(newTodoList)
	if err2 != nil {
		return []Todo{}, err2
	}
	if err3 := WriteFile(path, todolistJson); err3 != nil {
		return []Todo{}, err3
	}
	return newTodoList, nil
}

func ClearTodoList(path string) error {
	return WriteFile(path, []byte("[]"))
}
