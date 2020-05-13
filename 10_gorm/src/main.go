package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func gormOpen() (*gorm.DB, error) {
	return gorm.Open(
		"postgres",
		"host=database port=5432 user=dbuser password=dbpassword sslmode=disable",
	)
}

func dbHealthcheck() bool {
	db, err := gormOpen()
	defer func() {
		db.Close()
	}()
	return err == nil
}

type User struct {
	ID   uint
	Name string
	Age  uint
}

// Repository
////////////////////////////////////////////////////////////////////////////////
type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo() *UserRepo {
	db, _ := gormOpen()
	return &UserRepo{
		db: db,
	}
}

func (self UserRepo) Migrate() {
	self.db.AutoMigrate(&User{})
}

func (self UserRepo) Create(user *User) {
	self.db.Create(user)
}

func (self UserRepo) Find(id uint) *User {
	user := User{}
	self.db.First(&user, id)
	return &user
}

func (self UserRepo) FindAll() []User {
	users := []User{}
	self.db.Find(&users)
	return users
}

func (self UserRepo) Delete(id uint) {
	user := self.Find(id)
	self.db.Delete(user)
}

func (self UserRepo) Close() {
	self.db.Close()
}

// Entry point
////////////////////////////////////////////////////////////////////////////////
func main() {
	dbHealthcheck()
	repo := NewUserRepo()
	defer repo.Close()

	repo.Migrate()

	user := User{
		Name: "GormTutorialUser",
		Age:  17,
	}
	repo.Create(&user)

	u := repo.Find(user.ID)
	fmt.Printf("Created User : %v\n", u)

	users := repo.FindAll()
	for idx, usr := range users {
		fmt.Printf("All Users[%d] : %v\n", idx, usr)
	}

	repo.Delete(u.ID)
}
