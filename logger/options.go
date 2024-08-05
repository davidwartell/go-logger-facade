package logger

import "time"

type LoggingOption func(o *Options)

type Options struct {
	productNameShort string
	samplingEnabled  bool
	samplingOptions  SamplingOptions
}

func (o *Options) clone() *Options {
	return &Options{
		productNameShort: o.productNameShort,
		samplingEnabled:  o.samplingEnabled,
		samplingOptions:  o.samplingOptions,
	}
}

type SamplingOptions struct {
	Tick       time.Duration
	First      int
	Thereafter int
}

func (so *SamplingOptions) clone() *SamplingOptions {
	return &SamplingOptions{
		Tick:       so.Tick,
		First:      so.First,
		Thereafter: so.Thereafter,
	}
}

// WithProductNameShort sets the name of your product used in log file names.
// example: logger.Instance().StartTask(logger.WithProductNameShort("your-product-name-here"))
//
//goland:noinspection GoUnusedExportedFunction
func WithProductNameShort(productNameShort string) LoggingOption {
	return func(o *Options) {
		o.productNameShort = productNameShort
	}
}

//goland:noinspection GoUnusedExportedFunction
func WithSampling(samplingOptions SamplingOptions) LoggingOption {
	return func(o *Options) {
		o.samplingOptions = samplingOptions
		o.samplingEnabled = true
	}
}
