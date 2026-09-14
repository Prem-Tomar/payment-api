package httpapi

type ReadinessChecker interface {
	Check() error
}

type DefaultReadinessChecker struct{}

func (DefaultReadinessChecker) Check() error {
	return nil
}
