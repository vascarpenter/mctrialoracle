package routes

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type registerHTMLtemplate struct {
	Title       string
	EmailExists string
	CSS         string
}

// RegisterRouter  GET "/register" を処理
func RegisterRouter(c echo.Context) error {

	htmlvariable := registerHTMLtemplate{
		Title:       "病院登録",
		EmailExists: "",
		CSS:         "/css/register.css",
	}
	return c.Render(http.StatusOK, "register", htmlvariable)
}

// RegisterRouterPost  POST "/register" を処理
func RegisterRouterPost(c echo.Context) error {
	db := Repository()
	defer db.Close()
	ctx := context.Background()

	htmlvariable := registerHTMLtemplate{
		Title:       "病院登録",
		EmailExists: "",
		CSS:         "/css/register.css",
	}

	hospName := c.FormValue("hospname")
	userid := c.FormValue("userid")
	userpass := c.FormValue("userpass")
	mailaddress := c.FormValue("mailaddress")
	createdAt := time.Now()

	// 'SELECT * FROM hospitals WHERE userid = ? LIMIT 1'
	var usercnt int
	err := db.QueryRowContext(ctx, "SELECT COUNT(*) FROM hospitals WHERE userid = :1", userid).Scan(&usercnt)

	if err == nil && usercnt > 0 {
		htmlvariable.EmailExists = "同じユーザーIDがすでに存在します"
		return c.Render(http.StatusOK, "register", htmlvariable)
	}

	userpasscrypt, err := bcrypt.GenerateFromPassword([]byte(userpass), 10)
	if err := db.QueryRowContext(ctx, "SELECT COUNT(*) FROM hospitals", userid).Scan(&usercnt); err != nil {
		panic(err)
	}

	tx, _ := db.BeginTx(ctx, nil)
	if _, err := tx.Exec(`INSERT INTO hospitals (hospital_id,"NAME",created_at,userid,userpass,mailaddress)
		values (:1,:2,TO_DATE(:3,'YYYY/MM/DD'),:4,:5,:6)`,
		usercnt+1,                      // 1: hospital_id
		hospName,                       // 2: NAME
		createdAt.Format("2006/01/02"), // 3: created_at
		userid,                         // 4: userid
		string(userpasscrypt),                  // 5: userpass
		mailaddress,                    // 6: mailaddress
	); err != nil {
		panic(err)
	}
	if err := tx.Commit(); err != nil {
		panic(err)
	}
	htmlvariable.EmailExists = "追加しました"
	return c.Render(http.StatusOK, "register", htmlvariable)
}
