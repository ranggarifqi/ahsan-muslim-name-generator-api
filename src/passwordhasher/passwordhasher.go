package passwordhasher

type IPasswordHasherService interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword string, password string) (bool, error)
}
