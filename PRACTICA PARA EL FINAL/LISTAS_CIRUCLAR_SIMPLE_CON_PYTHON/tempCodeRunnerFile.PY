
aux = lista
while aux.sig != lista:
    aux = aux.sig

aux.sig = nuevo
nuevo.sig = lista