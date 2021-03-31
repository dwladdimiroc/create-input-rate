set term postscript eps color blacktext "Helvetica" 26 enhanced
set output "poisson.eps" 
set size nosquare 1.15,1
set key top right
set xlabel 'Time (s)'
set xrange [0:4000]
set xtics 500
set ylabel "Samples"
set yrange [0:8000]
set ytics 2000 nomirror
set datafile separator ","
plot "poisson.csv" using 1:2 title 'Poisson' with linespoints pi 25 lw 6
exit