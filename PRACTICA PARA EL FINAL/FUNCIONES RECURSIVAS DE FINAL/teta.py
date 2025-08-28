def nro_valido(a, suma, i):
    # Caso base
    if i == 1:
        return (suma % 5) == 0
    else:
        if i % 2 == 1:  # posición impar (contando desde el final)
            aux = a[i-1] * 2  # Python indexa desde 0
            print("num", a[i])
            if aux > 9:
                unidad = aux % 10
                decena = aux // 10
                return nro_valido(a, suma + (unidad + decena), i - 1)
            else:
                return nro_valido(a, suma + aux, i - 1)
        else:  # posición par
            return nro_valido(a, suma + a[i-1], i - 1)


# Ejemplo de tarjeta con 20 dígitos (ficticia)
#tarjeta = [4, 5, 3, 2, 1, 7, 6, 2, 9, 0, 4, 3, 5, 8, 2, 1, 4, 7, 9, 3] tarjeta falsa
tarjeta = [4, 5, 3, 2, 1, 7, 6, 2, 9, 0, 4, 3, 5, 8, 2, 1, 4, 7, 9, 5]

# Probar la función
es_valida = nro_valido(tarjeta, 0, len(tarjeta))
print("Tarjeta válida:", es_valida)