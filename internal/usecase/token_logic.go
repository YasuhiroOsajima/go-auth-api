package usecase

type ITokenLogic interface {
	GenerateToken(userId int) (string, error)
	ExtractUserId(bearToken string) (int, error)
}
