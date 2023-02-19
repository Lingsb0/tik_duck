package uuservice_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/808-not-found/tik_duck/cmd/useruser/uuservice"
	"github.com/808-not-found/tik_duck/kitex_gen/useruser"
)

func TestUserRelationActionService(t *testing.T) {
	type args struct {
		ctx context.Context //nolint
		req *useruser.RelationActionRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *useruser.RelationActionResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := uuservice.UserRelationActionService(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRelationActionService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRelationActionService() = %v, want %v", got, tt.want)
			}
		})
	}
}
