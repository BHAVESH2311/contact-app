package user

import "fmt"

 

// User represents a user with UserID, Username, Password, and Role.

type User struct {

    UserID   int

    Username string

    Password string

    Role     string

}

 

// NewUser creates a new User instance.

func NewUser(userID int, username, password, role string) *User {

    return &User{

        UserID:   userID,

        Username: username,

        Password: password,

        Role:     role,

    }

}

 

// NewAdminUser creates a new Admin User instance.

func NewAdminUser(userID int, username, password string) *User {

    return NewUser(userID, username, password, "Admin")

}

 

// NewStaffUser creates a new Staff User instance.

func NewStaffUser(userID int, username, password string) *User {

    return NewUser(userID, username, password, "Staff")

}

 

// UserRepository is a placeholder for user storage.

type UserRepository struct {

    users []*User

}

 

// CreateUser creates a new user and stores it in the UserRepository.

func (ur *UserRepository) CreateUser(username, password, role string) *User {

    userID := len(ur.users) + 1

    newUser := NewUser(userID, username, password, role)

    ur.users = append(ur.users, newUser)

    return newUser

}

 

// UpdateUser updates user information based on UserID.

func (ur *UserRepository) UpdateUser(userID int, newUsername, newPassword, newRole string) {

    for _, user := range ur.users {

        if user.UserID == userID {

            user.Username = newUsername

            user.Password = newPassword

            user.Role = newRole

            return

        }

    }

}

 

// DeleteUser deletes a user based on UserID.

func (ur *UserRepository) DeleteUser(userID int,currentUserRole string) {

	if currentUserRole != "Admin" {         
		fmt.Println("You don't have permission to delete users.")         
		return    }

    for i, user := range ur.users {

        if user.UserID == userID {

            ur.users = append(ur.users[:i], ur.users[i+1:]...)

            return

        }

    }

}

 

// GetUserByID retrieves a user by UserID.

func (ur *UserRepository) GetUserByID(userID int) *User {

    for _, user := range ur.users {

        if user.UserID == userID {

            return user

        }

    }

    return nil // User not found

}

 