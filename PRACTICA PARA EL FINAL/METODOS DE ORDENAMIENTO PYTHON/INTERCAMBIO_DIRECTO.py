# Pedimos al usuario ingresar 5 números y los guardamos en una lista
A = []
print("Ingresa 5 números:")
for i in range(5):
    num = int(input(f"Número {i+1}: "))
    A.append(num)

# Inicializamos la bandera en False para entrar al ciclo
bandera = False

# Mientras la bandera sea False, seguimos intentando ordenar
while not bandera:
    bandera = True  # Asumimos que el arreglo ya está ordenado

    # Recorremos el arreglo desde la posición 0 hasta la penúltima
    for j in range(len(A) - 1):
        # Si el elemento actual es mayor que el siguiente, los intercambiamos
        if A[j] > A[j + 1]:
            # Guardamos temporalmente el valor de A[j]
            x = A[j]
            # Intercambiamos los valores
            A[j] = A[j + 1]
            A[j + 1] = x

            # Como hubo un intercambio, el arreglo no está ordenado aún
            bandera = False

# Mostramos el arreglo ya ordenado
print("Arreglo ordenado:", A)
