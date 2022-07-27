package infrastructure

import (
	"fmt"
	"perfect_todo/domain"

	"github.com/jinzhu/gorm"
)

func DbInit() *gorm.DB {
    db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/perfect_todo?parseTime=true")
    if err != nil {
        fmt.Errorf("could not open database")
    }
    db.AutoMigrate(&domain.Todo{})
    return db
}

func DbCreate(todo domain.Todo) {
    db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/perfect_todo?parseTime=true")
    if err != nil {
        fmt.Errorf("could not open database")
    }
    db.Create(&todo)
}

func DbRead(id ...int) []domain.Todo {
    db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/perfect_todo?parseTime=true")
    if err != nil {
        fmt.Errorf("could not open database")
    }
    var todos []domain.Todo
    db.Find(&todos)
    return todos
}

func DbUpdate(id int, text string, status domain.Status, deadline int) domain.Todo {
    db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/perfect_todo?parseTime=true")
    if err != nil {
        fmt.Errorf("could not open database")
    }
    var todo domain.Todo
    db.First(&todo, id)
    todo.Text = text
    todo.Status = status
    todo.Deadline = deadline
    db.Save(&todo)
    return todo
}

func DbDelete(id int) {
    db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/perfect_todo?parseTime=true")
    if err != nil {
        fmt.Errorf("could not open database")
    }
    var todo domain.Todo
    db.First(&todo, id)
    db.Delete(&todo)
}