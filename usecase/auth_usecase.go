package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/k-yomo/go_echo_boilerplate/internal/custom_context"
	"github.com/k-yomo/go_echo_boilerplate/model"
	"github.com/k-yomo/go_echo_boilerplate/pkg/jwt_generator"
	"github.com/k-yomo/go_echo_boilerplate/pkg/sms"
	"github.com/k-yomo/go_echo_boilerplate/repository"
	"github.com/k-yomo/go_echo_boilerplate/usecase/input"
	"github.com/k-yomo/go_echo_boilerplate/usecase/output"
	"github.com/pkg/errors"
	"github.com/ttacon/libphonenumber"
	"golang.org/x/crypto/bcrypt"
)

// AuthUsecase is an interface for Auth related usecase
type AuthUsecase interface {
	TempSignUp(ctx *custom_context.Context) (*output.TempUserOutput, error)
	ConfirmTempUser(ctx *custom_context.Context) (*output.AuthTokenOutput, error)
	SignUp(ctx *custom_context.Context) (*output.CurrentUserOutput, error)
	SignIn(ctx *custom_context.Context) (*output.CurrentUserOutput, error)
	UpdateUnconfirmedPhoneNumber(ctx *custom_context.Context) (*output.SMSReconfirmationOutput, error)
	ConfirmPhoneNumber(ctx *custom_context.Context) (*output.CurrentUserOutput, error)
}
type authUsecase struct {
	TempUserRepo repository.TempUserRepository
	UserRepo     repository.UserRepository
	SMSMessenger sms.SMSMessenger
}

// NewAuthUsecase returns authUsecase
func NewAuthUsecase(tempUserRepo repository.TempUserRepository, userRepo repository.UserRepository, smsMessenger sms.SMSMessenger) *authUsecase {
	return &authUsecase{TempUserRepo: tempUserRepo, UserRepo: userRepo, SMSMessenger: smsMessenger}
}

// TempSignUp registers temporary user
func (au *authUsecase) TempSignUp(ctx *custom_context.Context) (*output.TempUserOutput, error) {
	params := new(input.TempSignUpInput)
	if err := ctx.BindWithValidation(params); err != nil {
		return nil, err
	}

	normalizedPhoneNumber, _ := libphonenumber.Parse(params.PhoneNumber, params.Region)
	tempUser := model.NewTempUser(libphonenumber.Format(normalizedPhoneNumber, libphonenumber.E164))
	if user, err := au.UserRepo.FindByPhoneNumber(context.Background(), tempUser.PhoneNumber); err != nil {
		return nil, err
	} else if user != nil {
		return nil, newAlreadyTakenError("phoneNumber", params.PhoneNumber)
	}

	newTempUser, err := au.TempUserRepo.UpsertByPhoneNumber(context.Background(), tempUser)
	if err != nil {
		return nil, err
	}
	smsBody := fmt.Sprintf("【APP_TITLE】認証コード: %s\nコードは30分間有効です。", newTempUser.AuthCode)
	if err := au.SMSMessenger.SendSMS(newTempUser.PhoneNumber, smsBody); err != nil {
		return nil, err
	}
	return output.NewTempUserOutput(newTempUser), nil
}

// ConfirmTempUser confirms temp user's phone number by validating authCode sent to phone number and authToken sent to client
func (au *authUsecase) ConfirmTempUser(ctx *custom_context.Context) (*output.AuthTokenOutput, error) {
	params := new(input.ConfirmTempUserInput)
	if err := ctx.BindWithValidation(params); err != nil {
		return nil, err
	}

	tu, err := au.TempUserRepo.FindByID(context.Background(), params.TempUserID)
	if err != nil {
		return nil, err
	} else if tu == nil {
		return nil, newNotFoundError("TempUser", "id", params.TempUserID)
	}

	if tu.IsExpired() {
		return nil, newExpiredError(errors.New(fmt.Sprintf("SMS confirmation with id %d is expired", tu.ID)))
	}

	if !tu.ValidateAuthInfo(params.AuthCode, params.AuthKey) {
		return nil, newUnauthenticatedError()
	}

	authToken, err := jwt_generator.GenerateJwt(tu.ID)
	if err != nil {
		return nil, err
	}
	return output.NewAuthTokenOutput(authToken), nil
}

// SignUp registers user
func (au *authUsecase) SignUp(ctx *custom_context.Context) (*output.CurrentUserOutput, error) {
	tempUserId := getUserIdFromContext(ctx)
	params := new(input.SignUpInput)
	if err := ctx.BindWithValidation(params); err != nil {
		return nil, err
	}

	tu, err := au.TempUserRepo.FindByID(context.Background(), tempUserId)
	if err != nil {
		return nil, err
	}
	u := model.NewUser(tu.PhoneNumber, params.Email, params.Password, params.FirstName, params.LastName, model.StringToGender(params.Gender), params.DateOfBirth.ToNullTime())
	if err := au.UserRepo.Create(context.Background(), u); err != nil {
		return nil, err
	}

	return output.NewCurrentUserOutput(u), nil
}

// SignIn authenticates user and returns user if authenticated
func (au *authUsecase) SignIn(ctx *custom_context.Context) (*output.CurrentUserOutput, error) {
	params := new(input.SignInInput)
	if err := ctx.BindWithValidation(params); err != nil {
		return nil, err
	}

	normalizedPhoneNumber, _ := libphonenumber.Parse(params.PhoneNumber, params.Region)
	u, err := au.UserRepo.FindByPhoneNumber(context.Background(), libphonenumber.Format(normalizedPhoneNumber, libphonenumber.E164))
	if err != nil {
		return nil, err
	} else if u == nil {
		return nil, newUnauthenticatedError()
	}
	if err := bcrypt.CompareHashAndPassword(u.PasswordDigest, []byte(params.Password)); err != nil {
		return nil, newUnauthenticatedError()
	}

	return output.NewCurrentUserOutput(u), nil
}

// UpdateUnconfirmedPhoneNumber updates user's unconfirmed phone number
func (au *authUsecase) UpdateUnconfirmedPhoneNumber(ctx *custom_context.Context) (*output.SMSReconfirmationOutput, error) {
	userId := getUserIdFromContext(ctx)
	params := new(input.UpdateUnconfirmedPhoneNumberInput)
	if err := ctx.BindWithValidation(params); err != nil {
		return nil, err
	}

	user, err := au.UserRepo.FindByID(context.Background(), uint64(userId))
	if err != nil {
		return nil, err
	} else if user == nil {
		return nil, newNotFoundError("User", "id", userId)
	}

	normalizedPhoneNumber, _ := libphonenumber.Parse(params.PhoneNumber, params.Region)
	phoneNumber := libphonenumber.Format(normalizedPhoneNumber, libphonenumber.E164)
	u, err := au.UserRepo.FindByPhoneNumber(context.Background(), phoneNumber)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	} else if u != nil {
		return nil, newAlreadyTakenError("phoneNumber", params.PhoneNumber)
	}

	user.UnconfirmedPhoneNumber = sql.NullString{String: phoneNumber, Valid: true}
	err = au.UserRepo.Update(context.Background(), user)
	if err != nil {
		return nil, err
	}

	smsReconfirmation, err := au.UserRepo.FindSMSReconfirmationByPhoneNumber(context.Background(), phoneNumber)
	if err != nil {
		return nil, err
	} else if smsReconfirmation == nil {
		smsReconfirmation = model.NewSMSReconfirmation(userId, phoneNumber)
	}

	if err := au.UserRepo.SaveSMSReconfirmation(context.Background(), smsReconfirmation); err != nil {
		return nil, err
	}

	smsBody := fmt.Sprintf("【APP_TITLE】認証コード: %s\nコードは30分間有効です。", smsReconfirmation.AuthCode)
	if err := au.SMSMessenger.SendSMS(smsReconfirmation.PhoneNumber, smsBody); err != nil {
		return nil, err
	}

	return output.NewSMSReconfirmationOutput(smsReconfirmation), nil
}

// ConfirmPhoneNumber confirms unconfirmed phone number and set it to phone number
func (au *authUsecase) ConfirmPhoneNumber(ctx *custom_context.Context) (*output.CurrentUserOutput, error) {
	userId := getUserIdFromContext(ctx)
	params := new(input.ConfirmPhoneNumberInput)
	if err := ctx.BindWithValidation(params); err != nil {
		return nil, err
	}

	user, err := au.UserRepo.FindByID(context.Background(), uint64(userId))
	if err != nil {
		return nil, err
	} else if user == nil {
		return nil, newNotFoundError("User", "id", userId)
	}

	sr, err := au.UserRepo.FindSMSReconfirmationByUser(context.Background(), user)
	if err != nil {
		return nil, err
	} else if sr == nil {
		return nil, newNotFoundError("SMSReconfirmation", "user_id", user.ID)
	}

	if sr.IsExpired() {
		return nil, newExpiredError(errors.New(fmt.Sprintf("SMS reconfirmation with id %d is expired", sr.ID)))
	}
	if sr.AuthCode != params.AuthCode {
		return nil, newUnauthenticatedError()
	}
	if !user.UnconfirmedPhoneNumber.Valid {
		return nil, newBadRequestError(errors.New("unconfirmed phone number is not set"))
	}

	user.PhoneNumber = user.UnconfirmedPhoneNumber.String
	user.UnconfirmedPhoneNumber = sql.NullString{String: "", Valid: false}
	if err := au.UserRepo.Update(context.Background(), user); err != nil {
		return nil, err
	}
	if err := au.UserRepo.DestroySMSReconfirmation(context.Background(), sr); err != nil {
		return nil, err
	}

	return output.NewCurrentUserOutput(user), nil
}
