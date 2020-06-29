package main

import (
	"context"
	"database/sql"
	"html/template"
	"io"
	"net/http"
	"os"

	"mctrialoracle/m/routes"

	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-oci8"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Template はHTMLテンプレートを利用するためのRenderer Interfaceです。
type Template struct {
	templates *template.Template
}

// Render はHTMLテンプレートにデータを埋め込んだ結果をWriterに書き込みます。
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// setUserMiddleware　cookieを参照して、ユーザがログインしていればdbにアクセスし名前、IDをcontextに入れる
func setUserMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			db := routes.Repository()
			defer db.Close()

			ctx := context.Background()

			if session, err := session.Get("oursession", c); err == nil {
				if userid, ok := session.Values["userid"]; ok {
					// cookie exist
					if useridint := userid.(int); useridint > 0 {
						// get hospital name
						var hospname string
						err := db.QueryRowContext(ctx, "SELECT name FROM hospitals WHERE hospital_id = :1", useridint).Scan(&hospname)
						//fmt.Printf("%+v\n", hospname)

						switch {
						case err == sql.ErrNoRows:
							// cannot get; pass through
						case err != nil:
							panic(err)
						default:
							c.Set("UserName", hospname)
							c.Set("UserID", useridint)
						}
					}
				}
			}

			return next(c)
		}
	}
}

// redirectLoginWithoutAuth  contextにIDが入っていないか 0であった場合は、login画面にリダイレクトする
func redirectLoginWithoutAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userid := c.Get("UserID")
			if userid == 0 || userid == nil {
				// not login'd, go to login page
				return c.Redirect(http.StatusFound, "/login")
			}
			return next(c)
		}
	}
}

func main() {
	// Echo instance
	e := echo.New()

	funcMap := template.FuncMap{
		"safehtml": func(text string) template.HTML { return template.HTML(text) },
	}
	t := &Template{
		templates: template.Must(template.New("").Funcs(funcMap).ParseGlob("views/*.html")),
	}
	e.Renderer = t

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} uri=${uri} path=${path} status=${status}\n",
	}))
	e.Use(middleware.Recover())

	var keystore string
	if keystore = os.Getenv("COOKIE_SEED"); keystore == "" {
		keystore = "secret key"
	}
	var store = sessions.NewCookieStore([]byte(keystore))
	e.Use(session.Middleware(store))

	// Routes
	e.Static("/css", "./static/css")
	e.Static("/img", "./static/img")
	e.Static("/javascript", "./static/javascript")
	e.GET("/register", routes.RegisterRouter)
	e.POST("/register", routes.RegisterRouterPost)
	e.GET("/resetpass", routes.ResetPassRouter, setUserMiddleware(), redirectLoginWithoutAuth())
	e.POST("/resetpass", routes.ResetPassRouterPost, setUserMiddleware(), redirectLoginWithoutAuth())
	e.GET("/patient", routes.PatientRouter, setUserMiddleware(), redirectLoginWithoutAuth())
	e.POST("/patient", routes.PatientRouterPost, setUserMiddleware(), redirectLoginWithoutAuth())
	e.GET("/login", routes.LoginRouter)
	e.POST("/login", routes.LoginRouterPost)
	e.GET("/logout", routes.LogoutRouter)
	e.GET("/stat", routes.StatRouter)
	e.GET("/admin", routes.AdminRouter, setUserMiddleware(), redirectLoginWithoutAuth())
	e.GET("/admin/:func", routes.AdminAnalyzeRouter, setUserMiddleware(), redirectLoginWithoutAuth())
	e.POST("/admin", routes.AdminRouterPost, setUserMiddleware(), redirectLoginWithoutAuth())
	e.GET("/patientedit/:hosp/:ser", routes.PatientEditRouter, setUserMiddleware(), redirectLoginWithoutAuth())
	e.POST("/patientedit/:hosp/:ser", routes.PatientEditRouterPost, setUserMiddleware(), redirectLoginWithoutAuth())
	e.GET("/eventedit/:hosp/:ser/:ev", routes.EventEditRouter, setUserMiddleware(), redirectLoginWithoutAuth())
	e.POST("/eventedit/:hosp/:ser/:ev", routes.EventEditRouterPost, setUserMiddleware(), redirectLoginWithoutAuth())
	e.GET("/event/:hosp/:ser", routes.EventRouter, setUserMiddleware(), redirectLoginWithoutAuth())
	e.POST("/event/:hosp/:ser", routes.EventRouterPost, setUserMiddleware(), redirectLoginWithoutAuth())
	e.GET("/eventlist/:hosp/:ser", routes.EventlistRouter, setUserMiddleware(), redirectLoginWithoutAuth())
	e.GET("/", routes.IndexRouter, setUserMiddleware(), redirectLoginWithoutAuth())

	// handle error
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			if he.Code == 404 {
				c.Render(http.StatusNotFound, "404", nil)
			} else {
				c.Render(http.StatusInternalServerError, "500", nil)
			}
		}
	}

	// Start server
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "3000"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
