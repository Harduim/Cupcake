package controllers

import (
	"cupcake/app/database"
	"cupcake/app/domain"
	"cupcake/app/repositories"
	"cupcake/app/service"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"time"
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
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["user"] = user.ID

	tokenString, err := token.SignedString(secretKey)

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

func CreateAccount(db repositories.UserRepositoryDb, name string, email string) *domain.User {
	newUser, err := domain.NewUser(name, email)

	if err != nil {
		return nil
	}

	user, err := db.Insert(newUser)

	if err != nil {
		return nil
	}

	return user
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
		userEmail := msalToken.Email
		userName := msalToken.Name

		user, err := repo.FindByEmail(userEmail)

		if err != nil && user.Email == "" {
			return err
		} else {
			user = CreateAccount(repo, userEmail, userName)
		}

		accessToken, err := generateJWT(user, secretKey)

		if err != nil {
			err := ctx.Status(fiber.StatusUnauthorized).SendString(err.Error())
			if err != nil {
				return err
			}
		}
		response := AuthResponse{
			Token:     accessToken,
			TokenType: "bearer",
		}

		u, err := json.Marshal(response)

		if err != nil {
			return err
		}

		return ctx.Status(fiber.StatusOK).SendString(string(u))
	}
}
