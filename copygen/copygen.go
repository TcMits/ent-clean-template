// Code generated by github.com/switchupcb/copygen
// DO NOT EDIT.

/* Specify the name of the generated file's package. */
package copygen

import (
	"github.com/TcMits/ent-clean-template/ent"
	"github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
)

// LoginInputToUserWhereInput copies a *usecase.LoginInput to a *ent.UserWhereInput.
func LoginInputToUserWhereInput(tU *ent.UserWhereInput, fL *usecase.LoginInput) {
	// *ent.UserWhereInput fields
	tU.Username = &fL.Username
}
