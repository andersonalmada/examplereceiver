// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package memoryscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/examplereceiver/internal/scraper/memoryscraper"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/examplereceiver/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/examplereceiver/internal/scraper/memoryscraper/internal/metadata"
)

// Config relating to Memory Metric Scraper.
type Config struct {
	metadata.MetricsBuilderConfig `mapstructure:",squash"`
	internal.ScraperConfig
}
