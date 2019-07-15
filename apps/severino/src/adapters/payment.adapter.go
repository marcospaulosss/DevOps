package adapters

import (
	"backend/apps/severino/src/structs"
	pb "backend/proto"
)

func ToDomainPaymentsTypes(pt []*pb.PaymentType) []structs.PaymentType {
	items := []structs.PaymentType{}
	for _, i := range pt {
		if i != nil {
			paymentType := ToDomainPaymentType(i)
			items = append(items, paymentType)
		}
	}
	return items
}

func ToDomainPaymentType(p *pb.PaymentType) structs.PaymentType {
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

func ToProtoPaymentsTypes(pt []structs.PaymentType) []*pb.PaymentType {
	paymentsTypes := make([]*pb.PaymentType, 0, len(pt))
	for _, i := range pt {
		paymentsTypes = append(paymentsTypes, ToProtoPaymentType(i))
	}
	return paymentsTypes
}

func ToProtoPaymentType(paymentType structs.PaymentType) *pb.PaymentType {
	return &pb.PaymentType{
		Id:           paymentType.ID,
		Name:         paymentType.Name,
		Installments: paymentType.Installments,
		Price:        paymentType.Price,
	}
}
