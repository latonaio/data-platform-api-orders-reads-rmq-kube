package dpfm_api_output_formatter

import (
	"data-platform-api-orders-reads-rmq-kube/DPFM_API_Caller/requests"
	api_input_reader "data-platform-api-orders-reads-rmq-kube/DPFM_API_Input_Reader"
	"database/sql"
	"fmt"
)

func ConvertToHeader(sdc *api_input_reader.SDC, rows *sql.Rows) (*Header, error) {
	pm := &requests.Header{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.OrderID,
			&pm.OrderDate,
			&pm.OrderType,
			&pm.Buyer,
			&pm.Seller,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.ContractType,
			&pm.ValidityStartDate,
			&pm.ValidityEndDate,
			&pm.InvoicePeriodStartDate,
			&pm.InvoicePeriodEndDate,
			&pm.TotalNetAmount,
			&pm.TotalTaxAmount,
			&pm.TotalGrossAmount,
			&pm.OverallDeliveryStatus,
			&pm.TotalBlockStatus,
			&pm.OverallOrdReltdBillgStatus,
			&pm.OverallDocReferenceStatus,
			&pm.TransactionCurrency,
			&pm.PricingDate,
			&pm.PriceDetnExchangeRate,
			&pm.RequestedDeliveryDate,
			&pm.HeaderCompleteDeliveryIsDefined,
			&pm.HeaderBillingBlockReason,
			&pm.DeliveryBlockReason,
			&pm.Incoterms,
			&pm.PaymentTerms,
			&pm.PaymentMethod,
			&pm.ReferenceDocument,
			&pm.ReferenceDocumentItem,
			&pm.BPAccountAssignmentGroup,
			&pm.AccountingExchangeRate,
			&pm.BillingDocumentDate,
			&pm.IsExportImportDelivery,
			&pm.HeaderText,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	header := &Header{
		OrderID:                         data.OrderID,
		OrderDate:                       data.OrderDate,
		OrderType:                       data.OrderType,
		Buyer:                           data.Buyer,
		Seller:                          data.Seller,
		CreationDate:                    data.CreationDate,
		LastChangeDate:                  data.LastChangeDate,
		ContractType:                    data.ContractType,
		ValidityStartDate:               data.ValidityStartDate,
		ValidityEndDate:                 data.ValidityEndDate,
		InvoicePeriodStartDate:          data.InvoicePeriodStartDate,
		InvoicePeriodEndDate:            data.InvoicePeriodEndDate,
		TotalNetAmount:                  data.TotalNetAmount,
		TotalTaxAmount:                  data.TotalTaxAmount,
		TotalGrossAmount:                data.TotalGrossAmount,
		OverallDeliveryStatus:           data.OverallDeliveryStatus,
		TotalBlockStatus:                data.TotalBlockStatus,
		OverallOrdReltdBillgStatus:      data.OverallOrdReltdBillgStatus,
		OverallDocReferenceStatus:       data.OverallDocReferenceStatus,
		TransactionCurrency:             data.TransactionCurrency,
		PricingDate:                     data.PricingDate,
		PriceDetnExchangeRate:           data.PriceDetnExchangeRate,
		RequestedDeliveryDate:           data.RequestedDeliveryDate,
		HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
		HeaderBillingBlockReason:        data.HeaderBillingBlockReason,
		DeliveryBlockReason:             data.DeliveryBlockReason,
		Incoterms:                       data.Incoterms,
		PaymentTerms:                    data.PaymentTerms,
		PaymentMethod:                   data.PaymentMethod,
		ReferenceDocument:               data.ReferenceDocument,
		ReferenceDocumentItem:           data.ReferenceDocumentItem,
		BPAccountAssignmentGroup:        data.BPAccountAssignmentGroup,
		AccountingExchangeRate:          data.AccountingExchangeRate,
		BillingDocumentDate:             data.BillingDocumentDate,
		IsExportImportDelivery:          data.IsExportImportDelivery,
		HeaderText:                      data.HeaderText,
	}
	return header, nil
}

func ConvertToHeaderPartner(sdc *api_input_reader.SDC, rows *sql.Rows) (*HeaderPartner, error) {
	pm := &requests.HeaderPartner{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.OrderID,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
			&pm.BusinessPartnerFullName,
			&pm.BusinessPartnerName,
			&pm.Organization,
			&pm.Country,
			&pm.Language,
			&pm.Currency,
			&pm.ExternalDocumentID,
			&pm.AddressID,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	headerPartner := &HeaderPartner{
		OrderID:                 data.OrderID,
		PartnerFunction:         data.PartnerFunction,
		BusinessPartner:         data.BusinessPartner,
		BusinessPartnerFullName: data.BusinessPartnerFullName,
		BusinessPartnerName:     data.BusinessPartnerName,
		Organization:            data.Organization,
		Country:                 data.Country,
		Language:                data.Language,
		Currency:                data.Currency,
		ExternalDocumentID:      data.ExternalDocumentID,
		AddressID:               data.AddressID,
	}
	return headerPartner, nil
}

func ConvertToItem(sdc *api_input_reader.SDC, rows *sql.Rows) (*Item, error) {
	pm := &requests.Item{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.OrderID,
			&pm.OrderItem,
			&pm.OrderItemCategory,
			&pm.OrderItemText,
			&pm.OrderItemTextByBuyer,
			&pm.OrderItemTextBySeller,
			&pm.Product,
			&pm.ProductStandardID,
			&pm.ProductGroup,
			&pm.BaseUnit,
			&pm.PricingDate,
			&pm.PriceDetnExchangeRate,
			&pm.RequestedDeliveryDate,
			&pm.DeliverFrom,
			&pm.DeliverTo,
			&pm.StockConfirmationPartnerFunction,
			&pm.StockConfirmationBusinessPartner,
			&pm.StockConfirmationPlant,
			&pm.StockConfirmationPlantBatch,
			&pm.StockConfirmationPlantBatchValidityStartDate,
			&pm.StockConfirmationPlantBatchValidityEndDate,
			&pm.ProductIsBatchManagedInStockConfirmationPlant,
			&pm.ServicesRenderingDate,
			&pm.OrderQuantityInBaseUnit,
			&pm.OrderQuantityInIssuingUnit,
			&pm.OrderQuantityInReceivingUnit,
			&pm.OrderIssuingUnit,
			&pm.OrderReceivingUnit,
			&pm.StockConfirmationPolicy,
			&pm.StockConfirmationStatus,
			&pm.ConfirmedOrderQuantityInBaseUnit,
			&pm.ItemWeightUnit,
			&pm.ProductGrossWeight,
			&pm.ItemGrossWeight,
			&pm.ProductNetWeight,
			&pm.ItemNetWeight,
			&pm.NetAmount,
			&pm.TaxAmount,
			&pm.GrossAmount,
			&pm.BillingDocumentDate,
			&pm.ProductionPlantPartnerFunction,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.ProductionPlantTimeZone,
			&pm.ProductionPlantStorageLocation,
			&pm.IssuingPlantPartnerFunction,
			&pm.IssuingPlantBusinessPartner,
			&pm.IssuingPlant,
			&pm.IssuingPlantTimeZone,
			&pm.IssuingPlantStorageLocation,
			&pm.ReceivingPlantPartnerFunction,
			&pm.ReceivingPlantBusinessPartner,
			&pm.ReceivingPlant,
			&pm.ReceivingPlantTimeZone,
			&pm.ReceivingPlantStorageLocation,
			&pm.ProductIsBatchManagedInProductionPlant,
			&pm.ProductIsBatchManagedInIssuingPlant,
			&pm.ProductIsBatchManagedInReceivingPlant,
			&pm.BatchMgmtPolicyInProductionPlant,
			&pm.BatchMgmtPolicyInIssuingPlant,
			&pm.BatchMgmtPolicyInReceivingPlant,
			&pm.ProductionPlantBatch,
			&pm.IssuingPlantBatch,
			&pm.ReceivingPlantBatch,
			&pm.ProductionPlantBatchValidityStartDate,
			&pm.ProductionPlantBatchValidityEndDate,
			&pm.IssuingPlantBatchValidityStartDate,
			&pm.IssuingPlantBatchValidityEndDate,
			&pm.ReceivingPlantBatchValidityStartDate,
			&pm.ReceivingPlantBatchValidityEndDate,
			&pm.Incoterms,
			&pm.BPTaxClassification,
			&pm.ProductTaxClassification,
			&pm.BPAccountAssignmentGroup,
			&pm.ProductAccountAssignmentGroup,
			&pm.PaymentTerms,
			&pm.DueCalculationBaseDate,
			&pm.PaymentDueDate,
			&pm.NetPaymentDays,
			&pm.PaymentMethod,
			&pm.DocumentRjcnReason,
			&pm.ItemBillingBlockReason,
			&pm.Project,
			&pm.AccountingExchangeRate,
			&pm.ReferenceDocument,
			&pm.ReferenceDocumentItem,
			&pm.ItemCompleteDeliveryIsDefined,
			&pm.ItemDeliveryStatus,
			&pm.IssuingStatus,
			&pm.ReceivingStatus,
			&pm.BillingStatus,
			&pm.TaxCode,
			&pm.TaxRate,
			&pm.CountryOfOrigin,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	item := &Item{
		OrderID:                          data.OrderID,
		OrderItem:                        data.OrderItem,
		OrderItemCategory:                data.OrderItemCategory,
		OrderItemText:                    data.OrderItemText,
		OrderItemTextByBuyer:             data.OrderItemTextByBuyer,
		OrderItemTextBySeller:            data.OrderItemTextBySeller,
		Product:                          data.Product,
		ProductStandardID:                data.ProductStandardID,
		ProductGroup:                     data.ProductGroup,
		BaseUnit:                         data.BaseUnit,
		PricingDate:                      data.PricingDate,
		PriceDetnExchangeRate:            data.PriceDetnExchangeRate,
		RequestedDeliveryDate:            data.RequestedDeliveryDate,
		DeliverFrom:                      data.DeliverFrom,
		DeliverTo:                        data.DeliverTo,
		StockConfirmationPartnerFunction: data.StockConfirmationPartnerFunction,
		StockConfirmationBusinessPartner: data.StockConfirmationBusinessPartner,
		StockConfirmationPlant:           data.StockConfirmationPlant,
		StockConfirmationPlantBatch:      data.StockConfirmationPlantBatch,
		StockConfirmationPlantBatchValidityStartDate:  data.StockConfirmationPlantBatchValidityStartDate,
		StockConfirmationPlantBatchValidityEndDate:    data.StockConfirmationPlantBatchValidityEndDate,
		ProductIsBatchManagedInStockConfirmationPlant: data.ProductIsBatchManagedInStockConfirmationPlant,
		ServicesRenderingDate:                         data.ServicesRenderingDate,
		OrderQuantityInBaseUnit:                       data.OrderQuantityInBaseUnit,
		OrderQuantityInIssuingUnit:                    data.OrderQuantityInIssuingUnit,
		OrderQuantityInReceivingUnit:                  data.OrderQuantityInReceivingUnit,
		OrderIssuingUnit:                              data.OrderIssuingUnit,
		OrderReceivingUnit:                            data.OrderReceivingUnit,
		StockConfirmationPolicy:                       data.StockConfirmationPolicy,
		StockConfirmationStatus:                       data.StockConfirmationStatus,
		ConfirmedOrderQuantityInBaseUnit:              data.ConfirmedOrderQuantityInBaseUnit,
		ItemWeightUnit:                                data.ItemWeightUnit,
		ProductGrossWeight:                            data.ProductGrossWeight,
		ItemGrossWeight:                               data.ItemGrossWeight,
		ProductNetWeight:                              data.ProductNetWeight,
		ItemNetWeight:                                 data.ItemNetWeight,
		NetAmount:                                     data.NetAmount,
		TaxAmount:                                     data.TaxAmount,
		GrossAmount:                                   data.GrossAmount,
		BillingDocumentDate:                           data.BillingDocumentDate,
		ProductionPlantPartnerFunction:                data.ProductionPlantPartnerFunction,
		ProductionPlantBusinessPartner:                data.ProductionPlantBusinessPartner,
		ProductionPlant:                               data.ProductionPlant,
		ProductionPlantTimeZone:                       data.ProductionPlantTimeZone,
		ProductionPlantStorageLocation:                data.ProductionPlantStorageLocation,
		IssuingPlantPartnerFunction:                   data.IssuingPlantPartnerFunction,
		IssuingPlantBusinessPartner:                   data.IssuingPlantBusinessPartner,
		IssuingPlant:                                  data.IssuingPlant,
		IssuingPlantTimeZone:                          data.IssuingPlantTimeZone,
		IssuingPlantStorageLocation:                   data.IssuingPlantStorageLocation,
		ReceivingPlantPartnerFunction:                 data.ReceivingPlantPartnerFunction,
		ReceivingPlantBusinessPartner:                 data.ReceivingPlantBusinessPartner,
		ReceivingPlant:                                data.ReceivingPlant,
		ReceivingPlantTimeZone:                        data.ReceivingPlantTimeZone,
		ReceivingPlantStorageLocation:                 data.ReceivingPlantStorageLocation,
		ProductIsBatchManagedInProductionPlant:        data.ProductIsBatchManagedInProductionPlant,
		ProductIsBatchManagedInIssuingPlant:           data.ProductIsBatchManagedInIssuingPlant,
		ProductIsBatchManagedInReceivingPlant:         data.ProductIsBatchManagedInReceivingPlant,
		BatchMgmtPolicyInProductionPlant:              data.BatchMgmtPolicyInProductionPlant,
		BatchMgmtPolicyInIssuingPlant:                 data.BatchMgmtPolicyInIssuingPlant,
		BatchMgmtPolicyInReceivingPlant:               data.BatchMgmtPolicyInReceivingPlant,
		ProductionPlantBatch:                          data.ProductionPlantBatch,
		IssuingPlantBatch:                             data.IssuingPlantBatch,
		ReceivingPlantBatch:                           data.ReceivingPlantBatch,
		ProductionPlantBatchValidityStartDate:         data.ProductionPlantBatchValidityStartDate,
		ProductionPlantBatchValidityEndDate:           data.ProductionPlantBatchValidityEndDate,
		IssuingPlantBatchValidityStartDate:            data.IssuingPlantBatchValidityStartDate,
		IssuingPlantBatchValidityEndDate:              data.IssuingPlantBatchValidityEndDate,
		ReceivingPlantBatchValidityStartDate:          data.ReceivingPlantBatchValidityStartDate,
		ReceivingPlantBatchValidityEndDate:            data.ReceivingPlantBatchValidityEndDate,
		Incoterms:                                     data.Incoterms,
		BPTaxClassification:                           data.BPTaxClassification,
		ProductTaxClassification:                      data.ProductTaxClassification,
		BPAccountAssignmentGroup:                      data.BPAccountAssignmentGroup,
		ProductAccountAssignmentGroup:                 data.ProductAccountAssignmentGroup,
		PaymentTerms:                                  data.PaymentTerms,
		DueCalculationBaseDate:                        data.DueCalculationBaseDate,
		PaymentDueDate:                                data.PaymentDueDate,
		NetPaymentDays:                                data.NetPaymentDays,
		PaymentMethod:                                 data.PaymentMethod,
		DocumentRjcnReason:                            data.DocumentRjcnReason,
		ItemBillingBlockReason:                        data.ItemBillingBlockReason,
		Project:                                       data.Project,
		AccountingExchangeRate:                        data.AccountingExchangeRate,
		ReferenceDocument:                             data.ReferenceDocument,
		ReferenceDocumentItem:                         data.ReferenceDocumentItem,
		ItemCompleteDeliveryIsDefined:                 data.ItemCompleteDeliveryIsDefined,
		ItemDeliveryStatus:                            data.ItemDeliveryStatus,
		IssuingStatus:                                 data.IssuingStatus,
		ReceivingStatus:                               data.ReceivingStatus,
		BillingStatus:                                 data.BillingStatus,
		TaxCode:                                       data.TaxCode,
		TaxRate:                                       data.TaxRate,
		CountryOfOrigin:                               data.CountryOfOrigin,
	}
	return item, nil
}

// func (sdc *SDC) ConvertToBusinessPartner() *requests.BusinessPartner {
// 	dataGeneral := sdc.General
// 	data := sdc.BusinessPartner
// 	return &requests.BusinessPartner{
// 		Product:                dataGeneral.Product,
// 		BusinessPartner:        data.BusinessPartner,
// 		ValidityEndDate:        data.ValidityEndDate,
// 		ValidityStartDate:      data.ValidityStartDate,
// 		BusinessPartnerProduct: data.BusinessPartnerProduct,
// 		IsMarkedForDeletion:    data.IsMarkedForDeletion,
// 	}
// }

// func (sdc *SDC) ConvertToProcurement() *requests.Procurement {
// 	dataGeneral := sdc.General
// 	dataBusinessPartner := sdc.BusinessPartner
// 	data := sdc.Procurement
// 	return &requests.Procurement{
// 		Product:                     dataGeneral.Product,
// 		BusinessPartner:             dataBusinessPartner.BusinessPartner,
// 		Plant:                       data.Plant,
// 		Buyable:                     data.Buyable,
// 		IsAutoPurOrdCreationAllowed: data.IsAutoPurOrdCreationAllowed,
// 		IsSourceListRequired:        data.IsSourceListRequired,
// 		IsMarkedForDeletion:         data.IsMarkedForDeletion,
// 	}
// }

// func (sdc *SDC) ConvertToBPPlant(num int) *requests.BPPlant {
// 	dataGeneral := sdc.General
// 	dataBusinessPartner := sdc.BusinessPartner
// 	data := sdc.BusinessPartner.BPPlant[num]
// 	return &requests.BPPlant{
// 		Product:                   dataGeneral.Product,
// 		BusinessPartner:           dataBusinessPartner.BusinessPartner,
// 		Plant:                     data.Plant,
// 		Issuable:                  data.Issuable,
// 		Receivable:                data.Receivable,
// 		IssuingStorageLocation:    data.IssuingStorageLocation,
// 		ReceivingStorageLocation:  data.ReceivingStorageLocation,
// 		AvailabilityCheckType:     data.AvailabilityCheckType,
// 		ProfitCenter:              data.ProfitCenter,
// 		MRPType:                   data.MRPType,
// 		MRPResponsible:            data.MRPResponsible,
// 		MinimumLotSizeQuantity:    data.MinimumLotSizeQuantity,
// 		MaximumLotSizeQuantity:    data.MaximumLotSizeQuantity,
// 		FixedLotSizeQuantity:      data.FixedLotSizeQuantity,
// 		IsBatchManagementRequired: data.IsBatchManagementRequired,
// 		ProcurementType:           data.ProcurementType,
// 		InventoryUnit:             data.InventoryUnit,
// 		IsMarkedForDeletion:       data.IsMarkedForDeletion,
// 	}
// }

// func (sdc *SDC) ConvertToProductDescription() *requests.ProductDescription {
// 	dataGeneral := sdc.General
// 	data := sdc.ProductDescription
// 	return &requests.ProductDescription{
// 		Product:            dataGeneral.Product,
// 		Language:           data.Language,
// 		ProductDescription: data.ProductDescription,
// 	}
// }
