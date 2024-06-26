package sknlinechart_test

import (
	"reflect"
	"time"

	"fyne.io/fyne/v2/theme"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/skoona/sknlinechart"
)

var _ = Describe("Linechartinterfaces", func() {

	It("Ensure Interfaces are implemented", func() {
		point := sknlinechart.NewChartDatapoint(62.4, theme.ColorYellow, time.Now().Format(time.RFC1123))
		dIntf := reflect.TypeOf((*sknlinechart.ChartDatapoint)(nil)).Elem()

		var dataPoints map[string][]*sknlinechart.ChartDatapoint
		chart, _ := sknlinechart.NewLineChart("Title", "Footer", 1, 10, &dataPoints)
		cIntf := reflect.TypeOf((*sknlinechart.LineChart)(nil)).Elem()

		By("ChartDatapoint interface should be implemented")
		Expect(reflect.TypeOf(point).Implements(dIntf)).To(BeTrue())

		By("LineChart interface should be implemented")
		Expect(reflect.TypeOf(chart).Implements(cIntf)).To(BeTrue())
	})
})
