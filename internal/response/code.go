package response

const (
	ErrCodeSuccess = 20000

	ErrCodeInvalidParams = 40000

	// Auth Code
	ErrCodePendingVerification  = 41000
	ErrCodeEmailAlreadyExists   = 41001
	ErrCodeExpiredSession       = 41002
	ErrCodeOtpDoesNotMatch      = 41003
	ErrCodeAuthenticationFailed = 41004
	ErrCodeAccountNotVerified   = 41005
	ErrCodeInternalServer       = 50000
)
