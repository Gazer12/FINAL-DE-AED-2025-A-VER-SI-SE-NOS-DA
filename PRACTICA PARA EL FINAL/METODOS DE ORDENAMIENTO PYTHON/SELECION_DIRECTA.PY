# Método de Ordenamiento por Selección Directa en Python

# Paso 1: Definir una lista vacía para almacenar los números
numeros = []

# Paso 2: Solicitar al usuario que ingrese 5 números
# Usamos un bucle para pedir uno por uno
for i in range(5):
    num = int(input(f"Ingrese el número {i+1}: "))  # Convierte la entrada a entero
    numeros.append(num)  # Agrega el número a la lista

# Mostrar la lista antes de ordenar
print("\nLista antes de ordenar:", numeros)

# Paso 3: Algoritmo de Selección Directa
# Recorremos desde el primer elemento hasta el penúltimo
for i in range(len(numeros) - 1):#len(numeros) devuelve la cantidd de elemento de la lista es decir en este caso 5 
    # Guardamos el índice del elemento mínimo encontrado
    min_index = i  
    
    # Recorremos la parte no ordenada de la lista
    for j in range(i + 1, len(numeros)):
        if numeros[j] < numeros[min_index]:  # Si encontramos un número menor(el signo < o > lo usamos si queremos que se ordene de menor a mayor o alrevez)
            min_index = j  # Guardamos su posición

        # Guardamos el valor actual antes de intercambiar
        aux = numeros[i]
        numeros[i] = numeros[min_index]
        numeros[min_index] = aux

# Paso 4: Mostrar la lista ya ordenada
print("Lista después de ordenar:", numeros)
