package routes

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type patientHTMLtemplate struct {
	Title        string
	HospitalName string
	AllowDate    string
	StartDate    string
	CSS          string
}

// PatientRouter  GET "/patient" を処理
func PatientRouter(c echo.Context) error {

	curDate := time.Now().Format("2006-01-02")
	hospname := c.Get("UserName").(string)
	htmlvariable := patientHTMLtemplate{
		Title:        "新規患者の登録",
		HospitalName: hospname,
		AllowDate:    curDate,
		StartDate:    curDate,
		CSS:          "/css/register.css",
	}
	return c.Render(http.StatusOK, "patient", htmlvariable)

}

// PatientRouterPost  POST "/patient" を処理
func PatientRouterPost(c echo.Context) error {
	ctx := context.Background()
	db := Repository()
	defer db.Close()

	loc, _ := time.LoadLocation("Asia/Tokyo")
	birth, _ := time.ParseInLocation("2006-01-02", c.FormValue("birth"), loc)
	years, _, _, _, _, _ := DiffDate(time.Now(), birth)
	//fmt.Printf("%+v,%+v,%+v,%+v\n", years, c.FormValue("birth"), c.FormValue("allowdate"), c.FormValue("startdate"))
	hospid := c.Get("UserID").(int)

	female := 0
	if c.FormValue("sex") != "0" {
		female = 1
	}

	count := 0
	err := db.QueryRowContext(ctx, "SELECT COUNT(*) FROM patients WHERE hospital_id = :1", hospid).Scan(&count)
	if err == sql.ErrNoRows {
		count = 0
	} else if err != nil {
		panic(err)
	}

	tx, _ := db.BeginTx(ctx, nil)
	if _, err := tx.Exec(`INSERT INTO patients (ID,PATIENT_ID,HOSPITAL_ID,SERIALID,TRIALGROUP,"INITIAL",BIRTHDATE,FEMALE,AGE,ALLOWDATE,STARTDATE)
		VALUES (:1,:2,:3,:4,:5,:6,TO_DATE(:7,'YYYY-MM-DD'),:8,:9,
		TO_DATE(:10,'YYYY-MM-DD'),TO_DATE(:11,'YYYY-MM-DD') )`,
		count+1,                   // 1: ID
		c.FormValue("patientid"),  // 2: PATIENT_ID
		hospid,                    // 3: HOSPITAL_ID
		count+1,                   // 4: SERIALID
		c.FormValue("trialgroup"), // 5: TRIALGROUP
		c.FormValue("initial"),    // 6: INITIAL
		c.FormValue("birth"),      // 7: BIRTHDATE
		female,                    // 8: FEMALE
		years,                     // 9: AGE
		c.FormValue("allowdate"),  // 10: ALLOWDATE
		c.FormValue("startdate"),  // 11: STARTDATE
	); err != nil {
		panic(err)
	}
	if err := tx.Commit(); err != nil {
		panic(err)
	}
	return c.Redirect(http.StatusFound, "/")
}
