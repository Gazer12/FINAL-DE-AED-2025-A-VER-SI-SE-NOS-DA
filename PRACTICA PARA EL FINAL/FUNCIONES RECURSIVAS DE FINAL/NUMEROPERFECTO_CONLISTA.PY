# Traducción directa del pseudocódigo a Python (sin mejoras lógicas)
# Lista doblemente enlazada + función recursiva "pepe"

class Nodo:
    def __init__(self, dato):
        self.dato = dato
        self.prox = None
        self.ant = None

# ----- función recursiva tal como en el pseudocódigo -----
def pepe(x, d, suma):
    if d == 0:
        if suma == x:
            return True
        else:
            return False
    else:
        if x % d == 0:
            suma = suma + d
            return pepe(x, d-1, suma)
        else:
            return pepe(x, d-1, suma)

# ----- armar lista con 5 números (6 y 28 son perfectos) -----
valores = [4, 6, 10, 28, 9]

prim = None
ult = None

for v in valores:
    nuevo = Nodo(v)
    if prim is None:
        prim = nuevo
        ult = nuevo
    else:
        ult.prox = nuevo
        nuevo.ant = ult
        ult = nuevo

# mostrar lista inicial
def mostrar_lista(titulo):
    print(titulo)
    p = prim
    while p is not None:
        print(p.dato, end=" ")
        p = p.prox
    print("\n")

mostrar_lista("Lista original:")

# ----- procedimiento eliminar (adaptado a Python) -----
def eliminar(p):
    global prim, ult
    if p == prim:
        prim = p.prox
        if prim is not None:
            prim.ant = None
    elif p == ult:
        ult = p.ant
        ult.prox = None
    else:
        p.ant.prox = p.prox
        p.prox.ant = p.ant

# ----- proceso principal -----
p = prim
while p is not None:
    sig = p.prox  # guardar siguiente
    if pepe(p.dato, p.dato - 1, 0):
        eliminar(p)
    p = sig

# mostrar lista final
mostrar_lista("Lista luego de eliminar números perfectos:")
