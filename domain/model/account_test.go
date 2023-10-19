package model_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"

	"github.com/henriquevschroeder/code-pix/domain/model"
	"github.com/stretchr/testify/require"
)

func TestModel_NewAccount(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, _ := model.NewBank(code, name)

	accountNumber := "abcnumber"
	ownerName := "Henrique"
	account, err := model.NewAccount(bank, accountNumber, ownerName)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(account.ID))
	require.Equal(t, account.Number, accountNumber)

	_, err = model.NewAccount(bank, "", ownerName)
	require.NotNil(t, err)
}
