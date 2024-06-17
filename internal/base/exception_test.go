package base

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCatch(t *testing.T) {
	t.Run("unknown exception", func(t *testing.T) {
		exception := Catch(func() {
			panic(fmt.Errorf("test error"))
		})

		require.NotNil(t, exception)
		require.Equal(t, InternalError.Message, exception.Message)
	})

	t.Run("defined exception", func(t *testing.T) {
		exception := Catch(func() {
			panic(InvalidParamErr)
		})

		require.NotNil(t, exception)
		require.Equal(t, InvalidParamErr.Message, exception.Message)
	})

	t.Run("no exception", func(t *testing.T) {
		exception := Catch(func() {
			// ok
		})

		require.Nil(t, exception)
	})

}

func TestExceptionToError(t *testing.T) {
	// non-nil exception
	except := &Exception{}
	var err error
	err = except
	assert.NotNil(t, err)

	// nil exception
	except = nil
	err = except
	assert.Nil(t, err)

	// func return a non-nil exception
	except = func() *Exception {
		return &Exception{}
	}()
	err = except
	assert.NotNil(t, err)

	// func return a nil exception
	except = func() *Exception {
		return nil
	}()
	err = except
	assert.Nil(t, err)

	// func return a non-nil exception as error
	err = func() error {
		return &Exception{}
	}()
	assert.NotNil(t, err)

	// func return a nil exception as error
	err = func() error {
		var exception *Exception
		return exception
	}()
	assert.Nil(t, err)
}
