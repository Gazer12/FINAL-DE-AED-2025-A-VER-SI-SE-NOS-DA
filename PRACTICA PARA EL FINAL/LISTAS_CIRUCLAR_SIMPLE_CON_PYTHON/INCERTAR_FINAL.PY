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

# Insertar al final (D)
nuevo = Nodo("D")

aux = lista
while aux.sig != lista:
    aux = aux.sig

aux.sig = nuevo
nuevo.sig = lista

# Recorrido
print("Lista luego de insertar D al final:")
aux = lista
while aux.sig != lista:
    print(aux.dato)
    aux = aux.sig
print(aux.dato)
