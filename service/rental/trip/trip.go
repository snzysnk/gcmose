package trip

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	trippb "project/service/rental/api"
	"project/service/shared/mgutil"
	"project/service/shared/service"
	"time"
)

type TripService struct {
	trippb.UnimplementedTripServiceServer
	MgService Trip
	Lock      Lock
	Identity  Identity
	Help      Help
}

func (s *TripService) CreateTrip(ctx context.Context, request *trippb.CreateTripRequest) (*trippb.CreateTripResponse, error) {
	accountId, err := service.GetContextAccountId(ctx)
	if err != nil {
		return nil, err
	}

	err = s.Identity.Verify(ctx, accountId)

	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "没有创建行程的权限")
	}

	//创建行程前应验证汽车可用.

	tr, err := s.MgService.CreateTrip(accountId, request.CartId, trippb.LocationStatus{
		Location: request.Start,
		Name:     "未知地点",
		Fee:      0,
		Km:       0,
	})

	if err != nil {
		return nil, err
	}

	//开锁
	go func() {
		err := s.Lock.Unlock(context.Background(), request.CartId)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(fmt.Sprintf("unlock cartId:%d successful", request.CartId))
		}
	}()

	return &trippb.CreateTripResponse{
		TripId: tr.ID.Hex(),
		Trip:   tr.Trip,
	}, nil
}

func (s *TripService) UpdateTrip(ctx context.Context, request *trippb.UpdateTripRequest) (*trippb.UpdateTripResponse, error) {
	if request.Current == nil || request.Second == 0 {
		return nil, status.Error(codes.Unknown, "参数错误")
	}

	accountId, err := service.GetContextAccountId(ctx)
	if err != nil {
		return nil, err
	}
	trip, err := s.MgService.FindTrip(accountId)

	if err != nil {
		return nil, status.Error(codes.Unknown, "缺少进行中的行程")
	}

	before := trip.Trip.Current.Location

	positionName := s.Help.Resolve(request.Current)

	fee, km, err := s.Help.Calculate(before, request.Current, trip.Trip.Current.Second, request.Second)

	fmt.Println(fee, km)

	if err != nil {
		return nil, status.Error(codes.Unknown, "计算距离费用失败")
	}

	current := &trippb.LocationStatus{
		Location: request.Current,
		Name:     positionName,
		Fee:      fee,
		Km:       km,
		Second:   request.Second,
	}

	trip.Trip.Current = current

	if request.End {
		trip.Trip.End = current
		trip.Trip.Status = trippb.TripStatus_TRIP_END
	}

	err = s.MgService.UpdateTrip(trip.ID, trip.UpdateAt, trip.Trip)

	if err != nil {
		return nil, err
	}

	return &trippb.UpdateTripResponse{Trip: trip.Trip}, nil
}

type Record struct {
	mgutil.ObjectIdField `bson:"inline"`
	mgutil.UpdateAtField `bson:"inline"`
	Trip                 *trippb.Trip `bson:"trip"`
}

var UpdateAtFunc = func() int64 {
	return time.Now().UnixNano()
}

var StartStatus = trippb.TripStatus_TRIP_ING

var NewObjectIdFunc = primitive.NewObjectID
