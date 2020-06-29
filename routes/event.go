package routes

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type eventHTMLtemplate struct {
	Title        string
	HospitalName string
	Hospid       string
	Serid        string
	CSS          string
	TodayDate    string
}

// EventRouter  GET "/event/:hosp/:ser" を処理
func EventRouter(c echo.Context) error {
	ctx := context.Background()
	db := Repository()
	defer db.Close()

	hosp := c.Param("hosp")
	ser := c.Param("ser")

	hospid := c.Get("UserID").(uint)
	hospidfromparam, err := strconv.Atoi(hosp)
	if err != nil || hospid == 0 || hospidfromparam != int(hospid) {
		// invalid access, go to logout
		return c.Redirect(http.StatusFound, "/logout")
	}

	stmt, err := db.PrepareContext(ctx, "SELECT initial FROM patients WHERE hospital_id = ? AND serialid=?")
	defer stmt.Close()
	var editing string
	err = stmt.QueryRowContext(ctx, hosp, ser).Scan(&editing)
	switch {
	case err == sql.ErrNoRows:
		// cannot get; pass through
	case err != nil:
		panic(err)
	default:
	}

	editText := editing + "さんのイベントの追加"
	hospname := c.Get("UserName").(string)
	curDate := time.Now().Format("2006-01-02")

	htmlvariable := eventHTMLtemplate{
		Title:        editText,
		HospitalName: hospname,
		Hospid:       hosp,
		Serid:        ser,
		CSS:          "/css/event.css",
		TodayDate:    curDate,
	}
	return c.Render(http.StatusOK, "event", htmlvariable)
}

// EventRouterPost  POST "/event/:hosp/:ser" を処理
func EventRouterPost(c echo.Context) error {
	db := Repository()
	defer db.Close()
	ctx := context.Background()

	hosp := c.Param("hosp")
	ser := c.Param("ser")
	serInt, _ := strconv.Atoi(ser)

	hospid := c.Get("UserID").(uint)
	hospidfromparam, err := strconv.Atoi(hosp)
	if err != nil || hospid == 0 || hospidfromparam != int(hospid) {
		// invalid access, go to logout
		return c.Redirect(http.StatusFound, "/logout")
	}

	loc, _ := time.LoadLocation("Asia/Tokyo")
	eventdate, err := time.ParseInLocation("2006/01/02", string(c.FormValue("date")), loc)
	if err != nil {
		panic(err)
	}
	bodyheight, _ := strconv.Atoi(c.FormValue("bh"))
	bodyweight, _ := strconv.Atoi(c.FormValue("bw"))
	sbp, _ := strconv.Atoi(c.FormValue("sbp"))
	dbp, _ := strconv.Atoi(c.FormValue("dbp"))
	hr, _ := strconv.Atoi(c.FormValue("hr"))

	// query max count of events
	var count int
	if err := db.QueryRowContext(ctx, "SELECT COUNT(*) FROM events WHERE hospital_id=:1 AND serialid=:2", hosp, ser).Scan(&count); err != nil {
		panic(err)
	}

	tx, err := db.BeginTx(ctx, nil)
	_, err = tx.Exec(`INSERT INTO event (hospital_id,serialid,eventid,"DATE",alive,dropout,macce,bh,bw,sbp,dbp,hr,event)
	  VALUES (:1,:2,:3,TO_DATE(:4, 'YYYY/MM/DD'),:5,:6,:7,:8,:9,:10,:11,:12,:13)`,
		hospid,                           // 1: hospital_id
		serInt,                           // 2: serialid
		count+1,                          // 3: eventid
		eventdate.Format("'2006/01/02'"), // 4: DATE
		c.FormValue("alive") != "0",      // 5: alive
		c.FormValue("dropout") != "0",    // 6: dropout
		c.FormValue("macce") != "0",      // 7: macce
		bodyheight,                       // 8: bh
		bodyweight,                       // 9: bw
		sbp,                              // 10: sbp
		dbp,                              // 11: dbp
		hr,                               // 12: hr
		c.FormValue("eventtext"),         // 13: event
	)
	if err := tx.Commit(); err != nil {
		panic(err)
	}

	//	fmt.Printf("%+v\n", event)

	return c.Redirect(http.StatusFound, "/eventlist/"+hosp+"/"+ser)
}
