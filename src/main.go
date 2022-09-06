package main

type Product struct {
	Model
	Code  string
	Price uint
}

func main() {
	db, err := Open("mssql", "sqlserver://gonf:gonf@localhost:1433?database=gonf")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "L1212", Price: 1000})

	var product Product
	db.First(&product, 1)                   // find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212
	db.Model(&product).Update("Price", 2000)

	db.Delete(&product)
}
