package contracts

type SupplierOverview struct {
	TotalRecords     int64          `json:"totalRecords"`
	TotalBilling     float64        `json:"totalBilling"`
	TotalSubscribers int64          `json:"totalSubscribers"`
	TotalCustomers   int64          `json:"totalCustomers"`
	BillingByMonth   []MonthBilling `json:"billingByMonth"`
}

type MonthBilling struct {
	YearMonth string  `json:"yearMonth"`
	Total     float64 `json:"total"`
}
