ACCION SEGUIMIENTO_VENTAS ES
AMBIENTE
    VIAJANTE = REGISTRO
        CLAVE = REGISTRO
            n_legajo : N(8)
            mes : AN(30)
        FINREGISTRO
        cant_ventas : N(4)
        importe : R(8,2)
    FINREGISTRO

    arch : ARCHIVO DE VIAJANTE ORDENADO POR CLAVE
    reg : VIAJANTE

    filiacion = REGISTRO
        CLAVE = REGISTRO
            n_legajo : N(8)
        FINREGISTRO
        apeynom : AN(30)
        estado_civil : AN(30)
        direccion : AN(30)
    FINREGISTRO

    arch_i : ARCHIVO INDEXADO DE filiacion
    reg_i : filiacion

    nodo = REGISTRO
        nombre : AN(30)
        cant : N(6)
        prox : PUNTERO A nodo
    FINREGISTRO

    prim, p, q, ant : PUNTERO A nodo

    res_legajo : N(8)

    cant_ventas : N(6)
    can_importe : R(10,2)

    cant_oct, cant_nov, cant_dic : R(10,2)
    total_oct, total_nov, total_dic : N(6)

-------------------------------------------------

PROCEDIMIENTO inicializar() ES
    ABRIR E/(arch)
    ABRIR E/(arch_i)

    LEER(arch, reg)

    prim := NIL

    res_legajo := reg.CLAVE.n_legajo

    cant_ventas := 0
    can_importe := 0

    cant_oct := 0
    cant_nov := 0
    cant_dic := 0

    total_oct := 0
    total_nov := 0
    total_dic := 0
FINPROCEDIMIENTO

-------------------------------------------------

PROCEDIMIENTO insertar(x : PUNTERO A nodo) ES
    SI prim = NIL O x.cant > prim.cant ENTONCES
        x.prox := prim
        prim := x
    SINO
        p := prim
        MIENTRAS p.prox <> NIL Y p.prox.cant >= x.cant HACER
            p := p.prox
        FINMIENTRAS
        x.prox := p.prox
        p.prox := x
    FINSI
FINPROCEDIMIENTO

-------------------------------------------------

PROCEDIMIENTO mostrar() ES
    p := prim
    MIENTRAS p <> NIL HACER
        ESCRIBIR("Nombre: ", p.nombre, " Total ventas: ", p.cant)
        p := p.prox
    FINMIENTRAS
FINPROCEDIMIENTO

-------------------------------------------------

PROCEDIMIENTO corte() ES
    reg_i.CLAVE.n_legajo := res_legajo
    LEER(arch_i, reg_i)

    SI EXISTE EN arch_i ENTONCES
        ESCRIBIR("Viajante: ", reg_i.apeynom)
        ESCRIBIR("Ventas realizadas: ", cant_ventas)
        ESCRIBIR("Importe total: ", can_importe)

        NUEVO(q)
        q.nombre := reg_i.apeynom
        q.cant := cant_ventas
        q.prox := NIL

        insertar(q)
    SINO
        ESCRIBIR("No existe ese viajante")
    FINSI

    cant_ventas := 0
    can_importe := 0
    res_legajo := reg.CLAVE.n_legajo
FINPROCEDIMIENTO

-------------------------------------------------

PROCESO
    inicializar()

    MIENTRAS NFDA(arch) HACER

        SI res_legajo <> reg.CLAVE.n_legajo ENTONCES
            corte()
        FINSI

        cant_ventas := cant_ventas + reg.cant_ventas
        can_importe := can_importe + reg.importe

        SEGUN reg.CLAVE.mes HACER
            "octubre":
                cant_oct := cant_oct + reg.importe
                total_oct := total_oct + 1
            "noviembre":
                cant_nov := cant_nov + reg.importe
                total_nov := total_nov + 1
            "diciembre":
                cant_dic := cant_dic + reg.importe
                total_dic := total_dic + 1
        FINSEGUN

        LEER(arch, reg)
    FINMIENTRAS

    corte()

    ESCRIBIR("Promedio Octubre: ", cant_oct / total_oct)
    ESCRIBIR("Promedio Noviembre: ", cant_nov / total_nov)
    ESCRIBIR("Promedio Diciembre: ", cant_dic / total_dic)

    ESCRIBIR("RANKING")
    mostrar()

    CERRAR(arch)
    CERRAR(arch_i)
FINPROCESO


			
							
