package postgres

import (
	"database/sql"
	"fmt"
	"time"

	pb "github.com/Javokhdev/Auth-Service/genprotos"

	"github.com/google/uuid"
)

type UsersStorage struct {
	db *sql.DB
}

func NewUsersStorage(db *sql.DB) *UsersStorage {
	return &UsersStorage{db: db}
}

func (p *UsersStorage) CreateUser(user *pb.Users) (*pb.Void, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO users (id, username, email, password)
		VALUES ($1, $2, $3, $4)
	`
	_, err := p.db.Exec(query, id, user.Username, user.Email, user.Password)
	return nil, err
}

func (p *UsersStorage) GetByIdUser(id *pb.ById) (*pb.Users, error) {
	query := `
			SELECT username, email from users 
			where id =$1 and deleted_at=0
		`
	row := p.db.QueryRow(query, id.Id)

	var user pb.Users

	err := row.Scan(&user.Username,
		&user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (p *UsersStorage) GetAllUser(us *pb.Users) (*pb.GetAllUsers, error) {
	users := &pb.GetAllUsers{}
	var arr []interface{}
	count:=1
	
	query:=` SELECT username, email from users 
	where deleted_at=0 `
	if len(us.Email) > 0 {
		query += fmt.Sprintf(" and email=$%d", count)
		count++
		arr = append(arr, us.Email)
	}
	if len(us.Username) > 0 {
		query += fmt.Sprintf(" and username=$%d", count)
		count++
		arr = append(arr, us.Username)
	}

	if len(us.Id) > 0 {
		query += fmt.Sprintf(" and id=$%d", count)
		count++
		arr = append(arr, us.Id)
	}
	row, err := p.db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var user pb.Users
		err = row.Scan(&user.Username,
			&user.Email)
		if err != nil {
			return nil, err
		}
		users.Users = append(users.Users, &user)
	}
	return users, nil
}

func (p *UsersStorage) UpdateUser(user *pb.Users) (*pb.Void, error) {
	query := `
		UPDATE users
		SET username = $1, email = $2, password = $3
		WHERE id = $4 
	`
	_, err := p.db.Exec(query, user.Username, user.Email, user.Password, user.Id)
	return nil, err
}

func (p *UsersStorage) DeleteUser(id *pb.ById) (*pb.Void, error) {
	query := `
		update users set deleted_at=$2
		where id = $1
	`
	_, err := p.db.Exec(query, id.Id, time.Now().Unix())
	return nil, err
}

func (p *UsersStorage) LoginUser(userName *pb.Users) (*pb.Users, error) {
	query := `
			SELECT username, email from users 
			where username =$1 and deleted_at=0
		`
	row := p.db.QueryRow(query, userName.Username)

	var user pb.Users

	err := row.Scan(&user.Username,
					&user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}