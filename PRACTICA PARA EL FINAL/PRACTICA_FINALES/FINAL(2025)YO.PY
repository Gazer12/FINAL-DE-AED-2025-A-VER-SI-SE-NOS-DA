import tkinter as tk
from tkinter import filedialog

# --- Funciones según tu estilo ---
def ver(x):
    x = x.lower()
    if x == 'a': return 1
    elif x == 'b': return 2
    elif x == 'c': return 3
    elif x == 'd': return 4
    elif x == 'e': return 5
    elif x == 'f': return 6
    elif x == 'g': return 7
    elif x == 'h': return 8
    elif x == 'i': return 9
    elif x == 'j': return 10
    elif x == 'k': return 11
    elif x == 'l': return 12
    elif x == 'm': return 13
    elif x == 'n': return 14
    elif x == 'o': return 15
    elif x == 'p': return 16
    elif x == 'q': return 17
    elif x == 'r': return 18
    elif x == 's': return 19
    elif x == 't': return 20
    elif x == 'u': return 21
    elif x == 'v': return 22
    elif x == 'w': return 23
    elif x == 'x': return 24
    elif x == 'y': return 25
    elif x == 'z': return 26
    else: return 0

def num(x):
    if x == 1: return 'a'
    elif x == 2: return 'b'
    elif x == 3: return 'c'
    elif x == 4: return 'd'
    elif x == 5: return 'e'
    elif x == 6: return 'f'
    elif x == 7: return 'g'
    elif x == 8: return 'h'
    elif x == 9: return 'i'
    elif x == 10: return 'j'
    elif x == 11: return 'k'
    elif x == 12: return 'l'
    elif x == 13: return 'm'
    elif x == 14: return 'n'
    elif x == 15: return 'o'
    elif x == 16: return 'p'
    elif x == 17: return 'q'
    elif x == 18: return 'r'
    elif x == 19: return 's'
    elif x == 20: return 't'
    elif x == 21: return 'u'
    elif x == 22: return 'v'
    elif x == 23: return 'w'
    elif x == 24: return 'x'
    elif x == 25: return 'y'
    elif x == 26: return 'z'
    else: return ''

# --- Nodo para lista enlazada ---
class Nodo:
    def __init__(self, letra):
        self.letra = letra
        self.prox = None

# --- Procedimiento para insertar nodo al final ---
def insertar(prim, nodo):
    if prim is None:
        return nodo
    else:
        p = prim
        while p.prox:
            p = p.prox
        p.prox = nodo
        return prim

# --- Proceso principal ---
def codificar_archivo(ruta, k):
    arr = [num(i) for i in range(1, 27)]

    with open(ruta, 'r', encoding='utf-8') as f:
        contenido = f.read()

    prim = None
    for v in contenido:
        if v.isalpha():
            pepe = ver(v)
            tito = arr[(pepe + k - 1) % 26]
            q = Nodo(tito)
            prim = insertar(prim, q)

    p = prim
    mensaje_codificado = ""
    while p:
        mensaje_codificado += p.letra
        p = p.prox

    print("Mensaje codificado:", mensaje_codificado)

# --- Programa ---
root = tk.Tk()
root.withdraw()  # Oculta la ventana principal

print("Seleccione el archivo .txt que desea codificar")
ruta = filedialog.askopenfilename(filetypes=[("Archivos de texto", "*.txt")])
if not ruta:
    print("No se seleccionó ningún archivo.")
else:
    k = int(input("Ingrese el número de letras a desplazar: "))
    codificar_archivo(ruta, k)



