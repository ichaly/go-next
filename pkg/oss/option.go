package oss

type Options struct {
	size        int64
	contentType string
}

type UploadOption func(*Options)

func WithSize(size int64) UploadOption {
	return func(o *Options) {
		if size > 0 {
			o.size = size
		}
	}
}

func WithContentType(contentType string) UploadOption {
	return func(o *Options) {
		if contentType != "" {
			o.contentType = contentType
		}
	}
}
