# benchmark-go-echo

## Running the benchmark
- Make sure docker daemon is started, containers are up for a while and that CPU and memory usage are stable. You can use the Docker `Resource Usage` extension.
- Make sure that the `jmeter.log` and `docker-stats.csv` files are empty.
- Run the capture command with the jmeter load test using the below command:
```
./capture-docker-stats.sh & (sleep 5; jmeter -n -t benchmark.jmx -l results.jtl -e -o ./jmeter_results) & wait
```