package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-orders-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-orders-reads-rmq-kube/DPFM_API_Output_Formatter"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *dpfm_api_output_formatter.Header
	var headerPartner *dpfm_api_output_formatter.HeaderPartner
	var headerPartnerContact *dpfm_api_output_formatter.HeaderPartnerContact
	var headerPartnerPlant *dpfm_api_output_formatter.HeaderPartnerPlant
	var address *dpfm_api_output_formatter.Address
	var item *dpfm_api_output_formatter.Item
	var itemPartner *dpfm_api_output_formatter.ItemPartner
	var itemPartnerPlant *dpfm_api_output_formatter.ItemPartnerPlant
	var itemPricingElement *dpfm_api_output_formatter.ItemPricingElement
	var itemSchedulingLine *dpfm_api_output_formatter.ItemSchedulingLine
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				header = c.Header(mtx, input, output, errs, log)
			}()
		case "HeaderPartner":
			func() {
				headerPartner = c.HeaderPartner(mtx, input, output, errs, log)
			}()
		case "HeaderPartnerContact":
			func() {
				headerPartnerContact = c.HeaderPartnerContact(mtx, input, output, errs, log)
			}()
		case "HeaderPartnerPlant":
			func() {
				headerPartnerPlant = c.HeaderPartnerPlant(mtx, input, output, errs, log)
			}()
		case "Address":
			func() {
				address = c.Address(mtx, input, output, errs, log)
			}()
		case "Item":
			func() {
				item = c.Item(mtx, input, output, errs, log)
			}()
		case "ItemPartner":
			func() {
				itemPartner = c.ItemPartner(mtx, input, output, errs, log)
			}()
		case "ItemPartnerPlant":
			func() {
				itemPartnerPlant = c.ItemPartnerPlant(mtx, input, output, errs, log)
			}()
		case "ItemPricingElement":
			func() {
				itemPricingElement = c.ItemPricingElement(mtx, input, output, errs, log)
			}()
		case "ItemSchedulingLine":
			func() {
				itemSchedulingLine = c.ItemSchedulingLine(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:               header,
		HeaderPartner:        headerPartner,
		HeaderPartnerContact: headerPartnerContact,
		HeaderPartnerPlant:   headerPartnerPlant,
		Address:              address,
		Item:                 item,
		ItemPartner:          itemPartner,
		ItemPartnerPlant:     itemPartnerPlant,
		ItemPricingElement:   itemPricingElement,
		ItemSchedulingLine:   itemSchedulingLine,
	}

	return data
}

func (c *DPFMAPICaller) Header(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Header {
	order := input.Header.OrderID

	rows, err := c.db.Query(
		`SELECT OrderID, OrderDate, OrderType, Buyer, Seller, CreationDate, LastChangeDate, 
		ContractType, ValidityStartDate, ValidityEndDate, InvoicePeriodStartDate, InvoicePeriodEndDate, 
		TotalNetAmount, TotalTaxAmount, TotalGrossAmount, OverallDeliveryStatus, TotalBlockStatus, 
		OverallOrdReltdBillgStatus, OverallDocReferenceStatus, TransactionCurrency, PricingDate, 
		PriceDetnExchangeRate, RequestedDeliveryDate, HeaderCompleteDeliveryIsDefined, 
		HeaderBillingBlockReason, DeliveryBlockReason, Incoterms, PaymentTerms, PaymentMethod, 
		ReferenceDocument, ReferenceDocumentItem, BPAccountAssignmentGroup, AccountingExchangeRate, 
		BillingDocumentDate, IsExportImportDelivery, HeaderText
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_data
		WHERE OrderID = ?;`, order,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToHeader(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeaderPartner(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.HeaderPartner {
	order := input.Header.OrderID
	partnerFunction := input.Header.HeaderPartner.PartnerFunction
	businessPartner := input.Header.HeaderPartner.BusinessPartner

	rows, err := c.db.Query(
		`SELECT OrderID, PartnerFunction, BusinessPartner, BusinessPartnerFullName, BusinessPartnerName, 
		Organization, Country, Language, Currency, ExternalDocumentID, AddressID
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_partner_data
		WHERE (OrderID, PartnerFunction, BusinessPartner) = (?, ?, ?);`, order, partnerFunction, businessPartner,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToHeaderPartner(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeaderPartnerContact(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.HeaderPartnerContact {
	order := input.Header.OrderID
	partnerFunction := input.Header.HeaderPartner.PartnerFunction
	businessPartner := input.Header.HeaderPartner.BusinessPartner
	contactID := input.Header.HeaderPartner.HeaderPartnerContact.ContactID

	rows, err := c.db.Query(
		`SELECT OrderID, PartnerFunction, BusinessPartner, ContactID, ContactPersonName, EmailAddress, PhoneNumber, 
		MobilePhoneNumber, FaxNumber, ContactTag1, ContactTag2, ContactTag3, ContactTag4
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_partner_contact_data
		WHERE (OrderID, PartnerFunction, BusinessPartner, ContactID) = (?, ?, ?, ?);`, order, partnerFunction, businessPartner, contactID,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToHeaderPartnerContact(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeaderPartnerPlant(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.HeaderPartnerPlant {
	order := input.Header.OrderID
	partnerFunction := input.Header.HeaderPartner.PartnerFunction
	businessPartner := input.Header.HeaderPartner.BusinessPartner
	plant := input.Header.HeaderPartner.HeaderPartnerPlant.Plant

	rows, err := c.db.Query(
		`SELECT OrderID, PartnerFunction, BusinessPartner, Plant
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_partner_plant_data
		WHERE (OrderID, PartnerFunction, BusinessPartner, Plant) = (?, ?, ?, ?);`, order, partnerFunction, businessPartner, plant,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToHeaderPartnerPlant(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Address(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Address {
	order := input.Header.OrderID
	addressID := input.Header.Address.AddressID

	rows, err := c.db.Query(
		`SELECT OrderID, AddressID, PostalCode, LocalRegion, Country, District, StreetName, CityName, 
		Building, Floor, Room
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_address_data
		WHERE (OrderID, AddressID) = (?, ?);`, order, addressID,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToAddress(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Item(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Item {
	order := input.Header.OrderID
	orderItem := input.Header.Item.OrderItem

	rows, err := c.db.Query(
		`SELECT OrderID, OrderItem, OrderItemCategory, OrderItemText, OrderItemTextByBuyer, OrderItemTextBySeller, 
		Product, ProductStandardID, ProductGroup, BaseUnit, PricingDate, PriceDetnExchangeRate, RequestedDeliveryDate, 
		DeliverFrom, DeliverTo, StockConfirmationPartnerFunction, StockConfirmationBusinessPartner, StockConfirmationPlant, 
		StockConfirmationPlantBatch, StockConfirmationPlantBatchValidityStartDate, StockConfirmationPlantBatchValidityEndDate, 
		ProductIsBatchManagedInStockConfirmationPlant, ServicesRenderingDate, OrderQuantityInBaseUnit, OrderQuantityInIssuingUnit, 
		OrderQuantityInReceivingUnit, OrderIssuingUnit, OrderReceivingUnit, StockConfirmationPolicy, StockConfirmationStatus, 
		ConfirmedOrderQuantityInBaseUnit, ItemWeightUnit, ProductGrossWeight, ItemGrossWeight, ProductNetWeight, ItemNetWeight,
		NetAmount, TaxAmount, GrossAmount, BillingDocumentDate, ProductionPlantPartnerFunction, ProductionPlantBusinessPartner, 
		ProductionPlant, ProductionPlantTimeZone, ProductionPlantStorageLocation, IssuingPlantPartnerFunction, IssuingPlantBusinessPartner, 
		IssuingPlant, IssuingPlantTimeZone, IssuingPlantStorageLocation, ReceivingPlantPartnerFunction, ReceivingPlantBusinessPartner, ReceivingPlant, 
		ReceivingPlantTimeZone, ReceivingPlantStorageLocation, ProductIsBatchManagedInProductionPlant, ProductIsBatchManagedInIssuingPlant, 
		ProductIsBatchManagedInReceivingPlant, BatchMgmtPolicyInProductionPlant, BatchMgmtPolicyInIssuingPlant, BatchMgmtPolicyInReceivingPlant,
		ProductionPlantBatch, IssuingPlantBatch, ReceivingPlantBatch, ProductionPlantBatchValidityStartDate, ProductionPlantBatchValidityEndDate, 
		IssuingPlantBatchValidityStartDate, IssuingPlantBatchValidityEndDate, ReceivingPlantBatchValidityStartDate, ReceivingPlantBatchValidityEndDate, 
		Incoterms, BPTaxClassification, ProductTaxClassification, BPAccountAssignmentGroup, ProductAccountAssignmentGroup, PaymentTerms, DueCalculationBaseDate, 
		NetPaymentDays, PaymentDueDate, PaymentMethod, DocumentRjcnReason, ItemBillingBlockReason, Project, AccountingExchangeRate, ReferenceDocument, 
		ReferenceDocumentItem, ItemCompleteDeliveryIsDefined, ItemDeliveryStatus, IssuingStatus, ReceivingStatus, BillingStatus, TaxCode, TaxRate, CountryOfOrigin
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_product_desc_by_bp_data
		WHERE (OrderID, OrderItem) = (?, ?, ?);`, order, orderItem,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToItem(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemPartner(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.ItemPartner {
	order := input.Header.OrderID
	orderItem := input.Header.Item.OrderItem
	partnerFunction := input.Header.Item.ItemPartner.PartnerFunction
	businessPartner := input.Header.Item.ItemPartner.BusinessPartner

	rows, err := c.db.Query(
		`SELECT OrderID, OrderItem, PartnerFunction, BusinessPartner
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_partner_data
		WHERE (OrderID, OrderItem, PartnerFunction, BusinessPartner) = (?, ?, ?, ?);`, order, orderItem, partnerFunction, businessPartner,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToItemPartner(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemPartnerPlant(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.ItemPartnerPlant {
	order := input.Header.OrderID
	orderItem := input.Header.Item.OrderItem
	partnerFunction := input.Header.Item.ItemPartner.PartnerFunction
	businessPartner := input.Header.Item.ItemPartner.BusinessPartner
	plant := input.Header.Item.ItemPartner.ItemPartnerPlant.Plant

	rows, err := c.db.Query(
		`SELECT OrderID, OrderItem, PartnerFunction, BusinessPartner, Plant
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_partner_plant_data
		WHERE (OrderID, OrderItem, PartnerFunction, BusinessPartner, Plant) = (?, ?, ?, ?, ?);`, order, orderItem, partnerFunction, businessPartner, plant,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToItemPartnerPlant(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemPricingElement(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.ItemPricingElement {
	order := input.Header.OrderID
	orderItem := input.Header.Item.OrderItem
	pricingProcedureCounter := input.Header.Item.ItemPricingElement.PricingProcedureCounter

	rows, err := c.db.Query(
		`SELECT OrderID, OrderItem, PricingProcedureStep, PricingProcedureCounter, ConditionType, 
		PricingDate, ConditionRateValue, ConditionCurrency, ConditionQuantity, ConditionQuantityUnit, 
		ConditionRecord, ConditionSequentialNumber, TaxCode, ConditionAmount, TransactionCurrency, 
		ConditionIsManuallyChanged
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_pricing_element_data
		WHERE (OrderID, OrderItem, PricingProcedureCounter) = (?, ?, ?);`, order, orderItem, pricingProcedureCounter,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToItemPricingElement(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemSchedulingLine(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.ItemSchedulingLine {
	order := input.Header.OrderID
	orderItem := input.Header.Item.OrderItem
	scheduleLine := input.Header.Item.ItemSchedulingLine.ScheduleLine

	rows, err := c.db.Query(
		`SELECT OrderID, OrderItem, ScheduleLine, Product, StockConfirmationPartnerFunction, StockConfirmationBusinessPartner, 
		StockConfirmationPlant, StockConfirmationPlantBatch, StockConfirmationPlantBatchValidityStartDate, 
		StockConfirmationPlantBatchValidityEndDate, ConfirmedDeliveryDate, RequestedDeliveryDate, OrderQuantityInBaseUnit, 
		ConfdOrderQtyByPDTAvailCheck, DeliveredQtyInOrderQtyUnit, OpenConfdDelivQtyInOrdQtyUnit, 
		DelivBlockReasonForSchedLine, PlusMinusFlag
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_schedule_line_data
		WHERE (OrderID, OrderItem, ScheduleLine) = (?, ?, ?);`, order, orderItem, scheduleLine,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToItemSchedulingLine(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
