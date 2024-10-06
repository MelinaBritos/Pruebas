package rutas

import "errors"

func verificarAtributos(dni string, nombre string, apellido string) []error {

	var errorList []error
	err := verificarDni(dni)

	appendError := func(err error) {
		errorList = append(errorList, err)
	}

	if err != nil {
		appendError(err)
	}
	err = verificarNombre(nombre)

	if err != nil {
		appendError(err)
	}

	err = verificarApellido(apellido)

	if err != nil {
		appendError(err)
	}

	return errorList
}

func verificarDni(dni string) error {

	if len(dni) != 8 {
		err := errors.New("el dni no tiene 8 caracteres")
		return err
	}

	return nil

}

func verificarNombre(nombre string) error {
	if len(nombre) < 3 {
		err := errors.New("el nombre debe tener al menos 3 caracteres")
		return err
	}
	return nil
}

func verificarApellido(apellido string) error {
	if len(apellido) < 3 {
		err := errors.New("el apellido debe tener al menos 3 caracteres")
		return err
	}
	return nil
}