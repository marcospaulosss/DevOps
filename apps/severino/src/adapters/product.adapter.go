package adapters

import (
	"time"

	"backend/apps/severino/src/structs"
	"backend/libs/util"
	pb "backend/proto"
)

func ToDomainProducts(p []*pb.Product) []structs.Product {
	products := []structs.Product{}
	if p == nil {
		return products
	}
	for _, i := range p {
		products = append(products, ToDomainProduct(i))
	}
	return products
}

func ToDomainProduct(p *pb.Product) structs.Product {
	if p == nil {
		return structs.Product{}
	}
	return structs.Product{
		ID:                p.Id,
		CreatedAt:         util.StrToTime(p.CreatedAt),
		UpdatedAt:         util.StrToTime(p.UpdatedAt),
		StartedSellingAt:  util.StrToTime(p.StartedSellingAt),
		FinishedSellingAt: util.StrToTime(p.FinishedSellingAt),
		UsageExpiresAt:    util.StrToTime(p.UsageExpiresAt),
		Name:              p.Name,
		Description:       p.Description,
		Stock:             p.Stock,
		SKU:               p.Sku,
		Image:             p.Image,
		IsPublished:       p.IsPublished,
		Company:           p.Company,
		ProductType:       p.ProductType,
		Items:             util.StrToJSONArr(p.Items),
		PaymentsTypes:     ToDomainPaymentsTypes(p.PaymentsTypes),
	}
}

func ToProtoProduct(product structs.Product) *pb.Product {
	return &pb.Product{
		Id:                product.ID,
		StartedSellingAt:  product.StartedSellingAt.Format(time.RFC3339),
		FinishedSellingAt: product.FinishedSellingAt.Format(time.RFC3339),
		UsageExpiresAt:    product.UsageExpiresAt.Format(time.RFC3339),
		Name:              product.Name,
		Description:       product.Description,
		Stock:             product.Stock,
		Image:             product.Image,
		IsPublished:       product.IsPublished,
		Company:           product.Company,
		ProductType:       product.ProductType,
		Items:             util.JSONArrToStr(product.Items),
		PaymentsTypes:     ToProtoPaymentsTypes(product.PaymentsTypes),
	}
}
