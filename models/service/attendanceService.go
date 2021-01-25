//Attendanceservice

package service

import (
	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/Rob-a21/Cassiopeia/models"
)

type StudentAttendanceServiceImpl struct {
	attendanceRepo models.StudentAttendanceRepository
}

func NewStudentAttendanceServiceImpl(attRepo models.StudentAttendanceRepository) *StudentAttendanceServiceImpl {
	return &StudentAttendanceServiceImpl{attendanceRepo: attRepo}
}

func (at *StudentAttendanceServiceImpl) ShowAttendance() ([]entity.Attendance, error) {

	attendance, err := at.attendanceRepo.ShowAttendance()

	if err != nil {
		return nil, err
	}

	return attendance, nil
}

func (at *StudentAttendanceServiceImpl) CheckAttendance(id int) ([]entity.Attendance, error) {

	student, err := at.attendanceRepo.CheckAttendance(id)

	if err != nil {
		return student, err
	}

	return student, nil
}

func (at *StudentAttendanceServiceImpl) FillAttendance(student entity.Attendance) error {

	err := at.attendanceRepo.FillAttendance(student)

	if err != nil {
		return err
	}

	return nil
}
