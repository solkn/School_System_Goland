package service

import (
	"github.com/Rob-a21/Cassiopeia/entity"
	"github.com/Rob-a21/Cassiopeia/models"
)

type ProfileServiceImpl struct {
	profileRepository models.ProfileRepository
}

func (ss *ProfileServiceImpl) EmailExists(email string) bool {
	panic("implement me")
}

func NewProfileServiceImpl(profrepo models.ProfileRepository) *ProfileServiceImpl {
	return &ProfileServiceImpl{profileRepository: profrepo}
}

func (ss *ProfileServiceImpl) Students() ([]entity.Student, error) {

	students, err := ss.profileRepository.Students()

	if err != nil {
		return nil, err
	}

	return students, nil
}

func (ss *ProfileServiceImpl) Student(id int) ([]entity.Student, error) {

	student, err := ss.profileRepository.Student(id)

	if err != nil {
		panic(err)
	}

	return student, nil
}

func (ss *ProfileServiceImpl) EmailExist(email string) bool {

	exists := ss.profileRepository.EmailExists(email)

	return exists
}

func (fs *ProfileServiceImpl) Families() ([]entity.Family, error) {

	family, err := fs.profileRepository.Families()

	if err != nil {

		return nil, err
	}

	return family, nil

}

func (ts *ProfileServiceImpl) Teachers() ([]entity.Teacher, error) {

	teacher, err := ts.profileRepository.Teachers()

	if err != nil {

		return nil, err
	}

	return teacher, nil

}

func (ts *ProfileServiceImpl) Teacher(id int) ([]entity.Teacher, error) {

	teacher, err := ts.profileRepository.Teacher(id)

	if err != nil {
		panic(err)
	}

	return teacher, nil
}

func (ss *ProfileServiceImpl) Family(id int) ([]entity.Family, error) {

	family, err := ss.profileRepository.Family(id)

	if err != nil {
		return nil, err
	}

	return family, nil
}

func (ss *ProfileServiceImpl) NewYearRegistration(student entity.Student) error {

	ss.profileRepository.NewYearRegistration(student)

	return nil
}

func (ss *ProfileServiceImpl) Admin(id int) ([]entity.Admin, error) {

	admin, err := ss.profileRepository.Admin(id)

	if err != nil {
		panic(err)
	}

	return admin, nil
}

func (ss *ProfileServiceImpl) AdminByEmail(email string) (entity.Admin, error) {

	admin, err := ss.profileRepository.AdminByEmail(email)

	if err != nil {
		panic(err)
	}

	return admin, nil
}

func (prf *ProfileServiceImpl) Admins() ([]entity.Admin, error) {

	admin, err := prf.profileRepository.Admins()

	if err != nil {

		return nil, err
	}

	return admin, nil
}

func (prf *ProfileServiceImpl) DeleteStudent(id int) error {

	err := prf.profileRepository.DeleteStudent(id)
	if err != nil {
		return err
	}
	return nil
}

func (prf *ProfileServiceImpl) DeleteTeacher(id int) error {

	err := prf.profileRepository.DeleteTeacher(id)
	if err != nil {
		return err
	}
	return nil
}
