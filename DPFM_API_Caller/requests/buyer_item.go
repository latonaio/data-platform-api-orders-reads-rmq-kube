package requests

type BuyerItem struct {
	OrderID                               int     `json:"OrderID"`
	BusinessPartnerFullName               *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName                   *string `json:"BusinessPartnerName"`
	DeliverToPartyBusinessPartnerName     *string `json:"DeliverToPartyBusinessPartnerName"`
	DeliverToPartyBusinessPartnerFullName *string `json:"DeliverToPartyBusinessPartnerFullName"`
	HeaderDeliveryStatus                  *string `json:"HeaderDeliveryStatus"`
}
