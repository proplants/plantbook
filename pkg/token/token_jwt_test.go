package token

import (
	"context"
	"testing"
	"time"
)

const (
	root7Token        string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjI1MjQ2NDQwMDAsInVpZCI6NywidXNlciI6InJvb3QiLCJ1c2VyX3JvbGUiOjF9.-CoCagMgLCugmQ4NU6ZQbcw-CPDXOnBsVTlmiHEyjzo"
	root7ExpiredToken string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjk0NjcyMDgwMCwidWlkIjo3LCJ1c2VyIjoicm9vdCIsInVzZXJfcm9sZSI6MX0.nqp2bM1bnqyOqZmnXdCek6ELgvmsJ0f3rwrZSu9YzDA"
)

var (
	tExp, _  = time.Parse(time.RFC3339, "2050-01-01T10:00:00Z")
	tPast, _ = time.Parse(time.RFC3339, "2000-01-01T10:00:00Z")
)

func TestJwtToken_Create(t *testing.T) {
	type fields struct {
		Secret []byte
	}
	type args struct {
		username     string
		uid          int64
		tokenExpTime int64
		roleID       int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			"simple_jwt",
			fields{Secret: []byte("secret")},
			args{username: "root", uid: 7, tokenExpTime: tExp.Unix(), roleID: 1},
			root7Token,
			false,
		},
		{
			"expired_jwt",
			fields{Secret: []byte("secret")},
			args{username: "root", uid: 7, tokenExpTime: tPast.Unix(), roleID: 1},
			root7ExpiredToken,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tk := &JwtToken{
				Secret: tt.fields.Secret,
			}
			got, err := tk.Create(tt.args.username, tt.args.uid, tt.args.tokenExpTime, tt.args.roleID)
			if (err != nil) != tt.wantErr {
				t.Errorf("JwtToken.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("JwtToken.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJwtToken_Check(t *testing.T) {
	type fields struct {
		Secret []byte
	}
	type args struct {
		ctx        context.Context
		inputToken string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			"simple_jwt_check",
			fields{Secret: []byte("secret")},
			args{ctx: context.Background(), inputToken: root7Token},
			true,
			false,
		},
		{
			"expired_jwt",
			fields{Secret: []byte("secret")},
			args{ctx: context.Background(), inputToken: root7ExpiredToken},
			false,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tk := &JwtToken{
				Secret: tt.fields.Secret,
			}
			got, err := tk.Check(tt.args.ctx, tt.args.inputToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("JwtToken.Check() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("JwtToken.Check() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJwtToken_FindUserData(t *testing.T) {
	type fields struct {
		Secret []byte
	}
	type args struct {
		inputToken string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantUid        int64
		wantUsername   string
		wantUserRoleID int32
		wantErr        bool
	}{
		{"simple_token", fields{Secret: []byte("secret")}, args{inputToken: root7Token}, 7, "root", 1, false},
		{"simple_token2", fields{Secret: []byte("secret")}, args{inputToken: root7ExpiredToken}, 0, "", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tk := &JwtToken{
				Secret: tt.fields.Secret,
			}
			gotUid, gotUsername, gotUserRoleID, err := tk.FindUserData(tt.args.inputToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("JwtToken.FindUserData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotUid != tt.wantUid {
				t.Errorf("JwtToken.FindUserData() gotUid = %v, want %v", gotUid, tt.wantUid)
			}
			if gotUsername != tt.wantUsername {
				t.Errorf("JwtToken.FindUserData() gotUsername = %v, want %v", gotUsername, tt.wantUsername)
			}
			if gotUserRoleID != tt.wantUserRoleID {
				t.Errorf("JwtToken.FindUserData() gotUserRoleID = %v, want %v", gotUserRoleID, tt.wantUserRoleID)
			}
		})
	}
}
