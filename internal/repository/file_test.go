package repository

import (
	"bytes"
	"context"
	"io"
	"math/rand"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"go.beyondstorage.io/v5/pkg/randbytes"
	"go.beyondstorage.io/v5/types"

	"github.com/TcMits/ent-clean-template/internal/testutils"
)

func TestNewFileRepository(t *testing.T) {
	type args struct {
		storager types.Storager
	}

	storager := testutils.GetMemmoryStorager(t)

	tests := []struct {
		name string
		args args
		want FileRepository
	}{
		{
			name: "Success",
			args: args{storager: storager},
			want: &fileRepository{
				&readFileRepository{storager: storager},
				&existFileRepository{storager: storager},
				&writeFileRepository{storager: storager, writeSize: _DefaultWriteSize},
				&deleteFileRepository{storager: storager},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFileRepository(tt.args.storager); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFileRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readFileRepository_Read(t *testing.T) {
	type fields struct {
		storager types.Storager
	}
	type args struct {
		ctx    context.Context
		path   string
		offset int64
		size   int64
	}

	ctx := context.Background()
	filePath := "test.txt"
	storager := testutils.GetMemmoryStorager(t)
	size := rand.Int63n(4 * 1024 * 1024)
	r := io.LimitReader(randbytes.NewRand(), size)
	buf := bytes.Buffer{}
	tee := io.TeeReader(r, &buf)
	storager.Write(filePath, tee, size)
	wantW, err := io.ReadAll(&buf)
	require.NoError(t, err)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantW   string
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				storager: storager,
			},
			args: args{
				ctx:    ctx,
				path:   filePath,
				offset: 0,
				size:   size,
			},
			want:    size,
			wantW:   string(wantW),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &readFileRepository{
				storager: tt.fields.storager,
			}
			w := &bytes.Buffer{}
			got, err := r.Read(tt.args.ctx, tt.args.path, w, tt.args.offset, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("readFileRepository.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("readFileRepository.Read() = %v, want %v", got, tt.want)
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("readFileRepository.Read() = %v, wantW %v", gotW[:10], tt.wantW[:10])
			}
		})
	}
}

func Test_writeFileRepository_Write(t *testing.T) {
	type fields struct {
		storager  types.Storager
		writeSize int64
	}
	type args struct {
		ctx  context.Context
		path string
		re   io.Reader
		size int64
	}

	ctx := context.Background()
	filePath := "test.txt"
	storager := testutils.GetMemmoryStorager(t)
	size := rand.Int63n(4 * 1024 * 1024)
	r := io.LimitReader(randbytes.NewRand(), size)
	storager.Write(filePath, r, size)
	size = rand.Int63n(4 * 1024 * 1024)
	r = io.LimitReader(randbytes.NewRand(), size)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:   "Success",
			fields: fields{storager: storager, writeSize: _DefaultWriteSize},
			args: args{
				ctx:  ctx,
				path: "test2.txt",
				re:   r,
				size: size,
			},
			want: size,
		},
		{
			name:   "ExistError",
			fields: fields{storager: storager, writeSize: _DefaultWriteSize},
			args: args{
				ctx:  ctx,
				path: "test.txt",
				re:   r,
				size: size,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &writeFileRepository{
				storager:  tt.fields.storager,
				writeSize: tt.fields.writeSize,
			}
			got, err := r.Write(tt.args.ctx, tt.args.path, tt.args.re, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("writeFileRepository.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("writeFileRepository.Write() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deleteFileRepository_Delete(t *testing.T) {
	type fields struct {
		storager types.Storager
	}
	type args struct {
		ctx  context.Context
		path string
	}

	ctx := context.Background()
	filePath := "test.txt"
	storager := testutils.GetMemmoryStorager(t)
	size := rand.Int63n(4 * 1024 * 1024)
	r := io.LimitReader(randbytes.NewRand(), size)
	storager.Write(filePath, r, size)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "Success",
			fields: fields{storager: storager},
			args: args{
				ctx:  ctx,
				path: "test.txt",
			},
		},
		{
			name:   "NotExist",
			fields: fields{storager: storager},
			args: args{
				ctx:  ctx,
				path: "test2.txt",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &deleteFileRepository{
				storager: tt.fields.storager,
			}
			if err := r.Delete(tt.args.ctx, tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("deleteFileRepository.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_existFileRepository_Exist(t *testing.T) {
	type fields struct {
		storager types.Storager
	}
	type args struct {
		ctx  context.Context
		path string
	}

	ctx := context.Background()
	filePath := "test.txt"
	storager := testutils.GetMemmoryStorager(t)
	size := rand.Int63n(4 * 1024 * 1024)
	r := io.LimitReader(randbytes.NewRand(), size)
	storager.Write(filePath, r, size)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				storager: storager,
			},
			args: args{
				ctx:  ctx,
				path: "test.txt",
			},
			want: true,
		},
		{
			name: "NotExist",
			fields: fields{
				storager: storager,
			},
			args: args{
				ctx:  ctx,
				path: "test2.txt",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &existFileRepository{
				storager: tt.fields.storager,
			}
			got, err := r.Exist(tt.args.ctx, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("existFileRepository.Exist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("existFileRepository.Exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
