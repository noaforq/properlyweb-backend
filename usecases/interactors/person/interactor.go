package person

import (
	"strconv"
	"strings"
	"time"

	"properly.jp/properlyweb-backend/entities"
)

type PersonPort interface {
	FindById(id int) (entities.Person, error)
}

type PersonInteractor struct {
	PersonPort PersonPort
}

func (interactor *PersonInteractor) FindById(id int) (person entities.Person, err error) {
	person, err = interactor.PersonPort.FindById(id)
	person.FirstName = "*****"
	person.LastName = "*****"
	person.Age, _ = GetAge(person.Birthday)

	return
}

func GetAge(birthday string) (age int, e error) {
	now := time.Now().Format("20060102")
	timeparsedBirthday, _ := time.Parse("20060102", strings.Replace(birthday, "-", "", -1))
	target := timeparsedBirthday.Format("20060102")
	nowInt, _ := strconv.Atoi(now)
	targetInt, _ := strconv.Atoi(target)
	age = (nowInt - targetInt) / 10000
	return
}
