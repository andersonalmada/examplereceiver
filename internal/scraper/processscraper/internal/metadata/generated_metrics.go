// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	conventions "go.opentelemetry.io/collector/semconv/v1.9.0"
)

// AttributeContextSwitchType specifies the a value context_switch_type attribute.
type AttributeContextSwitchType int

const (
	_ AttributeContextSwitchType = iota
	AttributeContextSwitchTypeInvoluntary
	AttributeContextSwitchTypeVoluntary
)

// String returns the string representation of the AttributeContextSwitchType.
func (av AttributeContextSwitchType) String() string {
	switch av {
	case AttributeContextSwitchTypeInvoluntary:
		return "involuntary"
	case AttributeContextSwitchTypeVoluntary:
		return "voluntary"
	}
	return ""
}

// MapAttributeContextSwitchType is a helper map of string to AttributeContextSwitchType attribute value.
var MapAttributeContextSwitchType = map[string]AttributeContextSwitchType{
	"involuntary": AttributeContextSwitchTypeInvoluntary,
	"voluntary":   AttributeContextSwitchTypeVoluntary,
}

// AttributeDirection specifies the a value direction attribute.
type AttributeDirection int

const (
	_ AttributeDirection = iota
	AttributeDirectionRead
	AttributeDirectionWrite
)

// String returns the string representation of the AttributeDirection.
func (av AttributeDirection) String() string {
	switch av {
	case AttributeDirectionRead:
		return "read"
	case AttributeDirectionWrite:
		return "write"
	}
	return ""
}

// MapAttributeDirection is a helper map of string to AttributeDirection attribute value.
var MapAttributeDirection = map[string]AttributeDirection{
	"read":  AttributeDirectionRead,
	"write": AttributeDirectionWrite,
}

// AttributePagingFaultType specifies the a value paging_fault_type attribute.
type AttributePagingFaultType int

const (
	_ AttributePagingFaultType = iota
	AttributePagingFaultTypeMajor
	AttributePagingFaultTypeMinor
)

// String returns the string representation of the AttributePagingFaultType.
func (av AttributePagingFaultType) String() string {
	switch av {
	case AttributePagingFaultTypeMajor:
		return "major"
	case AttributePagingFaultTypeMinor:
		return "minor"
	}
	return ""
}

// MapAttributePagingFaultType is a helper map of string to AttributePagingFaultType attribute value.
var MapAttributePagingFaultType = map[string]AttributePagingFaultType{
	"major": AttributePagingFaultTypeMajor,
	"minor": AttributePagingFaultTypeMinor,
}

// AttributeState specifies the a value state attribute.
type AttributeState int

const (
	_ AttributeState = iota
	AttributeStateSystem
	AttributeStateUser
	AttributeStateWait
)

// String returns the string representation of the AttributeState.
func (av AttributeState) String() string {
	switch av {
	case AttributeStateSystem:
		return "system"
	case AttributeStateUser:
		return "user"
	case AttributeStateWait:
		return "wait"
	}
	return ""
}

// MapAttributeState is a helper map of string to AttributeState attribute value.
var MapAttributeState = map[string]AttributeState{
	"system": AttributeStateSystem,
	"user":   AttributeStateUser,
	"wait":   AttributeStateWait,
}

type metricProcessContextSwitches struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills process.context_switches metric with initial data.
func (m *metricProcessContextSwitches) init() {
	m.data.SetName("process.context_switches")
	m.data.SetDescription("Number of times the process has been context switched.")
	m.data.SetUnit("{count}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricProcessContextSwitches) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, contextSwitchTypeAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("type", contextSwitchTypeAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricProcessContextSwitches) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricProcessContextSwitches) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricProcessContextSwitches(cfg MetricConfig) metricProcessContextSwitches {
	m := metricProcessContextSwitches{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricProcessCPUTime struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills process.cpu.time metric with initial data.
func (m *metricProcessCPUTime) init() {
	m.data.SetName("process.cpu.time")
	m.data.SetDescription("Total CPU seconds broken down by different states.")
	m.data.SetUnit("s")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricProcessCPUTime) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64, stateAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
	dp.Attributes().PutStr("state", stateAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricProcessCPUTime) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricProcessCPUTime) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricProcessCPUTime(cfg MetricConfig) metricProcessCPUTime {
	m := metricProcessCPUTime{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricProcessCPUUtilization struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills process.cpu.utilization metric with initial data.
func (m *metricProcessCPUUtilization) init() {
	m.data.SetName("process.cpu.utilization")
	m.data.SetDescription("Percentage of total CPU time used by the process since last scrape, expressed as a value between 0 and 1. On the first scrape, no data point is emitted for this metric.")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricProcessCPUUtilization) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64, stateAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
	dp.Attributes().PutStr("state", stateAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricProcessCPUUtilization) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricProcessCPUUtilization) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricProcessCPUUtilization(cfg MetricConfig) metricProcessCPUUtilization {
	m := metricProcessCPUUtilization{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricProcessDiskIo struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills process.disk.io metric with initial data.
func (m *metricProcessDiskIo) init() {
	m.data.SetName("process.disk.io")
	m.data.SetDescription("Disk bytes transferred.")
	m.data.SetUnit("By")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricProcessDiskIo) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, directionAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("direction", directionAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricProcessDiskIo) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricProcessDiskIo) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricProcessDiskIo(cfg MetricConfig) metricProcessDiskIo {
	m := metricProcessDiskIo{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricProcessDiskOperations struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills process.disk.operations metric with initial data.
func (m *metricProcessDiskOperations) init() {
	m.data.SetName("process.disk.operations")
	m.data.SetDescription("Number of disk operations performed by the process.")
	m.data.SetUnit("{operations}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricProcessDiskOperations) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, directionAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("direction", directionAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricProcessDiskOperations) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricProcessDiskOperations) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricProcessDiskOperations(cfg MetricConfig) metricProcessDiskOperations {
	m := metricProcessDiskOperations{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricProcessHandles struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills process.handles metric with initial data.
func (m *metricProcessHandles) init() {
	m.data.SetName("process.handles")
	m.data.SetDescription("Number of handles held by the process.")
	m.data.SetUnit("{count}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricProcessHandles) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricProcessHandles) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricProcessHandles) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricProcessHandles(cfg MetricConfig) metricProcessHandles {
	m := metricProcessHandles{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricProcessMemoryUsage struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills process.memory.usage metric with initial data.
func (m *metricProcessMemoryUsage) init() {
	m.data.SetName("process.memory.usage")
	m.data.SetDescription("The amount of physical memory in use.")
	m.data.SetUnit("By")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricProcessMemoryUsage) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricProcessMemoryUsage) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricProcessMemoryUsage) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricProcessMemoryUsage(cfg MetricConfig) metricProcessMemoryUsage {
	m := metricProcessMemoryUsage{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricProcessMemoryUtilization struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills process.memory.utilization metric with initial data.
func (m *metricProcessMemoryUtilization) init() {
	m.data.SetName("process.memory.utilization")
	m.data.SetDescription("Percentage of total physical memory that is used by the process.")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
}

func (m *metricProcessMemoryUtilization) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricProcessMemoryUtilization) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricProcessMemoryUtilization) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricProcessMemoryUtilization(cfg MetricConfig) metricProcessMemoryUtilization {
	m := metricProcessMemoryUtilization{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricProcessMemoryVirtual struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills process.memory.virtual metric with initial data.
func (m *metricProcessMemoryVirtual) init() {
	m.data.SetName("process.memory.virtual")
	m.data.SetDescription("Virtual memory size.")
	m.data.SetUnit("By")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricProcessMemoryVirtual) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricProcessMemoryVirtual) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricProcessMemoryVirtual) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricProcessMemoryVirtual(cfg MetricConfig) metricProcessMemoryVirtual {
	m := metricProcessMemoryVirtual{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricProcessOpenFileDescriptors struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills process.open_file_descriptors metric with initial data.
func (m *metricProcessOpenFileDescriptors) init() {
	m.data.SetName("process.open_file_descriptors")
	m.data.SetDescription("Number of file descriptors in use by the process.")
	m.data.SetUnit("{count}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricProcessOpenFileDescriptors) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricProcessOpenFileDescriptors) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricProcessOpenFileDescriptors) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricProcessOpenFileDescriptors(cfg MetricConfig) metricProcessOpenFileDescriptors {
	m := metricProcessOpenFileDescriptors{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricProcessPagingFaults struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills process.paging.faults metric with initial data.
func (m *metricProcessPagingFaults) init() {
	m.data.SetName("process.paging.faults")
	m.data.SetDescription("Number of page faults the process has made.")
	m.data.SetUnit("{faults}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricProcessPagingFaults) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, pagingFaultTypeAttributeValue string) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("type", pagingFaultTypeAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricProcessPagingFaults) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricProcessPagingFaults) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricProcessPagingFaults(cfg MetricConfig) metricProcessPagingFaults {
	m := metricProcessPagingFaults{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricProcessSignalsPending struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills process.signals_pending metric with initial data.
func (m *metricProcessSignalsPending) init() {
	m.data.SetName("process.signals_pending")
	m.data.SetDescription("Number of pending signals for the process.")
	m.data.SetUnit("{signals}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricProcessSignalsPending) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricProcessSignalsPending) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricProcessSignalsPending) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricProcessSignalsPending(cfg MetricConfig) metricProcessSignalsPending {
	m := metricProcessSignalsPending{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricProcessThreads struct {
	data     pmetric.Metric // data buffer for generated metric.
	config   MetricConfig   // metric config provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills process.threads metric with initial data.
func (m *metricProcessThreads) init() {
	m.data.SetName("process.threads")
	m.data.SetDescription("Process threads count.")
	m.data.SetUnit("{threads}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricProcessThreads) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.config.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricProcessThreads) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricProcessThreads) emit(metrics pmetric.MetricSlice) {
	if m.config.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricProcessThreads(cfg MetricConfig) metricProcessThreads {
	m := metricProcessThreads{config: cfg}
	if cfg.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user config.
type MetricsBuilder struct {
	config                           MetricsBuilderConfig // config of the metrics builder.
	startTime                        pcommon.Timestamp    // start time that will be applied to all recorded data points.
	metricsCapacity                  int                  // maximum observed number of metrics per resource.
	metricsBuffer                    pmetric.Metrics      // accumulates metrics data before emitting.
	buildInfo                        component.BuildInfo  // contains version information.
	metricProcessContextSwitches     metricProcessContextSwitches
	metricProcessCPUTime             metricProcessCPUTime
	metricProcessCPUUtilization      metricProcessCPUUtilization
	metricProcessDiskIo              metricProcessDiskIo
	metricProcessDiskOperations      metricProcessDiskOperations
	metricProcessHandles             metricProcessHandles
	metricProcessMemoryUsage         metricProcessMemoryUsage
	metricProcessMemoryUtilization   metricProcessMemoryUtilization
	metricProcessMemoryVirtual       metricProcessMemoryVirtual
	metricProcessOpenFileDescriptors metricProcessOpenFileDescriptors
	metricProcessPagingFaults        metricProcessPagingFaults
	metricProcessSignalsPending      metricProcessSignalsPending
	metricProcessThreads             metricProcessThreads
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(mbc MetricsBuilderConfig, settings receiver.CreateSettings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		config:                           mbc,
		startTime:                        pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                    pmetric.NewMetrics(),
		buildInfo:                        settings.BuildInfo,
		metricProcessContextSwitches:     newMetricProcessContextSwitches(mbc.Metrics.ProcessContextSwitches),
		metricProcessCPUTime:             newMetricProcessCPUTime(mbc.Metrics.ProcessCPUTime),
		metricProcessCPUUtilization:      newMetricProcessCPUUtilization(mbc.Metrics.ProcessCPUUtilization),
		metricProcessDiskIo:              newMetricProcessDiskIo(mbc.Metrics.ProcessDiskIo),
		metricProcessDiskOperations:      newMetricProcessDiskOperations(mbc.Metrics.ProcessDiskOperations),
		metricProcessHandles:             newMetricProcessHandles(mbc.Metrics.ProcessHandles),
		metricProcessMemoryUsage:         newMetricProcessMemoryUsage(mbc.Metrics.ProcessMemoryUsage),
		metricProcessMemoryUtilization:   newMetricProcessMemoryUtilization(mbc.Metrics.ProcessMemoryUtilization),
		metricProcessMemoryVirtual:       newMetricProcessMemoryVirtual(mbc.Metrics.ProcessMemoryVirtual),
		metricProcessOpenFileDescriptors: newMetricProcessOpenFileDescriptors(mbc.Metrics.ProcessOpenFileDescriptors),
		metricProcessPagingFaults:        newMetricProcessPagingFaults(mbc.Metrics.ProcessPagingFaults),
		metricProcessSignalsPending:      newMetricProcessSignalsPending(mbc.Metrics.ProcessSignalsPending),
		metricProcessThreads:             newMetricProcessThreads(mbc.Metrics.ProcessThreads),
	}
	for _, op := range options {
		op(mb)
	}
	return mb
}

// NewResourceBuilder returns a new resource builder that should be used to build a resource associated with for the emitted metrics.
func (mb *MetricsBuilder) NewResourceBuilder() *ResourceBuilder {
	return NewResourceBuilder(mb.config.ResourceAttributes)
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pmetric.ResourceMetrics) {
	if mb.metricsCapacity < rm.ScopeMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.ScopeMetrics().At(0).Metrics().Len()
	}
}

// ResourceMetricsOption applies changes to provided resource metrics.
type ResourceMetricsOption func(pmetric.ResourceMetrics)

// WithResource sets the provided resource on the emitted ResourceMetrics.
// It's recommended to use ResourceBuilder to create the resource.
func WithResource(res pcommon.Resource) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		res.CopyTo(rm.Resource())
	}
}

// WithStartTimeOverride overrides start time for all the resource metrics data points.
// This option should be only used if different start time has to be set on metrics coming from different resources.
func WithStartTimeOverride(start pcommon.Timestamp) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		var dps pmetric.NumberDataPointSlice
		metrics := rm.ScopeMetrics().At(0).Metrics()
		for i := 0; i < metrics.Len(); i++ {
			switch metrics.At(i).Type() {
			case pmetric.MetricTypeGauge:
				dps = metrics.At(i).Gauge().DataPoints()
			case pmetric.MetricTypeSum:
				dps = metrics.At(i).Sum().DataPoints()
			}
			for j := 0; j < dps.Len(); j++ {
				dps.At(j).SetStartTimestamp(start)
			}
		}
	}
}

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead.
// Resource attributes should be provided as ResourceMetricsOption arguments.
func (mb *MetricsBuilder) EmitForResource(rmo ...ResourceMetricsOption) {
	rm := pmetric.NewResourceMetrics()
	rm.SetSchemaUrl(conventions.SchemaURL)
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol/examplereceiver/process")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricProcessContextSwitches.emit(ils.Metrics())
	mb.metricProcessCPUTime.emit(ils.Metrics())
	mb.metricProcessCPUUtilization.emit(ils.Metrics())
	mb.metricProcessDiskIo.emit(ils.Metrics())
	mb.metricProcessDiskOperations.emit(ils.Metrics())
	mb.metricProcessHandles.emit(ils.Metrics())
	mb.metricProcessMemoryUsage.emit(ils.Metrics())
	mb.metricProcessMemoryUtilization.emit(ils.Metrics())
	mb.metricProcessMemoryVirtual.emit(ils.Metrics())
	mb.metricProcessOpenFileDescriptors.emit(ils.Metrics())
	mb.metricProcessPagingFaults.emit(ils.Metrics())
	mb.metricProcessSignalsPending.emit(ils.Metrics())
	mb.metricProcessThreads.emit(ils.Metrics())

	for _, op := range rmo {
		op(rm)
	}
	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user config, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(rmo ...ResourceMetricsOption) pmetric.Metrics {
	mb.EmitForResource(rmo...)
	metrics := mb.metricsBuffer
	mb.metricsBuffer = pmetric.NewMetrics()
	return metrics
}

// RecordProcessContextSwitchesDataPoint adds a data point to process.context_switches metric.
func (mb *MetricsBuilder) RecordProcessContextSwitchesDataPoint(ts pcommon.Timestamp, val int64, contextSwitchTypeAttributeValue AttributeContextSwitchType) {
	mb.metricProcessContextSwitches.recordDataPoint(mb.startTime, ts, val, contextSwitchTypeAttributeValue.String())
}

// RecordProcessCPUTimeDataPoint adds a data point to process.cpu.time metric.
func (mb *MetricsBuilder) RecordProcessCPUTimeDataPoint(ts pcommon.Timestamp, val float64, stateAttributeValue AttributeState) {
	mb.metricProcessCPUTime.recordDataPoint(mb.startTime, ts, val, stateAttributeValue.String())
}

// RecordProcessCPUUtilizationDataPoint adds a data point to process.cpu.utilization metric.
func (mb *MetricsBuilder) RecordProcessCPUUtilizationDataPoint(ts pcommon.Timestamp, val float64, stateAttributeValue AttributeState) {
	mb.metricProcessCPUUtilization.recordDataPoint(mb.startTime, ts, val, stateAttributeValue.String())
}

// RecordProcessDiskIoDataPoint adds a data point to process.disk.io metric.
func (mb *MetricsBuilder) RecordProcessDiskIoDataPoint(ts pcommon.Timestamp, val int64, directionAttributeValue AttributeDirection) {
	mb.metricProcessDiskIo.recordDataPoint(mb.startTime, ts, val, directionAttributeValue.String())
}

// RecordProcessDiskOperationsDataPoint adds a data point to process.disk.operations metric.
func (mb *MetricsBuilder) RecordProcessDiskOperationsDataPoint(ts pcommon.Timestamp, val int64, directionAttributeValue AttributeDirection) {
	mb.metricProcessDiskOperations.recordDataPoint(mb.startTime, ts, val, directionAttributeValue.String())
}

// RecordProcessHandlesDataPoint adds a data point to process.handles metric.
func (mb *MetricsBuilder) RecordProcessHandlesDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricProcessHandles.recordDataPoint(mb.startTime, ts, val)
}

// RecordProcessMemoryUsageDataPoint adds a data point to process.memory.usage metric.
func (mb *MetricsBuilder) RecordProcessMemoryUsageDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricProcessMemoryUsage.recordDataPoint(mb.startTime, ts, val)
}

// RecordProcessMemoryUtilizationDataPoint adds a data point to process.memory.utilization metric.
func (mb *MetricsBuilder) RecordProcessMemoryUtilizationDataPoint(ts pcommon.Timestamp, val float64) {
	mb.metricProcessMemoryUtilization.recordDataPoint(mb.startTime, ts, val)
}

// RecordProcessMemoryVirtualDataPoint adds a data point to process.memory.virtual metric.
func (mb *MetricsBuilder) RecordProcessMemoryVirtualDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricProcessMemoryVirtual.recordDataPoint(mb.startTime, ts, val)
}

// RecordProcessOpenFileDescriptorsDataPoint adds a data point to process.open_file_descriptors metric.
func (mb *MetricsBuilder) RecordProcessOpenFileDescriptorsDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricProcessOpenFileDescriptors.recordDataPoint(mb.startTime, ts, val)
}

// RecordProcessPagingFaultsDataPoint adds a data point to process.paging.faults metric.
func (mb *MetricsBuilder) RecordProcessPagingFaultsDataPoint(ts pcommon.Timestamp, val int64, pagingFaultTypeAttributeValue AttributePagingFaultType) {
	mb.metricProcessPagingFaults.recordDataPoint(mb.startTime, ts, val, pagingFaultTypeAttributeValue.String())
}

// RecordProcessSignalsPendingDataPoint adds a data point to process.signals_pending metric.
func (mb *MetricsBuilder) RecordProcessSignalsPendingDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricProcessSignalsPending.recordDataPoint(mb.startTime, ts, val)
}

// RecordProcessThreadsDataPoint adds a data point to process.threads metric.
func (mb *MetricsBuilder) RecordProcessThreadsDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricProcessThreads.recordDataPoint(mb.startTime, ts, val)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}
