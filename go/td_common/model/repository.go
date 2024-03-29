package model

import (
	"context"
)

type IRepository[T any] interface {
	Save(ctx context.Context, entity T) (T, error)
	SaveAll(ctx context.Context, entity []T) ([]T, error)
	FindAll(ctx context.Context) ([]T, error)
	StreamAll(ctx context.Context) (<-chan T, <-chan error)
	FindByID(ctx context.Context, id string) (T, error)
	DeleteByID(ctx context.Context, id string) error
	DeleteAll(ctx context.Context) error
	Count(ctx context.Context) (int64, error)
	InTransaction(ctx context.Context, fn func(context.Context) error) error
}

type ExporterRepository interface {
	IRepository[*Exporter]
	FindByName(ctx context.Context, name string) (*Exporter, error)
}

type ExportRunRepository interface {
	IRepository[*ExportRun]
	CountByExporterID(ctx context.Context, exporterId string) (int64, error)
	FindByExporterID(ctx context.Context, exporterId string) ([]*ExportRun, error)
	ExistByExporterIDAndRecducer(ctx context.Context, exporterId, reducer string) (bool, error)
	DeleteAllByExporterID(ctx context.Context, id string) error
}

type RequestRepository interface {
	IRepository[*RequestData]
	StreamByDataset(ctx context.Context, dataset string) (<-chan *RequestData, <-chan error)
	CountByUrlLike(ctx context.Context, url string) (int64, error)
	FindAllByUrlLikePaged(ctx context.Context, url string, page, pageSize int) ([]*RequestData, error)
	DeleteAllByLabel(ctx context.Context, label string) error
}

type TrainingRunRepository interface {
	IRepository[*TrainingRun]
	CountByModelID(ctx context.Context, modelId string) (int64, error)
	CountByName(ctx context.Context, modelName string) (int64, error)
	FindByModelID(ctx context.Context, modelId string) ([]*TrainingRun, error)
	FindByModelName(ctx context.Context, modelId string) ([]*TrainingRun, error)
	DeleteAllByModelID(ctx context.Context, id string) error
	InTransaction(ctx context.Context, fn func(context.Context) error) error
}

type UserRepository interface {
	IRepository[*UserData]
	FindByEmail(ctx context.Context, email string) (*UserData, error)
}

type ModelRepository interface {
	IRepository[*Model]
	FindByName(ctx context.Context, name string) (*Model, error)
}

type DatasetRepository interface {
	IRepository[*Dataset]
	FindByName(ctx context.Context, name string) (*Dataset, error)
	FindByLabel(ctx context.Context, label string) (*Dataset, error)
}
