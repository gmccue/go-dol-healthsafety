# go-dol-healthsafety

[![Build Status](https://api.travis-ci.org/gmccue/go-dol-healthsafety.png?branch=master)](https://travis-ci.org/gmccue/go-dol-healthsafety)
[![GoDoc](https://godoc.org/github.com/gmccue/go-dol-healthsafety?status.svg)](https://godoc.org/github.com/gmccue/go-dol-healthsafety)

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
ccr, err := api.QueryCarClosedRepurposed()
```

### Basic Mine Information
These queries provide access to the [Basic Mine Information API endpoint](http://developer.dol.gov/health-and-safety/basic-mine-info/).

#### Address of record
```go
aor, err := api.QueryAddressOfRecord()
```

### DOL OSHA Compliance
These queries provide access to the [DOL OSHA Compliance API endpoint](http://developer.dol.gov/health-and-safety/osha-compliance/).

#### OSHA Fatalities
```go
of, err := api.QueryOSHAFatalities()
```

### Full Mine Information
These queries provide access to the the [Full Mine Information API endpoint](http://developer.dol.gov/health-and-safety/full-mine-info-mines/).

#### MSHA Mines
```go
fm, err := api.QueryFullMineInfo()
```

### Gulf Oil Spill
These queries provide access to the the [Gulf Oil Spill API endpoint](http://developer.dol.gov/health-and-safety/gulf-oil-spill/).

#### OSHA Direct Read Sampling
```go
dr, err := QueryOSHADirectReadSampling()
```

#### Noise Report
```go
on, err := QueryOSHANoiseReport()
```

### Mine Accident Injuries
These queries provide access to the the [Mine Accident Injuries API endpoint](http://developer.dol.gov/health-and-safety/mine-accident-injuries/).

#### Mine Accident, Injury and Illness Report
```go
ma, err := QueryMineAccidents()
```

### MSHA Employment Production
These queries provide access to the [DOL OSHA Compliance API endpoint](http://developer.dol.gov/health-and-safety/msha-employment-production/).

#### MSHA Employment Production
```go
ep, err := api.QueryMSHAEmploymentProduction()
```

### Mine Inspections
These queries provide access to the [Mine Inspections API endpoint](http://developer.dol.gov/health-and-safety/mine-inspections/).

#### Mine Inspections
```go
ep, err := api.QueryMineInspections()
```

### Mine Violation
These queries provide access to the [Mine Violation API endpoint](http://developer.dol.gov/health-and-safety/mine-violation/).

#### Mine Violation
```go
mv, err := api.QueryMineViolations()
```
