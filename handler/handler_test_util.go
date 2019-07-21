package handler

import (
	"encoding/json"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/k-yomo/go_echo_boilerplate/config"
	"github.com/k-yomo/go_echo_boilerplate/internal/error_handler"
	"github.com/k-yomo/go_echo_boilerplate/internal/test_util/mock"
	"github.com/k-yomo/go_echo_boilerplate/middleware/jwt_middleware"
	"github.com/k-yomo/go_echo_boilerplate/pkg/jwt_generator"
	"github.com/k-yomo/go_echo_boilerplate/pkg/params_validator"
	"github.com/k-yomo/go_echo_boilerplate/repository"
	"github.com/k-yomo/go_echo_boilerplate/usecase"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http/httptest"
	"strings"
	"testing"
)

// ErrorResponse represents api response
type ErrorResponse struct {
	Code   string
	Errors []string
}

// setupTestServer sets up for integration test, returns echo instance and db
func setupTestServer() (*echo.Echo, *sqlx.DB, func()) {
	mock.FakeClockNow()
	db, err := config.NewTestDB()
	if err != nil {
		panic(err)
	}
	e := echo.New()
	e.Validator = params_validator.NewValidator()
	e.HTTPErrorHandler = error_handler.HTTPErrorHandler
	setHandler(e, db)
	return e, db, func() {
		truncateTables(db)
	}
}

func setHandler(e *echo.Echo, db *sqlx.DB) {
	repo := repository.NewRepository(db)
	uc := usecase.NewUsecase(repo, &mock.SMSMessenger{})
	NewHandler(e, uc, jwt_middleware.NewJWTMiddleware())
}

func truncateTables(db *sqlx.DB) {
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		panic(errors.Wrap(err, "show tables failed"))
	}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		err = rows.Scan(&tableName)
		if err != nil {
			panic(errors.Wrap(err, "show table failed"))
			continue
		}

		cmds := []string{
			"SET FOREIGN_KEY_CHECKS = 0",
			fmt.Sprintf("TRUNCATE %s", tableName),
			"SET FOREIGN_KEY_CHECKS = 1",
		}
		for _, cmd := range cmds {
			if _, err := db.Exec(cmd); err != nil {
				panic(errors.Wrap(err, "truncate table failed"))
				continue
			}
		}
	}
}

// handlerTestRunner prepare empty tables
func handlerTestRunner(m *testing.M) int {
	dbConfig, err := config.NewTestDBConfig()
	if err != nil {
		panic(errors.Wrap(err, "initialize db config failed"))
	}
	mg, err := migrate.New("file://../migrations/", config.NewDsn(dbConfig))
	if err != nil {
		panic(errors.Wrap(err, "initialize migrate instance failed"))
	}
	if err := mg.Drop(); err != nil && err != migrate.ErrNoChange {
		panic(errors.Wrap(err, "drop database failed"))
	}
	// need to be renewed to create schema_migrations
	mg, _ = migrate.New("file://../migrations/", config.NewDsn(dbConfig))
	if err := mg.Up(); err != nil {
		panic(errors.Wrap(err, "run migrations failed"))
	}

	return m.Run()
}

// makeRequest generates new http request
func makeRequest(e *echo.Echo, method, path, payload string, body interface{}, headers map[string]string) (resCode int, resBody string) {
	req := httptest.NewRequest(method, path, strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	resBody = rec.Body.String()
	_ = json.NewDecoder(rec.Body).Decode(body)
	return rec.Code, resBody
}

func generateAuthHeader(t *testing.T, id uint64) map[string]string {
	token, err := jwt_generator.GenerateJwt(id)
	if err != nil {
		t.Error(errors.Wrap(err, "failed to generate token"))
	}
	return map[string]string{"Authorization": fmt.Sprintf("Bearer %s", token)}
}
