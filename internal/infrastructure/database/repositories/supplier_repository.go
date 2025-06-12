package repositories

import (
	"context"
	"fmt"

	"github.com/davigomesdev/reconfile/internal/domain/contracts"
	"github.com/davigomesdev/reconfile/internal/domain/entities"
	"github.com/davigomesdev/reconfile/internal/domain/errors"
	domain_repositories "github.com/davigomesdev/reconfile/internal/domain/repositories"
	"github.com/jackc/pgx/pgtype"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	jsoniter "github.com/json-iterator/go"
	"github.com/lib/pq"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type supplierRepository struct {
	pool *pgxpool.Pool
}

func NewSupplierRepository(pool *pgxpool.Pool) domain_repositories.SupplierRepositoryInterface {
	return &supplierRepository{pool: pool}
}

func (r *supplierRepository) toEntity(scanner interface {
	Scan(dest ...interface{}) error
}) (*entities.SupplierEntity, error) {
	entity := &entities.SupplierEntity{}

	var tagsRaw []byte
	var additionalInfoRaw []byte
	var deletedAt pgtype.Timestamptz

	if err := scanner.Scan(
		&entity.ID,
		&entity.PartnerId,
		&entity.PartnerName,
		&entity.CustomerId,
		&entity.CustomerName,
		&entity.CustomerDomainName,
		&entity.CustomerCountry,
		&entity.MpnId,
		&entity.Tier2MpnId,
		&entity.InvoiceNumber,
		&entity.ProductId,
		&entity.SkuId,
		&entity.AvailabilityId,
		&entity.SkuName,
		&entity.ProductName,
		&entity.PublisherName,
		&entity.PublisherId,
		&entity.SubscriptionDescription,
		&entity.SubscriptionId,
		&entity.ChargeStartDate,
		&entity.ChargeEndDate,
		&entity.UsageDate,
		&entity.MeterType,
		&entity.MeterCategory,
		&entity.MeterId,
		&entity.MeterSubCategory,
		&entity.MeterName,
		&entity.MeterRegion,
		&entity.Unit,
		&entity.ResourceLocation,
		&entity.ConsumedService,
		&entity.ResourceGroup,
		&entity.ResourceURI,
		&entity.ChargeType,
		&entity.UnitPrice,
		&entity.Quantity,
		&entity.UnitType,
		&entity.BillingPreTaxTotal,
		&entity.BillingCurrency,
		&entity.PricingPreTaxTotal,
		&entity.PricingCurrency,
		&entity.ServiceInfo1,
		&entity.ServiceInfo2,
		&tagsRaw,
		&additionalInfoRaw,
		&entity.EffectiveUnitPrice,
		&entity.PCToBCExchangeRate,
		&entity.PCToBCExchangeRateDate,
		&entity.EntitlementId,
		&entity.EntitlementDescription,
		&entity.PartnerEarnedCreditPercentage,
		&entity.CreditPercentage,
		&entity.CreditType,
		&entity.BenefitOrderId,
		&entity.BenefitId,
		&entity.BenefitType,
		&entity.CreatedAt,
		&entity.UpdatedAt,
		&deletedAt,
	); err != nil {
		return nil, err

	}

	if len(tagsRaw) > 0 {
		if err := json.Unmarshal(tagsRaw, &entity.Tags); err != nil {
			return nil, err
		}
	}

	if len(additionalInfoRaw) > 0 {
		if err := json.Unmarshal(additionalInfoRaw, &entity.AdditionalInfo); err != nil {
			return nil, err
		}
	}

	if deletedAt.Status == pgtype.Present {
		entity.DeletedAt = &deletedAt.Time
	}

	return entity, nil
}

func (r *supplierRepository) Get(ctx context.Context, id string) (*entities.SupplierEntity, error) {
	entity, err := r.toEntity(r.pool.QueryRow(ctx,
		`SELECT * FROM suppliers WHERE id = $1 AND deleted_at IS NULL`, id,
	))
	if err != nil {
		return nil, errors.NewNotFoundError("Fornecedor não encontrado.")
	}

	return entity, nil
}

func (r *supplierRepository) GetAll(ctx context.Context) ([]*entities.SupplierEntity, error) {
	rows, err := r.pool.Query(ctx, `SELECT * FROM suppliers WHERE deleted_at IS NULL`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entities []*entities.SupplierEntity
	for rows.Next() {
		e, err := r.toEntity(rows)
		if err != nil {
			return nil, err
		}

		entities = append(entities, e)
	}
	return entities, nil
}

func (r *supplierRepository) GetOverview(ctx context.Context) (*contracts.SupplierOverview, error) {
	ov := &contracts.SupplierOverview{}

	row := r.pool.QueryRow(ctx, `
        SELECT
            COUNT(*) AS total_records,
            SUM(billing_pre_tax_total) AS total_billing,
            COUNT(DISTINCT subscription_id) AS total_subscribers,
            COUNT(DISTINCT customer_id) AS total_customers
        FROM suppliers
        WHERE deleted_at IS NULL;
    `)
	if err := row.Scan(&ov.TotalRecords, &ov.TotalBilling, &ov.TotalSubscribers, &ov.TotalCustomers); err != nil {
		return nil, fmt.Errorf("erro ao obter overview geral: %w", err)
	}

	rows, err := r.pool.Query(ctx, `
        SELECT TO_CHAR(usage_date, 'YYYY-MM') AS year_month,
               SUM(billing_pre_tax_total)::float8 AS total
        FROM suppliers
        WHERE deleted_at IS NULL
        GROUP BY year_month
        ORDER BY year_month;
    `)
	if err != nil {
		return nil, errors.NewBadRequestError("Erro ao obter faturamento por mês")
	}
	defer rows.Close()

	ov.BillingByMonth = []contracts.MonthBilling{}
	for rows.Next() {
		var mb contracts.MonthBilling
		if err := rows.Scan(&mb.YearMonth, &mb.Total); err != nil {
			return nil, errors.NewBadRequestError("Erro no scan do faturamento por mês")
		}
		ov.BillingByMonth = append(ov.BillingByMonth, mb)
	}

	return ov, nil
}

func (r *supplierRepository) Search(ctx context.Context, input *contracts.SearchInput) (*contracts.SearchResult[*entities.SupplierEntity], error) {
	page, perPage := 1, 15
	if input.Page != nil && *input.Page > 0 {
		page = *input.Page
	}
	if input.PerPage != nil && *input.PerPage > 0 {
		perPage = *input.PerPage
	}
	offset := (page - 1) * perPage

	baseQuery := "FROM suppliers WHERE deleted_at IS NULL"
	whereClause := ""
	args := []interface{}{}
	argPos := 1

	if input.Filter != nil && *input.Filter != "" {
		whereClause = fmt.Sprintf(" AND customer_name ILIKE $%d", argPos)
		args = append(args, "%"+*input.Filter+"%")
		argPos++
	}

	query := "SELECT * " + baseQuery + whereClause

	if input.Sort != nil && *input.Sort != "" {
		dir := "ASC"
		if input.SortDir != nil && *input.SortDir == contracts.SortDesc {
			dir = "DESC"
		}
		query += fmt.Sprintf(" ORDER BY %s %s", pq.QuoteIdentifier(*input.Sort), dir)
	}

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argPos, argPos+1)
	args = append(args, perPage, offset)

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	var items []*entities.SupplierEntity
	for rows.Next() {
		e, err := r.toEntity(rows)
		if err != nil {
			return nil, fmt.Errorf("mapping error: %w", err)
		}
		items = append(items, e)
	}

	countQuery := "SELECT COUNT(*) " + baseQuery + whereClause
	var total int
	if err := r.pool.QueryRow(ctx, countQuery, args[:len(args)-2]...).Scan(&total); err != nil {
		return nil, fmt.Errorf("count error: %w", err)
	}

	return &contracts.SearchResult[*entities.SupplierEntity]{
		Items:       items,
		Total:       total,
		CurrentPage: page,
		PerPage:     perPage,
		Sort:        input.Sort,
		SortDir:     input.SortDir,
		Filter:      input.Filter,
	}, nil
}

func (r *supplierRepository) CreateMany(ctx context.Context, suppliers []*entities.SupplierEntity) error {
	rows := make([][]interface{}, len(suppliers))

	for i, s := range suppliers {
		tagsJSON, _ := json.Marshal(s.Tags)
		additionalInfoJSON, _ := json.Marshal(s.AdditionalInfo)

		rows[i] = []interface{}{
			s.ID, s.PartnerId, s.PartnerName, s.CustomerId, s.CustomerName,
			s.CustomerDomainName, s.CustomerCountry, s.MpnId, s.Tier2MpnId, s.InvoiceNumber,
			s.ProductId, s.SkuId, s.AvailabilityId, s.SkuName, s.ProductName,
			s.PublisherName, s.PublisherId, s.SubscriptionDescription, s.SubscriptionId,
			s.ChargeStartDate, s.ChargeEndDate, s.UsageDate, s.MeterType, s.MeterCategory,
			s.MeterId, s.MeterSubCategory, s.MeterName, s.MeterRegion, s.Unit,
			s.ResourceLocation, s.ConsumedService, s.ResourceGroup, s.ResourceURI,
			s.ChargeType, s.UnitPrice, s.Quantity, s.UnitType, s.BillingPreTaxTotal,
			s.BillingCurrency, s.PricingPreTaxTotal, s.PricingCurrency, s.ServiceInfo1,
			s.ServiceInfo2, tagsJSON, additionalInfoJSON, s.EffectiveUnitPrice, s.PCToBCExchangeRate,
			s.PCToBCExchangeRateDate, s.EntitlementId, s.EntitlementDescription,
			s.PartnerEarnedCreditPercentage, s.CreditPercentage, s.CreditType,
			s.BenefitOrderId, s.BenefitId, s.BenefitType, s.CreatedAt, s.UpdatedAt,
		}
	}

	columns := []string{
		"id", "partner_id", "partner_name", "customer_id", "customer_name",
		"customer_domain_name", "customer_country", "mpn_id", "tier2_mpn_id", "invoice_number",
		"product_id", "sku_id", "availability_id", "sku_name", "product_name",
		"publisher_name", "publisher_id", "subscription_description", "subscription_id",
		"charge_start_date", "charge_end_date", "usage_date", "meter_type", "meter_category",
		"meter_id", "meter_sub_category", "meter_name", "meter_region", "unit",
		"resource_location", "consumed_service", "resource_group", "resource_uri",
		"charge_type", "unit_price", "quantity", "unit_type", "billing_pre_tax_total",
		"billing_currency", "pricing_pre_tax_total", "pricing_currency", "service_info1",
		"service_info2", "tags", "additional_info", "effective_unit_price", "pc_to_bc_exchange_rate",
		"pc_to_bc_exchange_rate_date", "entitlement_id", "entitlement_description",
		"partner_earned_credit_percentage", "credit_percentage", "credit_type",
		"benefit_order_id", "benefit_id", "benefit_type", "created_at", "updated_at",
	}

	_, err := r.pool.CopyFrom(ctx, pgx.Identifier{"suppliers"}, columns, pgx.CopyFromRows(rows))
	return err
}
