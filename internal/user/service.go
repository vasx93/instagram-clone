package user

// UserService represents the service layer for user operations.
type UserService struct {
	repo *UserRepository // Embedding UserRepositoryInterface
}

// NewUserService creates a new instance of UserService.
func NewUserService(repo *UserRepository) *UserService {
	return &UserService{repo}
}

// GetAllUsers retrieves all users.
func (s *UserService) GetAll() ([]*User, error) {
	return s.repo.GetAll()
}
