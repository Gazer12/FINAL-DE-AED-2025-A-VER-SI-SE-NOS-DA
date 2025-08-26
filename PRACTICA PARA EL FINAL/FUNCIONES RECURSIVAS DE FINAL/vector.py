def ver(x, suma, i):
    if i == 1:
        return suma
    else:
        return ver(x, suma + x[i-1], i-1)


# Programa principal
vector = [1, 1, 1, 1, 1, 1, 1, 1, 1, 1]  # ejemplo
resultado = ver(vector, vector[0], len(vector))
print("La suma de los elementos del vector es:", resultado)

'''function ver(x:arreglo [1..10] de entero,suma,i:entero)->entero 
    si i=1 entonces 
        var:suma 
    sino 
        var(x,suma+x[i],i-1)
    finsi 
finfuncion'''

