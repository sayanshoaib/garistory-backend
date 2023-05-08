package apimodels

import "errors"

var (
	ErrPhoneNumberAlreadyRegistered = errors.New("phone number is already registerd")
	ErrLoginFailed                  = errors.New("phonenumber and password didn't match")
	ErrPinVerification              = errors.New("pin didn't match")
	ErrSomethingWentWrong           = errors.New("something went wrong")
	ErrNoDataMatched                = errors.New("no record found")
	ErrWrongOtp                     = errors.New("OTP did not match")
	ErrNoOtp                        = errors.New("no otp to verify")
	ErrOtpExpired                   = errors.New("OTP expired")
	ErrNoRows                       = errors.New("no match")
	ErrAccountLocked                = errors.New("account locked")
	ErrAccountUnverified            = errors.New("account unverified")
	ErrUserAlreadyVerified          = errors.New("account already verified")
	ErrInvalidRefreshToken          = errors.New("invalid refresh token")
	ErrOtpSent                      = errors.New("OTP sent")
	ErrCantSendOtp                  = errors.New("can not send OTP")
	ErrServiceUnavilable            = errors.New("service unavailable")
	ErrAdminAccess                  = errors.New("access permission denied")
	ErrPaginationRequest            = errors.New("invalid pagination request")
	ErrInvalidPhoneNumber           = errors.New("invalid phone number")
)
