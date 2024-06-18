package cy

func Preheat[T any, U any](dst *T, src *U) error {
	return Copy(dst, src, withPreheat())
}

func PreheatAndCopy[T any, U any](dst *T, preSrc, src *U) error {
	err := Copy(dst, preSrc, withPreheat())
	if err != nil {
		return err
	}
	return Copy(dst, src, withPreheat())
}
