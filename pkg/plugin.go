package pkg

import (
	"context"
	"time"

	. "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
)

type ExamplePlugin struct{}

func (dp *ExamplePlugin) ListAndWatch(e *Empty, s DevicePlugin_ListAndWatchServer) error {

	s.Send(&ListAndWatchResponse{Devices: []*Device{&Device{
		ID:     time.Now().String(),
		Health: Healthy,
	}}})

	for {
		time.Sleep(3 * time.Second)

		s.Send(&ListAndWatchResponse{Devices: []*Device{&Device{
			ID:     time.Now().String(),
			Health: Healthy,
		}}})
	}
}

func (dp *ExamplePlugin) GetPreferredAllocation(context.Context, *PreferredAllocationRequest) (*PreferredAllocationResponse, error) {

	return &PreferredAllocationResponse{}, nil
}

func (dp *ExamplePlugin) Allocate(c context.Context, r *AllocateRequest) (*AllocateResponse, error) {

	responses := []*ContainerAllocateResponse{{Envs: map[string]string{"k8s-device-plugin-example": time.Now().String()}}}

	return &AllocateResponse{ContainerResponses: responses}, nil

}

func (ExamplePlugin) GetDevicePluginOptions(context.Context, *Empty) (*DevicePluginOptions, error) {
	return nil, nil
}

func (ExamplePlugin) PreStartContainer(context.Context, *PreStartContainerRequest) (*PreStartContainerResponse, error) {
	return nil, nil
}
