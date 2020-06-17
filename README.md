# pipeline-example

[![Go Report Card](https://goreportcard.com/badge/github.com/andygeiss/pipeline-example)](https://goreportcard.com/report/github.com/andygeiss/pipeline-example)

Build your own data pipeline to gather, organize and transform data by using protobuf as an intermediate format.

## Installation

First install the [Protobuf Compiler](https://developers.google.com/protocol-buffers/docs/downloads) and the corresponding [Protobuf Go Plugin](https://developers.google.com/protocol-buffers/docs/gotutorial)
manually or use the following command:

    make setup

## Usage

Run the provided example with a second:

    make
 
## Structure

    .
    ├── cmd
    │   └── pipeline
    │       └── main.go
    ├── data
    │   ├── external
    │   │   └── iris.csv
    │   ├── interim
    │   └── processed
    ├── go.mod
    ├── go.sum
    ├── internal
    │   ├── api
    │   │   ├── interim.pb.go
    │   │   ├── interim.proto
    │   │   ├── processed.pb.go
    │   │   └── processed.proto
    │   └── iris
    │       ├── load.go
    │       ├── organize.go
    │       ├── plotScatter.go
    │       └── printStats.go
    ├── Makefile
    ├── models
    ├── README.md
    └── reports
        ├── plot_histogram.png
        └── plot_scatter.png
