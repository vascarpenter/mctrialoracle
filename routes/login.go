package routes

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/labstack/echo-contrib/session"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type loginHTMLtemplate struct {
	Title  string
	NoUser string
	CSS    string
}

// LoginRouter  GET "/login" を処理
func LoginRouter(c echo.Context) error {

	htmlvariable := loginHTMLtemplate{
		Title:  "多施設臨床トライアルシステム ログイン",
		NoUser: "",
		CSS:    "/css/login.css",
	}
	return c.Render(http.StatusOK, "login", htmlvariable)
}

// LoginRouterPost  POST "/login" を処理
func LoginRouterPost(c echo.Context) error {

	db := Repository()
	defer db.Close()
	ctx := context.Background()

	userID := c.FormValue("userid")
	pass := c.FormValue("password")

	errStr := "指定されたユーザIDが存在しません"

	// SQL: select * from hospitals where userid = req.userid limit 1
	var userpass string
	var newhospid int
	err := db.QueryRowContext(ctx, "SELECT userpass,hospital_id FROM HOSPITALS WHERE userid = :1", userID).Scan(&userpass, &newhospid)

	//fmt.Printf("ID:%+v,Pass:%+v,Err:%+v,InDBPass:%+v,InDBID:%+v\n", userID, pass, err, userpass, newhospid)

	if err == sql.ErrNoRows {
		//errStr = "指定されたユーザIDが存在しません"
	} else if err != nil {
		panic(err)
	} else if err = bcrypt.CompareHashAndPassword([]byte(userpass), []byte(pass)); err != nil {
		errStr = "パスワードが間違っています"
	} else {
		// login success; create session and redirect to "/"
		session, _ := session.Get("oursession", c)
		session.Values["userid"] = newhospid
		err = session.Save(c.Request(), c.Response())

		if newhospid == 1 {
			return c.Redirect(http.StatusFound, "/admin")
		}

		return c.Redirect(http.StatusFound, "/")
	}

	htmlvariable := loginHTMLtemplate{
		Title:  "多施設臨床トライアルシステム ログイン",
		NoUser: errStr,
		CSS:    "/css/login.css",
	}
	return c.Render(http.StatusOK, "login", htmlvariable)
}
