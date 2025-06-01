package repository

import(

	"github.com/nunusavi/task-manager/internal/model"
)

func CreateUser(user *model.User) error{
	query := `INSERT INTO users (email,password) VALUES ($1,$2) RETURNING id, created_at`
	return DB.QueryRowx(query, user.Email, user.Password).Scan(&user.ID, &user.CreatedAt)
}

func GetUserByEmail(email string) (*model.User, error){
	var user model.User
	query := `SELECT * FROM users WHERE email = $1 LIMIT 1`
	err := DB.Get(&user, query, email)
	if err != nil{
		return nil, err
	}
	return &user, nil
}