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

	hospid := c.Get("UserID").(int)
	hospidfromparam, err := strconv.Atoi(hosp)
	if err != nil || hospid == 0 || hospidfromparam != int(hospid) {
		// invalid access, go to logout
		return c.Redirect(http.StatusFound, "/logout")
	}

	var editing string
	err = db.QueryRowContext(ctx, `SELECT "INITIAL" FROM patients WHERE hospital_id = :1 AND serialid=:2`, hosp, ser).Scan(&editing)
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

	hospid := c.Get("UserID").(int)
	hospidfromparam, err := strconv.Atoi(hosp)
	if err != nil || hospid == 0 || hospidfromparam != int(hospid) {
		// invalid access, go to logout
		return c.Redirect(http.StatusFound, "/logout")
	}

	loc, _ := time.LoadLocation("Asia/Tokyo")
	eventdate, err := time.ParseInLocation("2006-01-02", string(c.FormValue("date")), loc)
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
	if _, err = tx.Exec(`INSERT INTO events (id,hospital_id,serialid,eventid,"DATE",alive,dropout,macce,bh,bw,sbp,dbp,hr,event)
	  VALUES (:1,:2,:3,:4,TO_DATE(:5, 'YYYY/MM/DD'),:6,:7,:8,:9,:10,:11,:12,:13,:14)`,
		count+1,                        // 1: ID
		hospid,                         // 2: hospital_id
		serInt,                         // 3: serialid
		count+1,                        // 4: eventid
		eventdate.Format("2006/01/02"), // 5: DATE
		c.FormValue("alive") != "0",    // 6: alive
		c.FormValue("dropout") != "0",  // 7: dropout
		c.FormValue("macce") != "0",    // 8: macce
		bodyheight,                     // 9: bh
		bodyweight,                     // 10: bw
		sbp,                            // 11: sbp
		dbp,                            // 12: dbp
		hr,                             // 13: hr
		c.FormValue("eventtext"),       // 14: event
	); err != nil {
		panic(err)
	}
	if err := tx.Commit(); err != nil {
		panic(err)
	}

	//	fmt.Printf("%+v\n", event)

	return c.Redirect(http.StatusFound, "/eventlist/"+hosp+"/"+ser)
}
