ACCION es
AMBIENTE
    sec : secuencia de caracter
    v, tipo_cli, e1, e2 : caracter
    pack : AN(30)
    edad, tipo : entero

    total_registros : entero
    suma_edades_nula : entero
    cant_nula : entero
    total_premium : entero
    premium_con_pref : entero

    nodo = REGISTRO
        cod_pack : AN(30)
        preferencia : arreglo [1..3] de entero   // 1=A, 2=D, 3=N
        prox : puntero a nodo
    FINREGISTRO

    prim, p, q, ant : puntero a nodo

FUNCION preferencia(x : caracter) : entero
    SEGUN x HACER
        "A" : preferencia := 1
        "D" : preferencia := 2
        "N" : preferencia := 3
    FINSEGUN
FINFUNCION

PROCEDIMIENTO cargar() ES
    si prim = nil entonces
        nuevo(q)
        *q.cod_pack := pack
        para i := 1 a 3 hacer
            *q.preferencia[i] := 0
        finpara
        *q.preferencia[tipo] := 1
        prim := q
        *q.prox := nil
    sino
        p := prim
        ant := nil
        mientras p <> nil y *p.cod_pack <> pack hacer
            ant := p
            p := *p.prox
        finmientras

        si p = nil entonces
            nuevo(q)
            *q.cod_pack := pack
            para i := 1 a 3 hacer
                *q.preferencia[i] := 0
            finpara
            *q.preferencia[tipo] := 1
            *ant.prox := q
            *q.prox := nil
        sino
            *p.preferencia[tipo] := *p.preferencia[tipo] + 1
        finsi
    finsi
FINPROCEDIMIENTO

PROCEDIMIENTO mostrar() ES
    p := prim
    mientras p <> nil hacer
        si *p.preferencia[1] > *p.preferencia[2] entonces
            esc("Pack:", *p.cod_pack)
            esc("Alta:", *p.preferencia[1])
            esc("Deseable:", *p.preferencia[2])
        finsi
        p := *p.prox
    finmientras

    si cant_nula > 0 entonces
        esc("Promedio edad preferencia Nula:",
            suma_edades_nula / cant_nula)
    finsi

    si total_premium > 0 entonces
        esc("Porcentaje Premium con preferencias:",
            (premium_con_pref ) / total_premium)
    finsi
FINPROCEDIMIENTO

PROCESO
    arr(sec, v)
    avz(sec, v)

    prim := nil
    total_registros := 0
    suma_edades_nula := 0
    cant_nula := 0
    total_premium := 0
    premium_con_pref := 0

    mientras nfds(sec) hacer
        tipo_cli := v          // B o P

        // Salteo cliente AN(6)
        para i := 1 a 6 hacer
            avz(sec, v)
        finpara

        e1 := v
        avz(sec, v)
        e2 := v
        edad := AGE(e1, e2)
        avz(sec, v)

        pack := ""
        mientras v <> "-" hacer
            pack := pack + v
            avz(sec, v)
        finmientras

        avz(sec, v)            // leo preferencia
        tipo := preferencia(v)

        si tipo = 3 entonces
            suma_edades_nula := suma_edades_nula + edad
            cant_nula := cant_nula + 1
        finsi

        si tipo_cli = "P" entonces
            total_premium := total_premium + 1
            si tipo <> 3 entonces
                premium_con_pref := premium_con_pref + 1
            finsi
        finsi

        total_registros := total_registros + 1
        cargar()
        avz(sec, v)
    finmientras

    mostrar()
FINPROCESO
FINACCION
