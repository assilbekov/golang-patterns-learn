package db

import (
	"fmt"
	"log"
	"shared-db/conf"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
		db *gorm.DB
}

func New(conf *conf.Config) *DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=test port=%s sslmode=disable",
		conf.DB_HOST,
		conf.DB_USER,
		conf.DB_PASSWORD,
		conf.DB_PORT)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return &DB{
		db: db,
	}
}

func (d *DB) Migrate() error {
	// Define models to auto-migrate
	if err := d.db.AutoMigrate(&User{}, &UserWorkHours{}); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}
	return nil
}

func (d *DB) GetUsers() ([]User, error) {
	var users []User
	if err := d.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (d *DB) GetUser(id uint) (*User, error) {
	var user User
	if err := d.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (d *DB) CreateUser(user *User) error {
	return d.db.Create(user).Error
}

func (d *DB) UpdateUser(user *User) error {
	return d.db.Model(&User{}).Where("id = ?", user.ID).Updates(user).Error
}

func (d *DB) GetUserWorkHours(id uint) (*UserWorkHours, error) {
	var userWorkHours UserWorkHours
	if err := d.db.First(&userWorkHours, id).Error; err != nil {
		return nil, err
	}
	return &userWorkHours, nil
}

func (d *DB) CreateUserWorkHours(userWorkHours *UserWorkHours) error {
	return d.db.Create(userWorkHours).Error
}

func (d *DB) UpdateUserWorkHours(userWorkHours *UserWorkHours) error {
	return d.db.Model(&UserWorkHours{}).Where("id = ?", userWorkHours.ID).Updates(userWorkHours).Error
}
