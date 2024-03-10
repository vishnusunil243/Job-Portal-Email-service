package service

import (
	"context"

	helper "github.com/vishnusunil243/Job-Portal-Email-service/internal/Helper"
	"github.com/vishnusunil243/Job-Portal-Email-service/internal/adapters"
	"github.com/vishnusunil243/Job-Portal-proto-files/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type EmailService struct {
	adapter adapters.EmailInterface
	pb.UnimplementedEmailServiceServer
}

func NewEmailService(adapter adapters.EmailInterface) *EmailService {
	return &EmailService{
		adapter: adapter,
	}
}
func (email *EmailService) SendOTP(ctx context.Context, req *pb.SendOtpRequest) (*emptypb.Empty, error) {
	helper.SendOTP(req.Email)
	return &emptypb.Empty{}, nil
}
func (email *EmailService) VerifyOTP(ctx context.Context, req *pb.VerifyOTPRequest) (*pb.VerifyOTPResponse, error) {
	verified := helper.VerifyOTP(req.Email, req.Otp)
	res := &pb.VerifyOTPResponse{
		Verified: verified,
	}
	return res, nil
}
