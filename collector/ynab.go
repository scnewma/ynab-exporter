package collector

import (
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"go.bmvs.io/ynab"
	"go.bmvs.io/ynab/api/budget"
)

var (
	upDesc = prometheus.NewDesc(
		"ynab_up",
		"Whether the scrape succeeded.",
		nil, nil)
)

// NewYNABCollector creates a new ynabCollector given a ynab client.
// This function will panic if the provided client is nil.
func NewYNABCollector(client ynab.ClientServicer) prometheus.Collector {
	if client == nil {
		panic("client must be provided")
	}

	return ynabCollector{client}
}

type ynabCollector struct {
	client ynab.ClientServicer
}

func (c ynabCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c ynabCollector) Collect(ch chan<- prometheus.Metric) {
	budgets, err := c.getBudgets()
	if err != nil {
		log.WithError(err).Warn("failed to get budgets")
		ch <- newUpMetric(false)
		return
	}

	for _, budget := range budgets {
		ch <- newBudgetInfoMetric(budget)

		for _, account := range budget.Accounts {
			ch <- newAccountInfoMetric(budget.ID, account)
			ch <- newAccountBalanceMetric(account)
			ch <- newClearedAccountBalanceMetric(account)
		}

		for _, category := range budget.Categories {
			ch <- newCategoryInfoMetric(budget.ID, category)
			ch <- newCategoryBudgetedMetric(category)
			ch <- newCategoryActivityMetric(category)
			ch <- newCategoryBalanceMetric(category)
		}

		for _, group := range budget.CategoryGroups {
			ch <- newCategoryGroupInfoMetric(budget.ID, group)
		}
	}

	rateLimit := c.client.RateLimit()
	ch <- newRateLimitUsedMetric(rateLimit)
	ch <- newRateLimitTotalMetric(rateLimit)

	ch <- newUpMetric(true)
}

func (c ynabCollector) getBudgets() ([]*budget.Budget, error) {
	summaries, err := c.client.Budget().GetBudgets()
	if err != nil {
		return nil, errors.Wrapf(err, "could not list budgets")
	}

	var budgets []*budget.Budget
	for _, summary := range summaries {
		budget, err := c.client.Budget().GetBudget(summary.ID, nil)
		if err != nil {
			return nil, errors.Wrapf(err, "could not get budget with id %q", summary.ID)
		}
		budgets = append(budgets, budget.Budget)
	}
	return budgets, nil
}

func newUpMetric(up bool) prometheus.Metric {
	val := 0
	if up {
		val = 1
	}

	return prometheus.MustNewConstMetric(upDesc, prometheus.GaugeValue, float64(val))
}
