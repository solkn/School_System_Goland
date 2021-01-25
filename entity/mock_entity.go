package entity

import "time"

var CourseMock = Course{

	CourseID:   1,
	CourseName: "Mock 01",
	Grade:      2,
}

var AttendanceMock = Attendance{
	Date:      time.Now(),
	StudentID: 1,
}

var AdminMock = Admin{
	UserName:  "admin",
	LastName:  "adminn",
	Password:  "password admin",
	FirstName: "admin",
	Email:     "someone@gmail.com",
	Image:     "C/Users/rob_a21/Downloads/Telegram Desktop/photo_2020-01-19_17-12-06.jpg",
}

var StudentMock = Student{
	UserName:  "username",
	Password:  "student password",
	FirstName: "student 01",
	LastName:  "stud.",
	ID:        0001,
	Email:     "studentEmail@gmail.com",
	Grade:     12,
	Image:     "C/Users/rob_a21/Downloads/Telegram Desktop/photo_2020-01-19_17-12-06.jpg",
}

var TeacherMock = Teacher{
	UserName:  "teacher",
	Password:  "teach021",
	Phone:     "09191919191",
	Email:     "teacher@gmail.com",
	FirstName: "teacher01",
	LastName:  "tch",
	TeacherID: 0001,
	Image:     "C/Users/rob_a21/Downloads/Telegram Desktop/photo_2020-01-19_17-12-06.jpg",
}

var FamilyMock = Family{
	FirstName: "family 01",
	LastName:  "fam",
	Username:  "fam01",
	Password:  "familypass",
	FamilyID:  0001,
	Phone:     "09191938134",
	Email:     "family@gmail.com",
	Image:     "C/Users/rob_a21/Downloads/Telegram Desktop/photo_2020-01-19_17-12-06.jpg",
}

var AssessmentMock = Assessment{
	Value:     99,
	SubjectID: 01,
	StudentID: 0001,
	Grade:     10,
}

var NotificationMock = Notification{
	Message:          "message 01",
	NotifyName:       "notifier",
	NotificationDate: time.Now(),
}

var SessionMock = Session{
	ID:      0001,
	UUID:    "uuid",
	Expires: 011120,
}
