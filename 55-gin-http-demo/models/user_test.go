package models_test

import (
	"demo/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserValidate(t *testing.T) {
	user := &models.User{Id: 101, Name: "Jiten", Email: "Jitenp@outlook.com", Mobile: "12312"}
	err := user.Validate() // should not give error
	if err != nil {
		t.Fail()
	}
	user = &models.User{Id: 101, Email: "Jitenp@outlook.com", Mobile: "12312"}
	err = user.Validate() // should give error
	if err == nil {
		t.Fail()
	}

}

func TestUserValidate2(t *testing.T) {
	user := &models.User{Id: 101, Name: "Jiten", Email: "Jitenp@outlook.com", Mobile: "12312"}
	err := user.Validate() // should not give error
	assert.NoError(t, err)
	//user = &models.User{Id: 101, Email: "Jitenp@outlook.com", Mobile: "12312"}
	//err = user.Validate() // should give error
	//assert.NoError(t, err)
}

func TestUserMultiValidate(t *testing.T) {

	testcases := make([]struct {
		name string
		user *models.User
	}, 3)
	testcases[0] = struct {
		name string
		user *models.User
	}{
		name: "test missing Name", user: &models.User{Id: 101, Email: "Jitenp@outlook.com", Mobile: "12312"}}

	testcases[1] = struct {
		name string
		user *models.User
	}{
		name: "test missing Email", user: &models.User{Id: 101, Name: "Jiten", Mobile: "12312"}}

	testcases[2] = struct {
		name string
		user *models.User
	}{
		name: "test missing Mobile", user: &models.User{Id: 101, Name: "Jiten", Email: "jitenp@outlook.com"}}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.user.Validate()
			if err == nil {
				t.Errorf("error cannot be nil:test case is %s", tc.name)
				t.Fail()
			}
		})
	}
}
