package retry

func Do(f func(lastErr error) error) error {
	var (
		err   error
		count int
	)
	for count < 3 {
		count++
		err = f(err)
		if err == nil {
			break
		}
	}

	return err
}
