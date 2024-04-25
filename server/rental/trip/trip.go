package trip

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	Logger *zap.Logger
}

// CreateTrip(context.Context, *CreateTripRequest) (*CreateTripResponse, error)
//
//	func (s *Service) CreateTrip(c context.Context, req *rentalpb.CreateTripRequest) (*rentalpb.CreateTripResponse, error) {
//		aid, err := auth.AccountIdFromContext(c)
//		if err != nil {
//			return nil, err
//		}
//		s.Logger.Info("create trip", zap.String("start", req.Start), zap.String("account id ", aid.String()))
//		return nil, status.Error(codes.Unimplemented, "")
//	}
//
// CreateTrip(context.Context, *CreateTripRequest) (*CreateTripResponse, error)
//
//	GetTrip(context.Context, *GetTripRequest) (*Trip, error)
//	GetTrips(context.Context, *GetTripsRequest) (*GetTripsResponse, error)
//	UpdateTrip(context.Context, *UpdateTripRequest) (*Trip, error)
func (s *Service) CreateTrip(ctx context.Context, in *rentalpb.CreateTripRequest) (*rentalpb.CreateTripResponse, error) {
	return nil, status.Error(codes.Unauthenticated, "")
}

func (s *Service) GetTrip(ctx context.Context, in *rentalpb.GetTripRequest) (*rentalpb.Trip, error) {
	return nil, status.Error(codes.Unauthenticated, "")
}

func (s *Service) GetTrips(ctx context.Context, in *rentalpb.GetTripsRequest) (*rentalpb.GetTripsResponse, error) {
	return nil, status.Error(codes.Unauthenticated, "")
}

func (s *Service) UpdateTrip(ctx context.Context, in *rentalpb.UpdateTripRequest) (*rentalpb.Trip, error) {
	return nil, status.Error(codes.Unauthenticated, "")
}
