package transaction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StartTransaction(t *testing.T) {
	root := &Transactions{}
	root.StartTransaction()
	assert.Nil(t, root.curr.parent)
}

func Test_Commit(t *testing.T) {
	root := &Transactions{}
	root.Commit()
	err := root.Write("10", "20")
	assert.Nil(t, err)
	read, err := root.Read("10")
	assert.Nil(t, err)
	assert.Equal(t, "20", read)
	root.StartTransaction()
	root.Write("10", "40")
	read, err = root.Read("10")
	assert.Nil(t, err)
	assert.Equal(t, "40", read)
	root.Commit()
	read, err = root.Read("10")
	assert.Nil(t, err) //test persisted parent store
	assert.Equal(t, "40", read)
	err = root.Commit()
	assert.NotNil(t, err)
	root = nil
	err = root.Write("1000", "1000")
	assert.NotNil(t, err)
}

func Test_Delete(t *testing.T) {
	root := &Transactions{}
	root.Write("10", "20")
	root.Delete("10")
	val, err := root.Read("10")
	assert.NotNil(t, err)
	assert.Empty(t, val)
	root = nil
	root.Delete("10")
	assert.NotNil(t, err)
}

func Test_Abort(t *testing.T) {
	root := &Transactions{}
	err := root.Abort()
	assert.NotNil(t, err)
	root.StartTransaction()
	root.Write("1", "2")
	err = root.Abort()
	assert.Nil(t, err)
	val, err := root.Read("1")
	assert.NotNil(t, err)
	assert.Empty(t, val)
}

func Test_Read(t *testing.T) {
	root := &Transactions{}
	root.Write("10", "20")
	val, err := root.Read("10")
	assert.Nil(t, err)
	assert.Equal(t, "20", val)
	root.StartTransaction()
	val, err = root.Read("10")
	assert.Nil(t, err)
	assert.Equal(t, "20", val)
	root.StartTransaction()
	root.Delete("10")
	root.Abort()
	val, err = root.Read("10")
	assert.Nil(t, err)
	assert.Equal(t, "20", val)
	root.StartTransaction()
	root.Write("20", "40")
	root.StartTransaction()
	val, err = root.Read("20")
	assert.Nil(t, err)
	assert.Equal(t, "40", val)
}
