package errorkind

func New(description string) Kind {
	return &kind{
		description: description,
	}
}

type Kind interface {
	KindDescription() string
}

type kind struct {
	description string
}

func (m *kind) KindDescription() string {
	return m.description
}

func Wrap(err error, kind Kind) error {
	if err == nil || kind == nil {
		// There is no reason to return kind if there is no error
		return err
	}

	return &wrappedError{
		kind:          kind,
		originalError: err,
	}
}

type wrappedError struct {
	kind          Kind
	originalError error
}

func (m *wrappedError) Error() string {
	return m.kind.KindDescription() + ": " + m.originalError.Error()
}

func (m *wrappedError) Kind() Kind {
	return m.kind
}

func (m *wrappedError) OriginalError() error {
	return m.originalError
}

func Unwrap(err error) (originalError error, kind Kind) {
	wrappedError, ok := err.(interface {
		Kind() Kind
		OriginalError() error
	})
	if ok {
		return wrappedError.OriginalError(), wrappedError.Kind()
	}

	return err, nil
}

func GetKind(err error) (kind Kind) {
	wrappedError, ok := err.(interface {
		Kind() Kind
	})
	if ok {
		return wrappedError.Kind()
	}

	return nil
}
