package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strconv"

	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/julienschmidt/httprouter"

	"github.com/Rob-a21/Cassiopeia/models"
)

type ProfileHandler struct {
	tmpl           *template.Template
	profileService models.ProfileService
}

func NewProfileHandler(T *template.Template, NS models.ProfileService) *ProfileHandler {
	return &ProfileHandler{tmpl: T, profileService: NS}
}

func (prf *ProfileHandler) StudentsProfile(w http.ResponseWriter, r *http.Request) {

	students, err := prf.profileService.Students()

	if err != nil {

		return
	}
	_ = prf.tmpl.ExecuteTemplate(w, "student.profile.layout", students)

}

func (prf *ProfileHandler) StudentProfile(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		fmt.Println("id", idRaw)

		id, err := strconv.Atoi(idRaw)

		if err != nil {

			fmt.Println("conversion error!")

			return
		}

		student, err := prf.profileService.Student(id)

		if err != nil {

			fmt.Println("not conversion error!")

			return

		}
		fmt.Println("successful!")

		prf.tmpl.ExecuteTemplate(w, "student.profile.layout", student)

	}

}

func (prf *ProfileHandler) TeachersProfile(w http.ResponseWriter, r *http.Request) {

	teacher, err := prf.profileService.Teachers()
	if err != nil {
		panic(err)
	}
	prf.tmpl.ExecuteTemplate(w, "teacher.profile.layout", teacher)

}

func (prf *ProfileHandler) TeacherProfile(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		fmt.Println("id: ", idRaw)

		id, err := strconv.Atoi(idRaw)

		if err != nil {

			fmt.Println("conversion Error!")

			return

		}

		fmt.Println("successful!")

		teacher, err := prf.profileService.Teacher(id)

		if err != nil {

			return
		}

		prf.tmpl.ExecuteTemplate(w, "teacher.profile.layout", teacher)

	}

}
func (prf *ProfileHandler) AdminProfile(w http.ResponseWriter, r *http.Request) {

	admin, err := prf.profileService.Admins()
	if err != nil {
		panic(err)
	}
	prf.tmpl.ExecuteTemplate(w, "admin.profile.layout", admin)

}

func (prf *ProfileHandler) FamiliesProfile(w http.ResponseWriter, r *http.Request) {

	family, err := prf.profileService.Families()
	if err != nil {
		panic(err)
	}
	prf.tmpl.ExecuteTemplate(w, "family.profile.layout", family)

}
func (prf *ProfileHandler) FamilyProfile(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {

			panic(err)
		}

		student, err := prf.profileService.Family(id)

		if err != nil {
			panic(err)
		}

		prf.tmpl.ExecuteTemplate(w, "family.profile.layout", student)

	}
}
func (prf *ProfileHandler) EmailExists(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		email := r.URL.Query().Get("id")

		prf.profileService.EmailExists(email)

	}

	http.Redirect(w, r, "/admin/student", http.StatusSeeOther)
}

func (prf *ProfileHandler) AdminGetStudent(w http.ResponseWriter, r *http.Request) {

	students, err := prf.profileService.Students()
	if err != nil {
		panic(err)
	}

	_ = prf.tmpl.ExecuteTemplate(w, "admin.view.student.layout", students)

}

func (prf *ProfileHandler) AdminDeleteStudent(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		err = prf.profileService.DeleteStudent(id)

		if err != nil {
			panic(err)
		}

	}

	http.Redirect(w, r, "/admin/student", http.StatusSeeOther)
}

func (prf *ProfileHandler) AdminGetTeacher(w http.ResponseWriter, r *http.Request) {

	teachers, err := prf.profileService.Teachers()
	if err != nil {
		panic(err)
	}

	prf.tmpl.ExecuteTemplate(w, "admin.view.teacher.layout", teachers)

}

func (prf *ProfileHandler) AdminDeleteTeacher(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		prf.profileService.DeleteTeacher(id)

	}

	http.Redirect(w, r, "/admin/teacher", http.StatusSeeOther)
}

func (prf *ProfileHandler) ApiAdminProfile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	id, err := strconv.Atoi(path.Base(r.URL.Path))

	if err != nil {

		return
	}

	admin := entity.Admin{}

	prf.profileService.Admin(id)

	output, err := json.MarshalIndent(&admin, "", "\t\t")

	if err != nil {

		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(output)

	return
}

func (prf *ProfileHandler) ApiStudentProfile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	id, err := strconv.Atoi(path.Base(r.URL.Path))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	student, errs := prf.profileService.Student(id)

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(student, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

	// id, err := strconv.Atoi(path.Base(r.URL.Path))

	// if err != nil {

	// 	return
	// }

	// student := entity.Student{}

	// prf.profileService.Student(id)

	// output, err := json.MarshalIndent(&student, "", "\t\t")

	// if err != nil {

	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")

	// w.Write(output)

	// return
}

func (prf *ProfileHandler) ApiTeacherProfile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	id, err := strconv.Atoi(path.Base(r.URL.Path))

	if err != nil {

		return
	}

	teacher := entity.Teacher{}

	prf.profileService.Teacher(id)

	output, err := json.MarshalIndent(&teacher, "", "\t\t")

	if err != nil {

		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(output)

	return
}

func (prf *ProfileHandler) ApiFamilyProfile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	id, err := strconv.Atoi(path.Base(r.URL.Path))

	if err != nil {

		return
	}

	family := entity.Family{}

	prf.profileService.Family(id)

	output, err := json.MarshalIndent(&family, "", "\t\t")

	if err != nil {

		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(output)

	return
}

func (crs *ProfileHandler) ApiStudentsProfile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	Students, errs := crs.profileService.Students()
	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(Students, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

func (prf *ProfileHandler) NewYearRegistration(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		if err != nil {

			return
		}

		student, err := prf.profileService.Student(id)

		if err != nil {

			return
		}

		prf.tmpl.ExecuteTemplate(w, "student.year.registration.layout", student)

	} else if r.Method == http.MethodPost {

		stdt := entity.Student{}

		stdt.Grade, _ = strconv.Atoi(r.FormValue("grade"))

		prf.profileService.NewYearRegistration(stdt)

		prf.tmpl.ExecuteTemplate(w, "student.year.registration.layout", stdt)

		http.Redirect(w, r, "/student", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/student", http.StatusSeeOther)
	}
	return
}
