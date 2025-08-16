# Algoritmo de Inserción Directa en Python
# El usuario ingresa 5 números y los ordenamos paso a paso

# Paso 1: Definir la lista y pedir datos al usuario
A = []
for i in range(5):
    num = int(input(f"Ingrese el número {i+1}: "))
    A.append(num)

print("\nLista antes de ordenar:", A)

# Paso 2: Algoritmo de Inserción Directa
# Comenzamos desde la segunda posición (índice 1 en Python) hasta el final
for i in range(1, len(A)):
    X = A[i]      # Guardamos el elemento actual que vamos a insertar en su lugar correcto
    J = i - 1     # Posición del elemento anterior a 'X'
    
    # Mientras no salgamos de la lista (J >= 0) y el elemento 'X' sea menor que A[J]
    while J >= 0 and X < A[J]:
        A[J + 1] = A[J]  # Desplazamos hacia la derecha el elemento A[J]
        J = J - 1        # Avanzamos hacia la izquierda para seguir comparando
    
    # Cuando encontramos la posición correcta, insertamos el valor de 'X'
    A[J + 1] = X

# Paso 3: Mostrar lista ordenada
print("Lista después de ordenar:", A)
