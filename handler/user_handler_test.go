package handler

import (
	"github.com/k-yomo/go_echo_boilerplate/internal/test_util/fixture"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v3"
	"net/http"
	"testing"
	"time"
)

func TestGetProfile(t *testing.T) {
	path := "/v1/users/self"
	e, db, tearDown := setupTestServer()
	defer tearDown()
	user := fixture.CreateUser(t, db, nil)

	t.Run("with valid auth header", func(t *testing.T) {
		resBody := new(userResponse)
		resCode, _ := makeRequest(e, http.MethodGet, path, "", resBody, generateAuthHeader(t, user.ID))

		expectedBody := &userResponse{
			Email:       user.Email,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Gender:      user.GenderString(),
			DateOfBirth: null.NewString(user.DateOfBirth.Time.Format(time.RFC3339), user.DateOfBirth.Valid),
		}
		assert.Equal(t, http.StatusOK, resCode)
		assert.Equal(t, expectedBody.Email, resBody.Email)
		assert.Equal(t, expectedBody.FirstName, resBody.FirstName)
		assert.Equal(t, expectedBody.LastName, resBody.LastName)
		assert.Equal(t, expectedBody.Gender, resBody.Gender)
		assert.NotEmpty(t, resBody.AuthToken)
	})

	t.Run("with invalid auth header", func(t *testing.T) {
		resBody := new(ErrorResponse)
		resCode, _ := makeRequest(e, http.MethodGet, path, "", resBody, nil)

		expectedBody := &ErrorResponse{Code: "Unauthenticated", Errors: []string{"Unauthenticated"}}
		assert.Equal(t, http.StatusUnauthorized, resCode)
		assert.Equal(t, expectedBody.Code, resBody.Code)
		assert.Equal(t, expectedBody.Errors, resBody.Errors)
	})
}

func TestUpdateProfile(t *testing.T) {
	path := "/v1/users/self"
	e, db, tearDown := setupTestServer()
	defer tearDown()
	user := fixture.CreateUser(t, db, nil)

	t.Run("with valid body and auth header", func(t *testing.T) {
		reqBody := `{"firstName": "new", "lastName": "user", "gender":"male", "dateOfBirth":"1900-01-01"}`
		resBody := new(userResponse)
		resCode, _ := makeRequest(e, http.MethodPatch, path, reqBody, resBody, generateAuthHeader(t, user.ID))

		expectedBody := &userResponse{Email: user.Email, FirstName: "new", LastName: "user", Gender: "male", DateOfBirth: null.NewString("1900-01-01", true)}
		assert.Equal(t, http.StatusOK, resCode)
		assert.Equal(t, expectedBody.Email, resBody.Email)
		assert.Equal(t, expectedBody.FirstName, resBody.FirstName)
		assert.Equal(t, expectedBody.LastName, resBody.LastName)
		assert.Equal(t, expectedBody.Gender, resBody.Gender)
		assert.NotEmpty(t, resBody.AuthToken)
	})

	t.Run("with invalid body", func(t *testing.T) {
		reqBody := `{"firstName": "new", "lastName": "user", "gender": "invalid", "dateOfBirth": "1900-01-01"}`
		resBody := new(ErrorResponse)
		resCode, _ := makeRequest(e, http.MethodPatch, path, reqBody, resBody, generateAuthHeader(t, user.ID))

		expectedBody := &ErrorResponse{Code: "InvalidParams", Errors: []string{"Key: 'UpdateProfileInput.Gender' Error:Field validation for 'Gender' failed on the 'oneof' tag"}}
		assert.Equal(t, http.StatusUnprocessableEntity, resCode)
		assert.Equal(t, expectedBody.Code, resBody.Code)
		assert.Equal(t, expectedBody.Errors, resBody.Errors)
	})

	t.Run("with invalid auth header", func(t *testing.T) {
		resBody := new(ErrorResponse)
		resCode, _ := makeRequest(e, http.MethodPatch, path, "", resBody, nil)

		expectedBody := &ErrorResponse{Code: "Unauthenticated", Errors: []string{"Unauthenticated"}}
		assert.Equal(t, http.StatusUnauthorized, resCode)
		assert.Equal(t, expectedBody.Code, resBody.Code)
		assert.Equal(t, expectedBody.Errors, resBody.Errors)
	})
}

func TestUpdateEmail(t *testing.T) {
	path := "/v1/users/self/email"
	e, db, tearDown := setupTestServer()
	defer tearDown()
	user := fixture.CreateUser(t, db, nil)

	t.Run("with valid body and auth header", func(t *testing.T) {
		reqBody := `{"email": "new@mail.com"}`
		resBody := new(userResponse)
		resCode, _ := makeRequest(e, http.MethodPatch, path, reqBody, resBody, generateAuthHeader(t, user.ID))

		expectedBody := &userResponse{
			Email:       "new@mail.com",
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Gender:      user.GenderString(),
			DateOfBirth: null.NewString(user.DateOfBirth.Time.Format(time.RFC3339), user.DateOfBirth.Valid),
		}
		assert.Equal(t, http.StatusOK, resCode)
		assert.Equal(t, expectedBody.Email, resBody.Email)
		assert.Equal(t, expectedBody.FirstName, resBody.FirstName)
		assert.Equal(t, expectedBody.LastName, resBody.LastName)
		assert.Equal(t, expectedBody.Gender, resBody.Gender)
		assert.NotEmpty(t, resBody.AuthToken)
	})

	t.Run("when user not found", func(t *testing.T) {
		reqBody := `{"email": "new@mail.com"}`
		resBody := new(ErrorResponse)
		resCode, _ := makeRequest(e, http.MethodPatch, path, reqBody, resBody, generateAuthHeader(t, 10000))

		expectedBody := &ErrorResponse{Code: "NotFound", Errors: []string{"User with id = 10000 is not found"}}
		assert.Equal(t, http.StatusNotFound, resCode)
		assert.Equal(t, expectedBody.Code, resBody.Code)
		assert.Equal(t, expectedBody.Errors, resBody.Errors)
	})

	t.Run("with invalid body", func(t *testing.T) {
		reqBody := `{"email": "invalid.com"}`
		resBody := new(ErrorResponse)
		resCode, _ := makeRequest(e, http.MethodPatch, path, reqBody, resBody, generateAuthHeader(t, user.ID))

		expectedBody := &ErrorResponse{Code: "InvalidParams", Errors: []string{"Key: 'UpdateEmailInput.Email' Error:Field validation for 'Email' failed on the 'email' tag"}}
		assert.Equal(t, http.StatusUnprocessableEntity, resCode)
		assert.Equal(t, expectedBody.Code, resBody.Code)
		assert.Equal(t, expectedBody.Errors, resBody.Errors)
	})

	t.Run("with invalid auth header", func(t *testing.T) {
		resBody := new(ErrorResponse)
		resCode, _ := makeRequest(e, http.MethodPatch, path, "", resBody, nil)

		expectedBody := &ErrorResponse{Code: "Unauthenticated", Errors: []string{"Unauthenticated"}}
		assert.Equal(t, http.StatusUnauthorized, resCode)
		assert.Equal(t, expectedBody.Code, resBody.Code)
		assert.Equal(t, expectedBody.Errors, resBody.Errors)
	})
}
