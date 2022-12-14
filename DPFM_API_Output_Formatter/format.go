package dpfm_api_output_formatter

import (
	"data-platform-api-orders-reads-rmq-kube/DPFM_API_Caller/requests"
	api_input_reader "data-platform-api-orders-reads-rmq-kube/DPFM_API_Input_Reader"
	"database/sql"
	"fmt"
)

func ConvertToHeader(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]Header, error) {
	var header []Header

	for i := 0; true; i++ {
		pm := &requests.Header{}
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_orders_header_data'テーブルに対象のレコードが存在しません。")
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
			&pm.HeaderDeliveryStatus,
			&pm.HeaderBlockStatus,
			&pm.HeaderBillingStatus,
			&pm.HeaderDocReferenceStatus,
			&pm.TransactionCurrency,
			&pm.PricingDate,
			&pm.PriceDetnExchangeRate,
			&pm.RequestedDeliveryDate,
			&pm.HeaderCompleteDeliveryIsDefined,
			&pm.HeaderBillingBlockStatus,
			&pm.HeaderDeliveryBlockStatus,
			&pm.Incoterms,
			&pm.BillFromParty,
			&pm.BillToParty,
			&pm.BillFromCountry,
			&pm.BillToCountry,
			&pm.Payer,
			&pm.Payee,
			&pm.PaymentTerms,
			&pm.PaymentMethod,
			&pm.ReferenceDocument,
			&pm.ReferenceDocumentItem,
			&pm.BPAccountAssignmentGroup,
			&pm.AccountingExchangeRate,
			&pm.InvoiceDocumentDate,
			&pm.IsExportImportDelivery,
			&pm.HeaderText,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		header = append(header, Header{
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
			HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
			HeaderBlockStatus:               data.HeaderBlockStatus,
			HeaderBillingStatus:             data.HeaderBillingStatus,
			HeaderDocReferenceStatus:        data.HeaderDocReferenceStatus,
			TransactionCurrency:             data.TransactionCurrency,
			PricingDate:                     data.PricingDate,
			PriceDetnExchangeRate:           data.PriceDetnExchangeRate,
			RequestedDeliveryDate:           data.RequestedDeliveryDate,
			HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
			HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
			HeaderDeliveryBlockStatus:       data.HeaderDeliveryBlockStatus,
			Incoterms:                       data.Incoterms,
			BillFromParty:                   data.BillFromParty,
			BillToParty:                     data.BillToParty,
			BillFromCountry:                 data.BillFromCountry,
			BillToCountry:                   data.BillToCountry,
			Payer:                           data.Payer,
			Payee:                           data.Payee,
			PaymentTerms:                    data.PaymentTerms,
			PaymentMethod:                   data.PaymentMethod,
			ReferenceDocument:               data.ReferenceDocument,
			ReferenceDocumentItem:           data.ReferenceDocumentItem,
			BPAccountAssignmentGroup:        data.BPAccountAssignmentGroup,
			AccountingExchangeRate:          data.AccountingExchangeRate,
			InvoiceDocumentDate:             data.InvoiceDocumentDate,
			IsExportImportDelivery:          data.IsExportImportDelivery,
			HeaderText:                      data.HeaderText,
		})
	}

	return &header, nil
}

func ConvertToHeaderPartner(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]HeaderPartner, error) {
	var headerPartner []HeaderPartner

	for i := 0; true; i++ {
		pm := &requests.HeaderPartner{}
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_orders_header_partner_data'テーブルに対象のレコードが存在しません。")
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

		data := pm
		headerPartner = append(headerPartner, HeaderPartner{
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
		})
	}

	return &headerPartner, nil
}

func ConvertToHeaderPartnerContact(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]HeaderPartnerContact, error) {
	var headerPartnerContact []HeaderPartnerContact

	for i := 0; true; i++ {
		pm := &requests.HeaderPartnerContact{}
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_orders_header_partner_contact_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.OrderID,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
			&pm.ContactID,
			&pm.ContactPersonName,
			&pm.EmailAddress,
			&pm.PhoneNumber,
			&pm.MobilePhoneNumber,
			&pm.FaxNumber,
			&pm.ContactTag1,
			&pm.ContactTag2,
			&pm.ContactTag3,
			&pm.ContactTag4,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		headerPartnerContact = append(headerPartnerContact, HeaderPartnerContact{
			OrderID:           data.OrderID,
			PartnerFunction:   data.PartnerFunction,
			BusinessPartner:   data.BusinessPartner,
			ContactID:         data.ContactID,
			ContactPersonName: data.ContactPersonName,
			EmailAddress:      data.EmailAddress,
			PhoneNumber:       data.PhoneNumber,
			MobilePhoneNumber: data.MobilePhoneNumber,
			FaxNumber:         data.FaxNumber,
			ContactTag1:       data.ContactTag1,
			ContactTag2:       data.ContactTag2,
			ContactTag3:       data.ContactTag3,
			ContactTag4:       data.ContactTag4,
		})
	}

	return &headerPartnerContact, nil
}

func ConvertToHeaderPartnerPlant(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]HeaderPartnerPlant, error) {
	var headerPartnerPlant []HeaderPartnerPlant

	for i := 0; true; i++ {
		pm := &requests.HeaderPartnerPlant{}
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_orders_header_partner_plant_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.OrderID,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
			&pm.Plant,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm

		headerPartnerPlant = append(headerPartnerPlant, HeaderPartnerPlant{
			OrderID:         data.OrderID,
			PartnerFunction: data.PartnerFunction,
			BusinessPartner: data.BusinessPartner,
			Plant:           data.Plant,
		})
	}

	return &headerPartnerPlant, nil
}

func ConvertToAddress(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]Address, error) {
	var address []Address

	for i := 0; true; i++ {
		pm := &requests.Address{}
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_orders_address_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.OrderID,
			&pm.AddressID,
			&pm.PostalCode,
			&pm.LocalRegion,
			&pm.Country,
			&pm.District,
			&pm.StreetName,
			&pm.CityName,
			&pm.Building,
			&pm.Floor,
			&pm.Room,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		address = append(address, Address{
			OrderID:     data.OrderID,
			AddressID:   data.AddressID,
			PostalCode:  data.PostalCode,
			LocalRegion: data.LocalRegion,
			Country:     data.Country,
			District:    data.District,
			StreetName:  data.StreetName,
			CityName:    data.CityName,
			Building:    data.Building,
			Floor:       data.Floor,
			Room:        data.Room,
		})
	}

	return &address, nil
}

func ConvertToItem(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]Item, error) {
	var item []Item

	for i := 0; true; i++ {
		pm := &requests.Item{}
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_orders_item_data'テーブルに対象のレコードが存在しません。")
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

		data := pm
		item = append(item, Item{
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
		})
	}

	return &item, nil
}

func ConvertToItemPartner(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]ItemPartner, error) {
	var itemPartner []ItemPartner

	for i := 0; true; i++ {
		pm := &requests.ItemPartner{}
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_orders_item_partner_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.OrderID,
			&pm.OrderItem,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		itemPartner = append(itemPartner, ItemPartner{
			OrderID:         data.OrderID,
			OrderItem:       data.OrderItem,
			PartnerFunction: data.PartnerFunction,
			BusinessPartner: data.BusinessPartner,
		})
	}

	return &itemPartner, nil
}

func ConvertToItemPartnerPlant(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]ItemPartnerPlant, error) {
	var itemPartnerPlant []ItemPartnerPlant

	for i := 0; true; i++ {
		pm := &requests.ItemPartnerPlant{}
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_orders_item_partner_plant_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.OrderID,
			&pm.OrderItem,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
			&pm.Plant,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		itemPartnerPlant = append(itemPartnerPlant, ItemPartnerPlant{
			OrderID:         data.OrderID,
			OrderItem:       data.OrderItem,
			PartnerFunction: data.PartnerFunction,
			BusinessPartner: data.BusinessPartner,
			Plant:           data.Plant,
		})
	}

	return &itemPartnerPlant, nil
}

func ConvertToItemPricingElement(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]ItemPricingElement, error) {
	var itemPricingElement []ItemPricingElement

	for i := 0; true; i++ {
		pm := &requests.ItemPricingElement{}
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_orders_item_pricing_element_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.OrderID,
			&pm.OrderItem,
			&pm.PricingProcedureStep,
			&pm.PricingProcedureStep,
			&pm.PricingProcedureCounter,
			&pm.ConditionType,
			&pm.PricingDate,
			&pm.ConditionRateValue,
			&pm.ConditionCurrency,
			&pm.ConditionQuantity,
			&pm.ConditionQuantityUnit,
			&pm.ConditionRecord,
			&pm.ConditionSequentialNumber,
			&pm.TaxCode,
			&pm.ConditionAmount,
			&pm.TransactionCurrency,
			&pm.ConditionIsManuallyChanged,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		itemPricingElement = append(itemPricingElement, ItemPricingElement{
			OrderID:                    data.OrderID,
			OrderItem:                  data.OrderItem,
			PricingProcedureStep:       data.PricingProcedureStep,
			PricingProcedureCounter:    data.PricingProcedureCounter,
			ConditionType:              data.ConditionType,
			PricingDate:                data.PricingDate,
			ConditionRateValue:         data.ConditionRateValue,
			ConditionCurrency:          data.ConditionCurrency,
			ConditionQuantity:          data.ConditionQuantity,
			ConditionQuantityUnit:      data.ConditionQuantityUnit,
			ConditionRecord:            data.ConditionRecord,
			ConditionSequentialNumber:  data.ConditionSequentialNumber,
			TaxCode:                    data.TaxCode,
			ConditionAmount:            data.ConditionAmount,
			TransactionCurrency:        data.TransactionCurrency,
			ConditionIsManuallyChanged: data.ConditionIsManuallyChanged,
		})
	}

	return &itemPricingElement, nil
}

func ConvertToItemSchedulingLine(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]ItemSchedulingLine, error) {
	var itemSchedulingLine []ItemSchedulingLine

	for i := 0; true; i++ {
		pm := &requests.ItemSchedulingLine{}
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_orders_item_schedule_line_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.OrderID,
			&pm.OrderItem,
			&pm.ScheduleLine,
			&pm.Product,
			&pm.StockConfirmationPartnerFunction,
			&pm.StockConfirmationBusinessPartner,
			&pm.StockConfirmationPlant,
			&pm.StockConfirmationPlantBatch,
			&pm.StockConfirmationPlantBatchValidityStartDate,
			&pm.StockConfirmationPlantBatchValidityEndDate,
			&pm.ConfirmedDeliveryDate,
			&pm.RequestedDeliveryDate,
			&pm.OrderQuantityInBaseUnit,
			&pm.ConfdOrderQtyByPDTAvailCheck,
			&pm.DeliveredQtyInOrderQtyUnit,
			&pm.OpenConfdDelivQtyInOrdQtyUnit,
			&pm.DelivBlockReasonForSchedLine,
			&pm.PlusMinusFlag,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		itemSchedulingLine = append(itemSchedulingLine, ItemSchedulingLine{
			OrderID:                          data.OrderID,
			OrderItem:                        data.OrderItem,
			ScheduleLine:                     data.ScheduleLine,
			Product:                          data.Product,
			StockConfirmationPartnerFunction: data.StockConfirmationPartnerFunction,
			StockConfirmationBusinessPartner: data.StockConfirmationBusinessPartner,
			StockConfirmationPlant:           data.StockConfirmationPlant,
			StockConfirmationPlantBatch:      data.StockConfirmationPlantBatch,
			StockConfirmationPlantBatchValidityStartDate: data.StockConfirmationPlantBatchValidityStartDate,
			StockConfirmationPlantBatchValidityEndDate:   data.StockConfirmationPlantBatchValidityEndDate,
			ConfirmedDeliveryDate:                        data.ConfirmedDeliveryDate,
			RequestedDeliveryDate:                        data.RequestedDeliveryDate,
			OrderQuantityInBaseUnit:                      data.OrderQuantityInBaseUnit,
			ConfdOrderQtyByPDTAvailCheck:                 data.ConfdOrderQtyByPDTAvailCheck,
			DeliveredQtyInOrderQtyUnit:                   data.DeliveredQtyInOrderQtyUnit,
			OpenConfdDelivQtyInOrdQtyUnit:                data.OpenConfdDelivQtyInOrdQtyUnit,
			DelivBlockReasonForSchedLine:                 data.DelivBlockReasonForSchedLine,
			PlusMinusFlag:                                data.PlusMinusFlag,
		})
	}

	return &itemSchedulingLine, nil
}
