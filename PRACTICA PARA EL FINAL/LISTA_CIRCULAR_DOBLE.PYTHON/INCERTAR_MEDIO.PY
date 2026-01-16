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

# Insertar D después de B
aux = lista
while True:
    if aux.dato == "B":
        nuevo = NodoDoble("D")

        nuevo.sig = aux.sig
        nuevo.ant = aux

        aux.sig.ant = nuevo
        aux.sig = nuevo
        break

    aux = aux.sig
    if aux == lista:
        break

# Recorrido
print("Lista doble luego de insertar D después de B:")
aux = lista
while aux.sig != lista:
    print(aux.dato)
    aux = aux.sig
print(aux.dato)
