(
echo "================================================================================"
echo "Dividend payouts coming up"
echo "================================================================================"
AGENT='Mozilla/5.0 (iPhone; U; CPU iPhone OS 4_3_3 like Mac OS X; en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5'
LINK='http://www.dividenddata.co.uk/dividend-payment-dates.py?m=alldividends'
curl -L -s -A "${AGENT}" "${LINK}" | html2text  | egrep -w "($(grep -w share db_export.sql | grep INSERT | sed 's/.*VALUES..\([A-Z]*\).*/\1/g' | tr '\n' '|' | sed 's/|$//'))"
echo "================================================================================"
echo "Ex-dividends coming up"
echo "================================================================================"
AGENT='Mozilla/5.0 (iPhone; U; CPU iPhone OS 4_3_3 like Mac OS X; en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5'
LINK='http://www.dividenddata.co.uk/exdividenddate.py?m=all_dividends&fm=&sc=xd&st=0'
sleep 10
curl -L -s -A "${AGENT}" "${LINK}" | html2text > out
curl -L -s -A "${AGENT}" "${LINK}" | html2text | cut -b 1-54,62-99 | egrep -w "($(grep -w share db_export.sql | grep INSERT | sed 's/.*VALUES..\([A-Z]*\).*/\1/g' | tr '\n' '|' | sed 's/|$//'))"
) > reports/dividends_coming_up.txt
