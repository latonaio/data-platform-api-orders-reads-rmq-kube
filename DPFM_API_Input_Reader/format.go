package dpfm_api_input_reader

// func (sdc *SDC) ConvertToHeader() *requests.Header {
// 	data := sdc.Header
// 	return &requests.Header{
// 		OrderID:                         data.OrderID,
// 		OrderDate:                       data.OrderDate,
// 		OrderType:                       data.OrderType,
// 		Buyer:                           data.Buyer,
// 		Seller:                          data.Seller,
// 		CreationDate:                    data.CreationDate,
// 		LastChangeDate:                  data.LastChangeDate,
// 		ContractType:                    data.ContractType,
// 		ValidityStartDate:               data.ValidityStartDate,
// 		ValidityEndDate:                 data.ValidityEndDate,
// 		InvoicePeriodStartDate:          data.InvoicePeriodStartDate,
// 		InvoicePeriodEndDate:            data.InvoicePeriodEndDate,
// 		TotalNetAmount:                  data.TotalNetAmount,
// 		TotalTaxAmount:                  data.TotalTaxAmount,
// 		TotalGrossAmount:                data.TotalGrossAmount,
// 		HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
// 		HeaderBlockStatus:               data.HeaderBlockStatus,
// 		HeaderBillingStatus:             data.HeaderBillingStatus,
// 		HeaderDocReferenceStatus:        data.HeaderDocReferenceStatus,
// 		TransactionCurrency:             data.TransactionCurrency,
// 		PricingDate:                     data.PricingDate,
// 		PriceDetnExchangeRate:           data.PriceDetnExchangeRate,
// 		RequestedDeliveryDate:           data.RequestedDeliveryDate,
// 		HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
// 		HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
// 		HeaderDeliveryBlockStatus:       data.HeaderDeliveryBlockStatus,
// 		Incoterms:                       data.Incoterms,
// 		BillFromParty:                   data.BillFromParty,
// 		BillToParty:                     data.BillToParty,
// 		BillFromCountry:                 data.BillFromCountry,
// 		BillToCountry:                   data.BillToCountry,
// 		Payer:                           data.Payer,
// 		Payee:                           data.Payee,
// 		PaymentTerms:                    data.PaymentTerms,
// 		PaymentMethod:                   data.PaymentMethod,
// 		ReferenceDocument:               data.ReferenceDocument,
// 		ReferenceDocumentItem:           data.ReferenceDocumentItem,
// 		BPAccountAssignmentGroup:        data.BPAccountAssignmentGroup,
// 		AccountingExchangeRate:          data.AccountingExchangeRate,
// 		InvoiceDocumentDate:             data.InvoiceDocumentDate,
// 		IsExportImportDelivery:          data.IsExportImportDelivery,
// 		HeaderText:                      data.HeaderText,
// 	}
// }

// func (sdc *SDC) ConvertToHeaderPartner(num int) *requests.HeaderPartner {
// 	dataHeader := sdc.Header
// 	data := sdc.Header.HeaderPartner[num]
// 	return &requests.HeaderPartner{
// 		OrderID:                 dataHeader.OrderID,
// 		PartnerFunction:         data.PartnerFunction,
// 		BusinessPartner:         data.BusinessPartner,
// 		BusinessPartnerFullName: data.BusinessPartnerFullName,
// 		BusinessPartnerName:     data.BusinessPartnerName,
// 		Organization:            data.Organization,
// 		Country:                 data.Country,
// 		Language:                data.Language,
// 		Currency:                data.Currency,
// 		ExternalDocumentID:      data.ExternalDocumentID,
// 		AddressID:               data.AddressID,
// 	}
// }

// func (sdc *SDC) ConvertToHeaderPartnerPlant(hpNum, hppNum int) *requests.HeaderPartnerPlant {
// 	dataHeader := sdc.Header
// 	dataHeaderPartner := sdc.Header.HeaderPartner[hpNum]
// 	data := dataHeaderPartner.HeaderPartnerPlant[hppNum]
// 	return &requests.HeaderPartnerPlant{
// 		OrderID:         dataHeader.OrderID,
// 		PartnerFunction: dataHeaderPartner.PartnerFunction,
// 		BusinessPartner: dataHeaderPartner.BusinessPartner,
// 		Plant:           data.Plant,
// 	}
// }

// func (sdc *SDC) ConvertToHeaderPartnerContact(hpNum, hpcNum int) *requests.HeaderPartnerContact {
// 	dataHeader := sdc.Header
// 	dataHeaderPartner := sdc.Header.HeaderPartner[hpNum]
// 	data := dataHeaderPartner.HeaderPartnerContact[hpcNum]
// 	return &requests.HeaderPartnerContact{
// 		OrderID:           dataHeader.OrderID,
// 		PartnerFunction:   dataHeaderPartner.PartnerFunction,
// 		BusinessPartner:   dataHeaderPartner.BusinessPartner,
// 		ContactID:         data.ContactID,
// 		ContactPersonName: data.ContactPersonName,
// 		EmailAddress:      data.EmailAddress,
// 		PhoneNumber:       data.PhoneNumber,
// 		MobilePhoneNumber: data.MobilePhoneNumber,
// 		FaxNumber:         data.FaxNumber,
// 		ContactTag1:       data.ContactTag1,
// 		ContactTag2:       data.ContactTag2,
// 		ContactTag3:       data.ContactTag3,
// 		ContactTag4:       data.ContactTag4,
// 	}
// }

// func (sdc *SDC) ConvertToAddress(num int) *requests.Address {
// 	dataHeader := sdc.Header
// 	data := sdc.Header.Address[num]
// 	return &requests.Address{
// 		OrderID:     dataHeader.OrderID,
// 		AddressID:   data.AddressID,
// 		PostalCode:  data.PostalCode,
// 		LocalRegion: data.LocalRegion,
// 		Country:     data.Country,
// 		District:    data.District,
// 		StreetName:  data.StreetName,
// 		CityName:    data.CityName,
// 		Building:    data.Building,
// 		Floor:       data.Floor,
// 		Room:        data.Room,
// 	}
// }

// func (sdc *SDC) ConvertToItem(num int) *requests.Item {
// 	dataHeader := sdc.Header
// 	data := sdc.Header.Item[num]
// 	return &requests.Item{
// 		OrderID:                          dataHeader.OrderID,
// 		OrderItem:                        data.OrderItem,
// 		OrderItemCategory:                data.OrderItemCategory,
// 		OrderItemText:                    data.OrderItemText,
// 		OrderItemTextByBuyer:             data.OrderItemTextByBuyer,
// 		OrderItemTextBySeller:            data.OrderItemTextBySeller,
// 		Product:                          data.Product,
// 		ProductStandardID:                data.ProductStandardID,
// 		ProductGroup:                     data.ProductGroup,
// 		BaseUnit:                         data.BaseUnit,
// 		PricingDate:                      data.PricingDate,
// 		PriceDetnExchangeRate:            data.PriceDetnExchangeRate,
// 		RequestedDeliveryDate:            data.RequestedDeliveryDate,
// 		DeliverFrom:                      data.DeliverFrom,
// 		DeliverTo:                        data.DeliverTo,
// 		StockConfirmationPartnerFunction: data.StockConfirmationPartnerFunction,
// 		StockConfirmationBusinessPartner: data.StockConfirmationBusinessPartner,
// 		StockConfirmationPlant:           data.StockConfirmationPlant,
// 		StockConfirmationPlantBatch:      data.StockConfirmationPlantBatch,
// 		StockConfirmationPlantBatchValidityStartDate:  data.StockConfirmationPlantBatchValidityStartDate,
// 		StockConfirmationPlantBatchValidityEndDate:    data.StockConfirmationPlantBatchValidityEndDate,
// 		ProductIsBatchManagedInStockConfirmationPlant: data.ProductIsBatchManagedInStockConfirmationPlant,
// 		ServicesRenderingDate:                         data.ServicesRenderingDate,
// 		OrderQuantityInBaseUnit:                       data.OrderQuantityInBaseUnit,
// 		OrderQuantityInIssuingUnit:                    data.OrderQuantityInIssuingUnit,
// 		OrderQuantityInReceivingUnit:                  data.OrderQuantityInReceivingUnit,
// 		OrderIssuingUnit:                              data.OrderIssuingUnit,
// 		OrderReceivingUnit:                            data.OrderReceivingUnit,
// 		StockConfirmationPolicy:                       data.StockConfirmationPolicy,
// 		StockConfirmationStatus:                       data.StockConfirmationStatus,
// 		ConfirmedOrderQuantityInBaseUnit:              data.ConfirmedOrderQuantityInBaseUnit,
// 		ItemWeightUnit:                                data.ItemWeightUnit,
// 		ProductGrossWeight:                            data.ProductGrossWeight,
// 		ItemGrossWeight:                               data.ItemGrossWeight,
// 		ProductNetWeight:                              data.ProductNetWeight,
// 		ItemNetWeight:                                 data.ItemNetWeight,
// 		NetAmount:                                     data.NetAmount,
// 		TaxAmount:                                     data.TaxAmount,
// 		GrossAmount:                                   data.GrossAmount,
// 		BillingDocumentDate:                           data.BillingDocumentDate,
// 		ProductionPlantPartnerFunction:                data.ProductionPlantPartnerFunction,
// 		ProductionPlantBusinessPartner:                data.ProductionPlantBusinessPartner,
// 		ProductionPlant:                               data.ProductionPlant,
// 		ProductionPlantTimeZone:                       data.ProductionPlantTimeZone,
// 		ProductionPlantStorageLocation:                data.ProductionPlantStorageLocation,
// 		IssuingPlantPartnerFunction:                   data.IssuingPlantPartnerFunction,
// 		IssuingPlantBusinessPartner:                   data.IssuingPlantBusinessPartner,
// 		IssuingPlant:                                  data.IssuingPlant,
// 		IssuingPlantTimeZone:                          data.IssuingPlantTimeZone,
// 		IssuingPlantStorageLocation:                   data.IssuingPlantStorageLocation,
// 		ReceivingPlantPartnerFunction:                 data.ReceivingPlantPartnerFunction,
// 		ReceivingPlantBusinessPartner:                 data.ReceivingPlantBusinessPartner,
// 		ReceivingPlant:                                data.ReceivingPlant,
// 		ReceivingPlantTimeZone:                        data.ReceivingPlantTimeZone,
// 		ReceivingPlantStorageLocation:                 data.ReceivingPlantStorageLocation,
// 		ProductIsBatchManagedInProductionPlant:        data.ProductIsBatchManagedInProductionPlant,
// 		ProductIsBatchManagedInIssuingPlant:           data.ProductIsBatchManagedInIssuingPlant,
// 		ProductIsBatchManagedInReceivingPlant:         data.ProductIsBatchManagedInReceivingPlant,
// 		BatchMgmtPolicyInProductionPlant:              data.BatchMgmtPolicyInProductionPlant,
// 		BatchMgmtPolicyInIssuingPlant:                 data.BatchMgmtPolicyInIssuingPlant,
// 		BatchMgmtPolicyInReceivingPlant:               data.BatchMgmtPolicyInReceivingPlant,
// 		ProductionPlantBatch:                          data.ProductionPlantBatch,
// 		IssuingPlantBatch:                             data.IssuingPlantBatch,
// 		ReceivingPlantBatch:                           data.ReceivingPlantBatch,
// 		ProductionPlantBatchValidityStartDate:         data.ProductionPlantBatchValidityStartDate,
// 		ProductionPlantBatchValidityEndDate:           data.ProductionPlantBatchValidityEndDate,
// 		IssuingPlantBatchValidityStartDate:            data.IssuingPlantBatchValidityStartDate,
// 		IssuingPlantBatchValidityEndDate:              data.IssuingPlantBatchValidityEndDate,
// 		ReceivingPlantBatchValidityStartDate:          data.ReceivingPlantBatchValidityStartDate,
// 		ReceivingPlantBatchValidityEndDate:            data.ReceivingPlantBatchValidityEndDate,
// 		Incoterms:                                     data.Incoterms,
// 		BPTaxClassification:                           data.BPTaxClassification,
// 		ProductTaxClassification:                      data.ProductTaxClassification,
// 		BPAccountAssignmentGroup:                      data.BPAccountAssignmentGroup,
// 		ProductAccountAssignmentGroup:                 data.ProductAccountAssignmentGroup,
// 		PaymentTerms:                                  data.PaymentTerms,
// 		DueCalculationBaseDate:                        data.DueCalculationBaseDate,
// 		PaymentDueDate:                                data.PaymentDueDate,
// 		NetPaymentDays:                                data.NetPaymentDays,
// 		PaymentMethod:                                 data.PaymentMethod,
// 		DocumentRjcnReason:                            data.DocumentRjcnReason,
// 		ItemBillingBlockReason:                        data.ItemBillingBlockReason,
// 		Project:                                       data.Project,
// 		AccountingExchangeRate:                        data.AccountingExchangeRate,
// 		ReferenceDocument:                             data.ReferenceDocument,
// 		ReferenceDocumentItem:                         data.ReferenceDocumentItem,
// 		ItemCompleteDeliveryIsDefined:                 data.ItemCompleteDeliveryIsDefined,
// 		ItemDeliveryStatus:                            data.ItemDeliveryStatus,
// 		IssuingStatus:                                 data.IssuingStatus,
// 		ReceivingStatus:                               data.ReceivingStatus,
// 		BillingStatus:                                 data.BillingStatus,
// 		TaxCode:                                       data.TaxCode,
// 		TaxRate:                                       data.TaxRate,
// 		CountryOfOrigin:                               data.CountryOfOrigin,
// 	}
// }

// func (sdc *SDC) ConvertToItemPartner(iNum, ipNum int) *requests.ItemPartner {
// 	dataHeader := sdc.Header
// 	dataItem := sdc.Header.Item[iNum]
// 	data := dataItem.ItemPartner[ipNum]
// 	return &requests.ItemPartner{
// 		OrderID:         dataHeader.OrderID,
// 		OrderItem:       dataItem.OrderItem,
// 		PartnerFunction: data.PartnerFunction,
// 		BusinessPartner: data.BusinessPartner,
// 	}
// }

// func (sdc *SDC) ConvertToItemPricingElement(iNum, ipeNum int) *requests.ItemPricingElement {
// 	dataHeader := sdc.Header
// 	dataItem := sdc.Header.Item[iNum]
// 	data := dataItem.ItemPricingElement[ipeNum]
// 	return &requests.ItemPricingElement{
// 		OrderID:                    dataHeader.OrderID,
// 		OrderItem:                  dataItem.OrderItem,
// 		PricingProcedureStep:       data.PricingProcedureStep,
// 		PricingProcedureCounter:    data.PricingProcedureCounter,
// 		ConditionType:              data.ConditionType,
// 		PricingDate:                data.PricingDate,
// 		ConditionRateValue:         data.ConditionRateValue,
// 		ConditionCurrency:          data.ConditionCurrency,
// 		ConditionQuantity:          data.ConditionQuantity,
// 		ConditionQuantityUnit:      data.ConditionQuantityUnit,
// 		ConditionRecord:            data.ConditionRecord,
// 		ConditionSequentialNumber:  data.ConditionSequentialNumber,
// 		TaxCode:                    data.TaxCode,
// 		ConditionAmount:            data.ConditionAmount,
// 		TransactionCurrency:        data.TransactionCurrency,
// 		ConditionIsManuallyChanged: data.ConditionIsManuallyChanged,
// 	}
// }

// func (sdc *SDC) ConvertToItemSchedulingLine(iNum, islNum int) *requests.ItemSchedulingLine {
// 	dataHeader := sdc.Header
// 	dataItem := sdc.Header.Item[iNum]
// 	data := dataItem.ItemSchedulingLine[islNum]
// 	return &requests.ItemSchedulingLine{
// 		OrderID:                          dataHeader.OrderID,
// 		OrderItem:                        dataItem.OrderItem,
// 		ScheduleLine:                     data.ScheduleLine,
// 		Product:                          data.Product,
// 		StockConfirmationPartnerFunction: data.StockConfirmationPartnerFunction,
// 		StockConfirmationBusinessPartner: data.StockConfirmationBusinessPartner,
// 		StockConfirmationPlant:           data.StockConfirmationPlant,
// 		StockConfirmationPlantBatch:      data.StockConfirmationPlantBatch,
// 		StockConfirmationPlantBatchValidityStartDate: data.StockConfirmationPlantBatchValidityStartDate,
// 		StockConfirmationPlantBatchValidityEndDate:   data.StockConfirmationPlantBatchValidityEndDate,
// 		ConfirmedDeliveryDate:                        data.ConfirmedDeliveryDate,
// 		RequestedDeliveryDate:                        data.RequestedDeliveryDate,
// 		OrderQuantityInBaseUnit:                      data.OrderQuantityInBaseUnit,
// 		ConfdOrderQtyByPDTAvailCheck:                 data.ConfdOrderQtyByPDTAvailCheck,
// 		DeliveredQtyInOrderQtyUnit:                   data.DeliveredQtyInOrderQtyUnit,
// 		OpenConfdDelivQtyInOrdQtyUnit:                data.OpenConfdDelivQtyInOrdQtyUnit,
// 		DelivBlockReasonForSchedLine:                 data.DelivBlockReasonForSchedLine,
// 		PlusMinusFlag:                                data.PlusMinusFlag,
// 	}
// }
