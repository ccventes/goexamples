package main

import (
	"fmt"
)

type Arbol struct {
	Raiz      int64
	Izquierdo *Nodoizquierdo
	Derecho   *NodoDerecho
}

type Nodoizquierdo struct {
	Valor     int64
	Eshoja    bool
	Izquierdo *Nodoizquierdo
	Derecho   *NodoDerecho
}
type NodoDerecho struct {
	Valor     int64
	Eshoja    bool
	Izquierdo *Nodoizquierdo
	Derecho   *NodoDerecho
}

func CrearNodoIzquierdo(vnodo int64) *Nodoizquierdo {
	return &Nodoizquierdo{ //cuando ya estoy creando un objeto debo separarlos con comas
		Valor:  vnodo,
		Eshoja: true, //
	}

}
func CrearNodoDerecho(vnodo int64) *NodoDerecho {
	return &NodoDerecho{ //cuando ya estoy creando un objeto debo separarlos con comas
		Valor:  vnodo,
		Eshoja: true, //
	}

}

func CrearArbol(vraiz int64, izq int64, der int64) Arbol {

	I := CrearNodoIzquierdo(min(izq, der))
	D := CrearNodoDerecho(max(izq, der))
	arbol := Arbol{
		Raiz:      vraiz,
		Izquierdo: I,
		Derecho:   D,
	}
	return arbol

}

func InsertarArbol(A Arbol, nodo int64, actual int64) {

	//fmt.Println(A.Derecho.Valor)
	//if(nodo >= A.Derecho.Valor )

	//CASO DE FINALIZACIÓN
	//Primero es la raiz  actual = raiz
	if actual == A.Raiz {
		fmt.Println("-----------entró al if-----------")
		fmt.Println("nodo", nodo, "valor izquierdo", A.Izquierdo.Valor, "es menor ?", nodo < A.Derecho.Valor)

		if nodo < A.Izquierdo.Valor {

			A.Izquierdo = InsertarNodoI(nodo, A.Izquierdo) // esta necesario retornar el nuevo nodo
		}
		if nodo >= A.Izquierdo.Valor {

			A.Derecho = InsertarNodoD(nodo, A.Derecho) // esta necesario retornar el nuevo nodo
		}

	}

}

func InsertarNodoI(valor int64, I *Nodoizquierdo) *Nodoizquierdo {
	if I == nil {
		// Si el nodo izquierdo es nil, crea un nuevo nodo izquierdo
		NuevoIzquierdo := CrearNodoIzquierdo(valor)
		fmt.Println("nodo izquierdo creado con valor", NuevoIzquierdo.Valor)
		return NuevoIzquierdo
	}

	// Si el nodo izquierdo no es nil, simplemente actualiza su valor
	I.Eshoja = false
	I.Izquierdo = InsertarNodoI(valor, I.Izquierdo)

	return I
}
func InsertarNodoD(valor int64, D *NodoDerecho) *NodoDerecho { //yo le haria refactor a esto pero bueno
	if D == nil {
		// Si el nodo izquierdo es nil, crea un nuevo nodo izquierdo
		NuevoDerecho := CrearNodoDerecho(valor)
		fmt.Println("nodo derecho creado con valor", NuevoDerecho.Valor)
		return NuevoDerecho
	} //verifica que no sea null el nodo (no hay nodo), si no hay nodo lo crea  y vuelve a llamar a la función insertar
	D.Eshoja = false
	D.Derecho = InsertarNodoD(valor, D.Derecho) //segun yo esto lo esta haciendo dos veces
	return D

}

func main() {

	A := CrearArbol(20, 19, 21) //raiz con dos hijos
	fmt.Println("Arbol en la dirección ", A, "con raiz", A.Raiz, "nodo Izquierdo",
		A.Izquierdo.Valor, "nodo derecho", A.Derecho.Valor, "El hijo de la derecha",
		A.Derecho.Derecho)

	//vamos a insertar un nuevo nodo
	fmt.Println("aqui inserto una hoja")
	InsertarArbol(A, 18, A.Raiz)
	fmt.Println("Acabo de insertar un nuevo nodo de valor",
		A.Izquierdo.Izquierdo.Valor)

	fmt.Println("-------------------------Nuevamente--------------------")
	fmt.Println("Arbol en la dirección: ", A, "con raiz: ", A.Raiz,
		"nodo Izquierdo: ", A.Izquierdo.Valor, "Nuevo HIjo Izquierdo:",
		A.Izquierdo.Izquierdo.Valor, "\n--------------------------------")

}
