package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strconv"
	"time"

	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/Rob-a21/Cassiopeia/models"
	"github.com/julienschmidt/httprouter"
)

type AttendanceHandler struct {
	tmpl              *template.Template
	attendanceService models.StudentAttendanceService
}

func NewAttendanceHandler(T *template.Template, NS models.StudentAttendanceService) *AttendanceHandler {
	return &AttendanceHandler{tmpl: T, attendanceService: NS}
}

func (at *AttendanceHandler) FillStudentAttendance(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		attendance := entity.Attendance{}
		attendance.Date = time.Now()
		attendance.StudentID, _ = strconv.Atoi(r.FormValue("id"))

		_ = at.attendanceService.FillAttendance(attendance)

		http.Redirect(w, r, "/student", http.StatusSeeOther)

	}

	_ = at.tmpl.ExecuteTemplate(w, "student.attendance.new.layout", nil)

}

func (at *AttendanceHandler) CheckStudentAttendance(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		fmt.Println("id:", idRaw)

		if err != nil {

			return
		}

		attendance, err := at.attendanceService.CheckAttendance(id)

		if err != nil {
			return
		}

		at.tmpl.ExecuteTemplate(w, "student.attendance.layout", attendance)

	}

}

func (at *AttendanceHandler) ShowStudentsAttendance(w http.ResponseWriter, r *http.Request) {

	attendances, err := at.attendanceService.ShowAttendance()
	if err != nil {
		panic(err)
	}

	_ = at.tmpl.ExecuteTemplate(w, "student.attendance.layout", attendances)

}

func (at *AttendanceHandler) ApiStudentPostAttendance(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	len := r.ContentLength

	body := make([]byte, len)

	_, _ = r.Body.Read(body)

	attendance := entity.Attendance{}

	json.Unmarshal(body, &attendance)

	at.attendanceService.FillAttendance(attendance)

	w.WriteHeader(200)

	return
}

func (at *AttendanceHandler) ApiStudentCheckAttendance(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	id, err := strconv.Atoi(path.Base(r.URL.Path))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	attendance, errs := at.attendanceService.CheckAttendance(id)

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(attendance, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

//
func (att *AttendanceHandler) ApiStudentShowAttendance(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	attendances, errs := att.attendanceService.ShowAttendance()
	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(attendances, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}
