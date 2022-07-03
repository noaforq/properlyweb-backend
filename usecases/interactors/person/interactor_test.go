package person

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"properly.jp/properlyweb-backend/entities"
)

// 0. 実装されている
// 1. 個人情報を返却する
// 2. 氏名項目をマスクして返却する
// 3. 生年月日から年齢を計算し年齢項目に格納して返却する

type MockPersonPort struct {
	MockFindById func(id int) (entities.Person, error)
}

func (mock *MockPersonPort) FindById(id int) (person entities.Person, err error) {
	return mock.MockFindById(id)
}

// 0. 実装されている
func Test実装されている(t *testing.T) {
	interactor := &PersonInteractor{}
	interactor.PersonPort = &MockPersonPort{
		MockFindById: func(id int) (entities.Person, error) { return entities.Person{}, nil },
	}
	interactor.FindById(1)
}

// 1. 個人情報を返却する
func Test個人情報を返却する(t *testing.T) {
	interactor := &PersonInteractor{}
	interactor.PersonPort = &MockPersonPort{
		MockFindById: func(id int) (entities.Person, error) { return entities.Person{}, nil },
	}
	person, _ := interactor.FindById(1)

	assert.NotNil(t, person)
}

// 2. 氏名項目をマスクして返却する
func Test氏名項目をマスクして返却する(t *testing.T) {
	interactor := &PersonInteractor{}
	interactor.PersonPort = &MockPersonPort{
		MockFindById: func(id int) (entities.Person, error) {
			return entities.Person{
				Id:        1,
				FirstName: "Nobuhiko",
				LastName:  "Asano",
				Gender:    1,
				Birthday:  "1973-04-09",
			}, nil
		},
	}
	person, _ := interactor.FindById(1)
	assert.Equal(t, 1, person.Id)
	assert.Equal(t, "*****", person.FirstName)
	assert.Equal(t, "*****", person.LastName)
	assert.Equal(t, 1, person.Gender)
	assert.Equal(t, "1973-04-09", person.Birthday)
}

// 3. 氏生年月日から年齢を計算し年齢項目に格納して返却する
func Test生年月日から年齢を計算し年齢項目に格納して返却する(t *testing.T) {
	interactor := &PersonInteractor{}
	interactor.PersonPort = &MockPersonPort{
		MockFindById: func(id int) (entities.Person, error) {
			return entities.Person{
				Id:        1,
				FirstName: "Nobuhiko",
				LastName:  "Asano",
				Gender:    1,
				Birthday:  "1973-04-09",
			}, nil
		},
	}
	person, _ := interactor.FindById(1)
	assert.Equal(t, 49, person.Age)
}
