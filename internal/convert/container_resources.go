package convert

import (
	"github.com/pkg/errors"
	scoretypes "github.com/score-spec/score-go/types"
	coreV1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

func convertContainerResources(resources *scoretypes.ContainerResources) (coreV1.ResourceRequirements, error) {
	var out coreV1.ResourceRequirements
	var err error
	if resources != nil {
		if resources.Requests != nil {
			out.Requests, err = buildResourceList(resources.Requests)
			if err != nil {
				return out, errors.Wrap(err, "requests: failed to convert")
			}
			out.Limits, err = buildResourceList(resources.Limits)
			if err != nil {
				return out, errors.Wrap(err, "limits: failed to convert")
			}
		}
	}
	return out, nil
}

func buildResourceList(input *scoretypes.ResourcesLimits) (coreV1.ResourceList, error) {
	var err error
	output := make(coreV1.ResourceList)
	if input.Cpu != nil {
		output["cpu"], err = resource.ParseQuantity(*input.Cpu)
		if err != nil {
			return nil, errors.Wrapf(err, "cpu: failed to parse")
		}
	}
	if input.Memory != nil {
		output["memory"], err = resource.ParseQuantity(*input.Memory)
		if err != nil {
			return nil, errors.Wrapf(err, "memory: failed to parse")
		}
	}
	return output, nil
}
