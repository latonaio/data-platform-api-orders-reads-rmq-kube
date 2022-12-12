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
	var item *dpfm_api_output_formatter.Item
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
		case "Item":
			func() {
				item = c.Item(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:        header,
		HeaderPartner: headerPartner,
		Item:          item,
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
		`SELECT OrderID, PartnerFunction, BusinessPartner, BusinessPartnerFullName, BusinessPartnerName, Organization, Country, Language, Currency, ExternalDocumentID, AddressID
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_product_desc_by_bp_data
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
