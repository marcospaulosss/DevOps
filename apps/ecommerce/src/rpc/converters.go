package rpc

import (
	"time"

	"backend/apps/ecommerce/src/structs"
	"backend/libs/json"
	"backend/libs/util"
	pb "backend/proto"
)

func ToPbProduct(product structs.Product) *pb.Product {
	return &pb.Product{
		Id:                product.ID,
		CreatedAt:         product.CreatedAt.Format(time.RFC3339),
		UpdatedAt:         product.UpdatedAt.Format(time.RFC3339),
		StartedSellingAt:  product.StartedSellingAt.Format(time.RFC3339),
		FinishedSellingAt: product.FinishedSellingAt.Format(time.RFC3339),
		UsageExpiresAt:    product.UsageExpiresAt.Format(time.RFC3339),
		Name:              product.Name,
		Description:       product.Description,
		Stock:             product.Stock,
		Sku:               product.SKU,
		Image:             product.Image,
		IsPublished:       product.IsPublished,
		Company:           product.Company,
		ProductType:       product.ProductType,
		Items:             util.PtrToStr(product.Items),
		PaymentsTypes:     ToPbPaymentsTypes(product.PaymentsTypes),
		History:           util.PtrToStr(product.History),
	}
}

func ToPbProducts(pds []structs.Product) []*pb.Product {
	products := make([]*pb.Product, 0, len(pds))
	for _, i := range pds {
		products = append(products, ToPbProduct(i))
	}
	return products
}

func FromPbSearch(s *pb.Search) structs.Search {
	if s == nil {
		return structs.Search{}
	}
	return structs.Search{
		Pagination: FromPbPagination(s.GetPagination()),
	}
}

func FromPbPagination(p *pb.Pagination) structs.Pagination {
	if p == nil {
		return structs.Pagination{}
	}
	return structs.Pagination{
		Page:    p.GetPage(),
		PerPage: p.GetPerPage(),
		Order:   p.GetOrder(),
		SortBy:  p.GetSortBy(),
	}
}

func ToPbPagination(p structs.Pagination) *pb.Pagination {
	return &pb.Pagination{
		Page:    p.Page,
		PerPage: p.PerPage,
		Order:   p.Order,
		SortBy:  p.SortBy,
		Total:   p.Total,
	}
}

func FromPbProduct(product *pb.Product) structs.Product {
	if product == nil {
		return structs.Product{}
	}
	return structs.Product{
		ID:                product.Id,
		StartedSellingAt:  util.StrToTime(product.StartedSellingAt),
		FinishedSellingAt: util.StrToTime(product.FinishedSellingAt),
		UsageExpiresAt:    util.StrToTime(product.UsageExpiresAt),
		Name:              product.Name,
		Description:       product.Description,
		Stock:             product.Stock,
		Image:             product.Image,
		IsPublished:       product.IsPublished,
		Company:           product.Company,
		ProductType:       product.ProductType,
		Items:             util.StrToPtr(product.Items),
		PaymentsTypes:     FromPbPaymentsTypes(product.PaymentsTypes),
	}
}

func ToPbPaymentsTypes(pt *string) []*pb.PaymentType {
	var paymentsTypes []structs.PaymentType
	if pt == nil {
		return []*pb.PaymentType{}
	}
	err := json.Unmarshal([]byte(*pt), &paymentsTypes)
	if err != nil {
		return []*pb.PaymentType{}
	}
	list := make([]*pb.PaymentType, 0, len(paymentsTypes))
	for _, i := range paymentsTypes {
		p := ToPbPaymentType(i)
		list = append(list, p)
	}
	return list
}

func FromPbPaymentsTypes(pt []*pb.PaymentType) *string {
	items := make([]structs.PaymentType, 0, len(pt))
	for _, i := range pt {
		if i != nil {
			paymentType := FromPbPaymentType(i)
			items = append(items, paymentType)
		}
	}
	encoded, err := json.Marshal(items)
	result := "[]"
	if err != nil {
		return &result
	}
	result = string(encoded)
	return &result
}

func FromPbPaymentType(p *pb.PaymentType) structs.PaymentType {
	if p == nil {
		return structs.PaymentType{}
	}
	return structs.PaymentType{
		ID:           p.GetId(),
		Name:         p.GetName(),
		Price:        p.GetPrice(),
		Installments: p.GetInstallments(),
	}
}

func ToPbPaymentType(paymentType structs.PaymentType) *pb.PaymentType {
	return &pb.PaymentType{
		Id:           paymentType.ID,
		Name:         paymentType.Name,
		Installments: paymentType.Installments,
		Price:        paymentType.Price,
	}
}
