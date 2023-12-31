// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pagingscraper // import "github.com/andersonalmada/examplereceiver/internal/scraper/pagingscraper"

import (
	"github.com/andersonalmada/examplereceiver/internal"
	"github.com/andersonalmada/examplereceiver/internal/scraper/pagingscraper/internal/metadata"
)

// Config relating to Paging Metric Scraper.
type Config struct {
	// MetricsBuilderConfig allows customizing scraped metrics/attributes representation.
	metadata.MetricsBuilderConfig `mapstructure:",squash"`
	internal.ScraperConfig
}
