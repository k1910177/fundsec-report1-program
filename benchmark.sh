mkdir -p result

go test -bench=BenchmarkAES128SBox . -benchtime=10000000x -cpu=1 -memprofile=./result/mem128sbox.out -cpuprofile=./result/cpu128sbox.out -o=result/pprof128sbox.bin
go tool pprof --pdf result/pprof128sbox.bin result/mem128sbox.out > result/mem128sbox.pdf
go tool pprof --pdf result/pprof128sbox.bin result/cpu128sbox.out > result/cpu128sbox.pdf

go test -bench=BenchmarkAES128TTable . -benchtime=10000000x -cpu=1 -memprofile=./result/mem128ttable.out -cpuprofile=./result/cpu128ttable.out -o=result/pprof128ttable.bin
go tool pprof --pdf result/pprof128ttable.bin result/mem128ttable.out > result/mem128ttable.pdf
go tool pprof --pdf result/pprof128ttable.bin result/cpu128ttable.out > result/cpu128ttable.pdf

rm result/*.out result/*.bin
