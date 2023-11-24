package token

type Token interface {
	GenerateToken(id int) (string, error)
	ExtractUserId(bearToken string) (int, error)
}
