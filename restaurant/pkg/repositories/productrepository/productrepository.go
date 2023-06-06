package productrepository

import (
	"context"
	"fmt"
	"github.com/comp1x/final-task/restaurant/pkg/models"
	_ "github.com/google/uuid"
	restaurant "gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/restaurant"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type ProductService struct {
	restaurant.UnimplementedProductServiceServer

	db *gorm.DB
}

func New(dbURL string) (*ProductService, error) {
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("ошибка при подключении к базе данных: %w", err)
	}

	return &ProductService{
		db: db,
	}, nil
}

func (s *ProductService) CreateProduct(
	ctx context.Context, request *restaurant.CreateProductRequest,
) (*restaurant.CreateProductResponse, error) {
	if err := request.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	product := &models.Product{
		Name:        request.Name,
		Description: request.Description,
		Type:        request.Type,
		Weight:      request.Weight,
		Price:       request.Price,
	}

	if err := s.db.Create(product).Error; err != nil {
		log.Printf("ошибка при создании продукта в базе данных: %v", err)
		return nil, fmt.Errorf("ошибка при создании продукта")
	}

	return &restaurant.CreateProductResponse{}, nil
}

func (s *ProductService) GetProductList(
	ctx context.Context, request *restaurant.GetProductListRequest,
) (*restaurant.GetProductListResponse, error) {
	if err := request.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var products []models.Product
	if err := s.db.Find(&products).Error; err != nil {
		log.Printf("ошибка при получении списка продуктов из базы данных: %v", err)
		return nil, fmt.Errorf("ошибка при получении списка продуктов")
	}

	apiProducts := make([]*restaurant.Product, 0, len(products))
	for _, product := range products {
		apiProduct := &restaurant.Product{
			Uuid:        product.ID.String(),
			Name:        product.Name,
			Description: product.Description,
			Type:        product.Type,
			Weight:      product.Weight,
			Price:       product.Price,
			CreatedAt:   timestamppb.New(product.CreatedAt),
		}
		apiProducts = append(apiProducts, apiProduct)
	}

	return &restaurant.GetProductListResponse{
		Result: apiProducts,
	}, nil
}
