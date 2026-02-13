ACCION es
AMBIENTE
    viajes = REGISTRO
        id_viaje : N(8)
        provincia : 1..23
        mes : 1..12
        cant_dia : N(2)
        cant_personas : N(2)
        alojamiento : ("H","D","NI")
        deuda : N(6,2)
        credito : N(6,2)
    FINREGISTRO

    arch : archivo de viajes
    reg : viajes

    c1, c2, c3 : entero
    max, min : entero
    total, person : entero
    pro, mesMax : entero

    provincias : arreglo [1..23] de cadena
    personasProv : arreglo [1..23] de entero
    c : arreglo [1..23,1..12] de real        // deuda
    d : arreglo [1..23,1..12] de entero      // cantidad de viajes

PROCEDIMIENTO contar() ES
    SI reg.cant_dia <= 3 ENTONCES
        c1 := c1 + 1
    SINO
        SI reg.cant_dia < 10 ENTONCES
            c2 := c2 + 1
        SINO
            c3 := c3 + 1
        FINSI
    FINSI
FINPROCEDIMIENTO


PROCEDIMIENTO mostrar() ES
    // provincia con más y menos personas
    max := personasProv[1]
    min := personasProv[1]
    pro := 1

    PARA i := 1 HASTA 23 HACER
        SI personasProv[i] > max ENTONCES
            max := personasProv[i]
            pro := i
        FINSI
        SI personasProv[i] < min ENTONCES
            min := personasProv[i]
        FINSI
    FINPARA

    // deuda por provincia y mes
    PARA i := 1 HASTA 23 HACER
        PARA j := 1 HASTA 12 HACER
            SI c[i,j] > 0 ENTONCES
                ESC("En el mes ", j, " la provincia ",provincias[i], " debe $", c[i,j])
            FINSI
        FINPARA
    FINPARA

    // mayor cantidad de viajes
    max := 0
    PARA i := 1 HASTA 23 HACER
        PARA j := 1 HASTA 12 HACER
            SI d[i,j] > max ENTONCES
                max := d[i,j]
                pro := i
                mesMax := j
            FINSI
        FINPARA
    FINPARA

    // combinaciones sin viajes
    PARA i := 1 HASTA 23 HACER
        PARA j := 1 HASTA 12 HACER
            SI d[i,j] = 0 ENTONCES
                ESC("La provincia ", provincias[i]," en el mes ", j, " no tuvo viajes")
            FINSI
        FINPARA
    FINPARA

    ESC("En el mes ", mesMax, " la provincia ",provincias[pro], " tuvo la mayor cantidad de viajes")
    ESC("Provincia con más personas: ", provincias[pro])
    ESC("Provincia con menos personas: ", min)
    ESC("Cantidad total de viajes: ", total)
    ESC("Cantidad de personas que viajaron en diciembre: ", person)
FINPROCEDIMIENTO


PROCESO
    ABRIR E/(arch)
	LEER(arch, reg)


    // inicializaciones
    c1 := 0
    c2 := 0
    c3 := 0
    total := 0
    person := 0

    PARA i := 1 HASTA 23 HACER
        personasProv[i] := 0
        PARA j := 1 HASTA 12 HACER
            c[i,j] := 0
            d[i,j] := 0
        FINPARA
    FINPARA

    MIENTRAS NFDA(arch) HACER
        total := total + 1
        contar()

        personasProv[reg.provincia] :=personasProv[reg.provincia] + reg.cant_personas

        SI reg.deuda > 0 ENTONCES
            c[reg.provincia, reg.mes] :=c[reg.provincia, reg.mes] + reg.deuda
        FINSI
        d[reg.provincia, reg.mes] :=d[reg.provincia, reg.mes] + 1

        SI reg.mes = 12 ENTONCES
            person := person + reg.cant_personas
        FINSI
    FINMIENTRAS

    CERRAR(arch)
    mostrar()
FINACCION
