class Nodo:
    def __init__(self, dato):
        self.dato = dato
        self.sig = None


# Lista vacía
lista = None

# Insertar primer nodo
nuevo = Nodo("A")
nuevo.sig = nuevo
lista = nuevo

# Recorrido
print("Lista circular simple (un solo elemento):")
aux = lista
print(aux.dato)
