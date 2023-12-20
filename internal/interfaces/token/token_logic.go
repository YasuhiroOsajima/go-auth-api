package token

type TokenLogic struct {
	token Token
}

func NewTokenLogic(token Token) *TokenLogic {
	return &TokenLogic{token}
}

func (t *TokenLogic) GenerateToken(userId int) (string, error) {
	return t.token.GenerateToken(userId)
}

func (t *TokenLogic) ExtractUserId(token string) (int, error) {
	return t.token.ExtractUserId(token)
}
