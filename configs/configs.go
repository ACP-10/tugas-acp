package configs

import (
	"fmt"
	"tugas-acp/models/cart"
	"tugas-acp/models/category"
	"tugas-acp/models/customer"
	"tugas-acp/models/payment"

	cartitem "tugas-acp/models/cartItem"
	"tugas-acp/models/product"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type ConfigDB struct {
	DB_Username string
	DB_Password string
	DB_Host     string
	DB_Port     string
	DB_Database string
}

func InitConfigDB() ConfigDB {
	var configDB = ConfigDB{
		DB_Username: "root",
		DB_Password: "",
		// DB_Password: "Ikhda123", // comment this
		DB_Host:     "localhost",
		DB_Port:     "3306",
		DB_Database: "acp10",
	}
	return configDB
}

func InitDB() {
	configDB := InitConfigDB()

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		configDB.DB_Username,
		configDB.DB_Password,
		configDB.DB_Host,
		configDB.DB_Port,
		configDB.DB_Database)

	var error error
	DB, error = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if error != nil {
		panic("Database failed connection : " + error.Error())
	}
	Migration()
}

func Migration() {
	DB.AutoMigrate(&cart.Cart{})
	DB.AutoMigrate(&cartitem.CartItem{})
	DB.AutoMigrate(&category.Category{})
	DB.AutoMigrate(&customer.Customer{})
	DB.AutoMigrate(&product.Product{})
	DB.AutoMigrate(&payment.Payment{})
}

//========Test

func InitConfigDBTest() ConfigDB {
	var configDB = ConfigDB{
		DB_Username: "root",
		DB_Password: "",
		DB_Host:     "127.0.0.1",
		DB_Port:     "3306",
		DB_Database: "acp10_test",
	}
	return configDB
}

func InitDBTest() {
	configDB := InitConfigDBTest()

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		configDB.DB_Username,
		configDB.DB_Password,
		configDB.DB_Host,
		configDB.DB_Port,
		configDB.DB_Database)

	var error error
	DB, error = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if error != nil {
		panic("Database failed connection : " + error.Error())
	}
	MigrationTest()
}

func MigrationTest() {

	DB.Migrator().DropTable(
		&cart.Cart{},
		&cartitem.CartItem{},
		&customer.Customer{},
		&product.Product{},
		&category.Category{})
	DB.AutoMigrate(
		&cart.Cart{},
		&cartitem.CartItem{},
		&customer.Customer{},
		&product.Product{},
		&category.Category{})
}
