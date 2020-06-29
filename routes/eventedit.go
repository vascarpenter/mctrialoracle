package routes

import (
	"net/http"

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
	/*
		ctx := context.Background()
		db := Repository()
		defer db.Close()

		hosp := c.Param("hosp")
		ser := c.Param("ser")
		ev := c.Param("ev")

		hospid := c.Get("UserID").(uint)
		hospidfromparam, err := strconv.Atoi(hosp)
		if err != nil || hospid == 0 || hospidfromparam != int(hospid) {
			// invalid access, go to logout
			return c.Redirect(http.StatusFound, "/logout")
		}
			patient, err := models.Patients(qm.Where("hospital_id=? AND serialid=?", hosp, ser)).One(ctx, db)
			if err != nil {
				// invalid access, go to logout
				return c.Redirect(http.StatusFound, "/logout")
			}
			hospname := c.Get("UserName").(string)
			editText := hospname + " ID:" + patient.PatientID.String + " " + patient.Initial.String + "さんのイベントの編集"

			event, err := models.Events(qm.Where("hospital_id=? AND serialid=? AND eventid=?", hosp, ser, ev)).One(ctx, db)
			if err != nil {
				// invalid access, go to logout
				return c.Redirect(http.StatusFound, "/logout")
			}
		htmlvariable := eventeditHTMLtemplate{
			Title:        editText,
			HospitalName: hospname,
			Event:        *event,
			CSS:          "/css/event.css",
		}
		return c.Render(http.StatusOK, "eventedit", htmlvariable)
	*/
	return c.Redirect(http.StatusFound, "/logout")
}

// EventEditRouterPost  POST "/eventedit/:hosp/:ser/:ev" を処理
func EventEditRouterPost(c echo.Context) error {
	/*
			db := Repository()
			defer db.Close()
			ctx := context.Background()

			hosp := c.Param("hosp")
			ser := c.Param("ser")
			ev := c.Param("ev")

			hospid := c.Get("UserID").(uint)
			hospidfromparam, err := strconv.Atoi(hosp)
			if err != nil || hospid == 0 || hospidfromparam != int(hospid) {
				// invalid access, go to logout
				return c.Redirect(http.StatusFound, "/logout")
			}
				event, err := models.Events(qm.Where("hospital_id=? AND serialid=? AND eventid=?", hosp, ser, ev)).One(ctx, db)
				if err != nil {
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

				event.Alive = null.BoolFrom(c.FormValue("alive") != "0")
				event.Dropout = null.BoolFrom(c.FormValue("dropout") != "0")
				event.Macce = null.BoolFrom(c.FormValue("macce") != "0")
				event.Date = null.TimeFrom(eventdate)
				event.BH = null.IntFrom(bodyheight)
				event.BW = null.IntFrom(bodyweight)
				event.SBP = null.IntFrom(sbp)
				event.DBP = null.IntFrom(dbp)
				event.HR = null.IntFrom(hr)
				event.Event = null.StringFrom(c.FormValue("eventtext"))

				//	fmt.Printf("%+v\n", event)

				event.Update(ctx, db, boil.Infer())
		return c.Redirect(http.StatusFound, "/eventlist/"+hosp+"/"+ser)
	*/
	return c.Redirect(http.StatusFound, "/logout")
}
