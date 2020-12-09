package routes

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type resetpassHTMLtemplate struct {
	Title        string
	HospitalName string
	ErrPassRest  string
	CSS          string
}

// ResetPassRouter  GET "/resetpass" を処理
func ResetPassRouter(c echo.Context) error {

	hospname := c.Get("UserName").(string)
	htmlvariable := resetpassHTMLtemplate{
		Title:        "パスワードの変更",
		HospitalName: hospname,
		ErrPassRest:  "",
		CSS:          "/css/resetpass.css",
	}
	return c.Render(http.StatusOK, "resetpass", htmlvariable)
}

// ResetPassRouterPost  POST "/resetpass" を処理
func ResetPassRouterPost(c echo.Context) error {
	hospid := c.Get("UserID").(int)
	hospname := c.Get("UserName").(string)
	htmlvariable := resetpassHTMLtemplate{
		Title:        "パスワードの変更",
		HospitalName: hospname,
		ErrPassRest:  "",
		CSS:          "/css/resetpass.css",
	}

	db := Repository()
	defer db.Close()
	ctx := context.Background()

	oldpass := c.FormValue("oldpass")
	newpass := c.FormValue("newpass")

	var userpass string
	err := db.QueryRowContext(ctx, "SELECT userpass FROM hospitals WHERE hospital_id = :1", hospid).Scan(&userpass)

	if err != nil || err == sql.ErrNoRows {
		htmlvariable.ErrPassRest = "ユーザが存在しません"
		return c.Render(http.StatusOK, "resetpass", htmlvariable)
	}
	if err = bcrypt.CompareHashAndPassword([]byte(userpass), []byte(oldpass)); err != nil {
		htmlvariable.ErrPassRest = "パスワードが間違っています"
		return c.Render(http.StatusOK, "resetpass", htmlvariable)
	}

	newpasscrypt, _ := bcrypt.GenerateFromPassword([]byte(newpass), 10)

	tx, _ := db.BeginTx(ctx, nil)
	if _, err := tx.Exec("UPDATE hospitals SET userpass = :1 WHERE hospital_id = :2", string(newpasscrypt), hospid); err != nil {
		panic(err)
	}
	if err := tx.Commit(); err != nil {
		panic(err)
	}

	htmlvariable.ErrPassRest = "パスワードが変更されました"
	return c.Render(http.StatusOK, "resetpass", htmlvariable)
}
