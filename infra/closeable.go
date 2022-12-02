package infra

type Closeable interface {
	Close() error
}

func Close(c Closeable) (err error) {
	return c.Close()
}

func CloseOrTerminate(c Closeable) {
	if err := c.Close(); err != nil {
		panic(err)
	}
}
