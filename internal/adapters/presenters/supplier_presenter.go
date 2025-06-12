package presenters

import (
	"time"

	"github.com/davigomesdev/reconfile/internal/adapters/collections"
	"github.com/davigomesdev/reconfile/internal/domain/contracts"
	"github.com/davigomesdev/reconfile/internal/domain/entities"
)

type SupplierOutput struct {
	ID                            string             `json:"id"`
	PartnerId                     string             `json:"partnerId"`
	PartnerName                   string             `json:"partnerName"`
	CustomerId                    string             `json:"customerId"`
	CustomerName                  string             `json:"customerName"`
	CustomerDomainName            string             `json:"customerDomainName"`
	CustomerCountry               string             `json:"country"`
	MpnId                         int                `json:"mpnId"`
	Tier2MpnId                    int                `json:"tier2MpnId"`
	InvoiceNumber                 string             `json:"invoiceNumber"`
	ProductId                     string             `json:"productId"`
	SkuId                         string             `json:"skuId"`
	AvailabilityId                string             `json:"availabilityId"`
	SkuName                       string             `json:"skuName"`
	ProductName                   string             `json:"productName"`
	PublisherName                 string             `json:"publisherName"`
	PublisherId                   *string            `json:"publisherId"`
	SubscriptionDescription       *string            `json:"subscriptionDescription"`
	SubscriptionId                string             `json:"subscriptionId"`
	ChargeStartDate               time.Time          `json:"chargeStartDate"`
	ChargeEndDate                 time.Time          `json:"chargeEndDate"`
	UsageDate                     time.Time          `json:"usageDate"`
	MeterType                     string             `json:"meterType"`
	MeterCategory                 string             `json:"meterCategory"`
	MeterId                       string             `json:"meterId"`
	MeterSubCategory              string             `json:"meterSubCategory"`
	MeterName                     string             `json:"meterName"`
	MeterRegion                   *string            `json:"meterRegion"`
	Unit                          string             `json:"unit"`
	ResourceLocation              string             `json:"resourceLocation"`
	ConsumedService               string             `json:"consumedService"`
	ResourceGroup                 string             `json:"resourceGroup"`
	ResourceURI                   string             `json:"resourceURI"`
	ChargeType                    string             `json:"chargeType"`
	UnitPrice                     float64            `json:"unitPrice"`
	Quantity                      float64            `json:"quantity"`
	UnitType                      string             `json:"unitType"`
	BillingPreTaxTotal            float64            `json:"billingPreTaxTotal"`
	BillingCurrency               string             `json:"billingCurrency"`
	PricingPreTaxTotal            float64            `json:"pricingPreTaxTotal"`
	PricingCurrency               string             `json:"pricingCurrency"`
	ServiceInfo1                  *string            `json:"serviceInfo1"`
	ServiceInfo2                  *string            `json:"serviceInfo2"`
	Tags                          *map[string]string `json:"tags"`
	AdditionalInfo                *map[string]string `json:"additionalInfo"`
	EffectiveUnitPrice            float64            `json:"effectiveUnitPrice"`
	PCToBCExchangeRate            int                `json:"pctoBCExchangeRate"`
	PCToBCExchangeRateDate        time.Time          `json:"pctoBCExchangeRateDate"`
	EntitlementId                 string             `json:"entitlementId"`
	EntitlementDescription        string             `json:"entitlementDescription"`
	PartnerEarnedCreditPercentage int                `json:"partnerEarnedCreditPercentage"`
	CreditPercentage              int                `json:"creditPercentage"`
	CreditType                    string             `json:"creditType"`
	BenefitOrderId                *string            `json:"benefitOrderId"`
	BenefitId                     *string            `json:"benefitId"`
	BenefitType                   *string            `json:"benefitType"`
	CreatedAt                     time.Time          `json:"createdAt"`
	UpdatedAt                     time.Time          `json:"updatedAt"`
}

func newSupplierOutput(s *entities.SupplierEntity) *SupplierOutput {
	return &SupplierOutput{
		ID:                            s.ID,
		PartnerId:                     s.PartnerId,
		PartnerName:                   s.PartnerName,
		CustomerId:                    s.CustomerId,
		CustomerName:                  s.CustomerName,
		CustomerDomainName:            s.CustomerDomainName,
		CustomerCountry:               s.CustomerCountry,
		MpnId:                         s.MpnId,
		Tier2MpnId:                    s.Tier2MpnId,
		InvoiceNumber:                 s.InvoiceNumber,
		ProductId:                     s.ProductId,
		SkuId:                         s.SkuId,
		AvailabilityId:                s.AvailabilityId,
		SkuName:                       s.SkuName,
		ProductName:                   s.ProductName,
		PublisherName:                 s.PublisherName,
		PublisherId:                   s.PublisherId,
		SubscriptionDescription:       s.SubscriptionDescription,
		SubscriptionId:                s.SubscriptionId,
		ChargeStartDate:               s.ChargeStartDate,
		ChargeEndDate:                 s.ChargeEndDate,
		UsageDate:                     s.UsageDate,
		MeterType:                     s.MeterType,
		MeterCategory:                 s.MeterCategory,
		MeterId:                       s.MeterId,
		MeterSubCategory:              s.MeterSubCategory,
		MeterName:                     s.MeterName,
		MeterRegion:                   s.MeterRegion,
		Unit:                          s.Unit,
		ResourceLocation:              s.ResourceLocation,
		ConsumedService:               s.ConsumedService,
		ResourceGroup:                 s.ResourceGroup,
		ResourceURI:                   s.ResourceURI,
		ChargeType:                    s.ChargeType,
		UnitPrice:                     s.UnitPrice,
		Quantity:                      s.Quantity,
		UnitType:                      s.UnitType,
		BillingPreTaxTotal:            s.BillingPreTaxTotal,
		BillingCurrency:               s.BillingCurrency,
		PricingPreTaxTotal:            s.PricingPreTaxTotal,
		PricingCurrency:               s.PricingCurrency,
		ServiceInfo1:                  s.ServiceInfo1,
		ServiceInfo2:                  s.ServiceInfo2,
		Tags:                          s.Tags,
		AdditionalInfo:                s.AdditionalInfo,
		EffectiveUnitPrice:            s.EffectiveUnitPrice,
		PCToBCExchangeRate:            s.PCToBCExchangeRate,
		PCToBCExchangeRateDate:        s.PCToBCExchangeRateDate,
		EntitlementId:                 s.EntitlementId,
		EntitlementDescription:        s.EntitlementDescription,
		PartnerEarnedCreditPercentage: s.PartnerEarnedCreditPercentage,
		CreditPercentage:              s.CreditPercentage,
		CreditType:                    s.CreditType,
		BenefitOrderId:                s.BenefitOrderId,
		BenefitId:                     s.BenefitId,
		BenefitType:                   s.BenefitType,
		CreatedAt:                     s.CreatedAt,
		UpdatedAt:                     s.UpdatedAt,
	}
}

func NewSupplierPresenter(s *entities.SupplierEntity) *collections.Presenter[*SupplierOutput] {
	return &collections.Presenter[*SupplierOutput]{
		Data: newSupplierOutput(s),
	}
}

func NewSupplierCollectionPresenter(sr *contracts.SearchResult[*entities.SupplierEntity]) *collections.CollectionPresenter[*SupplierOutput] {
	data := make([]*SupplierOutput, 0, len(sr.Items))

	for _, supplier := range sr.Items {
		if p := newSupplierOutput(supplier); p != nil {
			data = append(data, p)
		}
	}

	return &collections.CollectionPresenter[*SupplierOutput]{
		Data: data,
		Meta: collections.PaginationPresenter{
			CurrentPage: sr.CurrentPage,
			PerPage:     sr.PerPage,
			LastPage:    sr.LastPage(),
			Total:       sr.Total,
		},
	}
}
