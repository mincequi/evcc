package meter

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decorateSMA(meter api.Meter, battery func() (float64, error)) api.Meter {
	switch {
	case battery == nil:
		return &struct {
			api.Meter
		}{
			Meter: meter,
		}

	case battery != nil:
		return &struct {
			api.Meter
			api.Battery
		}{
			Meter: meter,
			Battery: &decorateSMABatteryImpl{
				battery: battery,
			},
		}
	}

	return nil
}

type decorateSMABatteryImpl struct {
	battery func() (float64, error)
}

func (impl *decorateSMABatteryImpl) SoC() (float64, error) {
	return impl.battery()
}
