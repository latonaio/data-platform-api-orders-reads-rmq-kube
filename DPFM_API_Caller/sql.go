package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-orders-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-orders-reads-rmq-kube/DPFM_API_Output_Formatter"
	"strings"
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
	var header *[]dpfm_api_output_formatter.Header
	var headerPartner *[]dpfm_api_output_formatter.HeaderPartner
	var headerPartnerContact *[]dpfm_api_output_formatter.HeaderPartnerContact
	var headerPartnerPlant *[]dpfm_api_output_formatter.HeaderPartnerPlant
	var address *[]dpfm_api_output_formatter.Address
	var item *[]dpfm_api_output_formatter.Item
	var itemPartner *[]dpfm_api_output_formatter.ItemPartner
	var itemPartnerPlant *[]dpfm_api_output_formatter.ItemPartnerPlant
	var itemPricingElement *[]dpfm_api_output_formatter.ItemPricingElement
	var itemSchedulingLine *[]dpfm_api_output_formatter.ItemSchedulingLine
	var sellerItems *[]dpfm_api_output_formatter.SellerItems
	var buyerItems *[]dpfm_api_output_formatter.BuyerItems
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				header = c.Header(mtx, input, output, errs, log)
			}()
		case "HeaderByBuyer":
			func() {
				header = c.HeaderByBuyer(mtx, input, output, errs, log)
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
		case "SellerItems":
			func() {
				sellerItems = c.SellerItems(mtx, input, output, errs, log)
			}()
		case "BuyerItems":
			func() {
				buyerItems = c.BuyerItems(mtx, input, output, errs, log)
			}()
		default:
		}
		if len(*errs) != 0 {
			break
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
		SellerItems:          sellerItems,
		BuyerItems:           buyerItems,
	}

	return data
}

func (c *DPFMAPICaller) Header(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	orderID := input.Header.OrderID

	rows, err := c.db.Query(
		`SELECT OrderID, OrderDate, OrderType, Buyer, Seller, CreationDate, LastChangeDate, ContractType, ValidityStartDate,
		ValidityEndDate, InvoicePeriodStartDate, InvoicePeriodEndDate, TotalNetAmount, TotalTaxAmount, TotalGrossAmount,
		HeaderDeliveryStatus, HeaderBlockStatus, HeaderBillingStatus, HeaderDocReferenceStatus, TransactionCurrency,
		PricingDate, PriceDetnExchangeRate, RequestedDeliveryDate, HeaderCompleteDeliveryIsDefined, HeaderBillingBlockStatus,
		HeaderDeliveryBlockStatus, Incoterms, BillFromParty, BillToParty, BillFromCountry, BillToCountry, Payer, Payee,
		PaymentTerms, PaymentMethod, ReferenceDocument, ReferenceDocumentItem, BPAccountAssignmentGroup, AccountingExchangeRate,
		InvoiceDocumentDate, IsExportImportDelivery, HeaderText
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_data
		WHERE OrderID = ?;`, orderID,
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

func (c *DPFMAPICaller) HeaderByBuyer(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	buyer := input.Header.Buyer

	rows, err := c.db.Query(
		`SELECT OrderID, OrderDate, OrderType, Buyer, Seller, CreationDate, LastChangeDate, ContractType, ValidityStartDate,
		ValidityEndDate, InvoicePeriodStartDate, InvoicePeriodEndDate, TotalNetAmount, TotalTaxAmount, TotalGrossAmount,
		HeaderDeliveryStatus, HeaderBlockStatus, HeaderBillingStatus, HeaderDocReferenceStatus, TransactionCurrency,
		PricingDate, PriceDetnExchangeRate, RequestedDeliveryDate, HeaderCompleteDeliveryIsDefined, HeaderBillingBlockStatus,
		HeaderDeliveryBlockStatus, Incoterms, BillFromParty, BillToParty, BillFromCountry, BillToCountry, Payer, Payee,
		PaymentTerms, PaymentMethod, ReferenceDocument, ReferenceDocumentItem, BPAccountAssignmentGroup, AccountingExchangeRate,
		InvoiceDocumentDate, IsExportImportDelivery, HeaderText
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_data
		WHERE Buyer = ?;`, buyer,
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
) *[]dpfm_api_output_formatter.HeaderPartner {
	var args []interface{}
	orderID := input.Header.OrderID
	headerPartner := input.Header.HeaderPartner

	cnt := 0
	for _, v := range headerPartner {
		if v.PartnerFunction == nil || v.BusinessPartner == nil {
			continue
		} else if *v.PartnerFunction == "" {
			continue
		}
		args = append(args, orderID, v.PartnerFunction, v.BusinessPartner)
		cnt++
	}
	repeat := strings.Repeat("(?,?,?),", cnt-1) + "(?,?,?)"

	rows, err := c.db.Query(
		`SELECT OrderID, PartnerFunction, BusinessPartner, BusinessPartnerFullName, BusinessPartnerName,
		Organization, Country, Language, Currency, ExternalDocumentID, AddressID
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_partner_data
		WHERE (OrderID, PartnerFunction, BusinessPartner) IN ( `+repeat+` );`, args...,
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
) *[]dpfm_api_output_formatter.HeaderPartnerContact {
	var args []interface{}
	orderID := input.Header.OrderID
	headerPartner := input.Header.HeaderPartner

	cnt := 0
	for _, v := range headerPartner {
		if v.PartnerFunction == nil || v.BusinessPartner == nil {
			continue
		} else if *v.PartnerFunction == "" {
			continue
		}
		headerPartnerContact := v.HeaderPartnerContact
		for _, w := range headerPartnerContact {
			if w.ContactID == nil {
				continue
			}
			args = append(args, orderID, v.PartnerFunction, v.BusinessPartner, w.ContactID)
		}
		cnt++
	}
	repeat := strings.Repeat("(?,?,?,?),", cnt-1) + "(?,?,?,?)"

	rows, err := c.db.Query(
		`SELECT OrderID, PartnerFunction, BusinessPartner, ContactID, ContactPersonName, EmailAddress, PhoneNumber,
		MobilePhoneNumber, FaxNumber, ContactTag1, ContactTag2, ContactTag3, ContactTag4
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_partner_contact_data
		WHERE (OrderID, PartnerFunction, BusinessPartner, ContactID) IN ( `+repeat+` );`, args...,
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
) *[]dpfm_api_output_formatter.HeaderPartnerPlant {
	var args []interface{}
	orderID := input.Header.OrderID
	headerPartner := input.Header.HeaderPartner

	cnt := 0
	for _, v := range headerPartner {
		if v.PartnerFunction == nil || v.BusinessPartner == nil {
			continue
		} else if *v.PartnerFunction == "" {
			continue
		}
		headerPartnerPlant := v.HeaderPartnerPlant
		for _, w := range headerPartnerPlant {
			if w.Plant == nil {
				continue
			}
			args = append(args, orderID, v.PartnerFunction, v.BusinessPartner, w.Plant)
		}
		cnt++
	}
	repeat := strings.Repeat("(?,?,?,?),", cnt-1) + "(?,?,?,?)"

	rows, err := c.db.Query(
		`SELECT OrderID, PartnerFunction, BusinessPartner, Plant
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_partner_plant_data
		WHERE (OrderID, PartnerFunction, BusinessPartner, Plant) IN ( `+repeat+` );`, args...,
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
) *[]dpfm_api_output_formatter.Address {
	var args []interface{}
	orderID := input.Header.OrderID
	address := input.Header.Address

	cnt := 0
	for _, v := range address {
		if v.AddressID == nil {
			continue
		}
		args = append(args, orderID, v.AddressID)
		cnt++
	}
	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"

	rows, err := c.db.Query(
		`SELECT OrderID, AddressID, PostalCode, LocalRegion, Country, District, StreetName, CityName,
		Building, Floor, Room
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_address_data
		WHERE (OrderID, AddressID) IN ( `+repeat+` );`, args...,
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
) *[]dpfm_api_output_formatter.Item {
	var args []interface{}
	orderID := input.Header.OrderID
	item := input.Header.Item

	cnt := 0
	for _, v := range item {
		if v.OrderItem == nil {
			continue
		}
		args = append(args, orderID, v.OrderItem)
		cnt++
	}
	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"

	rows, err := c.db.Query(
		`SELECT OrderID, OrderItem, OrderItemCategory, OrderItemText, OrderItemTextByBuyer, 
		OrderItemTextBySeller, Product, ProductStandardID, ProductGroup, BaseUnit, PricingDate, 
		PriceDetnExchangeRate, RequestedDeliveryDate, DeliverFrom, DeliverTo, 
		StockConfirmationPartnerFunction, StockConfirmationBusinessPartner, StockConfirmationPlant, 
		StockConfirmationPlantBatch, StockConfirmationPlantBatchValidityStartDate, StockConfirmationPlantBatchValidityEndDate, 
		ProductIsBatchManagedInStockConfirmationPlant, ServicesRenderingDate, OrderQuantityInBaseUnit, 
		OrderQuantityInDeliveryUnit, DeliveryUnit, StockConfirmationPolicy, StockConfirmationStatus, 
		ConfirmedOrderQuantityInBaseUnit, ItemWeightUnit, ProductGrossWeight, ItemGrossWeight, 
		ProductNetWeight, ItemNetWeight, NetAmount, TaxAmount, GrossAmount, BillingDocumentDate, 
		ProductionPlantPartnerFunction, ProductionPlantBusinessPartner, ProductionPlant, ProductionPlantTimeZone, 
		ProductionPlantStorageLocation, IssuingPlantPartnerFunction, IssuingPlantBusinessPartner, IssuingPlant, 
		IssuingPlantTimeZone, IssuingPlantStorageLocation, ReceivingPlantPartnerFunction, ReceivingPlantBusinessPartner, 
		ReceivingPlant, ReceivingPlantTimeZone, ReceivingPlantStorageLocation, ProductIsBatchManagedInProductionPlant, 
		ProductIsBatchManagedInIssuingPlant, ProductIsBatchManagedInReceivingPlant, BatchMgmtPolicyInProductionPlant, 
		BatchMgmtPolicyInIssuingPlant, BatchMgmtPolicyInReceivingPlant, ProductionPlantBatch, IssuingPlantBatch, 
		ReceivingPlantBatch, ProductionPlantBatchValidityStartDate, ProductionPlantBatchValidityEndDate, 
		IssuingPlantBatchValidityStartDate, IssuingPlantBatchValidityEndDate, ReceivingPlantBatchValidityStartDate, 
		ReceivingPlantBatchValidityEndDate, Incoterms, BPTaxClassification, ProductTaxClassification, 
		BPAccountAssignmentGroup, ProductAccountAssignmentGroup, PaymentTerms, DueCalculationBaseDate, 
		PaymentDueDate, NetPaymentDays, PaymentMethod, DocumentRjcnReason, ItemBillingBlockReason, Project, 
		AccountingExchangeRate, ReferenceDocument, ReferenceDocumentItem, ItemCompleteDeliveryIsDefined, 
		ItemDeliveryStatus, IssuingStatus, ReceivingStatus, BillingStatus, TaxCode, TaxRate, CountryOfOrigin
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_data
		WHERE (OrderID, OrderItem) IN ( `+repeat+` );`, args...,
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
) *[]dpfm_api_output_formatter.ItemPartner {
	var args []interface{}
	orderID := input.Header.OrderID
	item := input.Header.Item

	cnt := 0
	for _, v := range item {
		if v.OrderItem == nil {
			continue
		}
		itemPartner := v.ItemPartner
		for _, w := range itemPartner {
			if w.PartnerFunction == nil || w.BusinessPartner == nil {
				continue
			} else if *w.PartnerFunction == "" {
				continue
			}
			args = append(args, orderID, v.OrderItem, w.PartnerFunction, w.BusinessPartner)
		}
		cnt++
	}
	repeat := strings.Repeat("(?,?,?,?),", cnt-1) + "(?,?,?,?)"

	rows, err := c.db.Query(
		`SELECT OrderID, OrderItem, PartnerFunction, BusinessPartner
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_partner_data
		WHERE (OrderID, OrderItem, PartnerFunction, BusinessPartner) IN ( `+repeat+` );`, args...,
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
) *[]dpfm_api_output_formatter.ItemPartnerPlant {
	var args []interface{}
	orderID := input.Header.OrderID
	item := input.Header.Item

	cnt := 0
	for _, v := range item {
		if v.OrderItem == nil {
			continue
		}
		itemPartner := v.ItemPartner
		for _, w := range itemPartner {
			if w.PartnerFunction == nil || w.BusinessPartner == nil {
				continue
			} else if *w.PartnerFunction == "" {
				continue
			}
			itemPartnerPlant := w.ItemPartnerPlant
			if itemPartnerPlant.Plant == nil {
				continue
			} else if *itemPartnerPlant.Plant == "" {
				continue
			}
			args = append(args, orderID, v.OrderItem, w.PartnerFunction, w.BusinessPartner, itemPartnerPlant.Plant)
		}
		cnt++
	}
	repeat := strings.Repeat("(?,?,?,?,?),", cnt-1) + "(?,?,?,?,?)"

	rows, err := c.db.Query(
		`SELECT OrderID, OrderItem, PartnerFunction, BusinessPartner, Plant
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_partner_plant_data
		WHERE (OrderID, OrderItem, PartnerFunction, BusinessPartner, Plant) IN ( `+repeat+` );`, args...,
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
) *[]dpfm_api_output_formatter.ItemPricingElement {
	var args []interface{}
	orderID := input.Header.OrderID
	item := input.Header.Item

	cnt := 0
	for _, v := range item {
		if v.OrderItem == nil {
			continue
		}
		itemPricingElement := v.ItemPricingElement
		for _, w := range itemPricingElement {
			if w.PricingProcedureCounter == nil {
				continue
			}
			args = append(args, orderID, v.OrderItem, w.PricingProcedureCounter)
		}
		cnt++
	}
	repeat := strings.Repeat("(?,?,?),", cnt-1) + "(?,?,?)"

	rows, err := c.db.Query(
		`SELECT OrderID, OrderItem, PricingProcedureStep, PricingProcedureCounter, ConditionType,
		PricingDate, ConditionRateValue, ConditionCurrency, ConditionQuantity, ConditionQuantityUnit,
		ConditionRecord, ConditionSequentialNumber, TaxCode, ConditionAmount, TransactionCurrency,
		ConditionIsManuallyChanged
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_pricing_element_data
		WHERE (OrderID, OrderItem, PricingProcedureCounter) IN ( `+repeat+` );`, args...,
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
) *[]dpfm_api_output_formatter.ItemSchedulingLine {
	var args []interface{}
	orderID := input.Header.OrderID
	item := input.Header.Item

	cnt := 0
	for _, v := range item {
		if v.OrderItem == nil {
			continue
		}
		itemSchedulingLine := v.ItemSchedulingLine
		for _, w := range itemSchedulingLine {
			if w.ScheduleLine == nil {
				continue
			}
			args = append(args, orderID, v.OrderItem, w.ScheduleLine)
		}
		cnt++
	}
	repeat := strings.Repeat("(?,?,?),", cnt-1) + "(?,?,?)"

	rows, err := c.db.Query(
		`SELECT OrderID, OrderItem, ScheduleLine, Product, StockConfirmationPartnerFunction, StockConfirmationBusinessPartner,
		StockConfirmationPlant, StockConfirmationPlantBatch, StockConfirmationPlantBatchValidityStartDate,
		StockConfirmationPlantBatchValidityEndDate, ConfirmedDeliveryDate, RequestedDeliveryDate, OrderQuantityInBaseUnit,
		ConfdOrderQtyByPDTAvailCheck, DeliveredQtyInOrderQtyUnit, OpenConfdDelivQtyInOrdQtyUnit,
		DelivBlockReasonForSchedLine, PlusMinusFlag
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_schedule_line_data
		WHERE (OrderID, OrderItem, ScheduleLine) IN ( `+repeat+` );`, args...,
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

func (c *DPFMAPICaller) SellerItems(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.SellerItems {
	seller := input.Header.Seller

	rows, err := c.db.Query(
		`SELECT OrdersHeader.OrderID,
       BusinessPartnerGeneral.BusinessPartnerFullName, 
       BusinessPartnerGeneral.BusinessPartnerName,
       BusinessPartnerGeneralForDeliverToParty.BusinessPartnerFullName as DeliverToPartyBusinessPartnerFullName,
       BusinessPartnerGeneralForDeliverToParty.BusinessPartnerName as DeliverToPartyBusinessPartnerName,
       OrdersHeader.HeaderDeliveryStatus
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_data as OrdersHeader
		INNER JOIN DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_partner_data as OrdersHeaderPartner
		ON OrdersHeader.OrderID = OrdersHeaderPartner.OrderID and OrdersHeader.Buyer = OrdersHeaderPartner.BusinessPartner
		INNER JOIN DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_general_data as BusinessPartnerGeneral
		ON OrdersHeader.Seller = BusinessPartnerGeneral.BusinessPartner
		INNER JOIN DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_data as OrdersItem
		ON OrdersHeader.OrderID = OrdersItem.OrderID
		LEFT JOIN DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_general_data as BusinessPartnerGeneralForDeliverToParty
		ON OrdersItem.DeliverToParty = BusinessPartnerGeneralForDeliverToParty.BusinessPartner
		WHERE (OrdersHeader.Seller) = (?);`, seller,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToSellerItems(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) BuyerItems(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.BuyerItems {
	buyer := input.Header.Buyer

	rows, err := c.db.Query(
		`SELECT OrdersHeader.OrderID,
       BusinessPartnerGeneral.BusinessPartnerFullName, 
       BusinessPartnerGeneral.BusinessPartnerName,
       BusinessPartnerGeneralForDeliverToParty.BusinessPartnerFullName as DeliverToPartyBusinessPartnerFullName,
       BusinessPartnerGeneralForDeliverToParty.BusinessPartnerName as DeliverToPartyBusinessPartnerName,
       OrdersHeader.HeaderDeliveryStatus
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_data as OrdersHeader
		INNER JOIN DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_partner_data as OrdersHeaderPartner
		ON OrdersHeader.OrderID = OrdersHeaderPartner.OrderID and OrdersHeader.Buyer = OrdersHeaderPartner.BusinessPartner
		INNER JOIN DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_general_data as BusinessPartnerGeneral
		ON OrdersHeader.Seller = BusinessPartnerGeneral.BusinessPartner
		INNER JOIN DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_data as OrdersItem
		ON OrdersHeader.OrderID = OrdersItem.OrderID
		LEFT JOIN DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_general_data as BusinessPartnerGeneralForDeliverToParty
		ON OrdersItem.DeliverToParty = BusinessPartnerGeneralForDeliverToParty.BusinessPartner
		WHERE (OrdersHeader.Buyer) = (?);`, buyer,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToBuyerItems(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
