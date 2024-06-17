package libs

import (
	"testing"
)

func Test_RzLibs(t *testing.T) {
	// UseSoma
	sutUseSoma, _ := UseSoma(2, 2)
	resUseSoma := 4

	if sutUseSoma != resUseSoma {
		t.Errorf("Ops...Não esta funcionando a lib UseSoma() o resultado é %d", sutUseSoma)
	}

	//

}
