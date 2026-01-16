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

# Insertar D después de B
aux = lista
while True:
    if aux.dato == "B":
        nuevo = Nodo("D")
        nuevo.sig = aux.sig
        aux.sig = nuevo
        break
    aux = aux.sig
    if aux == lista:
        break

# Recorrido
print("Lista luego de insertar D después de B:")
aux = lista
while aux.sig != lista:
    print(aux.dato)
    aux = aux.sig
print(aux.dato)
