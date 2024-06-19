package postgres_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	pb "github.com/Javokhdev/Auth-Service/genprotos"
	"github.com/Javokhdev/Auth-Service/storage/postgres"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	storage := postgres.NewUsersStorage(db)

	mock.ExpectExec("INSERT INTO users").WithArgs(sqlmock.AnyArg(), "username1", "email1@example.com", "password1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	user := &pb.Users{
		Username: "username1",
		Email:    "email1@example.com",
		Password: "password1",
	}

	_, err = storage.CreateUser(user)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetByIdUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	storage := postgres.NewUsersStorage(db)

	rows := sqlmock.NewRows([]string{"username", "email"}).
		AddRow("username1", "email1@example.com")

	mock.ExpectQuery("SELECT username, email from users where id =").
		WithArgs("1").
		WillReturnRows(rows)

	id := &pb.ById{Id: "1"}
	user, err := storage.GetByIdUser(id)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "username1", user.Username)
	assert.Equal(t, "email1@example.com", user.Email)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	storage := postgres.NewUsersStorage(db)

	rows := sqlmock.NewRows([]string{"username", "email"}).
		AddRow("username1", "email1@example.com").
		AddRow("username2", "email2@example.com")

	mock.ExpectQuery("SELECT username, email from users where deleted_at=0").
		WillReturnRows(rows)

	users, err := storage.GetAllUser(&pb.Users{})
	assert.NoError(t, err)
	assert.NotNil(t, users)
	assert.Len(t, users.Users, 2)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	storage := postgres.NewUsersStorage(db)

	mock.ExpectExec("UPDATE users SET username =").WithArgs("username1", "email1@example.com", "password1", "1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	user := &pb.Users{
		Id:       "1",
		Username: "username1",
		Email:    "email1@example.com",
		Password: "password1",
	}

	_, err = storage.UpdateUser(user)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	storage := postgres.NewUsersStorage(db)

	mock.ExpectExec("update users set deleted_at=").WithArgs("1", sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	id := &pb.ById{Id: "1"}

	_, err = storage.DeleteUser(id)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestLoginUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	storage := postgres.NewUsersStorage(db)

	rows := sqlmock.NewRows([]string{"username", "email"}).
		AddRow("username1", "email1@example.com")

	mock.ExpectQuery("SELECT username, email from users where username =").
		WithArgs("username1").
		WillReturnRows(rows)

	user := &pb.Users{
		Username: "username1",
	}

	loginUser, err := storage.LoginUser(user)
	assert.NoError(t, err)
	assert.NotNil(t, loginUser)
	assert.Equal(t, "username1", loginUser.Username)
	assert.Equal(t, "email1@example.com", loginUser.Email)
	assert.NoError(t, mock.ExpectationsWereMet())
}
