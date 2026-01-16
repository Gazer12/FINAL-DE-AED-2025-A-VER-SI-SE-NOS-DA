# =========================
# Definición del Nodo
# =========================
class Nodo:
    def __init__(self, dato):
        self.dato = dato
        self.sig = None   # puntero al siguiente nodo


# =========================
# Crear nodos
# =========================
A = Nodo("A")
B = Nodo("B")
C = Nodo("C")
D = Nodo("D")

# =========================
# Conectar nodos (lista circular)
# A -> B -> C -> D -> A
# =========================
A.sig = B
B.sig = C
C.sig = D
D.sig = A   # cierre circular

# =========================
# Puntero a la lista (primer nodo)
# =========================
lista = A


# =========================
# Recorrer la lista circular
# (equivalente a: MIENTRAS aux.sig <> lista)
# =========================
print("Recorrido de la lista:")
aux = lista

while aux.sig != lista:
    print(aux.dato)
    aux = aux.sig

# imprimir el último nodo
print(aux.dato)

print("Aux quedó parado en:", aux.dato)


# =========================
# Insertar un nuevo nodo E al final
# =========================
print("\nInsertando E al final...")

E = Nodo("E")

aux = lista
while aux.sig != lista:
    aux = aux.sig

# aux está en el último nodo
aux.sig = E
E.sig = lista


# =========================
# Recorrer nuevamente la lista
# =========================
print("\nRecorrido luego de insertar E:")
aux = lista

while aux.sig != lista:
    print(aux.dato)
    aux = aux.sig

print(aux.dato)
print("Aux quedó parado en:", aux.dato)
