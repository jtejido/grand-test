# grand-test
Golang Rand's test wrapper for TestU01


This is a wrapper for L'Ecuyer's [TestU01](http://simul.iro.umontreal.ca/testu01/tu01.html).
In an ideal world, [they](https://en.wikipedia.org/wiki/TestU01) claim to have been able to run a BigCrush test at approximately 4 hours. But in my world, it was 10-11 hours (it varies on the Sources being tested, that, and I have a clunky hardware).


### Running the BigCrush
```
	wget http://simul.iro.umontreal.ca/testu01/TestU01.zip
	unzip TestU01.zip
	mv TestU01-*/ TestU01
	cd TestU01
	./configure
	make
	cd ..
	go build main.go
	./main >> results.txt
```