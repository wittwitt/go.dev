package err1

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	ErrNotFound      = errors.New("file not foudn")
	ErrSizeNotEnough = errors.New("siz not enough")
)

func Test_is(t *testing.T) {

	baseErr := errors.New("base err")

	// %w wrap
	err1 := errors.New("err1")
	err2 := fmt.Errorf("err2: %w", err1)
	err3 := fmt.Errorf("err3: %w", err2)
	//
	err5 := fmt.Errorf("err5: %w, base: %w", err1, baseErr)

	require.Equal(t, true, errors.Is(err2, err1))
	require.Equal(t, true, errors.Is(err3, err1))
	//
	require.Equal(t, true, errors.Is(err5, err1))
	require.Equal(t, true, errors.Is(err5, baseErr))

	require.Equal(t, err2, errors.Unwrap(err3))
	require.NotEqual(t, err1, errors.Unwrap(err3))

	// string wrap
	err4 := fmt.Errorf("err1")
	require.Equal(t, false, errors.Is(err4, err1))

	//
	//
	// %v warp
	err12 := fmt.Errorf("err2: %v", err1)
	err13 := fmt.Errorf("err3: %v", err2)

	require.Equal(t, false, errors.Is(err12, err1))
	require.Equal(t, false, errors.Is(err13, err1))
}

func Test_is2(t *testing.T) {
	e0 := isas(0)
	require.True(t, errors.Is(e0, ErrNotFound))

	e5 := isas(5)
	require.False(t, errors.Is(e5, ErrNotFound))
	require.True(t, errors.Is(e5, ErrSizeNotEnough))

	e11 := isas(11)
	require.False(t, errors.Is(e11, ErrNotFound))
	require.True(t, errors.Is(e11, ErrSizeNotEnough))

	e22 := isas(22)
	require.NoError(t, e22)
}

func isas(s int) error {
	if s == 0 {
		return ErrNotFound
	}

	if s < 10 {
		return ErrSizeNotEnough
	}

	if s < 20 {
		return fmt.Errorf("%w, %v", ErrSizeNotEnough, 20-s)
	}

	return nil
}

func Test_warp(t *testing.T) {
	e1 := isas(11)
	e2 := fmt.Errorf("%w", e1)
	e3 := fmt.Errorf("%w", e2)
	require.Equal(t, true, errors.Is(e2, ErrSizeNotEnough))
	require.Equal(t, true, errors.Is(e3, ErrSizeNotEnough))
}

type ErrSf struct {
	error
}

func Test_as(t *testing.T) {

	e1 := isas(11)
	e2 := fmt.Errorf("e1: %w", e1)
	e3 := fmt.Errorf("e2: %w", e2)

	es := ErrSf{error: fmt.Errorf("e3: %w", e3)}

	e4 := fmt.Errorf("es: %w,", es)

	var sErr ErrSf

	if errors.As(e4, &sErr) {
		t.Log(sErr.Error())
	} else {
		t.Log(1)
	}
}
