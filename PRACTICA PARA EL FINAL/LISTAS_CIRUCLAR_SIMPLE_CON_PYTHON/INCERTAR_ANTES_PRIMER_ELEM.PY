class Nodo:
    def __init__(self, dato):
        self.dato = dato
        self.sig = None


# Crear lista A -> B -> C -> A
A = Nodo("A")
B = Nodo("B")
C = Nodo("C")

A.sig = B
B.sig = C
C.sig = A

lista = A

# Insertar al principio (D)
nuevo = Nodo("D")
nuevo.sig = lista

aux = lista
while aux.sig != lista:
    aux = aux.sig

aux.sig = nuevo
lista = nuevo

# Recorrido
print("Lista luego de insertar D al principio:")
aux = lista
while aux.sig != lista:
    print(aux.dato)
    aux = aux.sig
print(aux.dato)
