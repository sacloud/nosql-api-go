// Copyright 2025- The sacloud/nosql-api-go authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nosql

import (
	"context"
	"errors"

	v1 "github.com/sacloud/nosql-api-go/apis/v1"
)

type InstanceAPI interface {
	GetVersion(ctx context.Context) (*v1.NosqlGetVersionResponseNosql, error)
	UpgradeVersion(ctx context.Context, version string) error
	GetParameters(ctx context.Context) ([]v1.NosqlGetParameter, error)
	SetParameters(ctx context.Context, params []v1.NosqlPutParameter) error
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

var _ InstanceAPI = (*instanceOp)(nil)

type instanceOp struct {
	client *v1.Client
	dbId   string
}

func NewInstanceOp(client *v1.Client, dbId string) InstanceAPI {
	return &instanceOp{client: client, dbId: dbId}
}

func (op *instanceOp) GetVersion(ctx context.Context) (*v1.NosqlGetVersionResponseNosql, error) {
	res, err := op.client.GetVersion(ctx, v1.GetVersionParams{ApplianceID: op.dbId})
	if err != nil {
		return nil, NewAPIError("Instance.GetVersion", 0, err)
	}

	switch p := res.(type) {
	case *v1.NosqlGetVersionResponse:
		return &p.Nosql.Value, nil
	case *v1.BadRequestResponse:
		return nil, NewAPIError("Instance.GetVersion", 400, errors.New(p.ErrorMsg.Value))
	case *v1.UnauthorizedResponse:
		return nil, NewAPIError("Instance.GetVersion", 401, errors.New(p.ErrorMsg.Value))
	case *v1.ServerErrorResponse:
		return nil, NewAPIError("Instance.GetVersion", 500, errors.New(p.ErrorMsg.Value))
	default:
		return nil, NewAPIError("Instance.GetVersion", 0, nil)
	}
}

func (op *instanceOp) UpgradeVersion(ctx context.Context, version string) error {
	res, err := op.client.PutVersion(ctx, &v1.NosqlPutVersionRequest{
		Nosql: v1.NosqlPutVersionRequestNosql{Version: version}},
		v1.PutVersionParams{ApplianceID: op.dbId})
	if err != nil {
		return NewAPIError("Instance.UpgradeVersion", 0, err)
	}

	switch p := res.(type) {
	case *v1.NosqlPutVersionResponse:
		return nil
	case *v1.BadRequestResponse:
		return NewAPIError("Instance.UpgradeVersion", 400, errors.New(p.ErrorMsg.Value))
	case *v1.UnauthorizedResponse:
		return NewAPIError("Instance.UpgradeVersion", 401, errors.New(p.ErrorMsg.Value))
	case *v1.ServerErrorResponse:
		return NewAPIError("Instance.UpgradeVersion", 500, errors.New(p.ErrorMsg.Value))
	default:
		return NewAPIError("Instance.UpgradeVersion", 0, nil)
	}
}

func (op *instanceOp) GetParameters(ctx context.Context) ([]v1.NosqlGetParameter, error) {
	res, err := op.client.GetParameter(ctx, v1.GetParameterParams{ApplianceID: op.dbId})
	if err != nil {
		return nil, NewAPIError("Instance.GetParameters", 0, err)
	}

	switch p := res.(type) {
	case *v1.GetParameterResponse:
		return p.Nosql.Value.Parameters, nil
	case *v1.BadRequestResponse:
		return nil, NewAPIError("Instance.GetParameters", 400, errors.New(p.ErrorMsg.Value))
	case *v1.UnauthorizedResponse:
		return nil, NewAPIError("Instance.GetParameters", 401, errors.New(p.ErrorMsg.Value))
	case *v1.ServerErrorResponse:
		return nil, NewAPIError("Instance.GetParameters", 500, errors.New(p.ErrorMsg.Value))
	default:
		return nil, NewAPIError("Instance.GetParameters", 0, nil)
	}
}

func (op *instanceOp) SetParameters(ctx context.Context, params []v1.NosqlPutParameter) error {
	res, err := op.client.PutParameter(ctx, &v1.PutParameterRequest{
		Nosql: v1.PutParameterRequestNosql{Parameters: params}},
		v1.PutParameterParams{ApplianceID: op.dbId})
	if err != nil {
		return NewAPIError("Instance.SetParameters", 0, err)
	}

	switch p := res.(type) {
	case *v1.PutParameterResponse:
		return nil
	case *v1.BadRequestResponse:
		return NewAPIError("Instance.SetParameters", 400, errors.New(p.ErrorMsg.Value))
	case *v1.UnauthorizedResponse:
		return NewAPIError("Instance.SetParameters", 401, errors.New(p.ErrorMsg.Value))
	case *v1.ServerErrorResponse:
		return NewAPIError("Instance.SetParameters", 500, errors.New(p.ErrorMsg.Value))
	default:
		return NewAPIError("Instance.SetParameters", 0, nil)
	}
}

func (op *instanceOp) Start(ctx context.Context) error {
	res, err := op.client.PutAppliancePower(ctx, v1.PutAppliancePowerParams{ApplianceID: op.dbId})
	if err != nil {
		return NewAPIError("Instance.Start", 0, err)
	}

	switch p := res.(type) {
	case *v1.SuccessResponse:
		return nil
	case *v1.BadRequestResponse:
		return NewAPIError("Instance.Start", 400, errors.New(p.ErrorMsg.Value))
	case *v1.UnauthorizedResponse:
		return NewAPIError("Instance.Start", 401, errors.New(p.ErrorMsg.Value))
	case *v1.NotFoundResponse:
		return NewAPIError("Instance.Start", 404, errors.New(p.ErrorMsg.Value))
	case *v1.ServerErrorResponse:
		return NewAPIError("Instance.Start", 500, errors.New(p.ErrorMsg.Value))
	default:
		return NewAPIError("Instance.Start", 0, nil)
	}
}

func (op *instanceOp) Stop(ctx context.Context) error {
	res, err := op.client.DeleteAppliancePower(ctx, v1.DeleteAppliancePowerParams{ApplianceID: op.dbId})
	if err != nil {
		return NewAPIError("Instance.Stop", 0, err)
	}

	switch p := res.(type) {
	case *v1.SuccessResponse:
		return nil
	case *v1.BadRequestResponse:
		return NewAPIError("Instance.Stop", 400, errors.New(p.ErrorMsg.Value))
	case *v1.UnauthorizedResponse:
		return NewAPIError("Instance.Stop", 401, errors.New(p.ErrorMsg.Value))
	case *v1.NotFoundResponse:
		return NewAPIError("Instance.Stop", 404, errors.New(p.ErrorMsg.Value))
	case *v1.ServerErrorResponse:
		return NewAPIError("Instance.Stop", 500, errors.New(p.ErrorMsg.Value))
	default:
		return NewAPIError("Instance.Stop", 0, nil)
	}
}
