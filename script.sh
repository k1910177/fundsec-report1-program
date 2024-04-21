mkdir result

go test -bench=BenchmarkAES256SBox . -benchtime=1000000x -cpu=1 -memprofile=./result/mem256sbox.out -cpuprofile=./result/cpu256sbox.out -o=result/pprof256sbox.bin
go tool pprof --png result/pprof256sbox.bin result/mem256sbox.out > result/mem256sbox.png
go tool pprof --png result/pprof256sbox.bin result/cpu256sbox.out > result/cpu256sbox.png

go test -bench=BenchmarkAES256TTable . -benchtime=1000000x -cpu=1 -memprofile=./result/mem256ttable.out -cpuprofile=./result/cpu256ttable.out -o=result/pprof256ttable.bin
go tool pprof --png result/pprof256ttable.bin result/mem256ttable.out > result/mem256ttable.png
go tool pprof --png result/pprof256ttable.bin result/cpu256ttable.out > result/cpu256ttable.png
