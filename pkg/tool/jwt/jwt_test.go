package jwt

import (
	"reflect"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func TestParseUnverifiedJWT(t *testing.T) {
	type args struct {
		token string
	}
	testMapClaims := jwt.MapClaims{
		"test": "test",
	}
	expiredToken, _ := NewToken(testMapClaims, "", -time.Hour)
	token, _ := NewToken(testMapClaims, "", time.Hour)

	tests := []struct {
		name    string
		args    args
		want    jwt.MapClaims
		wantErr bool
	}{
		{
			name: "CorrectMapClaims",
			args: args{
				token: token,
			},
			want: testMapClaims,
		},
		{
			name: "EmptyToken",
			args: args{
				token: "",
			},
			want:    jwt.MapClaims{},
			wantErr: true,
		},
		{
			name: "ExpiredToken",
			args: args{
				token: expiredToken,
			},
			want:    testMapClaims,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseUnverifiedJWT(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseUnverifiedJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			delete(got, "exp")
			delete(got, "iat")
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseUnverifiedJWT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseJWT(t *testing.T) {
	type args struct {
		token           string
		verificationKey string
	}
	key := "test"
	testMapClaims := jwt.MapClaims{
		"test": "test",
	}
	expiredToken, _ := NewToken(testMapClaims, key, -time.Hour)
	token, _ := NewToken(testMapClaims, key, time.Hour)

	tests := []struct {
		name    string
		args    args
		want    jwt.MapClaims
		wantErr bool
	}{
		{
			name: "CorrectMapClaims",
			args: args{
				token:           token,
				verificationKey: key,
			},
			want: testMapClaims,
		},
		{
			name: "EmptyTokenWithCorrectKey",
			args: args{
				token:           "",
				verificationKey: key,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "EmptyKeyWithCorrectToken",
			args: args{
				token:           token,
				verificationKey: "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ExpiredToken",
			args: args{
				token:           expiredToken,
				verificationKey: key,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ExpiredTokenWithEmptyKey",
			args: args{
				token:           expiredToken,
				verificationKey: "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseJWT(tt.args.token, tt.args.verificationKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			delete(got, "exp")
			delete(got, "iat")
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseJWT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewToken(t *testing.T) {
	type args struct {
		payload         jwt.MapClaims
		signingKey      string
		secondsDuration time.Duration
	}
	now := time.Now()
	testMapClaims := jwt.MapClaims{
		"test": "test",
		"exp":  now.Add(time.Hour).Unix(),
		"iat":  now.Unix(),
	}
	testKey := "test"
	testToken, _ := jwt.NewWithClaims(
		jwt.SigningMethodHS256, testMapClaims).
		SignedString([]byte(testKey))
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "CorrectToken",
			args: args{
				payload:         testMapClaims,
				signingKey:      testKey,
				secondsDuration: time.Hour,
			},
			want: testToken,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewToken(tt.args.payload, tt.args.signingKey, tt.args.secondsDuration)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
