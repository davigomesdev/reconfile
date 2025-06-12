package supplier

import (
	"context"
	"fmt"
	"runtime"
	"sync"

	supplierDTO "github.com/davigomesdev/reconfile/internal/application/dtos/supplier"
	"github.com/davigomesdev/reconfile/internal/application/providers"
	"github.com/davigomesdev/reconfile/internal/domain/entities"
	"github.com/davigomesdev/reconfile/internal/domain/errors"
	"github.com/davigomesdev/reconfile/internal/domain/repositories"
	"github.com/davigomesdev/reconfile/pkg/utils"
	"github.com/go-playground/validator/v10"
)

const maxBatchSize = 1000

type ImportSuppliersUseCase struct {
	supplierRepository repositories.SupplierRepositoryInterface
	xlsxParserProvider *providers.XLSXParserProvider
}

func NewImportSuppliersUseCase(
	supplierRepository repositories.SupplierRepositoryInterface,
	xlsxParserProvider *providers.XLSXParserProvider,
) *ImportSuppliersUseCase {
	return &ImportSuppliersUseCase{
		supplierRepository: supplierRepository,
		xlsxParserProvider: xlsxParserProvider,
	}
}

func (uc *ImportSuppliersUseCase) Execute(ctx context.Context, input supplierDTO.ImportSuppliersDTO) error {
	f, err := input.File.Open()
	if err != nil {
		return errors.NewBadRequestError("Falha ao abrir o arquivo.")
	}
	defer f.Close()

	rows, err := uc.xlsxParserProvider.ParseXLSXRows(f, input.File.Size)
	if err != nil {
		return errors.NewBadRequestError("Falha ao analisar o XML da planilha.")
	}
	if len(rows) <= 1 {
		return errors.NewBadRequestError("Planilha vazia ou sem dados.")
	}

	workerCount := runtime.NumCPU()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var wg sync.WaitGroup
	errChan := make(chan error, 1)
	rowsChan := make(chan []string, workerCount*maxBatchSize)
	sharedPool := sync.Pool{New: func() interface{} { return make([]string, providers.TotalColumns) }}

	worker := func() {
		defer wg.Done()
		batch := make([]*entities.SupplierEntity, 0, maxBatchSize)
		defer func() {
			if len(batch) > 0 {
				if err := uc.supplierRepository.CreateMany(ctx, batch); err != nil {
					select {
					case errChan <- err:
					default:
					}
				}
			}
		}()

		for row := range rowsChan {
			select {
			case <-ctx.Done():
				return
			default:
				ent, err := entities.NewSupplierEntity(entities.SupplierProps{
					PartnerId:                     utils.Safe(row, 0),
					PartnerName:                   utils.Safe(row, 1),
					CustomerId:                    utils.Safe(row, 2),
					CustomerName:                  utils.Safe(row, 3),
					CustomerDomainName:            utils.Safe(row, 4),
					CustomerCountry:               utils.Safe(row, 5),
					MpnId:                         utils.ToInt(utils.Safe(row, 6)),
					Tier2MpnId:                    utils.ToInt(utils.Safe(row, 7)),
					InvoiceNumber:                 utils.Safe(row, 8),
					ProductId:                     utils.Safe(row, 9),
					SkuId:                         utils.Safe(row, 10),
					AvailabilityId:                utils.Safe(row, 11),
					SkuName:                       utils.Safe(row, 12),
					ProductName:                   utils.Safe(row, 13),
					PublisherName:                 utils.Safe(row, 14),
					PublisherId:                   utils.ToStringPtr(utils.Safe(row, 15)),
					SubscriptionDescription:       utils.ToStringPtr(utils.Safe(row, 16)),
					SubscriptionId:                utils.Safe(row, 17),
					ChargeStartDate:               utils.ToDate(utils.Safe(row, 18)),
					ChargeEndDate:                 utils.ToDate(utils.Safe(row, 19)),
					UsageDate:                     utils.ToDate(utils.Safe(row, 20)),
					MeterType:                     utils.Safe(row, 21),
					MeterCategory:                 utils.Safe(row, 22),
					MeterId:                       utils.Safe(row, 23),
					MeterSubCategory:              utils.Safe(row, 24),
					MeterName:                     utils.Safe(row, 25),
					MeterRegion:                   utils.ToStringPtr(utils.Safe(row, 26)),
					Unit:                          utils.Safe(row, 27),
					ResourceLocation:              utils.Safe(row, 28),
					ConsumedService:               utils.Safe(row, 29),
					ResourceGroup:                 utils.Safe(row, 30),
					ResourceURI:                   utils.Safe(row, 31),
					ChargeType:                    utils.Safe(row, 32),
					UnitPrice:                     utils.ToFloat(utils.Safe(row, 33)),
					Quantity:                      utils.ToFloat(utils.Safe(row, 34)),
					UnitType:                      utils.Safe(row, 35),
					BillingPreTaxTotal:            utils.ToFloat(utils.Safe(row, 36)),
					BillingCurrency:               utils.Safe(row, 37),
					PricingPreTaxTotal:            utils.ToFloat(utils.Safe(row, 38)),
					PricingCurrency:               utils.Safe(row, 39),
					ServiceInfo1:                  utils.ToStringPtr(utils.Safe(row, 40)),
					ServiceInfo2:                  utils.ToStringPtr(utils.Safe(row, 41)),
					Tags:                          utils.ToMapPtr(utils.Safe(row, 42)),
					AdditionalInfo:                utils.ToMapPtr(utils.Safe(row, 43)),
					EffectiveUnitPrice:            utils.ToFloat(utils.Safe(row, 44)),
					PCToBCExchangeRate:            utils.ToInt(utils.Safe(row, 45)),
					PCToBCExchangeRateDate:        utils.ToDate(utils.Safe(row, 46)),
					EntitlementId:                 utils.Safe(row, 47),
					EntitlementDescription:        utils.Safe(row, 48),
					PartnerEarnedCreditPercentage: utils.ToInt(utils.Safe(row, 49)),
					CreditPercentage:              utils.ToInt(utils.Safe(row, 50)),
					CreditType:                    utils.Safe(row, 51),
					BenefitOrderId:                utils.ToStringPtr(utils.Safe(row, 52)),
					BenefitId:                     utils.ToStringPtr(utils.Safe(row, 53)),
					BenefitType:                   utils.ToStringPtr(utils.Safe(row, 54)),
				})
				if err != nil {
					select {
					case errChan <- err:
					default:
					}
					continue
				}

				batch = append(batch, ent)
				if len(batch) >= maxBatchSize {
					if err := uc.supplierRepository.CreateMany(ctx, batch); err != nil {
						select {
						case errChan <- err:
						default:
						}
					}
					batch = batch[:0]
				}
				row := sharedPool.Get().([]string)
				defer sharedPool.Put(row[:0])
			}
		}
	}

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker()
	}

	go func() {
		defer close(rowsChan)
		for i, row := range rows {
			if i == 0 {
				continue
			}
			select {
			case <-ctx.Done():
				return
			case rowsChan <- row:
			}
		}
	}()

	wg.Wait()
	select {
	case err := <-errChan:
		if ve, ok := err.(validator.ValidationErrors); ok {
			return errors.NewBadRequestError(fmt.Sprintf("Erro de validação no campo: '%s', rule %s", ve[0].Field(), ve[0].Tag()))
		}
		return errors.NewBadRequestError(fmt.Sprintf("Erro de importação: %v", err))
	default:
		return nil
	}
}
