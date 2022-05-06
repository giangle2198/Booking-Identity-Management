package api

import (
	"booking-identity-management/config"
	"booking-identity-management/internal/common"
	"booking-identity-management/internal/helper"
	"booking-identity-management/internal/usecase"
	"booking-identity-management/pb"
	"context"
	"errors"
	"reflect"

	"google.golang.org/grpc/metadata"
)

type api struct {
	cfg               *config.Config
	pbConverter       helper.PbConverter
	baseUsecase       usecase.BaseUsecase
	userActionUsecase usecase.UserActionUsecase
	userRoleUsecase   usecase.UserRoleUsecase
	roleUsecase       usecase.RoleUsecase
}

func NewAPI(cfg *config.Config,
	pbConverter helper.PbConverter,
	baseUsecase usecase.BaseUsecase,
	userActionUsecase usecase.UserActionUsecase,
	userRoleUsecase usecase.UserRoleUsecase,
	roleUsecase usecase.RoleUsecase) pb.BIMAPIServer {
	return &api{
		cfg:               cfg,
		pbConverter:       pbConverter,
		baseUsecase:       baseUsecase,
		userActionUsecase: userActionUsecase,
		userRoleUsecase:   userRoleUsecase,
		roleUsecase:       roleUsecase,
	}
}

func getFromMD(ctx context.Context, key string) string {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if domain, ok := md[key]; ok && len(domain) != 0 {
			return domain[0]
		}
	}
	return ""
}

func getCreatedByFromMD(ctx context.Context) string {
	return getFromMD(ctx, common.DomainMDKey)
}

func createAndSetToResponse(errString string, in interface{}) {
	err := errors.New(errString)
	setErrorToResponse(err, in)
}

func setErrorToResponse(err error, in interface{}) {
	if err == nil || reflect.TypeOf(in).Kind() != reflect.Ptr {
		return
	}
	reason := common.ParseError(err)
	values := reflect.ValueOf(in)
	vIndirect := reflect.Indirect(values)
	reasonCode := vIndirect.FieldByName("ReasonCode")
	status := vIndirect.FieldByName("StatusCode")
	reasonMessage := vIndirect.FieldByName("ReasonMessage")
	if reasonCode.CanSet() {
		reasonCode.SetString(reason.Code())
	}
	if status.CanSet() {
		status.SetString(common.FAILED)
	}
	if reasonMessage.CanSet() {
		reasonMessage.SetString(reason.Message())
	}
}
