package token

type Token interface {
	GenerateToken(id int) (string, error)
	ExtractUserId(token string) (int, error)
}
