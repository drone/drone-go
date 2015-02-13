package drone

import (
	"fmt"
)

type UserService struct {
	*Client
}

// GET /api/users/{host}/{login}
func (s *UserService) Get(remote, login string) (*User, error) {
	var path = fmt.Sprintf("/api/users/%s/%s", remote, login)
	var user = User{}
	var err = s.run("GET", path, nil, &user)
	return &user, err
}

// GET /api/user
func (s *UserService) GetCurrent() (*User, error) {
	var user = User{}
	var err = s.run("GET", "/api/user", nil, &user)
	return &user, err
}

// POST /api/users/{host}/{login}
func (s *UserService) Create(remote, login string) (*User, error) {
	var path = fmt.Sprintf("/api/users/%s/%s", remote, login)
	var user = User{}
	var err = s.run("POST", path, nil, &user)
	return &user, err
}

// DELETE /api/users/{host}/{login}
func (s *UserService) Delete(remote, login string) error {
	var path = fmt.Sprintf("/api/users/%s/%s", remote, login)
	return s.run("DELETE", path, nil, nil)
}

// GET /api/users
func (s *UserService) List() ([]*User, error) {
	var users []*User
	var err = s.run("GET", "/api/users", nil, &users)
	return users, err
}
