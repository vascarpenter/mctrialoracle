package routes

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type adminHTMLtemplate struct {
	Title        string
	HospitalName string
	ImageTag     string
	Message      string
	Redirect     string
	CSS          string
}

// GetHospAndUser  Hosp ID, Usernameをcontextから得る (global varの代用)
func GetHospAndUser(c echo.Context) (hospid int, username string) {
	hospid = c.Get("UserID").(int)
	if hospid > 0 {
		username = c.Get("UserName").(string)
	}
	return hospid, username
}

// AdminRouter  GET "/admin" を処理
func AdminRouter(c echo.Context) error {
	db := Repository()
	defer db.Close()

	hospid, username := GetHospAndUser(c)
	if hospid != 1 {
		// invalid access, go to logout
		return c.Redirect(http.StatusFound, "/logout")
	}

	htmlvariable := adminHTMLtemplate{
		Title:        "adminコンソール",
		HospitalName: username,
		Redirect:     "",
		Message:      "",
		ImageTag:     "",
		CSS:          "/css/index.css",
	}

	return c.Render(http.StatusOK, "admin", htmlvariable)

}

var wg sync.WaitGroup
var (
	m = &sync.Map{}
)
var id string

// StatRouter ダウンロードの進捗を JSON で返す
func StatRouter(c echo.Context) error {
	ck, err := c.Cookie("download-progress")
	if err != nil {
		log.Println(err)
		return err
	}
	progress := 0
	v, ok := m.Load(ck.Value)
	if ok {
		if vi, ok := v.(int); ok {
			progress = vi
		}
	}
	return c.JSON(http.StatusOK, &struct {
		Progress int `json:"progress"`
	}{
		Progress: progress,
	})
}

// AdminAnalyzeRouter  GET "/admin/:func" を処理
func AdminAnalyzeRouter(c echo.Context) error {
	db := Repository()
	defer db.Close()
	hospid, username := GetHospAndUser(c)
	if hospid != 1 {
		// invalid access, go to logout
		return c.Redirect(http.StatusFound, "/logout")
	}

	htmlvariable := adminHTMLtemplate{
		Title:        "adminコンソール",
		HospitalName: username,
		Redirect:     "",
		Message:      "",
		CSS:          "/css/index.css",
	}
	sel, _ := strconv.Atoi(c.Param("func"))
	switch sel {
	case 1:
		htmlvariable.Message = "now calculating"
		id = uuid.New().String()
		htmlvariable.Redirect = `<meta http-equiv="refresh" content="0;URL='/admin/2'" />`
		c.SetCookie(&http.Cookie{
			Name:  "download-progress",
			Value: id,
			Path:  "/",
		})
		// cookieを保存させる
	case 2:
		// https://mattn.kaoriya.net/software/lang/go/20170622160723.htm を参考にしました

		htmlvariable.Message = "now calculating"
		htmlvariable.Redirect = `
		<script>
		window.addEventListener('load', function() {
		  var prog=function progress() {
			  fetch("/stat", {
				'credentials': "same-origin"
			  }).then(function(response) {
				return response.json();
			  }).then(function(json) {
				document.querySelector('#progress').textContent = "progress: "+json.progress+"%";
			  if (json.progress >= 100) {
				  location.href = "/admin/3";
				  clearInterval();
			  }
			})
		  };
		  setInterval(prog,1000);
		}, false);
		</script>
		` // 起動時から 1000msec間隔で/statし、progressに表示、>=100なら自動起動を中止
		// 非同期で走らせる
		m.Store(id, 0)
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = analyze()
			m.Store(id, 100) // progress 100%
		}()
	case 3:
		wg.Wait()
		m.Store(id, nil)
		htmlvariable.Message = "finished"
		htmlvariable.ImageTag = "<img src=/img/test.png><br>"

		//delete cookie
		c.SetCookie(&http.Cookie{
			Name:    "download-progress",
			Value:   id,
			Path:    "/",
			Expires: time.Now().Add(-1 * time.Hour),
		})

	default:
		htmlvariable.Message = "エラー"
	}
	return c.Render(http.StatusOK, "admin", htmlvariable)

}

func analyze() string {
	db := Repository()
	defer db.Close()
	ctx := context.Background()

	output := ""
	rows2, _ := db.QueryContext(ctx, "SELECT hospital_id, serialid FROM patients")
	defer rows2.Close()

	for rows2.Next() {
		var hospid int
		var serialid int
		if err := rows2.Scan(&hospid, &serialid); err != nil {
			panic(err)
		}

		//fmt.Printf("%+v,%+v\n", hospid, serialid)
		rows, err := db.QueryContext(ctx, `SELECT dropout,alive,macce,TO_CHAR("DATE",'YYYY/MM/DD') from events WHERE hospital_id=:1 AND serialid=:2`, hospid, serialid)
		if err != nil {
			panic(err)
		}

		dropoutflag := false
		macceflag := false
		deadflag := false
		for rows.Next() {
			var dropout, alive, macce sql.NullInt32
			var eventdate sql.NullString
			if err := rows.Scan(&dropout, &alive, &macce, &eventdate); err != nil {
				panic(err)
			}

			if dropout.Valid && dropout.Int32 != 0 {
				dropoutflag = true
				tx, _ := db.BeginTx(ctx, nil)
				if _, err := tx.Exec("UPDATE patients SET dropdate = TO_DATE(:1,'YYYY/MM/DD') WHERE hospital_id = :2 AND serialid = :3", eventdate.String, hospid, serialid); err != nil {
					panic(err)
				}
				if err := tx.Commit(); err != nil {
					panic(err)
				}
				break
			}
			if alive.Valid && alive.Int32 == 0 {
				deadflag = true
				tx, _ := db.BeginTx(ctx, nil)
				if _, err := tx.Exec("UPDATE patients SET deaddate = TO_DATE(:1,'YYYY/MM/DD') WHERE hospital_id = :2 AND serialid = :3", eventdate.String, hospid, serialid); err != nil {
					panic(err)
				}
				if err := tx.Commit(); err != nil {
					panic(err)
				}
				break
			}
			if macce.Valid && macce.Int32 != 0 {
				macceflag = true
				tx, _ := db.BeginTx(ctx, nil)
				if _, err := tx.Exec("UPDATE patients SET maccedate = TO_DATE(:1,'YYYY/MM/DD') WHERE hospital_id = :2 AND serialid = :3", eventdate.String, hospid, serialid); err != nil {
					panic(err)
				}
				if err := tx.Commit(); err != nil {
					panic(err)
				}
				break
			}

		}
		if dropoutflag == false && deadflag == false && macceflag == false {
			datestr := time.Now().Format("2006/01/02")
			tx, _ := db.BeginTx(ctx, nil)
			if _, err := tx.Exec("UPDATE patients SET finishdate = TO_DATE(:1,'YYYY/MM/DD') WHERE hospital_id = :2 AND serialid = :3", datestr, hospid, serialid); err != nil {
				panic(err)
			}
			if err := tx.Commit(); err != nil {
				panic(err)
			}
		}

	}
	m.Store(id, 50) // progress 50%
	// 強引にRを動かし解析
	/*	var rscript string
		if rscript = os.Getenv("RSCRIPT"); rscript == "" {
			rscript = "/usr/local/bin/Rscript"
		}
		out, err := exec.Command(rscript, "analysis.R").CombinedOutput()
		if err != nil {
			output += err.Error()
		} else {
			output += string(out)
		}
	*/
	return output
}

// AdminRouterPost  POST "/admin" を処理
func AdminRouterPost(c echo.Context) error {
	db := Repository()
	defer db.Close()

	hospid, username := GetHospAndUser(c)
	if hospid != 1 {
		// invalid access, go to logout
		return c.Redirect(http.StatusFound, "/logout")
	}

	htmlvariable := adminHTMLtemplate{
		Title:        "adminコンソール",
		HospitalName: username,
		ImageTag:     "",
		Message:      "",
		Redirect:     "",
		CSS:          "/css/index.css",
	}

	return c.Render(http.StatusOK, "admin", htmlvariable)
}
