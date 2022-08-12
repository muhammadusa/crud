package database

type Category struct {
	Id   int
	Name string
}

type Product struct {
	Id       int
	Name     string
	Price    float64
	Category Category
}

type PostProductBody struct {
	Name       string
	Price      float64
	CategoryId int
}

type DeleteProductBody struct {
	Id int
}

type PutProductBody struct {
	Id         int
	Name       string
	Price      float64
	CategoryId int
}
