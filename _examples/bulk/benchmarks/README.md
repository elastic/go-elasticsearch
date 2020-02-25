# Bulk Indexer Benchmarks

The [`benchmarks.go`](benchmarks.go) file executes end-to-end benchmarks for `esutil.NewBulkIndexer`. It allows to configure indexer parameters, index settings, number of runs. See `go run benchmarks.go --help` for an overview of configuration options:

```
go run benchmarks.go --help
  -count int
    	Number of documents to generate (default 100000)
  -dataset string
    	Dataset to use for indexing (default "small")
  -debug
    	Enable logging output
  -easyjson
    	Use mailru/easyjson for JSON decoding
  -fasthttp
    	Use valyala/fasthttp for HTTP transport
  -flush value
    	Flush threshold in bytes (default 3MB)
  -index string
    	Index name (default "test-bulk-benchmarks")
  -mockserver
    	Measure added, not flushed items
  -replicas int
    	Number of index replicas (default 0)
  -runs int
    	Number of runs (default 10)
  -shards int
    	Number of index shards (default 3)
  -wait duration
    	Wait duration between runs (default 1s)
  -warmup int
    	Number of warmup runs (default 3)
  -workers int
    	Number of indexer workers (default 4)
```

## Small Document

The [`small`](data/small/document.json) dataset uses a small document (126B).

```
ELASTICSEARCH_URL=http://server:9200 go run benchmarks.go --dataset=small --count=1_000_000 --flush=2MB --shards=5 --replicas=0 --fasthttp=true --easyjson=true
small: run [10x] warmup [3x] shards [5] replicas [0] workers [8] flush [2.0 MB] wait [1s] fasthttp easyjson
▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔
   1) add=1M	flush=1M	fail=0	reqs=52	dur=3.6s  	277,392 docs/sec
   2) add=1M	flush=1M	fail=0	reqs=52	dur=3.51s 	284,819 docs/sec
   3) add=1M	flush=1M	fail=0	reqs=52	dur=3.49s 	286,123 docs/sec
   4) add=1M	flush=1M	fail=0	reqs=52	dur=3.45s 	289,268 docs/sec
   5) add=1M	flush=1M	fail=0	reqs=52	dur=3.45s 	289,603 docs/sec
   6) add=1M	flush=1M	fail=0	reqs=52	dur=3.49s 	286,450 docs/sec
   7) add=1M	flush=1M	fail=0	reqs=52	dur=3.46s 	288,350 docs/sec
   8) add=1M	flush=1M	fail=0	reqs=52	dur=3.52s 	284,010 docs/sec
   9) add=1M	flush=1M	fail=0	reqs=52	dur=3.63s 	275,330 docs/sec
  10) add=1M	flush=1M	fail=0	reqs=52	dur=3.48s 	286,861 docs/sec
▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔
docs/sec: min [275,330] max [289,603] mean [286,286]
```

## HTTP Log Event

The [`httplog`](data/httplog/document.json) dataset uses a bigger document (2.5K), corresponding to a log event gathered by [Filebeat](https://www.elastic.co/guide/en/beats/filebeat/current/filebeat-module-nginx.html) from Nginx.

```
ELASTICSEARCH_URL=http://server:9200 go run benchmarks.go --dataset=httplog --count=1_000_000 --flush=3MB --shards=5 --replicas=0 --fasthttp=true --easyjson=true
httplog: run [10x] warmup [3x] shards [5] replicas [0] workers [8] flush [3.0 MB] wait [1s]
▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔
   1) add=1M	flush=1M	fail=0	reqs=649	dur=34.47s	29,007 docs/sec
   2) add=1M	flush=1M	fail=0	reqs=649	dur=33.89s	29,505 docs/sec
   3) add=1M	flush=1M	fail=0	reqs=649	dur=32.27s	30,979 docs/sec
   4) add=1M	flush=1M	fail=0	reqs=649	dur=32.44s	30,819 docs/sec
   5) add=1M	flush=1M	fail=0	reqs=649	dur=34.33s	29,123 docs/sec
   6) add=1M	flush=1M	fail=0	reqs=649	dur=32.42s	30,838 docs/sec
   7) add=1M	flush=1M	fail=0	reqs=649	dur=34.04s	29,370 docs/sec
   8) add=1M	flush=1M	fail=0	reqs=649	dur=34.14s	29,287 docs/sec
   9) add=1M	flush=1M	fail=0	reqs=649	dur=34.01s	29,397 docs/sec
  10) add=1M	flush=1M	fail=0	reqs=649	dur=33.85s	29,541 docs/sec
▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔
docs/sec: min [29,007] max [30,979] mean [29,451]
```

## Mock Server

The `--mockserver` flag allows to run the benchmark against a "mock server", in this case Nginx, to understand a theoretical performance of the client, without the overhead of a real Elasticsearch cluster.

```
ELASTICSEARCH_URL=http://server:8000 go run benchmarks.go --dataset=small --count=1_000_000 --flush=2MB --mockserver
small: run [10x] warmup [3x] shards [3] replicas [0] workers [8] flush [2.0 MB] wait [1s]
▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔
   1) add=1M	flush=0	fail=0	reqs=56	dur=840ms 	1,190,476 docs/sec
   2) add=1M	flush=0	fail=0	reqs=56	dur=840ms 	1,180,637 docs/sec
   3) add=1M	flush=0	fail=0	reqs=56	dur=840ms 	1,189,060 docs/sec
   4) add=1M	flush=0	fail=0	reqs=56	dur=840ms 	1,180,637 docs/sec
   5) add=1M	flush=0	fail=0	reqs=56	dur=820ms 	1,206,272 docs/sec
   6) add=1M	flush=0	fail=0	reqs=56	dur=830ms 	1,196,172 docs/sec
   7) add=1M	flush=0	fail=0	reqs=56	dur=820ms 	1,207,729 docs/sec
   8) add=1M	flush=0	fail=0	reqs=56	dur=840ms 	1,186,239 docs/sec
   9) add=1M	flush=0	fail=0	reqs=56	dur=830ms 	1,203,369 docs/sec
  10) add=1M	flush=0	fail=0	reqs=56	dur=850ms 	1,168,224 docs/sec
▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔
docs/sec: min [1,168,224] max [1,207,729] mean [1,189,768]
```

## Environment

Please note that these results are only illustrative, and the real performance depends on many factors.

The benchmarks have been run in the following environment:

* OS: Debian GNU/Linux 9.12 (stretch) 4.9.0-12-amd64
* Client: A `n2-standard-8` GCP instance (8 vCPUs/32GB RAM)
* Server: A `n2-standard-16` GCP instance (16 vCPUs/64GB RAM)
* A single node Elasticsearch `7.6.0` cluster
* 4GB of RAM for the Elasticsearch heap size
* A 100 GB zonal SSD persistent disk formatted as `ext4` for Elasticsearch data (3,000 IOPS, 48 MB/s)
