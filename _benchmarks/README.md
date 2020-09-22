# Benchmarks

This folder contains code for running the [common Elasticsearch client benchmarks](https://clients-benchmarks.elastic.co).

The `benchmarks` package contains the general configuration for the benchmarks. The `runner` package provides the code for running and measuring the client actions, storing the results in an Elasticsearch cluster. The individual client interactions are defined by the `actions` package.

The code for setting up the benchmarks infrastructure is available at
<https://github.com/elastic/elasticsearch-clients-benchmarks>. It uses Terraform to create, setup and configure the "target" and "runner" system.

