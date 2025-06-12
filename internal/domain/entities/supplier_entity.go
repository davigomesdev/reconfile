package entities

import (
	"time"

	"github.com/davigomesdev/reconfile/internal/domain/contracts"
	"github.com/davigomesdev/reconfile/internal/domain/validators"
)

type SupplierProps struct {
	ID                            *string
	PartnerId                     string
	PartnerName                   string
	CustomerId                    string
	CustomerName                  string
	CustomerDomainName            string
	CustomerCountry               string
	MpnId                         int
	Tier2MpnId                    int
	InvoiceNumber                 string
	ProductId                     string
	SkuId                         string
	AvailabilityId                string
	SkuName                       string
	ProductName                   string
	PublisherName                 string
	PublisherId                   *string
	SubscriptionDescription       *string
	SubscriptionId                string
	ChargeStartDate               time.Time
	ChargeEndDate                 time.Time
	UsageDate                     time.Time
	MeterType                     string
	MeterCategory                 string
	MeterId                       string
	MeterSubCategory              string
	MeterName                     string
	MeterRegion                   *string
	Unit                          string
	ResourceLocation              string
	ConsumedService               string
	ResourceGroup                 string
	ResourceURI                   string
	ChargeType                    string
	UnitPrice                     float64
	Quantity                      float64
	UnitType                      string
	BillingPreTaxTotal            float64
	BillingCurrency               string
	PricingPreTaxTotal            float64
	PricingCurrency               string
	ServiceInfo1                  *string
	ServiceInfo2                  *string
	Tags                          *map[string]string
	AdditionalInfo                *map[string]string
	EffectiveUnitPrice            float64
	PCToBCExchangeRate            int
	PCToBCExchangeRateDate        time.Time
	EntitlementId                 string
	EntitlementDescription        string
	PartnerEarnedCreditPercentage int
	CreditPercentage              int
	CreditType                    string
	BenefitOrderId                *string
	BenefitId                     *string
	BenefitType                   *string
	CreatedAt                     *time.Time
	UpdatedAt                     *time.Time
	DeletedAt                     *time.Time
}

type SupplierEntity struct {
	contracts.Entity
	PartnerId                     string             `json:"partnerId" validate:"required"`
	PartnerName                   string             `json:"partnerName" validate:"required,max=255"`
	CustomerId                    string             `json:"customerId" validate:"required"`
	CustomerName                  string             `json:"customerName" validate:"required,max=255"`
	CustomerDomainName            string             `json:"customerDomainName" validate:"required,fqdn"`
	CustomerCountry               string             `json:"country" validate:"required,max=255"`
	MpnId                         int                `json:"mpnId" validate:"numeric"`
	Tier2MpnId                    int                `json:"tier2MpnId" validate:"numeric"`
	InvoiceNumber                 string             `json:"invoiceNumber" validate:"required"`
	ProductId                     string             `json:"productId" validate:"required"`
	SkuId                         string             `json:"skuId" validate:"required"`
	AvailabilityId                string             `json:"availabilityId" validate:"required"`
	SkuName                       string             `json:"skuName" validate:"required,max=255"`
	ProductName                   string             `json:"productName" validate:"required,max=255"`
	PublisherName                 string             `json:"publisherName" validate:"required,max=255"`
	PublisherId                   *string            `json:"publisherId" validate:"omitempty,max=255"`
	SubscriptionDescription       *string            `json:"subscriptionDescription" validate:"omitempty,max=255"`
	SubscriptionId                string             `json:"subscriptionId" validate:"required"`
	ChargeStartDate               time.Time          `json:"chargeStartDate" validate:"required"`
	ChargeEndDate                 time.Time          `json:"chargeEndDate" validate:"required"`
	UsageDate                     time.Time          `json:"usageDate" validate:"required"`
	MeterType                     string             `json:"meterType" validate:"required,max=255"`
	MeterCategory                 string             `json:"meterCategory" validate:"required,max=255"`
	MeterId                       string             `json:"meterId" validate:"required"`
	MeterSubCategory              string             `json:"meterSubCategory" validate:"required,max=255"`
	MeterName                     string             `json:"meterName" validate:"required,max=255"`
	MeterRegion                   *string            `json:"meterRegion" validate:"omitempty,max=255"`
	Unit                          string             `json:"unit" validate:"required,max=255"`
	ResourceLocation              string             `json:"resourceLocation" validate:"required,max=255"`
	ConsumedService               string             `json:"consumedService" validate:"required,max=255"`
	ResourceGroup                 string             `json:"resourceGroup" validate:"required,max=255"`
	ResourceURI                   string             `json:"resourceURI" validate:"required"`
	ChargeType                    string             `json:"chargeType" validate:"required,max=255"`
	UnitPrice                     float64            `json:"unitPrice" validate:"numeric"`
	Quantity                      float64            `json:"quantity" validate:"numeric"`
	UnitType                      string             `json:"unitType" validate:"required,max=255"`
	BillingPreTaxTotal            float64            `json:"billingPreTaxTotal" validate:"numeric"`
	BillingCurrency               string             `json:"billingCurrency" validate:"required,max=255"`
	PricingPreTaxTotal            float64            `json:"pricingPreTaxTotal" validate:"numeric"`
	PricingCurrency               string             `json:"pricingCurrency" validate:"required,max=255"`
	ServiceInfo1                  *string            `json:"serviceInfo1" validate:"omitempty,max=255"`
	ServiceInfo2                  *string            `json:"serviceInfo2" validate:"omitempty,max=255"`
	Tags                          *map[string]string `json:"tags"`
	AdditionalInfo                *map[string]string `json:"additionalInfo"`
	EffectiveUnitPrice            float64            `json:"effectiveUnitPrice" validate:"numeric"`
	PCToBCExchangeRate            int                `json:"pctoBCExchangeRate" validate:"numeric"`
	PCToBCExchangeRateDate        time.Time          `json:"pctoBCExchangeRateDate" validate:"required"`
	EntitlementId                 string             `json:"entitlementId" validate:"required"`
	EntitlementDescription        string             `json:"entitlementDescription" validate:"required,max=255"`
	PartnerEarnedCreditPercentage int                `json:"partnerEarnedCreditPercentage" validate:"numeric"`
	CreditPercentage              int                `json:"creditPercentage" validate:"numeric"`
	CreditType                    string             `json:"creditType" validate:"required,max=255"`
	BenefitOrderId                *string            `json:"benefitOrderId" validate:"omitempty,max=255"`
	BenefitId                     *string            `json:"benefitId" validate:"omitempty,max=255"`
	BenefitType                   *string            `json:"benefitType" validate:"omitempty,max=255"`
}

func NewSupplierEntity(props SupplierProps) (*SupplierEntity, error) {
	entity := &SupplierEntity{
		PartnerId:                     props.PartnerId,
		PartnerName:                   props.PartnerName,
		CustomerId:                    props.CustomerId,
		CustomerName:                  props.CustomerName,
		CustomerDomainName:            props.CustomerDomainName,
		CustomerCountry:               props.CustomerCountry,
		MpnId:                         props.MpnId,
		Tier2MpnId:                    props.Tier2MpnId,
		InvoiceNumber:                 props.InvoiceNumber,
		ProductId:                     props.ProductId,
		SkuId:                         props.SkuId,
		AvailabilityId:                props.AvailabilityId,
		SkuName:                       props.SkuName,
		ProductName:                   props.ProductName,
		PublisherName:                 props.PublisherName,
		PublisherId:                   props.PublisherId,
		SubscriptionDescription:       props.SubscriptionDescription,
		SubscriptionId:                props.SubscriptionId,
		ChargeStartDate:               props.ChargeStartDate,
		ChargeEndDate:                 props.ChargeEndDate,
		UsageDate:                     props.UsageDate,
		MeterType:                     props.MeterType,
		MeterCategory:                 props.MeterCategory,
		MeterId:                       props.MeterId,
		MeterSubCategory:              props.MeterSubCategory,
		MeterName:                     props.MeterName,
		MeterRegion:                   props.MeterRegion,
		Unit:                          props.Unit,
		ResourceLocation:              props.ResourceLocation,
		ConsumedService:               props.ConsumedService,
		ResourceGroup:                 props.ResourceGroup,
		ResourceURI:                   props.ResourceURI,
		ChargeType:                    props.ChargeType,
		UnitPrice:                     props.UnitPrice,
		Quantity:                      props.Quantity,
		UnitType:                      props.UnitType,
		BillingPreTaxTotal:            props.BillingPreTaxTotal,
		BillingCurrency:               props.BillingCurrency,
		PricingPreTaxTotal:            props.PricingPreTaxTotal,
		PricingCurrency:               props.PricingCurrency,
		ServiceInfo1:                  props.ServiceInfo1,
		ServiceInfo2:                  props.ServiceInfo2,
		Tags:                          props.Tags,
		AdditionalInfo:                props.AdditionalInfo,
		EffectiveUnitPrice:            props.EffectiveUnitPrice,
		PCToBCExchangeRate:            props.PCToBCExchangeRate,
		PCToBCExchangeRateDate:        props.PCToBCExchangeRateDate,
		EntitlementId:                 props.EntitlementId,
		EntitlementDescription:        props.EntitlementDescription,
		PartnerEarnedCreditPercentage: props.PartnerEarnedCreditPercentage,
		CreditPercentage:              props.CreditPercentage,
		CreditType:                    props.CreditType,
		BenefitOrderId:                props.BenefitOrderId,
		BenefitId:                     props.BenefitId,
		BenefitType:                   props.BenefitType,
	}

	entity.Init(props.ID, props.CreatedAt, props.UpdatedAt, props.DeletedAt)

	if err := validators.ValidatorFields(entity); err != nil {
		return nil, err
	}

	return entity, nil
}

func (s *SupplierEntity) Update(props SupplierProps) error {
	s.PartnerId = props.PartnerId
	s.PartnerName = props.PartnerName
	s.CustomerId = props.CustomerId
	s.CustomerName = props.CustomerName
	s.CustomerDomainName = props.CustomerDomainName
	s.CustomerCountry = props.CustomerCountry
	s.MpnId = props.MpnId
	s.Tier2MpnId = props.Tier2MpnId
	s.InvoiceNumber = props.InvoiceNumber
	s.ProductId = props.ProductId
	s.SkuId = props.SkuId
	s.AvailabilityId = props.AvailabilityId
	s.SkuName = props.SkuName
	s.ProductName = props.ProductName
	s.PublisherName = props.PublisherName
	s.PublisherId = props.PublisherId
	s.SubscriptionDescription = props.SubscriptionDescription
	s.SubscriptionId = props.SubscriptionId
	s.ChargeStartDate = props.ChargeStartDate
	s.ChargeEndDate = props.ChargeEndDate
	s.UsageDate = props.UsageDate
	s.MeterType = props.MeterType
	s.MeterCategory = props.MeterCategory
	s.MeterId = props.MeterId
	s.MeterSubCategory = props.MeterSubCategory
	s.MeterName = props.MeterName
	s.MeterRegion = props.MeterRegion
	s.Unit = props.Unit
	s.ResourceLocation = props.ResourceLocation
	s.ConsumedService = props.ConsumedService
	s.ResourceGroup = props.ResourceGroup
	s.ResourceURI = props.ResourceURI
	s.ChargeType = props.ChargeType
	s.UnitPrice = props.UnitPrice
	s.Quantity = props.Quantity
	s.UnitType = props.UnitType
	s.BillingPreTaxTotal = props.BillingPreTaxTotal
	s.BillingCurrency = props.BillingCurrency
	s.PricingPreTaxTotal = props.PricingPreTaxTotal
	s.PricingCurrency = props.PricingCurrency
	s.ServiceInfo1 = props.ServiceInfo1
	s.ServiceInfo2 = props.ServiceInfo2
	s.Tags = props.Tags
	s.AdditionalInfo = props.AdditionalInfo
	s.EffectiveUnitPrice = props.EffectiveUnitPrice
	s.PCToBCExchangeRate = props.PCToBCExchangeRate
	s.PCToBCExchangeRateDate = props.PCToBCExchangeRateDate
	s.EntitlementId = props.EntitlementId
	s.EntitlementDescription = props.EntitlementDescription
	s.PartnerEarnedCreditPercentage = props.PartnerEarnedCreditPercentage
	s.CreditPercentage = props.CreditPercentage
	s.CreditType = props.CreditType
	s.BenefitOrderId = props.BenefitOrderId
	s.BenefitId = props.BenefitId
	s.BenefitType = props.BenefitType
	s.UpdatedAt = time.Now()

	return validators.ValidatorFields(s)

}
