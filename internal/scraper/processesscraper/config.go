// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package processesscraper // import "github.com/andersonalmada/examplereceiver/internal/scraper/processesscraper"

import (
	"github.com/andersonalmada/examplereceiver/internal"
	"github.com/andersonalmada/examplereceiver/internal/scraper/processesscraper/internal/metadata"
)

// Config relating to Processes Metric Scraper.
type Config struct {
	// MetricsBuilderConfig allows customizing scraped metrics/attributes representation.
	metadata.MetricsBuilderConfig `mapstructure:",squash"`
	internal.ScraperConfig
}
