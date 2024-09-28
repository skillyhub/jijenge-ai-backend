package handler

import (
	"context"
	"errors"

	"github.com/gofrs/uuid"
	pb "github.com/jijengeai/jijengeai/systems/business/pb/gen"
	"github.com/jijengeai/jijengeai/systems/business/pkg/db/models"
	svc "github.com/jijengeai/jijengeai/systems/business/pkg/service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
    pb.UnimplementedBusinessServiceServer
    service *svc.Service
    logger  *logrus.Logger
}

func NewHandler(svc *svc.Service, logger *logrus.Logger) *Handler {
    return &Handler{service: svc, logger: logger}
}

func (h *Handler) CreateBusiness(ctx context.Context, req *pb.CreateBusinessRequest) (*pb.BusinessResponse, error) {
    business, err := h.service.CreateBusiness(req.Name, req.RegId, req.Email, req.Phone, models.BusinessType(req.Type))
    if err != nil {
        h.logger.WithError(err).Error("Failed to create business")
        return nil, status.Errorf(codes.Internal, "Failed to create business: %v", err)
    }
    return &pb.BusinessResponse{Business: convertBusinessToPb(business)}, nil
}

func (h *Handler) GetBusiness(ctx context.Context, req *pb.GetBusinessRequest) (*pb.BusinessResponse, error) {
    id, err := uuid.FromString(req.Id)
    if err != nil {
        return nil, status.Errorf(codes.InvalidArgument, "Invalid business ID: %v", err)
    }

    business, err := h.service.GetBusinessByID(id)
    if err != nil {
        if errors.Is(err, svc.ErrBusinessNotFound) {
            return nil, status.Errorf(codes.NotFound, "Business not found")
        }
        h.logger.WithError(err).Error("Failed to get business")
        return nil, status.Errorf(codes.Internal, "Failed to get business: %v", err)
    }
    return &pb.BusinessResponse{Business: convertBusinessToPb(business)}, nil
}

func (h *Handler) UpdateBusiness(ctx context.Context, req *pb.UpdateBusinessRequest) (*pb.BusinessResponse, error) {
    id, err := uuid.FromString(req.Id)
    if err != nil {
        return nil, status.Errorf(codes.InvalidArgument, "Invalid business ID: %v", err)
    }

    business, err := h.service.UpdateBusiness(id, req.Name, req.Email, req.Phone, models.BusinessType(req.Type))
    if err != nil {
        if errors.Is(err, svc.ErrBusinessNotFound) {
            return nil, status.Errorf(codes.NotFound, "Business not found")
        }
        h.logger.WithError(err).Error("Failed to update business")
        return nil, status.Errorf(codes.Internal, "Failed to update business: %v", err)
    }
    return &pb.BusinessResponse{Business: convertBusinessToPb(business)}, nil
}

func (h *Handler) DeleteBusiness(ctx context.Context, req *pb.DeleteBusinessRequest) (*pb.DeleteBusinessResponse, error) {
    id, err := uuid.FromString(req.Id)
    if err != nil {
        return nil, status.Errorf(codes.InvalidArgument, "Invalid business ID: %v", err)
    }

    err = h.service.DeleteBusiness(id)
    if err != nil {
        if errors.Is(err, svc.ErrBusinessNotFound) {
            return nil, status.Errorf(codes.NotFound, "Business not found")
        }
        h.logger.WithError(err).Error("Failed to delete business")
        return nil, status.Errorf(codes.Internal, "Failed to delete business: %v", err)
    }
    return &pb.DeleteBusinessResponse{Success: true}, nil
}

func (h *Handler) SearchBusinesses(ctx context.Context, req *pb.SearchBusinessesRequest) (*pb.ListBusinessesResponse, error) {
    businesses, err := h.service.SearchBusinesses(req.Query)
    if err != nil {
        h.logger.WithError(err).Error("Failed to search businesses")
        return nil, status.Errorf(codes.Internal, "Failed to search businesses: %v", err)
    }

    pbBusinesses := make([]*pb.Business, len(businesses))
    for i, business := range businesses {
        pbBusinesses[i] = convertBusinessToPb(business)
    }

    return &pb.ListBusinessesResponse{
        Businesses: pbBusinesses,
        Total:      int32(len(businesses)),
    }, nil
}

func convertBusinessToPb(business *models.Business) *pb.Business {
    return &pb.Business{
        Id:    business.Id.String(),
        Name:  business.Name,
        RegId: business.RegID,
        Email: business.Email,
        Phone: business.Phone,
        Type:  modelBusinessTypeToPb(business.Type),
    }
}

func modelBusinessTypeToPb(businessType models.BusinessType) pb.BusinessType {
    switch businessType {
    case models.TypeWholesaler:
        return pb.BusinessType_WHOLESALER
    case models.TypeRetailer:
        return pb.BusinessType_RETAILER
    case models.TypeManufacturer:
        return pb.BusinessType_MANUFACTURER
    case models.TypeServiceProvider:
        return pb.BusinessType_SERVICE_PROVIDER
    case models.TypeDistributor:
        return pb.BusinessType_DISTRIBUTOR
    case models.TypeOnlineStore:
        return pb.BusinessType_ONLINE_STORE
    case models.TypeFranchise:
        return pb.BusinessType_FRANCHISE
    default:
        return pb.BusinessType_WHOLESALER 
    }
}