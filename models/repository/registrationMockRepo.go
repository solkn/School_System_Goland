package repository

import (
	"database/sql"
	"errors"

	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/Rob-a21/Cassiopeia/models"
)

type registrationMockRepo struct {
	conn *sql.DB
}

func NewregistrationMockRepo(db *sql.DB) models.RegistrationRepository {
	return &registrationMockRepo{conn: db}
}

func (registrationMockRepo) RegisterStudent(st entity.Student) error {

	entity.StudentMock.UserName = st.UserName
	entity.StudentMock.Password = st.Password
	entity.StudentMock.FirstName = st.FirstName
	entity.StudentMock.LastName = st.LastName
	entity.StudentMock.ID = st.ID
	entity.StudentMock.Grade = st.Grade
	entity.StudentMock.Email = st.Email

	if st.ID != entity.StudentMock.ID {
		return errors.New("Failed Registration")
	}
	return nil
}

func (registrationMockRepo) RegisterFamily(family entity.Family) error {
	entity.FamilyMock.FirstName = family.FirstName
	entity.FamilyMock.LastName = family.LastName
	entity.FamilyMock.Username = family.Username
	entity.FamilyMock.Password = family.Password
	entity.FamilyMock.FamilyID = family.FamilyID
	entity.FamilyMock.Phone = family.Phone
	entity.FamilyMock.Email = family.Email

	if family.FamilyID != entity.FamilyMock.FamilyID {
		return errors.New("Failed Registration")
	}
	return nil
}

func (registrationMockRepo) RegisterTeacher(teacher entity.Teacher) error {
	entity.TeacherMock.UserName = teacher.UserName
	entity.TeacherMock.Password = teacher.Password
	entity.TeacherMock.Phone = teacher.Phone
	entity.TeacherMock.Email = teacher.Email
	entity.TeacherMock.FirstName = teacher.FirstName
	entity.TeacherMock.LastName = teacher.LastName
	entity.TeacherMock.TeacherID = teacher.TeacherID

	if teacher.TeacherID != entity.TeacherMock.TeacherID {
		return errors.New("Failed Registration")
	}
	return nil
}

func (registrationMockRepo) RegisterAdmin(admin entity.Admin) error {
	entity.AdminMock.UserName = admin.UserName
	entity.AdminMock.Password = admin.Password
	entity.AdminMock.FirstName = admin.FirstName
	entity.AdminMock.LastName = admin.LastName
	entity.AdminMock.Email = admin.Email
	entity.AdminMock.Image = admin.Image

	if admin.UserName != entity.AdminMock.UserName {
		return errors.New("Failed Registration of Admin")
	}
	return nil
}
