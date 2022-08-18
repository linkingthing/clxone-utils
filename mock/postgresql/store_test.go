package postgresql

import (
	"errors"
	"testing"

	restdb "github.com/linkingthing/gorest/db"
	restresource "github.com/linkingthing/gorest/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type User struct {
	restresource.ResourceBase `json:",inline"`
	Name                      string
	Age                       int
}

var TableUser = restdb.ResourceDBType(&User{})

func TestStore(t *testing.T) {
	db, mocker, err := NewMocker()
	require.NoError(t, err)
	u := &User{
		Name: "joker",
		Age:  12,
	}

	var expectFillUsers []*User
	expectFillUsers = append(expectFillUsers,
		&User{Name: "jj", Age: 20},
		&User{Name: "jj1", Age: 21},
	)
	mocker.ExpectExec(ExpectMethodFill).ReturnResources(&expectFillUsers).ReturnError(nil)
	mocker.ExpectExec(ExpectMethodExists).ReturnBool(false).ReturnError(nil)
	mocker.ExpectExec(ExpectMethodUpdate).ReturnInt64(1).ReturnError(nil)
	mocker.ExpectExec(ExpectMethodInsert).ReturnError(nil).ReturnError(nil)
	mocker.ExpectExec(ExpectMethodFillEx).ReturnResources(nil).ReturnError(errors.New("fill failed"))
	mocker.ExpectExec(ExpectMethodGet).ReturnResources(&expectFillUsers).ReturnError(nil)
	mocker.ExpectExec(ExpectMethodGetEx).ReturnResources(&expectFillUsers).ReturnError(nil)
	mocker.ExpectExec(ExpectMethodCount).ReturnInt64(1).ReturnError(nil)
	mocker.ExpectExec(ExpectMethodDelete).ReturnInt64(1).ReturnError(nil)

	var users []*User
	err = restdb.WithTx(db, func(tx restdb.Transaction) error {
		err := tx.Fill(nil, &users)
		assert.NoError(t, err)
		assert.Equal(t, 2, len(users))

		ok, err := tx.Exists(TableUser, nil)
		assert.NoError(t, err)
		assert.Equal(t, false, ok)

		c, err := tx.Update(TableUser,
			map[string]interface{}{"name": "joker_new"},
			map[string]interface{}{restdb.IDField: "123"})
		assert.NoError(t, err)
		assert.Equal(t, int64(1), c)

		_, err = tx.Insert(u)
		assert.NoError(t, err)

		err = tx.FillEx(&users, "select * from gr_user")
		assert.Equal(t, err, errors.New("fill failed"))

		users0, err := tx.Get(TableUser, map[string]interface{}{restdb.IDField: "123"})
		assert.NoError(t, err)
		assert.Equal(t, &expectFillUsers, users0)

		users1, err := tx.GetEx(TableUser, "select * from gr_user")
		assert.NoError(t, err)
		assert.Equal(t, &expectFillUsers, users1)

		c1, err := tx.Count(TableUser, nil)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), c1)

		c2, err := tx.Delete(TableUser, map[string]interface{}{restdb.IDField: "123"})
		assert.NoError(t, err)
		assert.Equal(t, int64(1), c2)

		return nil
	})

	require.NoError(t, err)
}
