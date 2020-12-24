package modules

// User structure
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

// GetUsers Gets all users
func GetUsers() []*User{
	return users
}

// AddUser to collection
func AddUser(u User) (User, error){
	u.ID = nextID
	nextID++
	users = append(users, &u)

	return u, nil
}


