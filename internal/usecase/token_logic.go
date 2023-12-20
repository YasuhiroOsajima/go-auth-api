package usecase

type ITokenLogic interface {
	GenerateToken(userId int) (string, error)
	ExtractUserId(token string) (int, error)
}
