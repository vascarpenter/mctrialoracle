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
	if _, err := tx.Exec(`INSERT INTO patients (PATIENT_ID,HOSPITAL_ID,SERIALID,TRIALGROUP,"INITIAL",BIRTHDATE,FEMALE,AGE,ALLOWDATE,STARTDATE)
		VALUES (:1,:2,:3,:4,:5,TO_DATE(:6,'YYYY-MM-DD'),:7,:8,
		TO_DATE(:9,'YYYY-MM-DD'),TO_DATE(:10,'YYYY-MM-DD') )`,
		c.FormValue("patientid"),  // 1: PATIENT_ID
		hospid,                    // 2: HOSPITAL_ID
		count+1,                   // 3: SERIALID
		c.FormValue("trialgroup"), // 4: TRIALGROUP
		c.FormValue("initial"),    // 5: INITIAL
		c.FormValue("birth"),      // 6: BIRTHDATE
		female,                    // 7: FEMALE
		years,                     // 8: AGE
		c.FormValue("allowdate"),  // 9: ALLOWDATE
		c.FormValue("startdate"),  // 10: STARTDATE
	); err != nil {
		panic(err)
	}
	if err := tx.Commit(); err != nil {
		panic(err)
	}
	return c.Redirect(http.StatusFound, "/")
}
