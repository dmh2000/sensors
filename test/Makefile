files:=sin1010.csv sin1020.csv sin2010.csv sin2020.csv sin0505.csv


all: $(files) plot
	
sin1010.csv : FORCE
	python3 config.py 1.0 1.0 >sin1010.csv

sin1020.csv : FORCE
	python3 config.py 1.0 2.0 >sin1020.csv

sin2010.csv : FORCE
	python3 config.py 2.0 1.0 >sin2010.csv

sin2020.csv : FORCE
	python3 config.py 2.0 2.0 >sin2020.csv

sin0505.csv : FORCE
	python3 config.py 0.5 0.5 >sin0505.csv

plot: 
	python3 plot.py $(files)

clean:
	rm *.csv *.csv
	
FORCE:	