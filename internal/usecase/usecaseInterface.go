package usecase

type EmailUsecases interface {
	SendOtp(email string) error
}
