set term postscript eps color blacktext "Helvetica" 26 enhanced
set output "dist.eps" 
set size nosquare 1.15,1
set key top right
set xlabel 'Time (s)'
set xrange [0:4000]
set xtics 500
set ylabel "Samples"
set yrange [0:8000]
set ytics 2000 nomirror
set datafile separator ","
plot "norm.csv" using 1:2 title 'Normal' with linespoints pi 25 lw 6, \
	"n-norm.csv" using 1:2 title 'N-Normal' with linespoints pi 25 lw 6, \
	"logarithm.csv" using 1:2 title 'Logarithm' with linespoints pi 25 lw 6, \
	"exponential.csv" using 1:2 title 'Exponential' with linespoints pi 25 lw 6, \
	"poisson.csv" using 1:2 title 'Poisson' with linespoints pi 25 lw 6, \
	"lineal.csv" using 1:2 title 'Lineal' with linespoints pi 25 lw 6
exit