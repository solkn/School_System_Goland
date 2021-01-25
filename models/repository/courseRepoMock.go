package repository

import (
	"database/sql"
	"errors"

	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/Rob-a21/Cassiopeia/models"
)

type courseRepoMock struct {
	conn *sql.DB
}

func NewcourseRepoMock(db *sql.DB) models.CourseRepository {
	return &courseRepoMock{conn: db}
}

func (aRepo *courseRepoMock) AddCourse(course entity.Course) error {
	course = entity.CourseMock

	return nil
}

func (aRepo *courseRepoMock) GetCourse() ([]entity.Course, error) {
	posts := []entity.Course{entity.CourseMock}

	return posts, nil
}

func (aRepo *courseRepoMock) Course(id int) (*entity.Course, error) {

	c := entity.Course{}

	if id == 0001 {
		return &entity.CourseMock, nil
	}

	return &c, nil

}

func (aRepo *courseRepoMock) UpdateCourse(course entity.Course) error {
	course = entity.CourseMock

	return nil
}

func (aRepo *courseRepoMock) DeleteCourse(id int) error {

	//id = uint(entity.CourseMock.CourseID)
	//post := entity.CourseMock

	if id != 0001 {
		return errors.New("Delete Test Failed")
	}

	return nil
}
