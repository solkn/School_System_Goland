package handler

import (
	"encoding/json"
	"html/template"
	"net/http"
	"time"

	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/Rob-a21/Cassiopeia/models"
	"github.com/julienschmidt/httprouter"
)

type NotificationHandler struct {
	tmpl                *template.Template
	notificationService models.NotificationService
}

func NewNotificationHandler(T *template.Template, NS models.NotificationService) *NotificationHandler {
	return &NotificationHandler{tmpl: T, notificationService: NS}
}

func (ntf *NotificationHandler) StudentGetNotification(w http.ResponseWriter, r *http.Request) {

	notf, err := ntf.notificationService.GetNotification()
	if err != nil {
		panic(err)
	}
	_ = ntf.tmpl.ExecuteTemplate(w, "student.notification.layout", notf)

}

func (ntf *NotificationHandler) TeacherAddNotification(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		notf := entity.Notification{}

		notf.NotifyName = r.FormValue("name")
		notf.Message = r.FormValue("message")
		notf.NotificationDate = time.Now()

		ntf.notificationService.AddNotification(notf)

		http.Redirect(w, r, "/teacher/register", http.StatusSeeOther)

	}

	ntf.tmpl.ExecuteTemplate(w, "teacher.notification.layout", nil)

}

func (ntf *NotificationHandler) ApiTeacherPostNotification(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	len := r.ContentLength

	body := make([]byte, len)

	_, _ = r.Body.Read(body)

	notification := entity.Notification{}

	json.Unmarshal(body, &notification)

	_ = ntf.notificationService.AddNotification(notification)

	w.WriteHeader(200)

	return
}

func (ntf *NotificationHandler) ApiStudentGetNotification(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	notification, errs := ntf.notificationService.GetNotification()

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(notification, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
