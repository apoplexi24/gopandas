# GPandas

GPandas is a data manipulation and analysis library written in Go. It provides high-performance, easy-to-use data structures and tools for working with structured data, inspired by Python's pandas library.

## Features

* Zero-copy data operations where possible
* High-performance data structures
* Generic type support
* Concurrent operations
* CSV/JSON import/export
* Data filtering and transformation
* Set operations
* Type-safe operations

## Getting Started

### Prerequisites

GPandas requires Go version 1.18 or above (for generics support).

### Installation

```bash
go get github.com/apoplexi24/gpandas
```

## Core Components

### DataFrame
The primary data structure for handling 2-dimensional data with labeled axes.

### Series
1-dimensional labeled array that can hold data of any type.

### Set
Generic implementation of set operations supporting any comparable type.

## Performance

GPandas is designed with performance in mind, utilizing:
- Generic implementations to avoid interface overhead
- Efficient memory management
- Concurrent operations where applicable
- Zero-copy operations when possible


### Development Setup

1. Clone the repository
2. Install dependencies

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Inspired by Python's pandas library
- Built with Go's powerful generic system
- Contributions from the Go community

## Status

GPandas is under active development. While it's suitable for production use, some features are still being added and enhanced.