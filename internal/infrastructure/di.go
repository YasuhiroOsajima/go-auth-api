package infrastructure

import (
	infDb "github.com/YasuhiroOsajima/go-auth-api/internal/infrastructure/database"
	infToken "github.com/YasuhiroOsajima/go-auth-api/internal/infrastructure/token"
	"github.com/YasuhiroOsajima/go-auth-api/internal/interfaces/controllers"
	intDb "github.com/YasuhiroOsajima/go-auth-api/internal/interfaces/database"
	intToken "github.com/YasuhiroOsajima/go-auth-api/internal/interfaces/token"
	"github.com/YasuhiroOsajima/go-auth-api/internal/usecase"
)

var orm *infDb.Orm
var tkn *infToken.Token
var dbRepo *intDb.DatabaseRepository
var tknLogic *intToken.TokenLogic
var authInteractor *usecase.AuthInteractor
var userInteractor *usecase.UserInteractor
var authCtrl *controllers.AuthController

func init() {
	// err := godotenv.Load()
	// if err != nil {
	// 	panic("Error loading .env file")
	// }
	infDb.ConnectDataBase()

	orm = infDb.NewOrm()
	tkn = infToken.NewToken()

	dbRepo = intDb.NewDatabaseRepository(orm)
	tknLogic = intToken.NewTokenLogic(tkn)

	authInteractor = usecase.NewAuthInteractor(dbRepo, tknLogic)
	userInteractor = usecase.NewUserInteractor(dbRepo)

	authCtrl = controllers.NewAuthController(authInteractor, userInteractor)
}
