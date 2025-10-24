package logmetrics

import (
	"fmt"
	"os"
	"sort"
)

func SaveStat(fileNameToSave string, metrics []SerializeMetric) error {
	fmt.Printf("fileNameToSave: %s\n\n", fileNameToSave)
	f, err := os.Create(fileNameToSave)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.WriteString("timeStart(ms) duration(us)\n"); err != nil {
		return err
	}

	if len(metrics) == 0 {
		return nil
	}
	sorted := make([]SerializeMetric, len(metrics))
	copy(sorted, metrics)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].SerializeStartTime.Before(sorted[j].SerializeStartTime)
	})

	startTimePoint := sorted[0].SerializeStartTime
	for _, v := range sorted {
		timeStart := v.SerializeStartTime.Sub(startTimePoint)
		duration := v.SerializeEndTime.Sub(v.SerializeStartTime)
		if _, err := f.WriteString(fmt.Sprintf("%d %d\n", timeStart.Milliseconds(), duration.Microseconds())); err != nil {
			return err
		}
	}
	return nil
}
