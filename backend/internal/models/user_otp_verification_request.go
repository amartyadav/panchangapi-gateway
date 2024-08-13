package models

type UserOtpVerificationRequest struct { // verifyOtp endpoint
	SessionToken string `json:"sessionToken" binding:"required"`
	Otp          string `json:"otp" binding:"required"`
}
