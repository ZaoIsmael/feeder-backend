# Deporvillage - Feeder service
![tests](https://github.com/ZaoIsmael/feeder-backend/actions/workflows/test.yml/badge.svg)

This repository is a TCP service where you can register products by SKU.

The project is made with `Go` and designed with `DDD` patterns and `hexagonal architecture`.

## Prerequisites

* Go 1.17

## Getting started

This repository is organized with the [Standard Go Project Layout](https://github.com/golang-standards/project-layout/blob/master/README_es.md), as you can see below. wherein `cmd` we have the
application, server, and the bootstrapping of this one.

In the `internal` directory, we have everything related to our domain. This directory is organized by a module for
each `aggregate` and also cross-cutting. each module is organized in the same way; domain, application, and infrastructure.

Directory tmp is used to store the log files.

```shell
.
├── cmd
│   └── feeder-service
│       ├── main.go
│       └── src
│           ├── controller
│           │   ├── controller.go
│           │   ├── product_controller.go
│           │   └── report_controller.go
│           ├── server
│           │   ├── bootstrap.go
│           │   └── server.go
│           └── util
│               └── util.go
├── internal
│   ├── cross-cutting
│   │   ├── domain
│   │   │   ├── aggregate_root.go
│   │   │   ├── event_bus.go
│   │   │   ├── event_domain.go
│   │   │   ├── event_handler.go
│   │   │   ├── sku.go
│   │   │   └── sku_test.go
│   │   └── infrastructure
│   │       └── in_memory_event_bus.go
│   ├── inventory
│   │   ├── application
│   │   │   ├── add_product.go
│   │   │   ├── add_product_test.go
│   │   │   └── event-handlers
│   │   │       ├── product_was_added.go
│   │   │       └── product_was_added_test.go
│   │   ├── domain
│   │   │   ├── events
│   │   │   │   ├── product_was_added.go
│   │   │   │   ├── product_was_duplicated.go
│   │   │   │   └── product_was_invalid.go
│   │   │   ├── inventory.go
│   │   │   ├── Inventory_id.go
│   │   │   ├── inventory_test.go
│   │   │   ├── product.go
│   │   │   ├── register_product.go
│   │   │   └── repository.go
│   │   └── infrastructure
│   │       ├── file_register_product.go
│   │       ├── in_memory_register_product.go
│   │       └── in_memory_repository.go
│   └── report
│       ├── application
│       │   ├── event-handlers
│       │   │   ├── product_was_added.go
│       │   │   ├── product_was_added_test.go
│       │   │   ├── product_was_duplicated.go
│       │   │   ├── product_was_duplicated_test.go
│       │   │   ├── product_was_invalid.go
│       │   │   └── product_was_invalid_test.go
│       │   ├── get.go
│       │   └── get_test.go
│       ├── domain
│       │   ├── report.go
│       │   ├── report_id.go
│       │   └── repository.go
│       └── infrastructure
│           └── in_memory_repository.go
└── tmp
```
