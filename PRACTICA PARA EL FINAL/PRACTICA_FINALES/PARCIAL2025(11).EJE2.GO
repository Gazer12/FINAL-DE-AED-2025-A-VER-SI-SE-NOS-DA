ACCION ES
AMBIENTE
    ARR : ARREGLO [1..5,1..4] DE ENTERO
    VALIDA : ARREGLO [1..5,1..3] DE ENTERO

    archivito = REGISTRO
        departamento : N(6)
        origen : 1..3
        nivel : 1..5
        validad : ("si","no")
    FINREGISTRO

    arch : ARCHIVO DE archivito
    reg : archivito

    i, j : ENTERO
    cant, total : ENTERO
    porcentaje : REAL

    min, max : ENTERO
    min_level, min_origen : ENTERO
    max_level : ENTERO

------------------------------------------------

PROCEDIMIENTO inicializar() ES
    abrir e/(arch)
    leer(arch,reg)

    PARA i:=1 A 5 HACER
        PARA j:=1 A 4 HACER
            ARR[i,j] := 0
        FINPARA
    FINPARA

    PARA i:=1 A 5 HACER
        PARA j:=1 A 3 HACER
            VALIDA[i,j] := 0
        FINPARA
    FINPARA

    cant := 0
    total := 0

    min := 999999
    max := -1
FINPROCEDIMIENTO

------------------------------------------------

PROCEDIMIENTO informar() ES

    min := ARR[1,1]

    PARA i:=1 A 5 HACER
        PARA j:=1 A 3 HACER
            SI ARR[i,j] < min ENTONCES
                min := ARR[i,j]
                min_level := i
                min_origen := j
            FINSI
        FINPARA
    FINPARA

    esc("Menor cantidad: Nivel ", min_level," Origen ", min_origen," con ", min)

    max := ARR[1,4]

    PARA i:=1 A 5 HACER
        SI ARR[i,4] > max ENTONCES
            max := ARR[i,4]
            max_level := i
        FINSI
    FINPARA

    esc("Nivel con mayor cantidad total: ",max_level, " con ", max)

    PARA i:=1 A 5 HACER
        esc("Total del nivel ", i, " = ", ARR[i,4])
    FINPARA

    porcentaje := (cant * 100) / total
    esc("Porcentaje de válidas: ", porcentaje)

    min := VALIDA[1,1]

    PARA j:=1 A 3 HACER
        suma := 0
        PARA i:=1 A 5 HACER
            suma := suma + VALIDA[i,j]
        FINPARA

        SI j = 1 O suma < min ENTONCES
            min := suma
            min_origen := j
        FINSI
    FINPARA

    esc("Origen con menos válidas: ",min_origen, " con ", min)

FINPROCEDIMIENTO

------------------------------------------------

PROCESO
    inicializar()

    MIENTRAS nfda(arch) HACER
        ARR[reg.nivel,reg.origen] :=ARR[reg.nivel,reg.origen] + 1
        ARR[reg.nivel,4] :=ARR[reg.nivel,4] + 1
        total := total + 1

        SI reg.validad = "si" ENTONCES
            VALIDA[reg.nivel,reg.origen] :=VALIDA[reg.nivel,reg.origen] + 1
            cant := cant + 1
        FINSI

        leer(arch,reg)
    FINMIENTRAS

    informar()

FINACCION
