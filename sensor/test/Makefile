images:=sin1010.jpg sin1020.jpg sin2010.jpg sin2020.jpg sin0505.jpg


all: $(images)
	
sin1010.jpg : FORCE
	python3 config.py 1.0 1.0 >sin1010.csv
	python3 plot.py sin1010.csv sin1010.jpg
	xdg-open sin1010.jpg

sin1020.jpg : FORCE
	python3 config.py 1.0 2.0 >sin1020.csv
	python3 plot.py sin1020.csv sin1020.jpg
	xdg-open sin1020.jpg


sin2010.jpg : FORCE
	python3 config.py 2.0 1.0 >sin2010.csv
	python3 plot.py sin2010.csv sin2010.jpg
	xdg-open sin2010.jpg

sin2020.jpg : FORCE
	python3 config.py 2.0 2.0 >sin2010.csv
	python3 plot.py sin2010.csv sin2010.jpg
	xdg-open sin2010.jpg

sin0505.jpg : FORCE
	python3 config.py 0.5 0.5 >sin0505.csv
	python3 plot.py sin0505.csv sin0505.jpg
	xdg-open sin0505.jpg

FORCE:	