package user

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gotcha/model"
	"reflect"
	"testing"
)

func TestCreateNewUser(t *testing.T) {
	type args struct {
		username string
		name     string
		password string
		email    string
		role     string
	}
	tests := []struct {
		name string
		args args
		err  error
		user model.User
	}{
		{
			name: "Create a user that does not exist",
			args: args{
				username: "waki",
				name:     "wiaan",
				password: "password1",
				email:    "info@example.com",
				role:     "user",
			},
			err: nil,
			user: model.User{
				Username: "waki",
				Name:     "wiaan",
				Password: "password1",
				Email:    "info@example.com",
				Role:     "user",
			},
		},
	}
	// Mock DB for tests
	mockDb, _, _ := sqlmock.New()
	translator := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(translator, &gorm.Config{})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, user := CreateNewUser(db, tt.args.username, tt.args.name, tt.args.password, tt.args.email, tt.args.role)

			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("CreateNewUser() 'err' = %v, want %v", err, tt.err)
			}
			if !reflect.DeepEqual(user, tt.user) {
				t.Errorf("CreateNewUser() 'user' = %v, want %v", user, tt.user)
			}
		})
	}
}
