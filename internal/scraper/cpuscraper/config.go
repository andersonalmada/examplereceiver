// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cpuscraper // import "github.com/andersonalmada/examplereceiver/internal/scraper/cpuscraper"

import (
	"github.com/andersonalmada/examplereceiver/internal"
	"github.com/andersonalmada/examplereceiver/internal/scraper/cpuscraper/internal/metadata"
)

// Config relating to CPU Metric Scraper.
type Config struct {
	metadata.MetricsBuilderConfig `mapstructure:",squash"`
	internal.ScraperConfig
}
