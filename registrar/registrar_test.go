package registrar

import "testing"

func TestNoFiresEventWhenPerformerAddingIsFail(t *testing.T) {
	reg := &Registrar{}
	_, err := reg.AddPerformer(t.Context(), Performer{})
	if err != nil {
		t.Error("failed try must return error, given nil")
		return
	}
}
