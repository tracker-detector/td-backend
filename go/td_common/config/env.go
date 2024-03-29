package config

import (
	"os"
)

func EnvMongoURI() string {
	return os.Getenv("MONGO_URI")
}

func EnvDBName() string {
	return os.Getenv("DB_NAME")
}

func EnvRequestCollection() string {
	return os.Getenv("REQUEST_COLLECTION")
}

func EnvUserCollection() string {
	return os.Getenv("USER_COLLECTION")
}

func EnvTrainingRunCollection() string {
	return os.Getenv("TRAINING_RUNS_COLLECTION")
}

func EnvExporterCollection() string {
	return os.Getenv("EXPORTER_COLLECTION")
}

func EnvExporterRunsCollection() string {
	return os.Getenv("EXPORTER_RUNS_COLLECTION")
}

func EnvDatasetCollection() string {
	return os.Getenv("DATASET_COLLECTION")
}

func EnvAdminEmail() string {
	return os.Getenv("EMAIL")
}

func EnvMinIoURI() string {
	return os.Getenv("MINIO_URI")
}

func EnvMinIoAccessKey() string {
	return os.Getenv("MINIO_ACCESS_KEY")
}

func EnvMinIoPrivateKey() string {
	return os.Getenv("MINIO_PRIVATE_KEY")
}

func EnvExportBucketName() string {
	return os.Getenv("EXPORT_BUCKET_NAME")
}

func EnvModelBucketName() string {
	return os.Getenv("MODEL_BUCKET_NAME")
}

func EnvExtractorBucketName() string {
	return os.Getenv("EXTRACTOR_BUCKET_NAME")
}

func EnvAdminApiKey() string {
	return os.Getenv("ADMIN_API_KEY")
}

func EnvModelCollection() string {
	return os.Getenv("MODELS_COLLECTION")
}

func EnvMQURI() string {
	return os.Getenv("RABBIT_URI")
}

func EnvTrainQueueName() string {
	return os.Getenv("TRAIN_QUEUE")
}

func EnvExportQueueName() string {
	return os.Getenv("EXPORT_QUEUE")
}
