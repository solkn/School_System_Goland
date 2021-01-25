package handler

import (
	"encoding/json"
	"html/template"
	"net/http"
	"path"
	"strconv"

	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/Rob-a21/Cassiopeia/models"
	"github.com/julienschmidt/httprouter"
)

type CourseHandler struct {
	tmpl       *template.Template
	crsService models.CourseService
}

func NewCourseHandler(T *template.Template, CS models.CourseService) *CourseHandler {
	return &CourseHandler{tmpl: T, crsService: CS}
}

func (crs *CourseHandler) AdminAddCourse(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		course := entity.Course{}
		course.CourseName = r.FormValue("coursename")
		course.CourseID, _ = strconv.Atoi(r.FormValue("courseid"))
		course.Grade, _ = strconv.Atoi(r.FormValue("grade"))

		_ = crs.crsService.AddCourse(course)

	}

	_ = crs.tmpl.ExecuteTemplate(w, "admin.course.new.layout", nil)

}

func (crs *CourseHandler) AdminGetCourse(w http.ResponseWriter, r *http.Request) {

	courses, err := crs.crsService.GetCourse()
	if err != nil {
		panic(err)
	}

	crs.tmpl.ExecuteTemplate(w, "admin.course.layout", courses)

}

func (crs *CourseHandler) AdminUpdateCourse(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("courseid")
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		courses, err := crs.crsService.Course(id)

		if err != nil {
			panic(err)
		}

		_ = crs.tmpl.ExecuteTemplate(w, "admin.course.update.layout", courses)

	} else if r.Method == http.MethodPost {

		course := entity.Course{}
		course.CourseID, _ = strconv.Atoi(r.FormValue("courseid"))
		course.CourseName = r.FormValue("coursename")
		course.Grade, _ = strconv.Atoi(r.FormValue("grade"))

		err := crs.crsService.UpdateCourse(course)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/admin/course", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/admin/course", http.StatusSeeOther)
	}

}

func (crs *CourseHandler) AdminDeleteCourse(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("courseid")

		id, err := strconv.ParseInt(idRaw, 0, 0)

		if err != nil {
			panic(err)
		}

		_ = crs.crsService.DeleteCourse(int(id))

	}

	http.Redirect(w, r, "/admin/course", http.StatusSeeOther)
}

func (crs *CourseHandler) StudentCourse(w http.ResponseWriter, r *http.Request) {

	courses, err := crs.crsService.GetCourse()
	if err != nil {
		panic(err)
	}

	_ = crs.tmpl.ExecuteTemplate(w, "student.course.layout", courses)

}

func (crs *CourseHandler) FamilyGetCourse(w http.ResponseWriter, r *http.Request) {

	courses, err := crs.crsService.GetCourse()
	if err != nil {
		panic(err)
	}

	_ = crs.tmpl.ExecuteTemplate(w, "family.course.layout", courses)

}

//
//func (crs *CourseHandler)ApiAdminPostCourse(w http.ResponseWriter,r *http.Request){
//
//	if r.Method == http.MethodPost {
//		len := r.ContentLength
//
//		body := make([]byte, len)
//
//		r.Body.Read(body)
//
//		course := entity.Course{}
//
//		json.Unmarshal(body, &course)
//
//		crs.crsService.AddCourse(course)
//
//		w.WriteHeader(200)
//	}
//	   return
//}
//
//
//func (crs *CourseHandler)ApiAdminGetCourses(w http.ResponseWriter,r *http.Request) {
//
//if r.Method == http.MethodGet {
//	course := entity.Course{}
//
//	crs.crsService.GetCourse()
//
//	output, err := json.MarshalIndent(&course, "", "\t\t")
//
//	if err != nil {
//
//		return
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//
//	w.Write(output)
//}
//	return
//}

//
//func (crs *CourseHandler)ApiStudentGetCourse(w http.ResponseWriter,r *http.Request) {
//
//	if r.Method == http.MethodGet {
//		id, err := strconv.Atoi(path.Base(r.URL.Path))
//
//		if err != nil {
//
//			return
//		}
//
//		course := entity.Course{}
//
//		crs.crsService.Course(id)
//
//		output, err := json.MarshalIndent(&course, "", "\t\t")
//
//		if err != nil {
//
//			return
//		}
//
//		w.Header().Set("Content-Type", "application/json")
//
//		w.Write(output)
//	}
//	return
//}

//func (crs *CourseHandler)ApiStudentGetCourses(w http.ResponseWriter,r *http.Request) {
//
//
//if r.Method == http.MethodGet{
//
//	course := entity.Course{}
//
//	crs.crsService.GetCourse()
//
//	output,err := json.MarshalIndent(&course,"","\t\t")
//
//	if err != nil{
//
//		return
//	}
//
//	w.Header().Set("Content-Type","application/json")
//
//	w.Write(output)
//
//
//}
//	return
//}

func (crs *CourseHandler) ApiAdminPostCourse(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if r.Method == http.MethodPost {
		len := r.ContentLength

		body := make([]byte, len)

		_, _ = r.Body.Read(body)

		course := entity.Course{}

		_ = json.Unmarshal(body, &course)

		_ = crs.crsService.AddCourse(course)

		w.WriteHeader(200)
	}
	return
}

func (crs *CourseHandler) ApiAdminGetCourses(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	courses, errs := crs.crsService.GetCourse()

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(courses, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)

	_, _ = w.Write(output)

	return
}

func (crs *CourseHandler) ApiAdminGetCourse(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	course, errs := crs.crsService.Course(id)

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(course, "", "\t\t")

	_ = crs.crsService.DeleteCourse(id)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func (crs *CourseHandler) ApiAdminDeleteCourse(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	course := crs.crsService.DeleteCourse(id)

	//if errs != nil {
	//	w.Header().Set("Content-Type", "application/json")
	//	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	//	return
	//}

	output, err := json.MarshalIndent(course, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func (crs *CourseHandler) ApiStudentGetCourses(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	courses, errs := crs.crsService.GetCourse()

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(courses, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

func (crs *CourseHandler) ApiStudentGetCourse(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	id, err := strconv.Atoi(path.Base(r.URL.Path))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	course, errs := crs.crsService.Course(id)

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(course, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
