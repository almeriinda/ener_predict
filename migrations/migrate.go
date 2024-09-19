package main

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Consumption struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	UserID          uint      `gorm:"not null" json:"user_id"`
	ConsumptionDate time.Time `json:"consumption_date"`
	Amount          float64   `gorm:"not null" json:"amount"`
	Month           int       `json:"month"`
	Day             int       `json:"day"`
	Temperature     float64   `json:"temperature"`
	UsageHours      float64   `json:"usage_hours"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	User            User      `gorm:"foreignKey:UserID"`
}

func main() {
	dsn := "host=localhost user=postgres password=password dbname=ener_predict port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar ao banco de dados:", err)
	}

	err = db.AutoMigrate(&User{}, &Consumption{})
	if err != nil {
		log.Fatal("Falha ao migrar banco de dados:", err)
	}

	log.Println("Migração concluída com sucesso!")
}
