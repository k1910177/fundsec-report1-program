mkdir result

go test -bench=BenchmarkAES128SBox . -benchtime=1000000x -cpu=1 -memprofile=./result/mem128sbox.out -cpuprofile=./result/cpu128sbox.out -o=result/pprof128sbox.bin
go tool pprof --png result/pprof128sbox.bin result/mem128sbox.out > result/mem128sbox.png
go tool pprof --png result/pprof128sbox.bin result/cpu128sbox.out > result/cpu128sbox.png

go test -bench=BenchmarkAES128TTable . -benchtime=1000000x -cpu=1 -memprofile=./result/mem128ttable.out -cpuprofile=./result/cpu128ttable.out -o=result/pprof128ttable.bin
go tool pprof --png result/pprof128ttable.bin result/mem128ttable.out > result/mem128ttable.png
go tool pprof --png result/pprof128ttable.bin result/cpu128ttable.out > result/cpu128ttable.png

go test -bench=BenchmarkAES192TTable . -benchtime=1000000x -cpu=1 -memprofile=./result/mem192ttable.out -cpuprofile=./result/cpu192ttable.out -o=result/pprof192ttable.bin
go tool pprof --png result/pprof192ttable.bin result/mem192ttable.out > result/mem192ttable.png
go tool pprof --png result/pprof192ttable.bin result/cpu192ttable.out > result/cpu192ttable.png

go test -bench=BenchmarkAES192SBox . -benchtime=1000000x -cpu=1 -memprofile=./result/mem192sbox.out -cpuprofile=./result/cpu192sbox.out -o=result/pprof192sbox.bin
go tool pprof --png result/pprof192sbox.bin result/mem192sbox.out > result/mem192sbox.png
go tool pprof --png result/pprof192sbox.bin result/cpu192sbox.out > result/cpu192sbox.png

go test -bench=BenchmarkAES256SBox . -benchtime=1000000x -cpu=1 -memprofile=./result/mem256sbox.out -cpuprofile=./result/cpu256sbox.out -o=result/pprof256sbox.bin
go tool pprof --png result/pprof256sbox.bin result/mem256sbox.out > result/mem256sbox.png
go tool pprof --png result/pprof256sbox.bin result/cpu256sbox.out > result/cpu256sbox.png

go test -bench=BenchmarkAES256TTable . -benchtime=1000000x -cpu=1 -memprofile=./result/mem256ttable.out -cpuprofile=./result/cpu256ttable.out -o=result/pprof256ttable.bin
go tool pprof --png result/pprof256ttable.bin result/mem256ttable.out > result/mem256ttable.png
go tool pprof --png result/pprof256ttable.bin result/cpu256ttable.out > result/cpu256ttable.png

rm result/*.out result/*.bin
