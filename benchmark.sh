mkdir -p result

go test -bench=BenchmarkAES128SBox . -benchtime=10000000x -cpu=1 -memprofile=./result/mem128sbox.out -cpuprofile=./result/cpu128sbox.out -o=result/pprof128sbox.bin
go tool pprof --png result/pprof128sbox.bin result/mem128sbox.out > result/mem128sbox.png
go tool pprof --png result/pprof128sbox.bin result/cpu128sbox.out > result/cpu128sbox.png

go test -bench=BenchmarkAES128TTable . -benchtime=10000000x -cpu=1 -memprofile=./result/mem128ttable.out -cpuprofile=./result/cpu128ttable.out -o=result/pprof128ttable.bin
go tool pprof --png result/pprof128ttable.bin result/mem128ttable.out > result/mem128ttable.png
go tool pprof --png result/pprof128ttable.bin result/cpu128ttable.out > result/cpu128ttable.png

rm result/*.out result/*.bin
