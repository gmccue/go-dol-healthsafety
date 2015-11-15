# go-dol-healthsafety

[![Build Status](https://api.travis-ci.org/gmccue/go-ilab-childlabor.png?branch=master)](https://travis-ci.org/gmccue/go-ilab-childlabor)
[![GoDoc](https://godoc.org/github.com/gmccue/go-ilab-childlabor?status.svg)](https://godoc.org/github.com/gmccue/go-ilab-childlabor)

go-dol-healthsafety provides programmatic access to the [United States Department of Labor Health & Safety datasets API](http://developer.dol.gov/health-and-safety).

This library is currently a work in progress and only provides access to the datasets listed in [Provided Endpoints](https://github.com/gmccue/go-dol-healthsafety#provided-endpoints).

## Installation
```
go get github.com/gmccue/go-dol-healthsafety
```

## Usage
In order to use this library, you must have a valid API token for the DOL Public API. You can register for an API key at the [US DOL website](https://devtools.dol.gov/developer/).

```go
import (
    "log"

    hstats "github.com/gmccue/go-dol-healthsafety"
)

func main() {
    api := hstats.New("{my_secret_key}")
    api.Debug = true

    basicMineInfo, err := api.QueryAddressOfRecord()
    if err != nil {
        log.Println(err)
    }

    log.Printf("%v", basicMineInfo)
}
```

## Provided Endpoints
### Auto Workers
These queries provide access to the [Auto Workers API endpoint](http://developer.dol.gov/health-and-safety/auto-workers/).

#### Car Closed Repurposed
```go
func CarClosedRepurposed() {
    ccr, err := api.QueryCarClosedRepurposed()
    if err != nil {
        log.Println(err)
    }

    log.Printf("%v", ccr)
}
```

### Basic Mine Information
These queries provide access to the [Basic Mine Information API endpoint](http://developer.dol.gov/health-and-safety/basic-mine-info/).

#### Address of record
```go
func AddressOfRecord() {
    aor, err := api.QueryAddressOfRecord()
    if err != nil {
        log.Println(err)
    }

    log.Printf("%v", aor)
}
```

### DOL OSHA Compliance
These queries provide access to the [DOL OSHA Compliance API endpoint](http://developer.dol.gov/health-and-safety/osha-compliance/).

#### OSHA Fatalities
```go
func OSHAFatalities() {
    of, err := api.QueryOSHAFatalities()
    if err != nil {
        log.Println(err)
    }

    log.Printf("%v", of)
}
```
