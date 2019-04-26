Serverless Share Dealing Database
=================================

A simple share tracking system using:

- Go (to build binary to construct db from source)
- SQLite3 (to provide SQL interface to data)
- make (to manage rebuilds)
- Shell scripts (to run SQL queries)
- ShutIt (for shell automation of graphs)
- gnuplot (for ascii graphs)

Why?
====

Google finance is shutting down late 2017.

All the free alternatives didn't do exactly what I wanted (though looked nice).
So I figured it was an opportunity to put into practice the model of serverless
development using git and make et al outlined here:

https://zwischenzugs.wordpress.com/2017/08/07/a-non-cloud-serverless-application-pattern-using-git-and-docker/

How it works
============

Transactions are recorded in transactions.go, in a simple format.

Go code turns the transactions into a transient database using SQLite3 bindings, and retrieves the latest share price data from yahoo (share.go).

The database is queried by a script (queries.sh), and text-only reports on dividends, profitability and overall position are placed in reports/

Ascii graphs (see below) are created using gnnuplot (graphs.py), driven by an automation tool (ShutIt), and deposited in graphs/

Running make builds and runs all of the above, and commits and pushes updates to the git repo.

Here is a run in action, with a view of the graphs and reports:

[![asciicast](https://asciinema.org/a/wjmqdqJJASzomRWsKUvArNDli.png)](https://asciinema.org/a/wjmqdqJJASzomRWsKUvArNDli)

SQLite DB
=========

SQLite database contains shares info

Tables:

share - each share owned

share_price - price of share

share_tx - share BUY or SELL (FK to share)

dividend - share dividend (FK to share)

Data is stored in git in db_export.sql and then imported afresh each time.


Adding a new share
==================

To add a new share:

- Add a row with the share code in db.go, eg to add Centrica, you would add the CNA.L line as per below:

```
func insert_shares() {
    insert_share("BATS.L","British American Tobacco","")
    insert_share("CNA.L","Centrica","")
```

- Add an entry in transactions.go representing the transaction:

```
    date, _ = time.Parse(layout,"19/09/2017")
    my_transactions = append(my_transactions,transaction{code: "CNA.L",  price: 19050, number_of_shares: 1000, transaction_type: "BUY", transaction_date: date, notes: ""})
```

- Run make

Adding a sale
=============

- Add an entry in transactions.go representing the transaction:

```
    date, _ = time.Parse(layout,"19/09/2017")
    my_transactions = append(my_transactions,transaction{code: "CNA.L",  price: 19050, number_of_shares: 1000, transaction_type: "SELL", transaction_date: date, notes: ""})
```

NOTE prices are in hundredths of a pence/cent.

Adding a dividend
=================

- Add an entry in transactions.go representing the dividend:

```
    date, _ = time.Parse(layout,"19/09/2017")
    my_transactions = append(my_transactions,transaction{code: "CNA.L",  price: 19050, number_of_shares: 1000, transaction_type: "BUY", transaction_date: date, notes: ""})
```

NOTE prices are in hundredths of a pence/cent.

Adding an interest payment
==========================

- Add an entry in transactions.go representing the interest payment:

```
    date, _ = time.Parse(layout,"19/09/2017")
    my_transactions = append(my_transactions,transaction{code: "INTEREST", price: 1227,    number_of_shares: 0, transaction_type: "INTEREST", transaction_date: date, notes: ""})
```

NOTE prices are in pence/cent.

Updating the record
===================

Run:

```
make
```

Graphs
======

Graphs are placed in the 'graphs' folder. They look like this:

```
  143500 +-+-----------+-------------+-------------+-------------+-------------+-------------+-------------+-------------+-------------+-------------+-------------+-----------+-+
         +                           +                           +                           +                           +                           +                           +
         |                                                 *                                                                       '../graph_data/profits.dat' using 1:2 ******* |
         |                                                * *                                                                                        f(x) = 0.00x + 6.42 ####### |
  143000 +-+                                             *  *                                                                                                                  +-+
         |                                              *    *                                                                                                                   |
         |                                              *     *                                                                                                                  |
         |                                             *      *                                                                                                                  |
         |                                            *        *                                                                                                             *   |
  142500 +-+                                     *   *         *                                                                                                             * +-+
         |                                       ** *          *                                                                                                             *   |
         |                                      * * *           *                                                                                                            *   |
         |                **************        *  *            *                                                                                                            *   |
  142000 +-+       *******              *****   *                *                                                                                                           * +-+
         |                                   ***                 *                                                                                                         ***   |
         |                                     *                 *                                                                                                         *     |
         |                                                        *                                                                                                       *      |
         |                                                        *                                                                                                       *      |
  141500 +-+                                                       *                                                                                                      *    +-+
         |                                                         *                                                                                                      *      |
         |                                                         *                                                                                                     *       |
         |                                                          *                                                              *                                ******       |
  141000 +-+                                                        *                                                             **                            ****           +-+
         |                                                           *                                                            * *                          *                 |
         |                                                           *     ***            #######################################*##*##########################*##############   |
         |         ##################################################*#####*##*###########                                      *   *                          *                 |
         |                                                            *    *  *                                                 *   *                         *                  |
  140500 +-+                                                          *   *    *                             ***     *         *     *                        *                +-+
         |                                                             *  *    *              **             *  **  * *      **      *                        *                  |
         |                                                             *  *     *            *  *            *    * * *     *         *                      *                   |
         |                                                             *  *     *           *    *          *      *   *  **          *                      *                   |
         |                                                              * *     *          *      *         *          ***             *                     *                   |
  140000 +-+                                                            * *     *         *        *        *                          *                    *                  +-+
         |                                                              **       *        *         *       *                           *                   *                    |
         |                                                               *       *       *           *     *                            *                   *                    |
         |                                                               *       *       *            * ****                             *                 *                     |
  139500 +-+                                                                      *     *              *                                 ********     ******                   +-+
         |                                                                         *    *                                                        *****                           |
         |                                                                          *   *                                                                                        |
         |                                                                           *  *                                                                                        |
  139000 +-+                                                                          **                                                                                       +-+
         |                                                                             *                                                                                         |
         |                                                                                                                                                                       |
         |                                                                                                                                                                       |
         +                           +                           +                           +                           +                           +                           +
  138500 +-+-----------+-------------+-------------+-------------+-------------+-------------+-------------+-------------+-------------+-------------+-------------+-----------+-+
       07/13                       07/27                       08/10                       08/24                       09/07                       09/21                       10/05
```

Reports
=======

Reports are placed in the reports folder. They include:

- dividends.txt - a report on dividends
- position.txt  - a report on your overall current position
- profits.txt   - a report on your profits

