// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: role_action.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/gogoproto"
	_ "github.com/mwitkow/go-proto-validators"
	regexp "regexp"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *RoleDetail) Validate() error {
	return nil
}
func (this *GetRolesRequest) Validate() error {
	return nil
}
func (this *GetRolesResponse) Validate() error {
	for _, item := range this.Roles {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Roles", err)
			}
		}
	}
	return nil
}
func (this *GetRoleDetailRequest) Validate() error {
	return nil
}
func (this *GetRoleDetailResponse) Validate() error {
	if this.RoleDetail != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RoleDetail); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RoleDetail", err)
		}
	}
	return nil
}

var _regex_CreateRoleRequest_RoleName = regexp.MustCompile(`^([^!@#$%&*+=~?;:<>{}]*)$`)
var _regex_CreateRoleRequest_RoleDescription = regexp.MustCompile(`^([^!@#$%&*+=~?;:<>{}]*)$`)
var _regex_CreateRoleRequest_RoleManagerDomain = regexp.MustCompile(`^(^[\w|\/])([\_\-\/a-zA-Z0-9\ ]*)([\w]$)$`)
var _regex_CreateRoleRequest_RoleAlias = regexp.MustCompile(`^(^[\w|\/])([\_\-\/a-zA-Z0-9\ ]*)([\w]$)$`)

func (this *CreateRoleRequest) Validate() error {
	if !_regex_CreateRoleRequest_RoleName.MatchString(this.RoleName) {
		return github_com_mwitkow_go_proto_validators.FieldError("RoleName", fmt.Errorf(`value '%v' must be a string conforming to regex "^([^!@#$%&*+=~?;:<>{}]*)$"`, this.RoleName))
	}
	if !_regex_CreateRoleRequest_RoleDescription.MatchString(this.RoleDescription) {
		return github_com_mwitkow_go_proto_validators.FieldError("RoleDescription", fmt.Errorf(`value '%v' must be a string conforming to regex "^([^!@#$%&*+=~?;:<>{}]*)$"`, this.RoleDescription))
	}
	if !_regex_CreateRoleRequest_RoleManagerDomain.MatchString(this.RoleManagerDomain) {
		return github_com_mwitkow_go_proto_validators.FieldError("RoleManagerDomain", fmt.Errorf(`value '%v' must be a string conforming to regex "^(^[\\w|\\/])([\\_\\-\\/a-zA-Z0-9\\ ]*)([\\w]$)$"`, this.RoleManagerDomain))
	}
	if !_regex_CreateRoleRequest_RoleAlias.MatchString(this.RoleAlias) {
		return github_com_mwitkow_go_proto_validators.FieldError("RoleAlias", fmt.Errorf(`value '%v' must be a string conforming to regex "^(^[\\w|\\/])([\\_\\-\\/a-zA-Z0-9\\ ]*)([\\w]$)$"`, this.RoleAlias))
	}
	return nil
}
func (this *CreateRoleResponse) Validate() error {
	return nil
}
func (this *UpdateRoleRequest) Validate() error {
	return nil
}
func (this *UpdateRoleResponse) Validate() error {
	return nil
}
func (this *DeleteRoleRequest) Validate() error {
	return nil
}
func (this *DeleteRoleResponse) Validate() error {
	return nil
}
func (this *GetUserRolesRequest) Validate() error {
	return nil
}
func (this *GetUserRolesResponse) Validate() error {
	return nil
}
func (this *GetRoleUsersRequest) Validate() error {
	return nil
}
func (this *GetRoleUsersResponse) Validate() error {
	return nil
}
func (this *SetUserRoleRequest) Validate() error {
	return nil
}
func (this *SetUserRoleResponse) Validate() error {
	return nil
}
func (this *SetRoleUserRequest) Validate() error {
	return nil
}
func (this *SetRoleUserResponse) Validate() error {
	return nil
}
