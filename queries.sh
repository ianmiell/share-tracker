#!/bin/bash
set -x
(
echo "================================================================================"
echo "Current price"
echo "================================================================================"
echo ".headers on
.mode column
select share.name as name, share_price.num_held as num_held, price/10000.00 * num_held as currvalue from share, share_price where share_price.code = share.code order by share.code; " | sqlite3 shares.db

echo "================================================================================"
echo "Shares bought"
echo "================================================================================"
echo ".headers on
.mode column
select (select sum(price / 10000.00 * number_of_shares) from share_tx where transaction_type = 'BUY') as shares_bought from share order by share.code limit 1;" | sqlite3 shares.db

echo "================================================================================"
echo "Shares sold"
echo "================================================================================"
echo ".headers on
.mode column
select (select sum(price / 10000.00 * number_of_shares) from share_tx where transaction_type = 'SELL') as shares_sold from share order by share.code limit 1;" | sqlite3 shares.db

echo "================================================================================"
echo "Shares held"
echo "================================================================================"
echo ".headers on
.mode column
select sum(price/10000.00 * num_held) as shares_held from share_price order by share_price.code limit 1;" | sqlite3 shares.db
) > reports/position.txt

(
echo "================================================================================"
echo "Monthly dividends"
echo "================================================================================"
echo ".headers on
select sum(dividend)/100.00, strftime('%Y-%m',date) as month from dividend group by 2 order by 2;" | sqlite3 shares.db
) > reports/monthly_dividends.txt

(
echo "================================================================================"
echo "Monthly dividends 12-month MA"
echo "================================================================================"
echo ".headers on
select
   coalesce((select round(sum(dividend)/100.00/12) from dividend d1 where d1.date <= d2.date and d1.date > date(d2.date,'-1 year')),0.0) as lastyrdivis,
   strftime('%Y-%m',date) as month
from dividend d2
group by 2
order by 2;" | sqlite3 shares.db
) > reports/monthly_dividends_12mth_ma.txt

./dividends.sh

(
echo "================================================================================"
echo "Profit per share"
echo "================================================================================"
echo ".headers on
.mode column
select
	code                                                                                                                                                                  as code,
	price/10000.00 * num_held                                                                                                                                             as curr_value,
    ifnull((select sum(dividend) / 100.00 from dividend where code = s.code), 0)                                                                                          as dividend,
    cast(printf('%.2f',ifnull((select sum(price / 10000.00 * number_of_shares) as profit from share_tx where transaction_type = 'SELL' and code = s.code), 0)) as double) as sales,
    cast(printf('%.2f',ifnull((select sum(price / 10000.00 * number_of_shares) from share_tx where transaction_type = 'BUY' and code = s.code), 0)) as double)            as buys,
	cast(printf('%.2f',(price/10000.00 * num_held + ifnull((select sum(dividend) / 100.00 from dividend where code = s.code), 0) + ifnull((select sum(price / 10000.00 * number_of_shares) as profit from share_tx where transaction_type = 'SELL' and code = s.code), 0) - ifnull((select sum(price / 10000.00 * number_of_shares) from share_tx where transaction_type = 'BUY' and code = s.code), 0))) as double)
                                                                                                                                                                         as profit,
	round((select julianday('now') from share limit 1) - julianday((select min(date) from share_tx s1 where s1.code = s.code and transaction_type = 'BUY')) - 0.5)       as days_held,
	cast(printf('%.2f',(price/10000.00 * num_held + ifnull((select sum(dividend) / 100.00 from dividend where code = s.code), 0) + ifnull((select sum(price / 10000.00 * number_of_shares) as profit from share_tx where transaction_type = 'SELL' and code = s.code), 0) - ifnull((select sum(price / 10000.00 * number_of_shares) from share_tx where transaction_type = 'BUY' and code = s.code), 0)) / (round((select julianday('now') from share limit 1) - julianday((select min(date) from share_tx s1 where s1.code = s.code and transaction_type = 'BUY')) - 0.5))) as double)
                                                                                                                                                                         as profit_per_day,
	(select count(*) from share_tx where code = s.code and transaction_type = 'SELL')                                                                                    as num_sales
from
	share_price s
order by
	8 desc,
	1 asc;" | sqlite3 shares.db
) > reports/profit_per_share.txt

(
echo "================================================================================"
echo "Overall profit"
echo "================================================================================"
echo ".headers on
.mode column
select
    ifnull((select sum(price/10000.00 * num_held) as currvalue from share_price),0) as current_value,
    ifnull((select sum(price/10000.00 * num_held) from share_price),0) +
      ifnull((select sum(price) / 100.00 from cash_tx),0) +
      ifnull((select sum(price / 10000.00 * number_of_shares) from share_tx where transaction_type = 'SELL'),0) +
      ifnull((select sum(dividend) / 100.00 from dividend),0) -
      ifnull((select sum(price / 10000.00 * number_of_shares) from share_tx where transaction_type = 'BUY'),0)
                                                                                    as profit
from
	share
order by
	share.code
limit 1;" | sqlite3 shares.db
) > reports/overall_profit.txt


## GRAPH DATA
# Get data for the 12 month rolling dividend graph
(
echo ".headers off
drop table if exists mthgraph;
create table mthgraph as select
   ifnull((select round(sum(dividend)/100.00/12) from dividend d1 where d1.date < d2.date and d1.date > date(d2.date,'-1 year')),0) as amount,
   strftime('%Y-%m',date) as month
from dividend d2
group by 2
order by 2;
select rowid, amount from mthgraph;" | sqlite3 -separator ' ' shares.db
) > graph_data/dividend_12mth.dat

echo 'drop table if exists mthgraph;' | sqlite3 shares.db

# Get data for the profits graph and add it iff it's different from yesterday, and today does not already exist
echo "insert into profit_history select
    strftime('%s','now') as time,
    ifnull((select sum(price/10000.00 * num_held) from share_price),0) +
      ifnull((select sum(price) / 100.00 from cash_tx),0) +
      ifnull((select sum(price / 10000.00 * number_of_shares) from share_tx where transaction_type = 'SELL'),0) +
      ifnull((select sum(dividend) / 100.00 from dividend),0) -
      ifnull((select sum(price / 10000.00 * number_of_shares) from share_tx where transaction_type = 'BUY'),0)
        as profit
from share
order by share.code
limit 1;" | sqlite3 -separator ' ' shares.db

(
echo ".headers off
select time, profit
from profit_history
order by 1;" | sqlite3 -separator ' ' shares.db
) > graph_data/profits.dat


# TODO: update counts of shares based on numbers bought and sold
#update share_price set num_held = (
#   select sum(
#       (
#          select number_of_shares from share_tx where s.code = code and transaction_type = 'BUY'
#       )
#       -
#       (
#          select number_of_shares from share_tx where s.code = code and transaction_type = 'SELL'
#       )
#   )
#   from share as s where code = share_price.code
#);
