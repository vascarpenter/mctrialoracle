package routes

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type eventeditHTMLtemplate struct {
	Title        string
	HospitalName string
	Event        Events
	CSS          string
}

// EventEditRouter  GET "/eventedit/:hosp/:ser/:ev" を処理
func EventEditRouter(c echo.Context) error {
	ctx := context.Background()
	db := Repository()
	defer db.Close()

	hosp := c.Param("hosp")
	ser := c.Param("ser")
	ev := c.Param("ev")

	hospid := c.Get("UserID").(int)
	hospidfromparam, err := strconv.Atoi(hosp)
	if err != nil || hospid == 0 || hospidfromparam != int(hospid) {
		// invalid access, go to logout
		return c.Redirect(http.StatusFound, "/logout")
	}

	var initial string
	var patientID string
	if err := db.QueryRowContext(ctx,
		`SELECT PATIENT_ID,"INITIAL"
			FROM PATIENTS WHERE hospital_id = :1 AND serialid = :2`, hosp, ser).Scan(
		&patientID, &initial); err != nil {
		panic(err)
	}
	hospname := c.Get("UserName").(string)
	editText := hospname + " ID:" + patientID + " " + initial + "さんのイベントの編集"

	var event Events
	event.HospitalID = int(hospid)
	err = db.QueryRowContext(ctx, `SELECT hospital_id,serialid,eventid,
				TO_CHAR("DATE",'YYYY-MM-DD'),
				alive,dropout,macce,bh,bw,sbp,dbp,hr,
				event
				FROM events WHERE hospital_id=:1 AND serialid=:2 AND eventid=:3`, hosp, ser, ev).Scan(
		&event.HospitalID, &event.SerialID, &event.EventID,
		&event.Date,
		&event.Alive, &event.Dropout, &event.Macce,
		&event.Bh, &event.Bw, &event.Sbp, &event.Dbp, &event.Hr,
		&event.Event)

	if err == sql.ErrNoRows {
		// invalid access, go to logout
		return c.Redirect(http.StatusFound, "/logout")
	}
	if err != nil {
		panic(err)
	}

	htmlvariable := eventeditHTMLtemplate{
		Title:        editText,
		HospitalName: hospname,
		Event:        event,
		CSS:          "/css/event.css",
	}
	return c.Render(http.StatusOK, "eventedit", htmlvariable)
}

// EventEditRouterPost  POST "/eventedit/:hosp/:ser/:ev" を処理
func EventEditRouterPost(c echo.Context) error {
	db := Repository()
	defer db.Close()
	ctx := context.Background()

	hosp := c.Param("hosp")
	ser := c.Param("ser")
	ev := c.Param("ev")

	hospid := c.Get("UserID").(int)
	hospidfromparam, err := strconv.Atoi(hosp)
	if err != nil || hospid == 0 || hospidfromparam != int(hospid) {
		// invalid access, go to logout
		return c.Redirect(http.StatusFound, "/logout")
	}

	bodyheight, _ := strconv.Atoi(c.FormValue("bh"))
	bodyweight, _ := strconv.Atoi(c.FormValue("bw"))
	sbp, _ := strconv.Atoi(c.FormValue("sbp"))
	dbp, _ := strconv.Atoi(c.FormValue("dbp"))
	hr, _ := strconv.Atoi(c.FormValue("hr"))

	tx, err := db.BeginTx(ctx, nil)
	if _, err = tx.Exec(`UPDATE events SET "DATE"=TO_DATE(:1,'YYYY/MM/DD'),
		alive=:2,dropout=:3,macce=:4,
		bh=:5,bw=:6,sbp=:7,dbp=:8,hr=:9,event=:10
	  	WHERE hospital_id=:11 AND serialid=:12 AND eventid=:13`,
		c.FormValue("date"),           // 1: DATE
		c.FormValue("alive") != "0",   // 2: alive
		c.FormValue("dropout") != "0", // 3: dropout
		c.FormValue("macce") != "0",   // 4: macce
		bodyheight,                    // 5: bh
		bodyweight,                    // 6: bw
		sbp,                           // 7: sbp
		dbp,                           // 8: dbp
		hr,                            // 9: hr
		c.FormValue("eventtext"),      // 10: event

		hosp, ser, ev); err != nil {
		panic(err)
	}
	if err := tx.Commit(); err != nil {
		panic(err)
	}

	//	fmt.Printf("%+v\n", event)

	return c.Redirect(http.StatusFound, "/eventlist/"+hosp+"/"+ser)
}
