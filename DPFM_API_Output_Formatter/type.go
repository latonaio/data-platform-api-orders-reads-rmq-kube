package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	Result              bool        `json:"result"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	SQLUpdateResult     *bool       `json:"sql_update_result"`
	SQLUpdateError      string      `json:"sql_update_error"`
	SubfuncResult       *bool       `json:"subfunc_result"`
	SubfuncError        string      `json:"subfunc_error"`
	ExconfResult        *bool       `json:"exconf_result"`
	ExconfError         string      `json:"exconf_error"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type Message struct {
	Header               *[]Header               `json:"Header"`
	HeaderPartner        *[]HeaderPartner        `json:"HeaderPartner"`
	HeaderPartnerContact *[]HeaderPartnerContact `json:"HeaderPartnerContact"`
	HeaderPartnerPlant   *[]HeaderPartnerPlant   `json:"HeaderPartnerPlant"`
	Address              *[]Address              `json:"Address"`
	Item                 *[]Item                 `json:"Item"`
	ItemPartner          *[]ItemPartner          `json:"ItemPartner"`
	ItemPartnerPlant     *[]ItemPartnerPlant     `json:"ItemPartnerPlant"`
	ItemPricingElement   *[]ItemPricingElement   `json:"ItemPricingElement"`
	ItemSchedulingLine   *[]ItemSchedulingLine   `json:"ItemSchedulingLine"`
	SellerItems          *[]SellerItems          `json:"SellerItems"`
	BuyerItems           *[]BuyerItems           `json:"BuyerItems"`
}

type Orders struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Product       string `json:"Product"`
	APISchema     string `json:"api_schema"`
	MaterialCode  string `json:"material_code"`
	Deleted       string `json:"deleted"`
}

type Header struct {
	OrderID                         int      `json:"OrderID"`
	OrderDate                       *string  `json:"OrderDate"`
	OrderType                       *string  `json:"OrderType"`
	Buyer                           *int     `json:"Buyer"`
	Seller                          *int     `json:"Seller"`
	CreationDate                    *string  `json:"CreationDate"`
	LastChangeDate                  *string  `json:"LastChangeDate"`
	ContractType                    *string  `json:"ContractType"`
	ValidityStartDate               *string  `json:"ValidityStartDate"`
	ValidityEndDate                 *string  `json:"ValidityEndDate"`
	InvoicePeriodStartDate          *string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate            *string  `json:"InvoicePeriodEndDate"`
	TotalNetAmount                  *float32 `json:"TotalNetAmount"`
	TotalTaxAmount                  *float32 `json:"TotalTaxAmount"`
	TotalGrossAmount                *float32 `json:"TotalGrossAmount"`
	HeaderDeliveryStatus            *string  `json:"HeaderDeliveryStatus"`
	HeaderBlockStatus               *bool    `json:"HeaderBlockStatus"`
	HeaderBillingStatus             *string  `json:"HeaderBillingStatus"`
	HeaderDocReferenceStatus        *string  `json:"HeaderDocReferenceStatus"`
	TransactionCurrency             *string  `json:"TransactionCurrency"`
	PricingDate                     *string  `json:"PricingDate"`
	PriceDetnExchangeRate           *float32 `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate           *string  `json:"RequestedDeliveryDate"`
	HeaderCompleteDeliveryIsDefined *bool    `json:"HeaderCompleteDeliveryIsDefined"`
	HeaderBillingBlockStatus        *bool    `json:"HeaderBillingBlockStatus"`
	HeaderDeliveryBlockStatus       *bool    `json:"HeaderDeliveryBlockStatus"`
	Incoterms                       *string  `json:"Incoterms"`
	BillFromParty                   *int     `json:"BillFromParty"`
	BillToParty                     *int     `json:"BillToParty"`
	BillFromCountry                 *string  `json:"BillFromCountry"`
	BillToCountry                   *string  `json:"BillToCountry"`
	Payer                           *int     `json:"Payer"`
	Payee                           *int     `json:"Payee"`
	PaymentTerms                    *string  `json:"PaymentTerms"`
	PaymentMethod                   *string  `json:"PaymentMethod"`
	ReferenceDocument               *int     `json:"ReferenceDocument"`
	ReferenceDocumentItem           *int     `json:"ReferenceDocumentItem"`
	BPAccountAssignmentGroup        *string  `json:"BPAccountAssignmentGroup"`
	AccountingExchangeRate          *float32 `json:"AccountingExchangeRate"`
	InvoiceDocumentDate             *string  `json:"InvoiceDocumentDate"`
	IsExportImportDelivery          *bool    `json:"IsExportImportDelivery"`
	HeaderText                      *string  `json:"HeaderText"`
}

type HeaderPartner struct {
	OrderID                 int     `json:"OrderID"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
	AddressID               *int    `json:"AddressID"`
}

type HeaderPartnerContact struct {
	OrderID           int     `json:"OrderID"`
	PartnerFunction   string  `json:"PartnerFunction"`
	BusinessPartner   int     `json:"BusinessPartner"`
	ContactID         int     `json:"ContactID"`
	ContactPersonName *string `json:"ContactPersonName"`
	EmailAddress      *string `json:"EmailAddress"`
	PhoneNumber       *string `json:"PhoneNumber"`
	MobilePhoneNumber *string `json:"MobilePhoneNumber"`
	FaxNumber         *string `json:"FaxNumber"`
	ContactTag1       *string `json:"ContactTag1"`
	ContactTag2       *string `json:"ContactTag2"`
	ContactTag3       *string `json:"ContactTag3"`
	ContactTag4       *string `json:"ContactTag4"`
}

type HeaderPartnerPlant struct {
	OrderID         int    `json:"OrderID"`
	PartnerFunction string `json:"PartnerFunction"`
	BusinessPartner int    `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
}

type Item struct {
	OrderID                                       int      `json:"OrderID"`
	OrderItem                                     int      `json:"OrderItem"`
	OrderItemCategory                             string   `json:"OrderItemCategory"`
	SupplyChainRelationshipID                     int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID             *int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID        *int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipStockConfPlantID       *int     `json:"SupplyChainRelationshipStockConfPlantID"`
	SupplyChainRelationshipProductionPlantID      *int     `json:"SupplyChainRelationshipProductionPlantID"`
	OrderItemText                                 string   `json:"OrderItemText"`
	OrderItemTextByBuyer                          string   `json:"OrderItemTextByBuyer"`
	OrderItemTextBySeller                         string   `json:"OrderItemTextBySeller"`
	Product                                       string   `json:"Product"`
	ProductStandardID                             string   `json:"ProductStandardID"`
	ProductGroup                                  *string  `json:"ProductGroup"`
	BaseUnit                                      string   `json:"BaseUnit"`
	PricingDate                                   string   `json:"PricingDate"`
	PriceDetnExchangeRate                         *float32 `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate                         string   `json:"RequestedDeliveryDate"`
	DeliverToParty                                *int     `json:"DeliverToParty"`
	DeliverFromParty                              *int     `json:"DeliverFromParty"`
	DeliverToPlant                                *string  `json:"DeliverToPlant"`
	DeliverToPlantTimeZone                        *string  `json:"DeliverToPlantTimeZone"`
	DeliverToPlantStorageLocation                 *string  `json:"DeliverToPlantStorageLocation"`
	ProductIsBatchManagedInDeliverToPlant         *bool    `json:"ProductIsBatchManagedInDeliverToPlant"`
	BatchMgmtPolicyInDeliverToPlant               *string  `json:"BatchMgmtPolicyInDeliverToPlant"`
	DeliverToPlantBatch                           *string  `json:"DeliverToPlantBatch"`
	DeliverToPlantBatchValidityStartDate          *string  `json:"DeliverToPlantBatchValidityStartDate"`
	DeliverToPlantBatchValidityEndDate            *string  `json:"DeliverToPlantBatchValidityEndDate"`
	DeliverFromPlant                              *string  `json:"DeliverFromPlant"`
	DeliverFromPlantTimeZone                      *string  `json:"DeliverFromPlantTimeZone"`
	DeliverFromPlantStorageLocation               *string  `json:"DeliverFromPlantStorageLocation"`
	ProductIsBatchManagedInDeliverFromPlant       *bool    `json:"ProductIsBatchManagedInDeliverFromPlant"`
	BatchMgmtPolicyInDeliverFromPlant             *string  `json:"BatchMgmtPolicyInDeliverFromPlant"`
	DeliverFromPlantBatch                         *string  `json:"DeliverFromPlantBatch"`
	DeliverFromPlantBatchValidityStartDate        *string  `json:"DeliverFromPlantBatchValidityStartDate"`
	DeliverFromPlantBatchValidityEndDate          *string  `json:"DeliverFromPlantBatchValidityEndDate"`
	DeliveryUnit                                  string   `json:"DeliveryUnit"`
	StockConfirmationBusinessPartner              *int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                        *string  `json:"StockConfirmationPlant"`
	StockConfirmationPlantTimeZone                *string  `json:"StockConfirmationPlantTimeZone"`
	StockConfirmationPlantBatch                   *string  `json:"StockConfirmationPlantBatch"`
	StockConfirmationPlantBatchValidityStartDate  *string  `json:"StockConfirmationPlantBatchValidityStartDate"`
	StockConfirmationPlantBatchValidityEndDate    *string  `json:"StockConfirmationPlantBatchValidityEndDate"`
	ProductIsBatchManagedInStockConfirmationPlant *bool    `json:"ProductIsBatchManagedInStockConfirmationPlant"`
	BatchMgmtPolicyStockConfirmationInPlant       *string  `json:"BatchMgmtPolicyStockConfirmationInPlant"`
	ServicesRenderingDate                         *string  `json:"ServicesRenderingDate"`
	OrderQuantityInBaseUnit                       float32  `json:"OrderQuantityInBaseUnit"`
	OrderQuantityInDeliveryUnit                   float32  `json:"OrderQuantityInDeliveryUnit"`
	StockConfirmationPolicy                       *string  `json:"StockConfirmationPolicy"`
	StockConfirmationStatus                       *string  `json:"StockConfirmationStatus"`
	ConfirmedOrderQuantityInBaseUnit              *float32 `json:"ConfirmedOrderQuantityInBaseUnit"`
	ItemWeightUnit                                *string  `json:"ItemWeightUnit"`
	ProductGrossWeight                            *float32 `json:"ProductGrossWeight"`
	ItemGrossWeight                               *float32 `json:"ItemGrossWeight"`
	ProductNetWeight                              *float32 `json:"ProductNetWeight"`
	ItemNetWeight                                 *float32 `json:"ItemNetWeight"`
	NetAmount                                     *float32 `json:"NetAmount"`
	TaxAmount                                     *float32 `json:"TaxAmount"`
	GrossAmount                                   *float32 `json:"GrossAmount"`
	InvoiceDocumentDate                           *string  `json:"InvoiceDocumentDate"`
	ProductionPlantBusinessPartner                *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                               *string  `json:"ProductionPlant"`
	ProductionPlantTimeZone                       *string  `json:"ProductionPlantTimeZone"`
	ProductionPlantStorageLocation                *string  `json:"ProductionPlantStorageLocation"`
	ProductIsBatchManagedInProductionPlant        *bool    `json:"ProductIsBatchManagedInProductionPlant"`
	BatchMgmtPolicyInProductionPlant              *string  `json:"BatchMgmtPolicyInProductionPlant"`
	ProductionPlantBatch                          *string  `json:"ProductionPlantBatch"`
	ProductionPlantBatchValidityStartDate         *string  `json:"ProductionPlantBatchValidityStartDate"`
	ProductionPlantBatchValidityEndDate           *string  `json:"ProductionPlantBatchValidityEndDate"`
	Incoterms                                     *string  `json:"Incoterms"`
	TransactionTaxClassification                  string   `json:"TransactionTaxClassification"`
	ProductTaxClassificationBillToCountry         string   `json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry       string   `json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification                      string   `json:"DefinedTaxClassification"`
	AccountAssignmentGroup                        string   `json:"AccountAssignmentGroup"`
	ProductAccountAssignmentGroup                 string   `json:"ProductAccountAssignmentGroup"`
	PaymentTerms                                  string   `json:"PaymentTerms"`
	DueCalculationBaseDate                        *string  `json:"DueCalculationBaseDate"`
	PaymentDueDate                                *string  `json:"PaymentDueDate"`
	NetPaymentDays                                *int     `json:"NetPaymentDays"`
	PaymentMethod                                 string   `json:"PaymentMethod"`
	Project                                       *string  `json:"Project"`
	AccountingExchangeRate                        *float32 `json:"AccountingExchangeRate"`
	ReferenceDocument                             *int     `json:"ReferenceDocument"`
	ReferenceDocumentItem                         *int     `json:"ReferenceDocumentItem"`
	ItemCompleteDeliveryIsDefined                 *bool    `json:"ItemCompleteDeliveryIsDefined"`
	ItemDeliveryStatus                            *string  `json:"ItemDeliveryStatus"`
	IssuingStatus                                 *string  `json:"IssuingStatus"`
	ReceivingStatus                               *string  `json:"ReceivingStatus"`
	ItemBillingStatus                             *string  `json:"ItemBillingStatus"`
	TaxCode                                       *string  `json:"TaxCode"`
	TaxRate                                       *float32 `json:"TaxRate"`
	CountryOfOrigin                               *string  `json:"CountryOfOrigin"`
	CountryOfOriginLanguage                       *string  `json:"CountryOfOriginLanguage"`
	ItemBlockStatus                               *bool    `json:"ItemBlockStatus"`
	ItemDeliveryBlockStatus                       *bool    `json:"ItemDeliveryBlockStatus"`
	ItemBillingBlockStatus                        *bool    `json:"ItemBillingBlockStatus"`
}

type ItemPartner struct {
	OrderID         int    `json:"OrderID"`
	OrderItem       int    `json:"OrderItem"`
	PartnerFunction string `json:"PartnerFunction"`
	BusinessPartner int    `json:"BusinessPartner"`
}

type ItemPartnerPlant struct {
	OrderID         int    `json:"OrderID"`
	OrderItem       int    `json:"OrderItem"`
	PartnerFunction string `json:"PartnerFunction"`
	BusinessPartner int    `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
}

type ItemPricingElement struct {
	OrderID                    int      `json:"OrderID"`
	OrderItem                  int      `json:"OrderItem"`
	PricingProcedureStep       *int     `json:"PricingProcedureStep"`
	PricingProcedureCounter    int      `json:"PricingProcedureCounter"`
	ConditionType              *string  `json:"ConditionType"`
	PricingDate                *string  `json:"PricingDate"`
	ConditionRateValue         *float32 `json:"ConditionRateValue"`
	ConditionCurrency          *string  `json:"ConditionCurrency"`
	ConditionQuantity          *float32 `json:"ConditionQuantity"`
	ConditionQuantityUnit      *string  `json:"ConditionQuantityUnit"`
	ConditionRecord            *int     `json:"ConditionRecord"`
	ConditionSequentialNumber  *int     `json:"ConditionSequentialNumber"`
	TaxCode                    *string  `json:"TaxCode"`
	ConditionAmount            *float32 `json:"ConditionAmount"`
	TransactionCurrency        *string  `json:"TransactionCurrency"`
	ConditionIsManuallyChanged *bool    `json:"ConditionIsManuallyChanged"`
}

type ItemSchedulingLine struct {
	OrderID                                      int      `json:"OrderID"`
	OrderItem                                    int      `json:"OrderItem"`
	ScheduleLine                                 int      `json:"ScheduleLine"`
	Product                                      *string  `json:"Product"`
	StockConfirmationPartnerFunction             *string  `json:"StockConfirmationPartnerFunction"`
	StockConfirmationBusinessPartner             *int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                       *string  `json:"StockConfirmationPlant"`
	StockConfirmationPlantBatch                  *string  `json:"StockConfirmationPlantBatch"`
	StockConfirmationPlantBatchValidityStartDate *string  `json:"StockConfirmationPlantBatchValidityStartDate"`
	StockConfirmationPlantBatchValidityEndDate   *string  `json:"StockConfirmationPlantBatchValidityEndDate"`
	ConfirmedDeliveryDate                        *string  `json:"ConfirmedDeliveryDate"`
	RequestedDeliveryDate                        *string  `json:"RequestedDeliveryDate"`
	OrderQuantityInBaseUnit                      *float32 `json:"OrderQuantityInBaseUnit"`
	ConfdOrderQtyByPDTAvailCheck                 *float32 `json:"ConfdOrderQtyByPDTAvailCheck"`
	DeliveredQtyInOrderQtyUnit                   *float32 `json:"DeliveredQtyInOrderQtyUnit"`
	OpenConfdDelivQtyInOrdQtyUnit                *float32 `json:"OpenConfdDelivQtyInOrdQtyUnit"`
	DelivBlockReasonForSchedLine                 *bool    `json:"DelivBlockReasonForSchedLine"`
	PlusMinusFlag                                *string  `json:"PlusMinusFlag"`
}

type SellerItems struct {
	OrderID                               int     `json:"OrderID"`
	BusinessPartnerFullName               *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName                   *string `json:"BusinessPartnerName"`
	DeliverToPartyBusinessPartnerName     *string `json:"DeliverToPartyBusinessPartnerName"`
	DeliverToPartyBusinessPartnerFullName *string `json:"DeliverToPartyBusinessPartnerFullName"`
	HeaderDeliveryStatus                  *string `json:"HeaderDeliveryStatus"`
}

type BuyerItems struct {
	OrderID                               int     `json:"OrderID"`
	BusinessPartnerFullName               *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName                   *string `json:"BusinessPartnerName"`
	DeliverToPartyBusinessPartnerName     *string `json:"DeliverToPartyBusinessPartnerName"`
	DeliverToPartyBusinessPartnerFullName *string `json:"DeliverToPartyBusinessPartnerFullName"`
	HeaderDeliveryStatus                  *string `json:"HeaderDeliveryStatus"`
}

type Address struct {
	OrderID     int     `json:"OrderID"`
	AddressID   int     `json:"AddressID"`
	PostalCode  *string `json:"PostalCode"`
	LocalRegion *string `json:"LocalRegion"`
	Country     *string `json:"Country"`
	District    *string `json:"District"`
	StreetName  *string `json:"StreetName"`
	CityName    *string `json:"CityName"`
	Building    *string `json:"Building"`
	Floor       *int    `json:"Floor"`
	Room        *int    `json:"Room"`
}
