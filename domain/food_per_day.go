package domain

type FoodPerDay struct {
	Id      int
	UsersID int
	Users   Users
	FoodsID int
	Foods   Foods
	Date    int
}
