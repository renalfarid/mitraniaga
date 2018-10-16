package main

import (
	"fmt"
	"net/http"
	"github.com/qor/qor"
	"github.com/qor/admin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Name string
}

type Product struct {
	gorm.Model
	Name string
	Description string
}

func main()  {

	DB, _ := gorm.Open("sqlite3", "demo.db")
	DB.AutoMigrate(&User{}, &Product{})

	Admin := admin.New(&admin.AdminConfig{DB: DB})

	Admin.AddResource(&User{})
	Admin.AddResource(&Product{})

	// Initalize an HTTP request multiplexer
	mux := http.NewServeMux()

	// Mount admin to the mux
	Admin.MountTo("/admin", mux)

	fmt.Println("Listening on: 9000")
	http.ListenAndServe(":9000", mux)

}






