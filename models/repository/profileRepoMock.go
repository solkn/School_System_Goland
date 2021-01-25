package repository

import (
	"database/sql"
	"errors"

	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/Rob-a21/Cassiopeia/models"
)

type profileRepoMock struct {
	conn *sql.DB
}

func NewprofileRepoMock(db *sql.DB) models.ProfileRepository {
	return &profileRepoMock{conn: db}
}

func (profileRepoMock) Students() ([]entity.Student, error) {

	students := []entity.Student{entity.StudentMock}

	return students, nil
}

func (profileRepoMock) Student(id int) ([]entity.Student, error) {
	//null := entity.Student{}
	student := []entity.Student{entity.StudentMock}

	if id == 0001 {
		return student, nil
	}

	return student, nil
}

func (profileRepoMock) EmailExists(email string) bool {

	if email == entity.TeacherMock.Email || email == entity.StudentMock.Email || email == entity.AdminMock.Email {
		return false
	}
	return true
}

func (profileRepoMock) DeleteStudent(id int) error {
	if id != 0001 {
		return errors.New("Delete Has Failed...!")
	}

	return nil
}

func (profileRepoMock) Families() ([]entity.Family, error) {
	families := []entity.Family{entity.FamilyMock}

	return families, nil
}

func (profileRepoMock) Family(id int) ([]entity.Family, error) {
	family := []entity.Family{entity.FamilyMock}

	return family, nil
}

func (profileRepoMock) Teachers() ([]entity.Teacher, error) {
	teachers := []entity.Teacher{entity.TeacherMock}

	return teachers, nil
}

func (profileRepoMock) Teacher(id int) ([]entity.Teacher, error) {
	teacher := []entity.Teacher{entity.TeacherMock}

	if id == 0001 {
		return teacher, nil
	}

	return teacher, nil
}

func (profileRepoMock) DeleteTeacher(id int) error {

	if id != 0001 {
		return errors.New("Delete Has Failed...!")
	}

	return nil
}

func (profileRepoMock) Admin(id int) ([]entity.Admin, error) {

	admins := []entity.Admin{entity.AdminMock}

	return admins, nil
}

func (profileRepoMock) AdminByEmail(email string) (entity.Admin, error) {

	null := entity.Admin{}

	if email == "someone@gmail.com" {
		return entity.AdminMock, nil
	}

	return null, nil
}

func (profileRepoMock) Admins() ([]entity.Admin, error) {
	admins := []entity.Admin{entity.AdminMock}

	return admins, nil
}
func (profileRepoMock) NewYearRegistration(student entity.Student) error {

	_ = entity.StudentMock

	return nil
}
