
### topN

```bash
# Generate Data
$ ./generate_data.sh 10 /tmp/randomInputData 

# Run tests & build topn for OSX
$ docker run -v "$PWD":/go golang:1.12 bash -c "go test -v"

$ docker run -v "$PWD":/go -e GOOS=darwin -e GOARCH=amd64 golang:1.12 bash -c "go build -o topn"

# Run
$ ./topn -topn 10 -filePath /tmp/randomInputData 
999999991517
999999794139
999999760922
999999717879
999999291233
999999185518
999999179731
999999071496
999999050489
999998727776
