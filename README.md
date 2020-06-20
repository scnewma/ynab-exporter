# YNAB Exporter

[![Go Report Card](https://goreportcard.com/badge/github.com/scnewma/ynab-exporter)](https://goreportcard.com/report/github.com/scnewma/ynab-exporter)

Prometheus exporter for YNAB budgets.

## Metrics Collected

| Name                         | Labels                                                           |
|------------------------------|------------------------------------------------------------------|
| ynab_budget_info             | budget_id, budget_name                                           |
| ynab_account_info            | budget_id, account_id, closed, deleted, name, on_budget, type    |
| ynab_category_info           | budget_id, category_group_id, category_id, deleted, hidden, name |
| ynab_category_group_info     | budget_id, category_group_id, deleted, hidden, name              |
| ynab_account_balance         | account_id                                                       |
| ynab_account_cleared_balance | account_id                                                       |
| ynab_category_activity       | category_id                                                      |
| ynab_category_balance        | category_id                                                      |
| ynab_category_budgeted       | category_id                                                      |
| ynab_up                      |                                                                  |
| ynab_ratelimit_total         |                                                                  |
| ynab_ratelimit_used          |                                                                  |

## Building and Running

Prerequisites:

* Go
* [YNAB API Token](https://api.youneedabudget.com/#personal-access-tokens)

Building:

```
make
./ynab-exporter --ynab.access-token=[TOKEN]
```

## Running Tests

```
make test
```

## Using Docker

```
# put YNAB token in .token file then
make docker-run

# OR

docker run -p 9721:9721 -d scnewma/ynab-exporter --ynab.access-token=[TOKEN]
```
