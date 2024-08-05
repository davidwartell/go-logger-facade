package logger

type loggerConfig struct {
	instances map[string]*LogInstance
	options   *Options
}

func (cfg *loggerConfig) clone() (clone *loggerConfig) {
	clone = new(loggerConfig)
	clone.instances = make(map[string]*LogInstance, len(cfg.instances))
	for k, v := range cfg.instances {
		clone.instances[k] = v
	}
	if cfg.options != nil {
		clone.options = cfg.options.clone()
	}
	return
}
