package service

import (
	"golang.org/x/crypto/bcrypt"
	"myapp/internal/models"
	"myapp/internal/repository"
	"myapp/internal/shared/payloads"
	"myapp/pkg/auth"
	"os"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

// HashPassword encrypts controllers_post password
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// CheckPassword checks user password
func checkPassword(passwordInDB, providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordInDB), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) SignUp(payload *payloads.SignUpPayload) (error) {
	hashedPass, err := hashPassword(payload.Password)
	if err != nil {
		return err
	}

	user := &models.User{
		FirstName: payload.FirstName, LastName: payload.LastName, Email: payload.Email, Password: hashedPass,
	}
	err = s.repo.CreateUserRecord(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) SignIn(payload *payloads.SignInPayload) (string, error) {
	user, err := s.repo.FindUserByEmail(&payload.Email)
	if err != nil {
		return "", err
	}

	err = checkPassword(user.Password, payload.Password)
	if err != nil {
		return "", err
	}

	jwtWrapper := auth.JwtWrapper{
		SecretKey:       os.Getenv("ACCESS_TOKEN_SECRET_KEY"),
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}
	signedToken, err := jwtWrapper.GenerateToken(user.Email, user.ID)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *UserService) GetUserProfile(email string) (*models.User, error) {
	return s.repo.FindUserByEmail(&email)
}
