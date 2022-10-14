package password

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestValidatePassword(t *testing.T) {
	type args struct {
		passwordHash string
		password     string
	}
	testPassword := "test"
	testPasswordHash, _ := GetHashPassword(testPassword)
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "WrongPassword",
			args: args{
				passwordHash: "test",
				password:     testPassword,
			},
			wantErr: true,
		},
		{
			name: "CorrectPassword",
			args: args{
				passwordHash: testPasswordHash,
				password:     testPassword,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidatePassword(tt.args.passwordHash, tt.args.password); (got != nil) != tt.wantErr {
				t.Errorf("ValidatePassword() = %v, wantErr %v", got, tt.wantErr)
			}
		})
	}
}

func TestGetHashPassword(t *testing.T) {
	type args struct {
		password string
	}
	testPassword := "test"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(testPassword), 13)
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "CorrectHashPassword",
			args: args{
				password: testPassword,
			},
			want: string(hashedPassword),
		},
		{
			name: "EmptyPassword",
			args: args{
				password: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetHashPassword(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (ValidatePassword(got, tt.args.password) != nil) != tt.wantErr {
				t.Errorf("GetHashPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
