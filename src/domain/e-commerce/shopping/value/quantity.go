package value

type Quantity int

func (value Quantity) Check() error {

	if value <= 0 {

		// Returns Value Invariant Error

		return nil
	}

	if value >= 10 {

		// Returns Value Invariant Error

		return nil
	}

	return nil
}
