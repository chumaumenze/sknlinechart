package sknlinechart

import "fyne.io/fyne/v2"

// LineChart feature list
type LineChart interface {
	// Chart Attributes
	IsDataPointMarkersEnabled() bool // mouse button 2 toggles
	IsHorizGridLinesEnabled() bool
	IsVertGridLinesEnabled() bool
	IsMousePointDisplayEnabled() bool // hoverable and mouse button one

	SetDataPointMarkers(enable bool)
	SetHorizGridLines(enable bool)
	SetVertGridLines(enable bool)
	SetMousePointDisplay(enable bool)

	// Info labels
	GetTopLeftLabel() string
	GetTitle() string
	GetTopRightLabel() string

	// Scale legend
	GetMiddleLeftLabel() string
	GetMiddleRightLabel() string

	// Info Labels
	GetBottomLeftLabel() string
	GetBottomCenteredLabel() string
	GetBottomRightLabel() string

	SetTopLeftLabel(newValue string)
	SetTitle(newValue string)
	SetTopRightLabel(newValue string)
	SetMiddleLeftLabel(newValue string)
	SetMiddleRightLabel(newValue string)
	SetBottomLeftLabel(newValue string)
	SetBottomCenteredLabel(newValue string)
	SetBottomRightLabel(newValue string)

	// ApplyDataSeries add a whole data series at once
	// expect this will rarely be used, since loading more than 120 point will raise error
	ApplyDataSeries(seriesName string, newSeries []LineChartDatapoint) error

	// ApplyDataPoint primary method to add another data point to any series
	// If series has more than 120 points, point 0 will be rolled out making room for this one
	ApplyDataPoint(seriesName string, newDataPoint LineChartDatapoint)

	// SetMinSize set minimum size of internal widget
	//  do not use
	//
	// Internal Use Only
	SetMinSize(s fyne.Size)

	// Refresh() expensive call to reload all on screen elements
	// prefer to use the containers Refresh() method
	//
	// Internal Use Only
	Refresh()

	// Resize Set a new size on the widget
	// use the Window Resize() method rather than this widgets
	//
	// Internal Use Only
	Resize(s fyne.Size)
}
