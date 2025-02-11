package iteracao

func Repetir(letra string, vezes int) (repetead string) {
	if vezes == 0 {
		vezes = 5
	}
	for i := 0; i < vezes; i++ {
		repetead += letra
	}
	return
}
