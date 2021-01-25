package main

import (
	"database/sql"

	"github.com/Rob-a21/Cassiopeia/models/repository"
	"github.com/Rob-a21/Cassiopeia/models/service"
	"github.com/julienschmidt/httprouter"

	"html/template"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/Rob-a21/Cassiopeia/delivery/http/handler"
)

var tmpl = template.Must(template.ParseGlob("c:/Users/solki/go/src/github.com/Rob-a21/Cassiopeia/delivery/web/templates/*"))

func main() {

	//csrfSignKey := []byte(token.GenerateRandomID(32))

	dbconn, err := sql.Open("postgres",
		"postgres://postgres:aait@localhost/school?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	if err := dbconn.Ping(); err != nil {
		panic(err)
	}


	profileRepository := repository.NewPsqlProfileRepositoryImpl(dbconn)
	profileService := service.NewProfileServiceImpl(profileRepository)
	profileHandler := handler.NewProfileHandler(tmpl, profileService)

	notificationRepository := repository.NewPsqlNotificationRepositoryImpl(dbconn)
	notificationService := service.NewNotificationServiceImpl(notificationRepository)
	notificationHandler := handler.NewNotificationHandler(tmpl, notificationService)

	courseRepository := repository.NewPsqlCourseRepositoryImpl(dbconn)
	courseService := service.NewCourseServiceImpl(courseRepository)
	courseHandler := handler.NewCourseHandler(tmpl, courseService)

	attendanceRepository := repository.NewStudentAttendanceRepositoryImpl(dbconn)
	attendanceService := service.NewStudentAttendanceServiceImpl(attendanceRepository)
	attendanceHandler := handler.NewAttendanceHandler(tmpl, attendanceService)

	assessmentRepository := repository.NewAssessmentRepositoryImpl(dbconn)
	assessmentService := service.NewAssessmentServiceImpl(assessmentRepository)
	assessmentHandler := handler.NewAssessmentHandler(tmpl, assessmentService)

	router := httprouter.New()



	// admin api handlers
	router.GET("/api/admin/profile/:id", profileHandler.ApiAdminProfile)
	router.GET("/api/admin/course", courseHandler.ApiAdminPostCourse)
	router.GET("/api/admin/course/:id", courseHandler.ApiAdminGetCourse)
	router.GET("/api/admin/courses", courseHandler.ApiAdminGetCourses)
	router.DELETE("/api/admin/course/delete/:id", courseHandler.ApiAdminDeleteCourse)

	// // student api handlers

	router.GET("/api/student/course/:id", courseHandler.ApiStudentGetCourse)
	router.GET("/api/student/profile/:id", profileHandler.ApiStudentProfile)
	router.GET("/api/student/courses", courseHandler.ApiStudentGetCourses)
	router.GET("/api/students/profile", profileHandler.ApiStudentsProfile)

	router.GET("/api/student/notification", notificationHandler.ApiStudentGetNotification)
	router.POST("/api/student/attendance/new", attendanceHandler.ApiStudentPostAttendance)
	router.GET("/api/student/attendance/check/:id", attendanceHandler.ApiStudentCheckAttendance)
	router.GET("/api/student/attendance/show", attendanceHandler.ApiStudentShowAttendance)

	// // teacher api handlers

	router.GET("/api/teacher/profile/:id", profileHandler.ApiTeacherProfile)
	router.POST("/api/teacher/assessment/new", assessmentHandler.ApiTeacherPostAssessment)
	router.POST("/api/teacher/notification", notificationHandler.ApiTeacherPostNotification)

	http.ListenAndServe(":8080", router)
}
