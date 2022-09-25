package repository

import (
	"context"
	"errors"
	"io"
	"math"

	"go.beyondstorage.io/v5/pairs"
	"go.beyondstorage.io/v5/services"
	"go.beyondstorage.io/v5/types"
)

const _DefaultWriteSize int64 = 1024 * 10

type readFileRepository struct {
	storager types.Storager
}

type existFileRepository struct {
	storager types.Storager
}

type writeFileRepository struct {
	storager  types.Storager
	writeSize int64
}

type deleteFileRepository struct {
	storager types.Storager
}

type fileRepository struct {
	*readFileRepository
	*existFileRepository
	*writeFileRepository
	*deleteFileRepository
}

func NewFileRepository(storager types.Storager) FileRepository {
	if storager == nil {
		panic("storager is required")
	}
	return &fileRepository{
		&readFileRepository{storager: storager},
		&existFileRepository{storager: storager},
		&writeFileRepository{storager: storager, writeSize: _DefaultWriteSize},
		&deleteFileRepository{storager: storager},
	}
}

func (r *readFileRepository) Read(
	ctx context.Context,
	path string,
	w io.Writer,
	offset int64,
	size int64,
) (int64, error) {
	return r.storager.ReadWithContext(ctx, path, w, pairs.WithOffset(offset), pairs.WithSize(size))
}

func (r *writeFileRepository) Write(
	ctx context.Context,
	path string,
	re io.Reader,
	size int64,
) (int64, error) {
	_, err := r.storager.StatWithContext(ctx, path)
	if err == nil || !errors.Is(err, services.ErrObjectNotExist) {
		return 0, errors.New("writeFileRepository - Write - r.storager.Stat: Object is exist")
	}
	// check if have write size
	storageWriteSize, haveWriteSize := r.storager.Metadata().GetWriteSizeMaximum()
	if !haveWriteSize {
		storageWriteSize = r.writeSize
	}

	writeSize := int64(math.Min(float64(r.writeSize), float64(storageWriteSize)))
	appendTimes := int64(math.Ceil(float64(size) / float64(writeSize)))
	count := int64(0)
	as := r.storager.(types.Appender)
	o, err := as.CreateAppendWithContext(ctx, path)

	// appending
	for i := int64(0); i < int64(appendTimes); i++ {
		n, err := as.WriteAppendWithContext(ctx, o, re, writeSize)
		if err != nil {
			return 0, err
		}
		count += n
	}
	err = as.CommitAppendWithContext(ctx, o)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *deleteFileRepository) Delete(ctx context.Context, path string) error {
	return r.storager.DeleteWithContext(ctx, path)
}

func (r *existFileRepository) Exist(ctx context.Context, path string) (bool, error) {
	_, err := r.storager.StatWithContext(ctx, path)
	if err != nil {
		if errors.Is(err, services.ErrObjectNotExist) {
			err = nil
		}
		return false, err
	}
	return true, nil
}
