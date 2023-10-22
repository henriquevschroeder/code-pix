package model_test

import (
	"testing"

	"github.com/henriquevschroeder/code-pix/domain/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestModel_NewPixKey(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, _ := model.NewBank(code, name)

	accountNumber := "abcd"
	ownerName := "Henrique"
	account, _ := model.NewAccount(bank, accountNumber, ownerName)

	kind := "email"
	key := "test@mail.com"
	pixKey, _ := model.NewPixKey(kind, account, key)

	require.NotEmpty(t, uuid.FromStringOrNil(pixKey.ID))
	require.Equal(t, pixKey.Kind, kind)
	require.Equal(t, pixKey.Status, "active")

	kind = "cpf"
	_, err := model.NewPixKey(kind, account, key)
	require.Nil(t, err)

	_, err = model.NewPixKey("nome", account, key)
	require.NotNil(t, err)
}
