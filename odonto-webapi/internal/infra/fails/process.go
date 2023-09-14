package fails

func ProcessStack(code int, msg string, err error) Stack {
	// without err
	if err == nil {
		return NewError(code, msg, nil)
	}
	// with err
	errStack, ok := err.(Stack)
	if !ok {
		errStack = NewError(code, err.Error(), nil)
	}
	return NewError(code, msg, errStack)
}

func ProcessErrStack(err error) Stack {
	// without err
	if err == nil {
		return nil
	}
	// with err
	errStack, ok := err.(Stack)
	if !ok {
		return NewError(0, err.Error(), nil)
	}
	return NewError(errStack.Code(), errStack.Message(), errStack)
}

func ProcessErrStackWithStatus(code int, err error) Stack {
	// without err
	if err == nil {
		return nil
	}
	// with err
	errStack, ok := err.(Stack)
	if !ok {
		return NewError(0, err.Error(), nil)
	}
	return NewError(code, errStack.Message(), errStack)
}
