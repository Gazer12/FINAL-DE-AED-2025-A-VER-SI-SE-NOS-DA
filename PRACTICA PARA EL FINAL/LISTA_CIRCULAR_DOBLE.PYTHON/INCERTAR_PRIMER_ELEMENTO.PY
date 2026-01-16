class NodoDoble:
    def __init__(self, dato):
        self.dato = dato
        self.sig = None
        self.ant = None


lista = None

nuevo = NodoDoble("A")
nuevo.sig = nuevo
nuevo.ant = nuevo
lista = nuevo

print("Lista circular doble (un solo elemento):")
print(lista.dato)
