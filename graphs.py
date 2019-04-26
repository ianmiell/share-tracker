import shutit

# Generates graphs of data stored within graph_data.

gnuplot_expect='gnuplot> '

s = shutit.create_session('bash',loglevel='info')

s.send('cd /space/git/shares/graphs')
s.send('gnuplot',expect=gnuplot_expect)
s.send('set terminal png 5120,3840',expect=gnuplot_expect)
s.send('set yrange [0:]',expect=gnuplot_expect)


#### FOR BEST FIT BEGIN
# Define the function a*x+b
s.send("f(x) = a*x + b",expect=gnuplot_expect)
# Set the title function based on a+b
s.send("title_f(a,b) = sprintf('f(x) = %.2fx + %.2f', a, b)",expect=gnuplot_expect)
#### FOR BEST FIT DONE

# Fit the function to the graph
#s.send('set timefmt "%Y-%m"',expect=gnuplot_expect)
s.send("fit f(x) '../graph_data/dividend_12mth.dat' using 1:2 via a, b",expect=gnuplot_expect)
s.send("set output 'dividends.png'",expect=gnuplot_expect)
s.send("plot '../graph_data/dividend_12mth.dat' using 1:2 with lines, f(x) title title_f(a,b)",expect=gnuplot_expect)

s.send('set xdata time',expect=gnuplot_expect)
s.send('set key left',expect=gnuplot_expect)
s.send('set timefmt "%s"',expect=gnuplot_expect)
s.send('set label "WMH+"  at 1334823036,2000 point',expect=gnuplot_expect)
# Fit the function to the graph
s.send("fit f(x) '../graph_data/profits.dat' using 1:2 via a, b",expect=gnuplot_expect)
s.send("set output 'profits.png'",expect=gnuplot_expect)

# Plot the graph with two lines
s.send("plot '../graph_data/profits.dat' using 1:2 with lines, f(x) title title_f(a,b)",expect=gnuplot_expect)
