package handler

import (
	"fmt"
	"github.com/k-yomo/go_echo_boilerplate/internal/test_util/fixture"
	"github.com/k-yomo/go_echo_boilerplate/internal/test_util/mock"
	"github.com/k-yomo/go_echo_boilerplate/model"
	"github.com/k-yomo/go_echo_boilerplate/pkg/clock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v3"
	"net/http"
	"testing"
	"time"
)

type tempSignUpResponse struct {
	Id          int
	PhoneNumber string
	AuthKey     string
	CreatedAt   string
}

func TestTempSignUp(t *testing.T) {
	path := "/v1/auth/temp_sign_up"

	t.Run("with valid request body", func(t *testing.T) {
		e, _, tearDown := setupTestServer()
		defer tearDown()
		resBody := new(tempSignUpResponse)
		resCode, _ := makeRequest(e, http.MethodPost, path, `{"phoneNumber": "08011112222","region":"JP"}`, resBody, nil)

		expectedBody := &tempSignUpResponse{Id: 1, PhoneNumber: "080-1111-2222", CreatedAt: clock.Now().Format(time.RFC3339)}
		assert.Equal(t, http.StatusCreated, resCode)
		assert.Equal(t, expectedBody.Id, resBody.Id)
		assert.Equal(t, expectedBody.PhoneNumber, resBody.PhoneNumber)
		assert.Equal(t, 20, len(resBody.AuthKey))
		assert.Equal(t, expectedBody.CreatedAt, resBody.CreatedAt)
	})

	t.Run("with invalid request body", func(t *testing.T) {
		e, db, tearDown := setupTestServer()
		defer tearDown()
		fixture.CreateUser(t, db, &model.User{PhoneNumber: "+818011112222"})

		testCases := map[string]struct {
			payload            string
			expectedStatusCode int
			expectedBody       *ErrorResponse
		}{
			"invalid phone number": {
				`{"phoneNumber": "abcdef", "region": "JP"}`,
				http.StatusUnprocessableEntity,
				&ErrorResponse{Code: "InvalidParams", Errors: []string{"Key: 'TempSignUpInput.PhoneNumber' Error:Field validation for 'PhoneNumber' failed on the 'phoneNumber' tag"}},
			},
			"invalid region format": {
				`{"phoneNumber": "08011112222", "region": "11"}`,
				http.StatusUnprocessableEntity,
				&ErrorResponse{Code: "InvalidParams", Errors: []string{"Key: 'TempSignUpInput.PhoneNumber' Error:Field validation for 'PhoneNumber' failed on the 'phoneNumber' tag", "Key: 'TempSignUpInput.Region' Error:Field validation for 'Region' failed on the 'phoneNumberRegion' tag"}},
			},
			"phone number is already taken": {
				`{"phoneNumber": "08011112222", "region": "JP"}`,
				http.StatusConflict,
				&ErrorResponse{Code: "AlreadyTaken", Errors: []string{"phoneNumber: 08011112222 is already taken"}},
			},
		}

		for _, tc := range testCases {
			body := new(ErrorResponse)
			resCode, _ := makeRequest(e, http.MethodPost, path, tc.payload, body, nil)
			assert.Equal(t, tc.expectedStatusCode, resCode)
			assert.Equal(t, tc.expectedBody.Code, body.Code)
			assert.EqualValues(t, tc.expectedBody.Errors, body.Errors)
		}
	})
}

type confirmTempUserResponse struct {
	AuthToken string
}

func TestConfirmTempUser(t *testing.T) {
	path := "/v1/auth/confirm"

	t.Run("with valid request body", func(t *testing.T) {
		e, db, tearDown := setupTestServer()
		defer tearDown()
		tempUser := fixture.CreateTempUser(t, db, "")

		reqBody := fmt.Sprintf(`{"tempUserId": %d, "authCode":"%s", "authKey": "%s"}`, tempUser.ID, tempUser.AuthCode, tempUser.AuthKey)
		resBody := new(confirmTempUserResponse)
		resCode, _ := makeRequest(e, http.MethodPost, path, reqBody, resBody, nil)

		assert.Equal(t, http.StatusOK, resCode)
		assert.NotEmpty(t, resBody.AuthToken)
	})

	t.Run("with invalid request body", func(t *testing.T) {
		e, db, tearDown := setupTestServer()
		defer tearDown()
		tempUser := fixture.CreateTempUser(t, db, "")

		testCases := map[string]struct {
			payload            string
			expectedStatusCode int
			expectedBody       *ErrorResponse
		}{
			"nonexistent temp user": {
				fmt.Sprintf(`{"tempUserId": 12345, "authCode":"%s", "authKey": "%s"}`, tempUser.AuthCode, tempUser.AuthKey),
				http.StatusNotFound,
				&ErrorResponse{Code: "NotFound", Errors: []string{"TempUser with id = 12345 is not found"}},
			},
			"incorrect auth code": {
				fmt.Sprintf(`{"tempUserId": %d, "authCode":"incorrect token", "authKey": "%s"}`, tempUser.ID, tempUser.AuthKey),
				http.StatusUnauthorized,
				&ErrorResponse{Code: "Unauthenticated", Errors: []string{"Unauthenticated"}},
			},
			"incorrect auth key": {
				fmt.Sprintf(`{"tempUserId": %d, "authCode":"%s", "authKey": "incorrect token"}`, tempUser.ID, tempUser.AuthCode),
				http.StatusUnauthorized,
				&ErrorResponse{Code: "Unauthenticated", Errors: []string{"Unauthenticated"}},
			},
		}

		for _, tc := range testCases {
			resBody := new(ErrorResponse)
			resCode, _ := makeRequest(e, http.MethodPost, path, tc.payload, resBody, nil)
			assert.Equal(t, tc.expectedStatusCode, resCode)
			assert.Equal(t, tc.expectedBody.Code, resBody.Code)
			assert.EqualValues(t, tc.expectedBody.Errors, resBody.Errors)
		}
	})

	t.Run("when request made after 30 minutes have passed", func(t *testing.T) {
		e, db, tearDown := setupTestServer()
		defer tearDown()
		tempUser := fixture.CreateTempUser(t, db, "")

		mock.FakeClockNowWithExtraTime(time.Minute * 31)
		reqBody := fmt.Sprintf(`{"tempUserId": %d, "authCode":"%s", "authKey": "%s"}`, tempUser.ID, tempUser.AuthCode, tempUser.AuthKey)
		resBody := new(ErrorResponse)
		resCode, _ := makeRequest(e, http.MethodPost, path, reqBody, resBody, nil)

		expectedBody := &ErrorResponse{Code: "Expired", Errors: []string{fmt.Sprintf("SMS confirmation with id %d is expired", tempUser.ID)}}
		assert.Equal(t, http.StatusUnauthorized, resCode)
		assert.Equal(t, expectedBody.Code, resBody.Code)
		assert.EqualValues(t, expectedBody.Errors, resBody.Errors)
	})
}

type userResponse struct {
	ID          uint64
	PhoneNumber string
	Email       string
	FirstName   string
	LastName    string
	DateOfBirth null.String
	Gender      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	AuthToken   string
}

func TestSignUp(t *testing.T) {
	path := "/v1/auth/sign_up"

	t.Run("with valid auth header", func(t *testing.T) {
		e, db, tearDown := setupTestServer()
		defer tearDown()
		tempUser := fixture.CreateTempUser(t, db, "")

		t.Run("with valid request body", func(t *testing.T) {
			reqBody := fmt.Sprintf(`{"firstName": "貫児", "lastName": "四方田", "email": "test@example.com", "password":"password", "gender":"male", "dateOfBirth":"1995-07-05"}`)
			resBody := new(userResponse)
			resCode, _ := makeRequest(e, http.MethodPost, path, reqBody, resBody, generateAuthHeader(t, tempUser.ID))

			expectedBody := &userResponse{Email: "test@example.com", FirstName: "貫児", LastName: "四方田", Gender: "male", DateOfBirth: null.NewString("1995-07-05", true)}
			assert.Equal(t, http.StatusCreated, resCode)
			assert.Equal(t, expectedBody.Email, resBody.Email)
			assert.Equal(t, expectedBody.FirstName, resBody.FirstName)
			assert.Equal(t, expectedBody.LastName, resBody.LastName)
			assert.Equal(t, expectedBody.Gender, resBody.Gender)
			assert.Equal(t, expectedBody.DateOfBirth, resBody.DateOfBirth)
			assert.NotEmpty(t, resBody.AuthToken)
		})

		t.Run("with invalid request body", func(t *testing.T) {
			reqBody := fmt.Sprintf(`{"firstName": "", "lastName": "", "email": "test.com", "password":"pass"}`)
			resBody := new(ErrorResponse)
			resCode, _ := makeRequest(e, http.MethodPost, path, reqBody, resBody, generateAuthHeader(t, tempUser.ID))

			expectedBody := &ErrorResponse{Code: "InvalidParams", Errors: []string{"Key: 'SignUpInput.Email' Error:Field validation for 'Email' failed on the 'email' tag", "Key: 'SignUpInput.Password' Error:Field validation for 'Password' failed on the 'min' tag", "Key: 'SignUpInput.FirstName' Error:Field validation for 'FirstName' failed on the 'required' tag", "Key: 'SignUpInput.LastName' Error:Field validation for 'LastName' failed on the 'required' tag"}}
			assert.Equal(t, http.StatusUnprocessableEntity, resCode)
			assert.Equal(t, expectedBody.Code, resBody.Code)
			assert.Equal(t, expectedBody.Errors, resBody.Errors)
		})
	})

	t.Run("with invalid auth header", func(t *testing.T) {
		e, _, tearDown := setupTestServer()
		defer tearDown()

		reqBody := `{"firstName": "貫児", "lastName": "四方田", "email": "test@example.com", "gender": "unknown"}`
		resBody := new(ErrorResponse)
		resCode, _ := makeRequest(e, http.MethodPost, path, reqBody, resBody, nil)

		expectedBody := &ErrorResponse{Code: "Unauthenticated", Errors: []string{"Unauthenticated"}}
		assert.Equal(t, http.StatusUnauthorized, resCode)
		assert.Equal(t, expectedBody.Code, resBody.Code)
		assert.Equal(t, expectedBody.Errors, resBody.Errors)
	})
}

func TestSignIn(t *testing.T) {
	path := "/v1/auth/sign_in"
	now := time.Now()
	clock.Now = func() time.Time { return now }

	e, db, tearDown := setupTestServer()
	defer tearDown()
	user := fixture.CreateUser(t, db, nil)

	t.Run("with valid sign in info", func(t *testing.T) {
		reqBody := fmt.Sprintf(`{"phoneNumber": "08011112222", "region":"JP", "password":"password"}`)
		resBody := new(userResponse)
		resCode, _ := makeRequest(e, http.MethodPost, path, reqBody, resBody, nil)

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

	t.Run("with invalid sign in info", func(t *testing.T) {
		reqBody := fmt.Sprintf(`{"phoneNumber": "08011112222", "region":"JP", "password":"wrongpassword"}`)
		resBody := new(ErrorResponse)
		resCode, _ := makeRequest(e, http.MethodPost, path, reqBody, resBody, nil)

		expectedBody := &ErrorResponse{Code: "Unauthenticated", Errors: []string{"Unauthenticated"}}
		assert.Equal(t, http.StatusUnauthorized, resCode)
		assert.Equal(t, expectedBody.Code, resBody.Code)
		assert.Equal(t, expectedBody.Errors, resBody.Errors)
	})
}

type smsReconfirmationResponse struct {
	ID          int
	PhoneNumber string
	CreatedAt   time.Time
}

func TestUpdateUnconfirmedPhoneNumber(t *testing.T) {
	path := "/v1/auth/phone_number"
	e, db, tearDown := setupTestServer()
	defer tearDown()
	user := fixture.CreateUser(t, db, nil)

	t.Run("with valid body and auth header", func(t *testing.T) {
		reqBody := `{"phoneNumber": "080-9999-8888", "region": "JP"}`
		resBody := new(smsReconfirmationResponse)
		resCode, _ := makeRequest(e, http.MethodPatch, path, reqBody, resBody, generateAuthHeader(t, user.ID))

		expectedBody := &smsReconfirmationResponse{PhoneNumber: "080-9999-8888"}
		assert.Equal(t, http.StatusOK, resCode)
		assert.Equal(t, expectedBody.PhoneNumber, resBody.PhoneNumber)
	})

	t.Run("with invalid phone number format", func(t *testing.T) {
		reqBody := `{"phoneNumber": "080-9999-88", "region": "JP"}`
		resBody := new(ErrorResponse)
		resCode, _ := makeRequest(e, http.MethodPatch, path, reqBody, resBody, generateAuthHeader(t, user.ID))

		expectedBody := &ErrorResponse{Code: "InvalidParams", Errors: []string{"Key: 'UpdateUnconfirmedPhoneNumberInput.PhoneNumber' Error:Field validation for 'PhoneNumber' failed on the 'phoneNumber' tag"}}
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

func TestConfirmPhoneNumber(t *testing.T) {
	path := "/v1/auth/phone_number/confirm"
	t.Run("when unconfirmed number is registered", func(t *testing.T) {
		t.Run("with valid request body and header", func(t *testing.T) {
			e, db, tearDown := setupTestServer()
			defer tearDown()
			sr := fixture.CreateSMSReconfirmation(t, db, nil, nil)

			reqBody := fmt.Sprintf(`{"authCode":"%s"}`, sr.AuthCode)
			resBody := new(userResponse)
			resCode, _ := makeRequest(e, http.MethodPost, path, reqBody, resBody, generateAuthHeader(t, sr.ID))

			assert.Equal(t, http.StatusOK, resCode)
			assert.NotEmpty(t, resBody.AuthToken)
		})

		t.Run("with incorrect auth code", func(t *testing.T) {
			e, db, tearDown := setupTestServer()
			defer tearDown()
			sr := fixture.CreateSMSReconfirmation(t, db, nil, nil)

			reqBody := `{"authCode":"invalid"}`
			resBody := new(ErrorResponse)
			resCode, _ := makeRequest(e, http.MethodPost, path, reqBody, resBody, generateAuthHeader(t, sr.UserID))

			expectedBody := &ErrorResponse{Code: "Unauthenticated", Errors: []string{"Unauthenticated"}}
			assert.Equal(t, http.StatusUnauthorized, resCode)
			assert.Equal(t, expectedBody.Code, resBody.Code)
			assert.EqualValues(t, expectedBody.Errors, resBody.Errors)
		})

		t.Run("when request made after 30 minutes have passed", func(t *testing.T) {
			e, db, tearDown := setupTestServer()
			defer tearDown()
			sr := fixture.CreateSMSReconfirmation(t, db, nil, nil)
			authHeader := generateAuthHeader(t, sr.UserID)
			mock.FakeClockNowWithExtraTime(time.Minute * 31)

			reqBody := fmt.Sprintf(`{"authCode":"%s"}`, sr.AuthCode)
			resBody := new(ErrorResponse)
			resCode, _ := makeRequest(e, http.MethodPost, path, reqBody, resBody, authHeader)

			expectedBody := &ErrorResponse{Code: "Expired", Errors: []string{fmt.Sprintf("SMS reconfirmation with id %d is expired", sr.ID)}}
			assert.Equal(t, http.StatusUnauthorized, resCode)
			assert.Equal(t, expectedBody.Code, resBody.Code)
			assert.EqualValues(t, expectedBody.Errors, resBody.Errors)
		})
	})
}
