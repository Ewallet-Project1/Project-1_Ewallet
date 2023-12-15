package entities

type User struct {
	ID        int
	FullName  string
	Phone     string
	Password  string
	Address   string
	Balance   uint64
	CreatedAt []uint8
}
