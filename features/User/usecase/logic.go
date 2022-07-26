package usecase

import (
	"backend/domain"
	"backend/features/User/data"
	"log"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userData domain.UserData
	validate *validator.Validate
}

func New(uuc domain.UserData, v *validator.Validate) domain.UserUseCase {
	return &userUseCase{
		userData: uuc,
		validate: v,
	}
}

func (uuc *userUseCase) SearchUser(username string) (domain.User, int) {
	search := uuc.userData.SearchUserData(username)

	if search.ID == 0 {
		log.Println("Data not found")
		return domain.User{}, 404
	}

	return search, 200
}

func (uuc *userUseCase) DeleteUser(id int) int {
	status := uuc.userData.DeleteUserData(id)

	if !status {
		log.Println("Data not found")
		return 404
	}

	return 200
}

// Register implements domain.UserUseCase
func (uuc *userUseCase) RegisterUser(newuser domain.User, cost int) int {
	var user = data.FromModel(newuser)
	validError := uuc.validate.Struct(user)

	if validError != nil {
		log.Println("Validation errror : ", validError.Error())
		return 400
	}

	hashed, hasherr := bcrypt.GenerateFromPassword([]byte(user.Password), cost)

	if hasherr != nil {
		log.Println("Cant encrypt: ", hasherr)
		return 500
	}
	user.Password = string(hashed)
	insert := uuc.userData.RegisterData(user.ToModel())

	if insert.ID == 0 {
		log.Println("Empty data")
		return 404
	}

	return 200
}

// UpdateUser implements domain.UserUseCase
func (uuc *userUseCase) UpdateUser(newuser domain.User, userid int, cost int) int {
	var user = data.FromModel(newuser)
	validError := uuc.validate.Struct(user)

	if validError != nil {
		log.Println("Validation errror : ", validError.Error())
		return 400
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), cost)

	if err != nil {
		log.Println("Error encrypt password", err)
		return 500
	}

	user.ID = uint(userid)
	user.Password = string(hashed)

	update := uuc.userData.UpdateUserData(user.ToModel())

	if update.ID == 0 {
		log.Println("Data not found")
		return 404
	}

	return 200
}
