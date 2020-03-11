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

Before running the benchmarks, install `easyjson` and generate the auxiliary files:

```
go mod download
go get -u github.com/mailru/easyjson/...
grep '~/go/bin' ~/.profile || echo 'export PATH=$PATH:~/go/bin' >> ~/.profile && source ~/.profile
go generate -v ./model
```

## Small Document

The [`small`](data/small/document.json) dataset uses a small document (126B).

```
ELASTICSEARCH_URL=http://server:9200 go run benchmarks.go --dataset=small --count=1_000_000 --flush=2MB --shards=5 --replicas=0 --fasthttp=true --easyjson=true
small: run [10x] warmup [3x] shards [5] replicas [0] workers [8] flush [2.0 MB] wait [1s] fasthttp easyjson
▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔
   1) add=1M  flush=1M  fail=0  reqs=52 dur=3.58s   279,173 docs/sec
   2) add=1M  flush=1M  fail=0  reqs=52 dur=3.52s   284,090 docs/sec
   3) add=1M  flush=1M  fail=0  reqs=52 dur=3.45s   289,351 docs/sec
   4) add=1M  flush=1M  fail=0  reqs=52 dur=3.49s   286,123 docs/sec
   5) add=1M  flush=1M  fail=0  reqs=52 dur=3.47s   287,852 docs/sec
   6) add=1M  flush=1M  fail=0  reqs=52 dur=3.47s   288,184 docs/sec
   7) add=1M  flush=1M  fail=0  reqs=52 dur=3.54s   282,246 docs/sec
   8) add=1M  flush=1M  fail=0  reqs=52 dur=3.47s   288,101 docs/sec
   9) add=1M  flush=1M  fail=0  reqs=52 dur=3.54s   282,485 docs/sec
  10) add=1M  flush=1M  fail=0  reqs=52 dur=3.46s   288,350 docs/sec
▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔
docs/sec: min [279,173] max [289,351] mean [286,987]
```

## HTTP Log Event

The [`httplog`](data/httplog/document.json) dataset uses a bigger document (2.5K), corresponding to a log event gathered by [Filebeat](https://www.elastic.co/guide/en/beats/filebeat/current/filebeat-module-nginx.html) from Nginx.

```
ELASTICSEARCH_URL=http://server:9200 go run benchmarks.go --dataset=httplog --count=1_000_000 --flush=3MB --shards=5 --replicas=0 --fasthttp=true --easyjson=true
httplog: run [10x] warmup [3x] shards [5] replicas [0] workers [8] flush [3.0 MB] wait [1s] fasthttp easyjson
▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔
   1) add=1M   flush=1M fail=0   reqs=649 dur=19.93s  50,165 docs/sec
   2) add=1M   flush=1M fail=0   reqs=649 dur=18.84s  53,072 docs/sec
   3) add=1M   flush=1M fail=0   reqs=649 dur=19.13s  52,249 docs/sec
   4) add=1M   flush=1M fail=0   reqs=649 dur=19.26s  51,912 docs/sec
   5) add=1M   flush=1M fail=0   reqs=649 dur=18.98s  52,662 docs/sec
   6) add=1M   flush=1M fail=0   reqs=649 dur=19.21s  52,056 docs/sec
   7) add=1M   flush=1M fail=0   reqs=649 dur=18.91s  52,865 docs/sec
   8) add=1M   flush=1M fail=0   reqs=649 dur=19.25s  51,934 docs/sec
   9) add=1M   flush=1M fail=0   reqs=649 dur=19.44s  51,440 docs/sec
  10) add=1M   flush=1M fail=0   reqs=649 dur=19.24s  51,966 docs/sec
▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔
docs/sec: min [50,165] max [53,072] mean [52,011]
```

## Mock Server

The `--mockserver` flag allows to run the benchmark against a "mock server", in this case Nginx, to understand a theoretical performance of the client, without the overhead of a real Elasticsearch cluster.

```
ELASTICSEARCH_URL=http://server:8000 go run benchmarks.go --dataset=small --count=1_000_000 --flush=2MB --warmup=0 --mockserver
small: run [10x] warmup [0x] shards [3] replicas [0] workers [8] flush [2.0 MB] wait [1s]
▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔
   1) add=1M   flush=0  fail=0   reqs=56  dur=810ms   1,222,493 docs/sec
   2) add=1M   flush=0  fail=0   reqs=56  dur=810ms   1,230,012 docs/sec
   3) add=1M   flush=0  fail=0   reqs=56  dur=790ms   1,251,564 docs/sec
   4) add=1M   flush=0  fail=0   reqs=56  dur=840ms   1,187,648 docs/sec
   5) add=1M   flush=0  fail=0   reqs=56  dur=800ms   1,237,623 docs/sec
   6) add=1M   flush=0  fail=0   reqs=56  dur=800ms   1,237,623 docs/sec
   7) add=1M   flush=0  fail=0   reqs=56  dur=800ms   1,240,694 docs/sec
   8) add=1M   flush=0  fail=0   reqs=56  dur=820ms   1,216,545 docs/sec
   9) add=1M   flush=0  fail=0   reqs=56  dur=790ms   1,253,132 docs/sec
  10) add=1M   flush=0  fail=0   reqs=56  dur=810ms   1,223,990 docs/sec
▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔
docs/sec: min [1,187,648] max [1,253,132] mean [1,233,818]
```

## Environment

Please note that these results are only illustrative, and the real performance depends on many factors:
the size and structure of your data, the index settings and mappings, the cluster setup or the hardware specification.

The benchmarks have been run in the following environment:

* OS: Ubuntu 18.04.4 LTS (5.0.0-1031-gcp)
* Client: A `n2-standard-8` [GCP instance](https://cloud.google.com/compute/docs/machine-types#n2_machine_types) (8 vCPUs/32GB RAM)
* Server: A `n2-standard-16` [GCP instance](https://cloud.google.com/compute/docs/machine-types#n2_machine_types) (16 vCPUs/64GB RAM)
* Disk: A [local SSD](https://cloud.google.com/compute/docs/disks#localssds) formatted as `ext4` on NVMe interface for Elasticsearch data
* A single-node Elasticsearch cluster, `7.6.0`, [default distribution](https://www.elastic.co/downloads/elasticsearch), installed from a TAR, with 4GB locked for heap
* Nginx 1.17.8 with [`nginx.conf`](etc/nginx.conf)
