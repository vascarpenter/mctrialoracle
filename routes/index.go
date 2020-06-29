package routes

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type indexHTMLtemplate struct {
	Title        string
	HospitalName string
	Patients     []Patients
	CSS          string
}

// IndexRouter  GET "/" を処理
func IndexRouter(c echo.Context) error {

	db := Repository()
	defer db.Close()
	ctx := context.Background()

	userid := c.Get("UserID") // useid is not nil, because if nil, middleware detects it
	useridint := userid.(int)
	if useridint == 0 {
		return c.Redirect(http.StatusFound, "/login")
	}
	if useridint == 1 {
		return c.Redirect(http.StatusFound, "/admin")
	}
	username := c.Get("UserName").(string)

	// SQL: SELECT * FROM patients WHERE hospital_id = ?
	rows2, _ := db.QueryContext(ctx,
		`SELECT PATIENT_ID,HOSPITAL_ID,SERIALID,TRIALGROUP,"INITIAL",
		TO_CHAR(BIRTHDATE,'YYYY/MM/DD'),FEMALE,AGE, 
		TO_CHAR(ALLOWDATE,'YYYY/MM/DD'), 
		TO_CHAR(STARTDATE,'YYYY/MM/DD'), 
		TO_CHAR(DROPDATE,'YYYY/MM/DD'), 
		TO_CHAR(MACCEDATE,'YYYY/MM/DD'), 
		TO_CHAR(DEADDATE,'YYYY/MM/DD'), 
		TO_CHAR(FINISHDATE,'YYYY/MM/DD'), 
		DIFFDAYS FROM PATIENTS WHERE HOSPITAL_ID=:1`, useridint)

	defer rows2.Close()

	count := 0
	var s []Patients
	for rows2.Next() {
		var p Patients
		if err := rows2.Scan(&p.PatientID, &p.HospitalID, &p.SerialID,
			&p.Trialgroup, &p.Initial,
			&p.Birthdate, &p.Female, &p.Age,
			&p.Allowdate, &p.Startdate, &p.Allowdate, &p.Maccedate,
			&p.Deaddate, &p.Finishdate, &p.Diffdays); err != nil {
			panic(err)
		}
		s = append(s, p)
		count++
	}
	// fmt.Printf("%+v\n", s)

	// fmt.Printf("%+v\n", s[0])

	// 差分を計算
	for i := 0; i < count; i++ {
		if s[i].Startdate.Valid {
			loc, _ := time.LoadLocation("Asia/Tokyo")
			t, _ := time.ParseInLocation("2006/01/02", s[i].Startdate.String, loc)

			diff := time.Now().Sub(t).Hours() / 24
			s[i].Diffdays = sql.NullInt32{int32(diff), true}
		}
	}

	htmlvariable := indexHTMLtemplate{
		Title:        "患者一覧",
		HospitalName: username,
		Patients:     s,
		CSS:          "/css/index.css",
	}
	return c.Render(http.StatusOK, "index", htmlvariable)
}
