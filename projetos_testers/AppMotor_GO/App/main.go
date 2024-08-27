package app

import (
	"fmt"

	"github.com/rzjprogramador/AppMotor_GO/App/externals/libs_rzj"
	testers "github.com/rzjprogramador/AppMotor_GO/App/testers"
)

func Main_App() {
	fmt.Println(" ========= Main do AppMotor ========= .")

	fmt.Println(" ------- VEICULO ------- ")
	testers.TesterVeiculo()

	fmt.Println(" ------- PESSOA ------- ")
	testers.TesterPessoa()

	fmt.Println(" ------- libs_rzj_TESTER ------- ")
	fmt.Println(libs_rzj.UseSoma(2, 2))
}
