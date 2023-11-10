package grpcGW

import (
	"context"
	pb "github.com/techvify-oliver/protoserver/generated"
	"grpc_server_testing/rabbit"
)

func (s *Server) SendMailToMailService(ctx context.Context, in *pb.EmailRequest) (*pb.EmailResponse, error) {
	isSent := s.Rabbitmq.SendEmail(rabbit.EmailRequest{
		From:    in.From,
		To:      in.To,
		Subject: in.Subject,
		Content: in.Content,
	})
	if !isSent {
		return &pb.EmailResponse{
			Message: "Send mail failed",
			IsSent:  false,
		}, nil
	}
	return &pb.EmailResponse{
		Message: "Send mail successfully",
		IsSent:  true,
	}, nil
}
