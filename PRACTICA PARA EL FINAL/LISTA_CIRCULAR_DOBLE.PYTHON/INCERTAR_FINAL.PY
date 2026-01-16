class NodoDoble:
    def __init__(self, dato):
        self.dato = dato
        self.sig = None
        self.ant = None


# Crear lista A <-> B <-> C <-> A
A = NodoDoble("A")
B = NodoDoble("B")
C = NodoDoble("C")

A.sig = B
B.ant = A
B.sig = C
C.ant = B
C.sig = A
A.ant = C

lista = A

# Insertar D al final
nuevo = NodoDoble("D")
ultimo = lista.ant

nuevo.sig = lista
nuevo.ant = ultimo

ultimo.sig = nuevo
lista.ant = nuevo

# Recorrido
print("Lista doble luego de insertar D al final:")
aux = lista
while aux.sig != lista:
    print(aux.dato)
    aux = aux.sig
print(aux.dato)
