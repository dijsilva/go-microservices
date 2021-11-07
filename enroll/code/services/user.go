package services

import (
	"enroll/appErrors"
	"enroll/database"
	"enroll/database/entities"
	"enroll/helpers"
	"enroll/providers"
	"fmt"

	"github.com/google/uuid"
)

type UserService struct{}

type CreateUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginUserInput struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserReturn struct {
	Id       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
	Name     string    `json:"name"`
	Profile  string    `json:"profile"`
	Token    string    `json:"token"`
}

func findByEmail(email string) appErrors.ErrorResponse {
	databaseConnection := database.Database.Connection
	var user entities.User
	var err appErrors.ErrorResponse
	databaseConnection.Find(&user, map[string]interface{}{"email": email})

	if user != (entities.User{}) {
		return appErrors.AlreadyExists("User already exists")
	}
	return err
}

func findByUserName(username string) appErrors.ErrorResponse {
	databaseConnection := database.Database.Connection
	var user entities.User
	var err appErrors.ErrorResponse
	databaseConnection.Find(&user, map[string]interface{}{"username": username})

	if user != (entities.User{}) {
		return appErrors.AlreadyExists("User already exists")
	}
	return err
}

func findByUserNameOrEmail(username string, email string) appErrors.ErrorResponse {
	databaseConnection := database.Database.Connection
	var user entities.User
	var err appErrors.ErrorResponse
	databaseConnection.Where("email = ?", email).Or("username = ?", username).First(&user)

	if user != (entities.User{}) {
		return appErrors.AlreadyExists("User already exists")
	}
	return err
}

func (us *UserService) CreateUser(input *CreateUserInput) appErrors.ErrorResponse {
	databaseConnection := database.Database.Connection
	var user entities.User
	var err appErrors.ErrorResponse

	errorAlreadyExists := findByUserNameOrEmail(input.Username, input.Email)

	if errorAlreadyExists.Message != "" {
		return errorAlreadyExists
	}

	var hashPass string
	var errHash error
	hashPass, errHash = helpers.GenerateHash(input.Password)
	if errHash != nil {
		return appErrors.InternalServerError("")
	}
	user = entities.User{
		Id:        uuid.New(),
		Name:      input.Name,
		Username:  input.Username,
		Email:     input.Email,
		Password:  hashPass,
		ProfileID: 2,
	}
	databaseConnection.Create(&user)
	return err
}

func (us *UserService) LoginUser(input *LoginUserInput) (LoginUserReturn, appErrors.ErrorResponse) {
	databaseConnection := database.Database.Connection
	user := entities.User{}
	var userReturn LoginUserReturn
	var err appErrors.ErrorResponse
	if input.Email != "" {
		if databaseConnection.Joins("Profile").Where(&entities.User{
			Email: input.Email,
		}).First(&user).Error == nil {
			if helpers.CheckPassWord(input.Password, user.Password) {
				authControl := providers.AuthControl{}
				tokenResp, errorAuthControl := authControl.GetToken(providers.CreateTokenInput{
					UserId:    user.Id.String(),
					Profile:   user.Profile.ProfileName,
					UserMail:  user.Email,
					UserName:  user.Username,
					TokenKind: fmt.Sprintf("LOGIN_%s", user.Profile.ProfileName),
				})
				if errorAuthControl.Message != "" {
					err = errorAuthControl
				}
				userReturn = LoginUserReturn{
					Id:       user.Id,
					Email:    user.Email,
					Username: user.Username,
					Name:     user.Name,
					Profile:  user.Profile.ProfileName,
					Token:    tokenResp.Data.Token,
				}
			} else {
				err = appErrors.IncorrectCredentials("")
			}
		} else {
			err = appErrors.NotFound("User not found")

		}
	} else if input.Username != "" {
		if databaseConnection.Preload("Profiles").Where(&entities.User{
			Username: input.Username,
		}).Find(&user).Error != nil {
			if helpers.CheckPassWord(input.Password, user.Password) {
				// var token string
				// token, err = helpers.GenerateJwtToken(user)
				userReturn = LoginUserReturn{
					Id:       user.Id,
					Email:    user.Email,
					Username: user.Username,
					Name:     user.Name,
					Profile:  user.Profile.ProfileName,
					Token:    "token",
				}
			} else {
				err = appErrors.IncorrectCredentials("Incorrect credentials")
			}
		} else {
			err = appErrors.NotFound("User not found")
		}
	} else {
		err = appErrors.BadInput("Username or email should be informed")
	}

	return userReturn, err
}

type User struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Username    string    `json:"username"`
	ProfileName string    `json:"profile_name"`
	ProfileId   uint      `json:"profile_id"`
}

func (us *UserService) ListUsers() []User {
	databaseConnection := database.Database.Connection
	var users []User

	databaseConnection.
		Select("users.id,users.name,users.email,users.username,users.profile_id, profiles.profile_name, profiles.id as profile_id").
		Joins("left join profiles on profiles.id = users.profile_id").Find(&users)
	return users
}
