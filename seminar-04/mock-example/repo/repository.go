package repo

type User struct {
	ID   int
	Name string
}

type UserRepository interface {
	GetUserByID(id int) (*User, error)
}

func GetUserName(repo UserRepository, id int) (string, error) {
	user, err := repo.GetUserByID(id)
	if err != nil {
		return "", err
	}
	return user.Name, nil
}
