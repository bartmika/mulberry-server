package repositories

import (
	"encoding/json"

	"github.com/sdomino/scribble"
	"github.com/bartmika/mulberry-server/pkg/models"
)

// UserRepo implements models.UserRepository
type UserRepo struct {
	db *scribble.Driver
}

func NewUserRepo(db *scribble.Driver) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) Create(uuid string, name string, email string, passwordHash string) error {
	u := models.User{
		Uuid: uuid,
		Name: name,
		Email: email,
		PasswordHash: passwordHash,
	}
	if err := r.db.Write("users", uuid, u); err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) FindByUuid(uuid string) (*models.User, error) {
    u := models.User{}
	if err := r.db.Read("users", uuid, &u); err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepo) FindByEmail(email string) (*models.User, error) {
	records, err := r.db.ReadAll("users")
	if err != nil {
        return nil, err
	}

    for _, f := range records {
		userFound := models.User{}
        if err := json.Unmarshal([]byte(f), &userFound); err != nil {
            return nil, err
		}
		if userFound.Email == email {
			return &userFound, nil
		}
	}
	return nil, nil
}

func (r *UserRepo) Save(user *models.User) error {
	if err := r.db.Write("users", user.Uuid, user); err != nil {
		return err
	}
	return nil
}
