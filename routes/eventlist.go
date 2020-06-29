package routes

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type eventlistHTMLtemplate struct {
	Title        string
	HospitalName string
	Person       Patients
	Events       []Events
	CSS          string
}

// EventlistRouter  GET "/eventlist/:hosp/:ser" を処理
func EventlistRouter(c echo.Context) error {

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

	var p Patients
	if err := db.QueryRowContext(ctx,
		`SELECT PATIENT_ID,HOSPITAL_ID,SERIALID,TRIALGROUP,"INITIAL",
		TO_CHAR(BIRTHDATE,'YYYY/MM/DD'),FEMALE,AGE, 
		TO_CHAR(ALLOWDATE,'YYYY/MM/DD'), 
		TO_CHAR(STARTDATE,'YYYY/MM/DD'), 
		TO_CHAR(DROPDATE,'YYYY/MM/DD'), 
		TO_CHAR(MACCEDATE,'YYYY/MM/DD'), 
		TO_CHAR(DEADDATE,'YYYY/MM/DD'), 
		TO_CHAR(FINISHDATE,'YYYY/MM/DD'), 
		DIFFDAYS FROM PATIENTS WHERE hospital_id = :1 AND serialid = :2`, hosp, ser).Scan(
		&p.PatientID, &p.HospitalID, &p.SerialID,
		&p.Trialgroup, &p.Initial,
		&p.Birthdate, &p.Female, &p.Age,
		&p.Allowdate, &p.Startdate, &p.Allowdate, &p.Maccedate,
		&p.Deaddate, &p.Finishdate, &p.Diffdays); err != nil {
		panic(err)
	}
	// get first line and quit
	//fmt.Printf("%+v\n", p)

	rows2, _ := db.QueryContext(ctx,
		`SELECT HOSPITAL_ID,SERIALID,EVENTID,
			TO_CHAR("DATE",'YYYY/MM/DD'),
			ALIVE, DROPOUT, MACCE, BH, BW, SBP, DBP, HR, EVENT,
			DIFFDAYS FROM EVENTS WHERE HOSPITAL_ID=:1 AND SERIALID=:2`, hosp, ser)

	defer rows2.Close()

	count := 0
	var s []Events
	for rows2.Next() {
		var e Events
		if err := rows2.Scan(&e.HospitalID, &e.SerialID, &e.EventID,
			&e.Date,
			&e.Alive, &e.Dropout, &e.Macce, &e.Bh, &e.Bw,
			&e.Sbp, &e.Dbp, &e.Hr, &e.Event,
			&e.Diffdays); err != nil {
			panic(err)
		}
		s = append(s, e)
		count++
	}

	// 差分を計算し代入(temp)
	if p.Startdate.Valid {
		loc, _ := time.LoadLocation("Asia/Tokyo")
		stdate, _ := time.ParseInLocation("2006/01/02", p.Startdate.String, loc)
		diff := time.Now().Sub(stdate).Hours() / 24
		p.Diffdays = sql.NullInt32{Int32: int32(diff), Valid: true}
		for ind, e := range s {
			if e.Date.Valid {
				evdate, _ := time.ParseInLocation("2006/01/02", e.Date.String, loc)
				diff := evdate.Sub(stdate).Hours() / 24
				s[ind].Diffdays = sql.NullInt32{Int32: int32(diff), Valid: true} // eを編集しても反映されない
				//fmt.Printf("%+v\n", s[ind])
			}
		}
	}

	username := c.Get("UserName").(string)

	htmlvariable := eventlistHTMLtemplate{
		Title:        "イベント一覧",
		HospitalName: username,
		Person:       p,
		Events:       s,
		CSS:          "/css/index.css",
	}
	return c.Render(http.StatusOK, "eventlist", htmlvariable)
}
