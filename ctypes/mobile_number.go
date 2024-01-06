package ctypes

import (
	"log/slog"
	"strconv"
)

func ConvertToUint(s string) uint64 {
	parseUint, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		slog.Error("Failed to convert the mobile number",
			slog.String("str", s),
			slog.String("err", err.Error()))
		panic(err)
	}
	return parseUint
}
