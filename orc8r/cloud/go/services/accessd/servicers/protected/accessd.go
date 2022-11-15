/*
Copyright 2020 The Magma Authors.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/*
Accessd service provides the gRPC interface for cloud services & REST server to
check & manage Identity access permissions.
*/
package servicers

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	accessprotos "magma/orc8r/cloud/go/services/accessd/protos"
	"magma/orc8r/cloud/go/services/accessd/storage"
	"magma/orc8r/lib/go/protos"
)

type AccessControlServer struct {
	store storage.AccessdStorage
}

func NewAccessdServer(store storage.AccessdStorage) *AccessControlServer {
	return &AccessControlServer{store}
}

// SetOperator Overwrites Permissions for operator Identity to manage others
// Request includes ACL to add for the Operator
func (srv *AccessControlServer) SetOperator(ctx context.Context, req *accessprotos.AccessControl_ListRequest) (*protos.Void, error) {
	err := accessprotos.VerifyACLRequest(req)
	if err != nil {
		return nil, err
	}

	acl := &accessprotos.AccessControl_List{
		Operator: req.Operator,
		Entities: map[string]*accessprotos.AccessControl_Entity{},
	}
	err = accessprotos.AddToACL(acl, req.Entities)
	if err != nil {
		return nil, err
	}

	err = srv.store.PutACL(req.Operator, acl)
	if err != nil {
		return nil, err
	}

	return &protos.Void{}, nil
}

// AddACL Adds Permissions for one Identity to manage others
// Request includes ACL to add for the Operator
func (srv *AccessControlServer) UpdateOperator(ctx context.Context, req *accessprotos.AccessControl_ListRequest) (*protos.Void, error) {
	err := accessprotos.VerifyACLRequest(req)
	if err != nil {
		return nil, err
	}

	err = srv.store.UpdateACLWithEntities(req.Operator, req.Entities)
	if err != nil {
		return nil, err
	}

	return &protos.Void{}, nil
}

// DeleteOperator Removes all operator's permissions (the entire operator's ACL)
func (srv *AccessControlServer) DeleteOperator(ctx context.Context, oper *protos.Identity) (*protos.Void, error) {
	return &protos.Void{}, srv.store.DeleteACL(oper)
}

// GetOperatorACL Returns the managing Identity's permissions list
func (srv *AccessControlServer) GetOperatorACL(ctx context.Context, oper *protos.Identity) (*accessprotos.AccessControl_List, error) {
	return srv.store.GetACL(oper)
}

// GetOperatorsACLs Returns the managing Identities' permissions list
func (srv *AccessControlServer) GetOperatorsACLs(
	ctx context.Context, opers *protos.Identity_List,
) (*accessprotos.AccessControl_Lists, error) {
	res, err := srv.store.GetManyACL(opers.GetList())
	return &accessprotos.AccessControl_Lists{Acls: res}, err
}

// Returns the managing Identity's permissions for a given entity
// NOTE: Takes into account wildcards for the entity's type in the ACL
func (srv *AccessControlServer) GetPermissions(
	ctx context.Context,
	req *accessprotos.AccessControl_PermissionsRequest,
) (*accessprotos.AccessControl_Entity, error) {
	res := &accessprotos.AccessControl_Entity{}
	err := accessprotos.VerifyPermissionsRequest(req)
	if err != nil {
		return res, err
	}
	acl, err := srv.store.GetACL(req.Operator)
	if err != nil {
		return res, err
	}
	res.Id = req.Entity
	res.Permissions = accessprotos.GetEntityPermissions(acl, res.Id) // Aggregated entity permissions
	return res, nil
}

// Returns the managing Identity's permissions for a given entity
// NOTE: Takes into account wildcards for the entity's type in the ACL
func (srv *AccessControlServer) CheckPermissions(
	ctx context.Context,
	req *accessprotos.AccessControl_ListRequest,
) (*protos.Void, error) {
	err := accessprotos.VerifyACLRequest(req)
	if err != nil {
		return nil, err
	}

	acl, err := srv.store.GetACL(req.Operator)
	if err != nil {
		return nil, err
	}

	return &protos.Void{}, accessprotos.CheckEntitiesPermissions(acl, req.Entities)
}

// Lists all globally registered operators on the cloud
func (srv *AccessControlServer) ListOperators(ctx context.Context, _ *protos.Void) (*protos.Identity_List, error) {
	res := new(protos.Identity_List)
	operators, err := srv.store.ListAllIdentity()
	res.List = operators
	return res, err
}

// Cleanup a given entity from all Operators' ACLs
// TBD: This needs to be implemented to avoid security venerability when deleting
// a network with customer selected ID (vs. generated by the cloud ID)
func (srv *AccessControlServer) DeleteEntity(ctx context.Context, ent *protos.Identity) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "Not Implemented")
}
