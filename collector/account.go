package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"go.bmvs.io/ynab/api/account"
)

var (
	accountInfoDesc = prometheus.NewDesc(
		"ynab_account_info",
		"Information about the YNAB account",
		[]string{"budget_id", "account_id", "name", "type", "on_budget", "closed", "deleted"}, nil)

	accountBalanceDesc = prometheus.NewDesc(
		"ynab_account_balance",
		"The total balance of the given account.",
		[]string{"account_id"}, nil)

	accountClearedBalanceDesc = prometheus.NewDesc(
		"ynab_account_cleared_balance",
		"The total cleared balance of the given account.",
		[]string{"account_id"}, nil)
)

func newAccountInfoMetric(budgetID string, acc *account.Account) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		accountInfoDesc,
		prometheus.GaugeValue,
		1,
		budgetID, acc.ID, acc.Name, string(acc.Type), bool2str(acc.OnBudget), bool2str(acc.Closed), bool2str(acc.Deleted))
}

func newAccountBalanceMetric(acc *account.Account) prometheus.Metric {
	return prometheus.MustNewConstMetric(accountBalanceDesc, prometheus.GaugeValue,
		dollars(acc.Balance), acc.ID)
}

func newClearedAccountBalanceMetric(acc *account.Account) prometheus.Metric {
	return prometheus.MustNewConstMetric(accountClearedBalanceDesc, prometheus.GaugeValue,
		dollars(acc.ClearedBalance), acc.ID)
}
