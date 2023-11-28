package main

import (
	"bufio"
	"fmt"
	"os"
)

type Arbol struct {
	Raiz *Nodo
}

type Nodo struct {
	Valor     int64
	Eshoja    bool
	Izquierdo *Nodo
	Derecho   *Nodo
}

func CrearArbol() Arbol {

	return Arbol{

		Raiz: nil,
	}
}

// Hijo -> Padre
func Dir(Actual Nodo, Compare Nodo) bool {
	//Derecha true
	//izquierda false
	fmt.Println("::Calcilando Dir para:")
	fmt.Println("::Actual::", Actual.Valor)
	fmt.Println("::compare::", Compare.Valor)
	if Actual.Valor >= Compare.Valor {
		fmt.Println("//Dir retornara false::///")
		return false
	} else if Actual.Valor < Compare.Valor {
		fmt.Println("//Dir retornara true::///")
		return true
	}
	return true

}
func Move(Padre Nodo, Hijo Nodo) Nodo { // función para cambiar recursivamente a que nodo me quiero mover

	fmt.Println("Presiona Enter")
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
	fmt.Println("Continua")
	fmt.Println("Entro a Move", Padre.Valor)
	fmt.Println("valor nodo", Padre.Valor)
	dir := Dir(Padre, Hijo)
	fmt.Println("dir", dir)

	if Padre.Izquierdo == nil && dir == false {

		fmt.Println("caso 1 JAJAJAJJA", Padre.Izquierdo)
		Padre.Izquierdo = &Hijo
		fmt.Println("caso 1 JEJEJEJEJEJEJEJEJE", Padre.Izquierdo)

	} else if Padre.Derecho == nil && dir == true {

		Padre.Derecho = &Hijo

	} else if Padre.Izquierdo != nil && dir == false {

		fmt.Println("VA a entrar a la recursividad izquierdada")
		*Padre.Izquierdo = Move(*Padre.Izquierdo, Hijo)

	} else if Padre.Derecho != nil && dir == true {

		fmt.Println("VA a entrar a la recursividad Derecha")
		*Padre.Derecho = Move(*Padre.Derecho, Hijo)

	}
	fmt.Println("SIN acumulator Padre es", Padre)
	return Padre

}
func InsertarArbol(Arb Arbol, n Nodo) Arbol {
	fmt.Println("|||||||||||||Insertando Nodo de valor : ", n.Valor)
	if Arb.Raiz == nil { // Caso 1, el arbol esta vacio
		Arb.Raiz = &n
		return Arb
	} else if Arb.Raiz != nil && (Arb.Raiz.Izquierdo == nil || Arb.Raiz.Derecho == nil) {
		// caso 2 el arbol no esta vacio, pero solo contiene un elemento
		fmt.Println("+IF caso 2+")
		if Dir(*Arb.Raiz, n) == true {
			fmt.Println("+Insertar a la derecha+")
			Arb.Raiz.Derecho = &n
			fmt.Println("+Arbol insertado a la derecha+", Arb.Raiz.Derecho)
			return Arb

		} else {
			fmt.Println("+Insertar a la izquierda+")
			Arb.Raiz.Izquierdo = &n
			fmt.Println("+Arbol insertado a la derecha+", Arb.Raiz.Izquierdo)
			return Arb
		}

	} else if Arb.Raiz != nil && (Arb.Raiz.Izquierdo != nil && Arb.Raiz.Derecho != nil) {
		fmt.Println("Entro Al if de Move")
		NodeChain := Move(*Arb.Raiz, n)
		fmt.Println("NodeChain", NodeChain.Izquierdo)
		Arb.Raiz = &NodeChain
	}
	return Arb

}
func crearNodo(valor int64) Nodo {

	return Nodo{

		Valor: valor,
	}
}

func main() {

	A := CrearArbol()
	fmt.Println(" He creado un arbol ta vacio? ", A.Raiz)
	//cadena := [...]int64{16, 19, 7, 1, 13, 9, 3}
	cadena := [...]int64{16, 19, 7, 6, 5, 20, 9, 3}
	// 16
	fmt.Println("Primer valor de la cadena", cadena[0])
	fmt.Println("----creando un  nodo, el primer valor de la cadena----")
	fmt.Println("--Primer Nodo insertado con valor:")
	n := crearNodo(cadena[0])
	A = InsertarArbol(A, n)
	fmt.Println("--Nodo insertado con valor:", A, A.Raiz.Valor)
	fmt.Println("")

	//19
	n2 := crearNodo(cadena[1])
	fmt.Println("----insertando nodo en el arbol----")
	A = InsertarArbol(A, n2)
	fmt.Println("--Nodo insertado con valor:", A.Raiz.Derecho.Valor)
	fmt.Println("")

	//7
	fmt.Println("----creando un  nodo, el tercer valor de la cadena----")
	n3 := crearNodo(cadena[2])
	fmt.Println("----insertando nodo en el arbol----")
	A = InsertarArbol(A, n3)
	fmt.Println("--  Nodo insertado con valor:", A.Raiz.Izquierdo.Valor)
	fmt.Println("")

	//6
	fmt.Println("----creando un  nodo, el cuarto valor de la cadena----")
	n4 := crearNodo(cadena[3])
	A = InsertarArbol(A, n4)
	fmt.Println("----( PON CUIDADO AQUI MAMAHUEVO)///insertando nodo en el arbol----")
	fmt.Println("--(GLUGLUGLU)Nodo insertado con valor:", A.Raiz.Izquierdo.Izquierdo) //

	fmt.Println("----creando un  nodo, el Quinto valor de la cadena----")
	n5 := crearNodo(cadena[4])
	fmt.Println("----insertando nodo en el arbol----")
	A = InsertarArbol(A, n5)
	fmt.Println("--GLO Nodo insertado con valor:", A.Raiz.Izquierdo.Izquierdo.Izquierdo) //

	fmt.Println("")

	fmt.Println("----creando un  nodo, el Sexto valor de la cadena----")
	n6 := crearNodo(cadena[5])
	fmt.Println("----insertando nodo en el arbol----")
	A = InsertarArbol(A, n6)
	fmt.Println("--GLO Nodo insertado con valor:", A.Raiz.Derecho.Derecho) //

}
