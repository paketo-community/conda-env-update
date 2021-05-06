package fakes

import (
	"sync"

	"github.com/paketo-buildpacks/packit"
)

type Planner struct {
	MergeLayerTypesCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			String                  string
			BuildpackPlanEntrySlice []packit.BuildpackPlanEntry
		}
		Returns struct {
			Launch bool
			Build  bool
		}
		Stub func(string, []packit.BuildpackPlanEntry) (bool, bool)
	}
}

func (f *Planner) MergeLayerTypes(param1 string, param2 []packit.BuildpackPlanEntry) (bool, bool) {
	f.MergeLayerTypesCall.Lock()
	defer f.MergeLayerTypesCall.Unlock()
	f.MergeLayerTypesCall.CallCount++
	f.MergeLayerTypesCall.Receives.String = param1
	f.MergeLayerTypesCall.Receives.BuildpackPlanEntrySlice = param2
	if f.MergeLayerTypesCall.Stub != nil {
		return f.MergeLayerTypesCall.Stub(param1, param2)
	}
	return f.MergeLayerTypesCall.Returns.Launch, f.MergeLayerTypesCall.Returns.Build
}
