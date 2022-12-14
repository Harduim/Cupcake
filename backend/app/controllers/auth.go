package controllers

import (
	"cupcake/app/database"
	domain "cupcake/app/models"
	"cupcake/app/repositories"
	"cupcake/app/service"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type MsalToken struct {
	jwt.Claims
	Aud        string   `json:"aud"`
	Iss        string   `json:"iss"`
	Iat        string   `json:"iat"`
	Nbf        string   `json:"nbf"`
	Exp        int      `json:"exp"`
	Acr        string   `json:"acr"`
	Aio        string   `json:"aio"`
	Amr        []string `json:"amr"`
	Appid      string   `json:"appid"`
	Appidacr   string   `json:"appidacr"`
	Email      string   `json:"email"`
	FamilyName string   `json:"family_name"`
	GivenName  string   `json:"given_name"`
	Idp        string   `json:"idp"`
	Ipaddr     string   `json:"ipaddr"`
	Name       string   `json:"name"`
	Oid        string   `json:"oid"`
	Rh         string   `json:"rh"`
	Scp        string   `json:"scp"`
	Sub        string   `json:"sub"`
	Tid        string   `json:"tid"`
	UniqueName string   `json:"unique_name"`
	Uti        string   `json:"uti"`
	Ver        string   `json:"ver"`
}

type AuthResponse struct {
	Token     string `json:"access_token"`
	TokenType string `json:"token_type"`
}

func generateJWT(user *domain.User, secretKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * 86400).Unix()
	claims["admin"] = user.IsAdmin
	claims["user"] = user.ID
	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func getInfoFromSSOToken(token string) *MsalToken {
	msalToken := MsalToken{}
	_, _, _ = new(jwt.Parser).ParseUnverified(token, &msalToken)
	return &msalToken
}

func CreateAccount(db repositories.UserRepositoryDb, name string, oid string) (*domain.User, error) {
	isAdmin := false
	newUser, err := domain.NewUser(name, oid, &isAdmin)

	if err != nil {
		return nil, err
	}

	user, err := db.Insert(newUser)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func AuthenticateSSO(sso *service.SSOClient) fiber.Handler {

	return func(ctx *fiber.Ctx) error {

		url, err := sso.AuthCodeURL(ctx.Context())

		if err != nil {
			return err
		}

		return ctx.Redirect(url, 302)
	}
}

func Token(sso *service.SSOClient, db *database.Database, secretKey string) fiber.Handler {
	repo := repositories.UserRepositoryDb{Db: db}

	return func(ctx *fiber.Ctx) error {

		code := ctx.Query("code")
		result, err := sso.App.AcquireTokenByAuthCode(ctx.Context(), code, sso.RedirectUrl, sso.Scopes)
		if err != nil {
			err := ctx.Status(fiber.StatusUnauthorized).SendString(err.Error())
			if err != nil {
				return err
			}
		}

		if err != nil {
			return err
		}

		msalToken := getInfoFromSSOToken(result.AccessToken)
		userOid := msalToken.Oid
		userName := msalToken.Name

		user, err := repo.Find(userOid)

		if err != nil {
			user, err = CreateAccount(repo, userName, userOid)
		}

		if err != nil {
			return err
		}

		accessToken, err := generateJWT(user, secretKey)

		if err != nil {
			err := ctx.Status(fiber.StatusUnauthorized).SendString(err.Error())
			if err != nil {
				return err
			}
		}

		return ctx.Status(fiber.StatusOK).Redirect("/login?tkn=" + accessToken)
	}
}
