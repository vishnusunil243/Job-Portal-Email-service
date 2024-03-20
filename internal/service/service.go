package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/vishnusunil243/Job-Portal-Email-service/internal/adapters"
	"github.com/vishnusunil243/Job-Portal-Email-service/internal/helper/otp"
	"github.com/vishnusunil243/Job-Portal-proto-files/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	otp.SendOTP(req.Email)
	return &emptypb.Empty{}, nil
}
func (email *EmailService) VerifyOTP(ctx context.Context, req *pb.VerifyOTPRequest) (*pb.VerifyOTPResponse, error) {
	verified := otp.VerifyOTP(req.Email, req.Otp)
	res := &pb.VerifyOTPResponse{
		Verified: verified,
	}
	return res, nil
}
func (email *EmailService) AddNotification(ctx context.Context, req *pb.AddNotificationRequest) (*emptypb.Empty, error) {
	if req.UserId == "" {
		return nil, fmt.Errorf("please provide a valid userID")
	}
	var message primitive.M
	if err := json.Unmarshal([]byte(req.Message), &message); err != nil {
		return nil, fmt.Errorf("failed to parse message JSON: %v", err)
	}
	if err := email.adapter.AddNotification(req.UserId, message); err != nil {
		return nil, err
	}
	return nil, nil
}
func (email *EmailService) GetAllNotifications(req *pb.GetNotificationsByUserId, srv pb.EmailService_GetAllNotificationsServer) error {
	notifications, err := email.adapter.GetAllNotifications(req.UserId)
	if err != nil {
		return err
	}
	for _, notification := range notifications {
		message, ok := notification["message"].(string)
		if !ok {
			return fmt.Errorf("message field is not a string in notification: %v", notification)
		}
		res := &pb.NotificationResponse{
			Message: message,
		}
		if err := srv.Send(res); err != nil {
			return err
		}
	}
	return nil
}
