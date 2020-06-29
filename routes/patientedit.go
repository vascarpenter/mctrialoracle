package routes

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type patienteditHTMLtemplate struct {
	Title        string
	HospitalName string
	Person       Patients
	CSS          string
}

// PatientEditRouter  GET "/patientedit/:hosp/:ser" を処理
func PatientEditRouter(c echo.Context) error {
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

	hospname := c.Get("UserName").(string)

	var patient Patients
	err = db.QueryRowContext(ctx, `SELECT
			patient_id,hospital_id,serialid,trialgroup,"INITIAL",TO_CHAR(birthdate,'YYYY-MM-DD'),female,age,
			TO_CHAR(allowdate,'YYYY-MM-DD'),
			TO_CHAR(startdate,'YYYY-MM-DD'),
			TO_CHAR(dropdate,'YYYY-MM-DD'),
			TO_CHAR(maccedate,'YYYY-MM-DD'),
			TO_CHAR(deaddate,'YYYY-MM-DD'),
			TO_CHAR(finishdate,'YYYY-MM-DD')
			FROM patients WHERE hospital_id=:1 AND serialid=:2`, hosp, ser).Scan(&patient.PatientID,
		&patient.HospitalID, &patient.SerialID,
		&patient.Trialgroup,
		&patient.Initial, &patient.Birthdate, &patient.Female, &patient.Age,
		&patient.Allowdate,
		&patient.Startdate,
		&patient.Dropdate,
		&patient.Maccedate,
		&patient.Deaddate,
		&patient.Finishdate)
	if err != nil || err == sql.ErrNoRows {
		panic(err)
	}

	//fmt.Printf("%+v\n", patient)
	editText := hospname + " " + string(patient.Initial.String) + "さんの基本情報の編集"

	htmlvariable := patienteditHTMLtemplate{
		Title:        editText,
		HospitalName: hospname,
		Person:       patient,
		CSS:          "/css/register.css",
	}
	return c.Render(http.StatusOK, "patientedit", htmlvariable)

}

// PatientEditRouterPost  POST "/patientedit/:hosp/:ser" を処理
func PatientEditRouterPost(c echo.Context) error {
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
	loc, _ := time.LoadLocation("Asia/Tokyo")
	birth, _ := time.ParseInLocation("2006-01-02", c.FormValue("birth"), loc)
	years, _, _, _, _, _ := DiffDate(time.Now(), birth)
	female := 0
	if c.FormValue("sex") != "0" {
		female = 1
	}
	trialgroup, err := strconv.Atoi(c.FormValue("trialgroup"))

	tx, _ := db.BeginTx(ctx, nil)
	if _, err := tx.Exec(`UPDATE patients
	SET PATIENT_ID=:1,TRIALGROUP=:2,"INITIAL"=:3,
		BIRTHDATE=TO_DATE(:4,'YYYY-MM-DD'),FEMALE=:5,AGE=:6,
		ALLOWDATE=TO_DATE(:7,'YYYY-MM-DD'),
		STARTDATE=TO_DATE(:8,'YYYY-MM-DD')
		WHERE hospital_id=:9 AND serialid=:10`,
		c.FormValue("patientid"), // 1: PATIENT_ID
		trialgroup,               // 2: TRIALGROUP
		c.FormValue("initial"),   // 3: INITIAL
		c.FormValue("birth"),     // 4: BIRTHDATE
		female,                   // 5: FEMALE
		years,                    // 6: AGE
		c.FormValue("allowdate"), // 7: ALLOWDATE
		c.FormValue("startdate"), // 8: STARTDATE

		hosp, ser); err != nil {
		panic(err)
	}
	if err := tx.Commit(); err != nil {
		panic(err)
	}
	return c.Redirect(http.StatusFound, "/")
}
