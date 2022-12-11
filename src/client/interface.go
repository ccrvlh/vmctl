package client

import (
	flAPI "github.com/weaveworks-liquidmetal/flintlock/api/services/microvm/v1alpha1"
	flTypes "github.com/weaveworks-liquidmetal/flintlock/api/types"
	"google.golang.org/protobuf/types/known/emptypb"
)

//counterfeiter:generate -o fakeclient/ . FlintlockClient
type FlintlockClient interface {
	Create(mvm *flTypes.MicroVMSpec) (*flAPI.CreateMicroVMResponse, error)
	Get(uid string) (*flAPI.GetMicroVMResponse, error)
	List(name, ns string) (*flAPI.ListMicroVMsResponse, error)
	Delete(uid string) (*emptypb.Empty, error)
	Stop(uid string) (*emptypb.Empty, error)
	Start(uid string) (*emptypb.Empty, error)
	Snapshot(uid string) (*emptypb.Empty, error)
	SSH(uid string) (*emptypb.Empty, error)
	Events(uid string) (*emptypb.Empty, error)
	Inspect(uid string) (*emptypb.Empty, error)
	Logs(uid string) (*emptypb.Empty, error)
	Launch(uid string) (*emptypb.Empty, error)
	Network(uid string) (*emptypb.Empty, error)
	Storage(uid string) (*emptypb.Empty, error)
	Close() error
}
