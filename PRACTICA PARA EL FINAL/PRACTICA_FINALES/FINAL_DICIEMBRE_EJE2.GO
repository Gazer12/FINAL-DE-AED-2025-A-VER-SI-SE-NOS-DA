ACCION ESTADISTICAS_ATLETISMO ES
AMBIENTE
    info = REGISTRO
        clave = REGISTRO
            pais : AN(30)
            equipo : AN(30)
        FINREGISTRO
        atleta : AN(30)
        disciplina : AN(10)
        tiempo : N(4)
        descalificado : AN(2)
    FINREGISTRO

    arch : ARCHIVO DE info ORDENADO POR clave
    reg : info

    nodo = REGISTRO
        equipo : AN(30)
        cantidad : N(4)
        prox : PUNTERO A nodo
    FINREGISTRO

    prim, p, q : PUNTERO A nodo

    salida = REGISTRO
        pais : AN(30)
        tiempo : N(4)
    FINREGISTRO

    arch_sal : ARCHIVO DE salida

    res_pais, res_equipo : AN(30)
    acum_equipo, acum_pais, total_general : N(6)
    equipo_descalificado : LOGICO


PROCEDIMIENTO inicializar() ES
    ABRIR E/(arch)
    ABRIR S/(arch_sal)
    LEER(arch, reg)
    res_pais := reg.clave.pais
    res_equipo := reg.clave.equipo
    prim := NIL
    total_general := 0
FINPROCEDIMIENTO


PROCEDIMIENTO insertar(q : PUNTERO A nodo) ES
    SI prim = NIL ENTONCES
        prim := q
    SINO
        p := prim
        MIENTRAS p.prox <> NIL HACER
            p := *p.prox
        FINMIENTRAS
        *p.prox := q
    FINSI
FINPROCEDIMIENTO


PROCEDIMIENTO contar_descalificado(e : AN(30)) ES
    p := prim
    MIENTRAS (p <> NIL) Y (*p.equipo <> e) HACER
        p := *p.prox
    FINMIENTRAS

    SI p = NIL ENTONCES//la condicion tiene que ser nil ya que si no esta el nombre del equivo en la lista avaza hasta el final lo que ahce que p=nil
        NUEVO(q)
        *q.equipo := e
        *q.cantidad := 1
        *q.prox := NIL
        insertar(q)
    SINO
        p.cantidad := p.cantidad + 1
    FINSI
FINPROCEDIMIENTO


PROCEDIMIENTO corte_equipo() ES
    ESCRIBIR("Equipo: ", res_equipo,
             " - Tiempo total: ", acum_equipo)

    SI (NO equipo_descalificado) Y ((reg.disciplina = "4x100") O (reg.disciplina = "4x400")) Y (acum_equipo < 160) ENTONCES
        salida.pais := res_pais
        salida.tiempo := acum_equipo
        GRABAR(arch_sal, salida)
    FINSI

    acum_pais := acum_pais + acum_equipo
    total_general := total_general + acum_equipo

    acum_equipo := 0
    equipo_descalificado := FALSO
    res_equipo := reg.clave.equipo
FINPROCEDIMIENTO


PROCEDIMIENTO corte_pais() ES
    corte_equipo()
    ESCRIBIR("País: ", res_pais,
             " - Tiempo total: ", acum_pais)
    acum_pais := 0
    res_pais := reg.clave.pais
FINPROCEDIMIENTO


PROCESO
    inicializar()
    acum_equipo := 0
    acum_pais := 0
    equipo_descalificado := FALSO

    MIENTRAS NFDA(arch) HACER
        SI res_pais <> reg.clave.pais ENTONCES
            corte_pais()
        SINO
            SI res_equipo <> reg.clave.equipo ENTONCES
                corte_equipo()
            FINSI
        FINSI

        SI reg.descalificado = "Si" ENTONCES
            equipo_descalificado := VERDADERO
            contar_descalificado(reg.clave.equipo)
        SINO
            acum_equipo := acum_equipo + reg.tiempo
        FINSI

        LEER(arch, reg)
    FINMIENTRAS

    corte_pais()
    ESCRIBIR("Total general del torneo: ", total_general)
	P:=prim 
	mientras p<>nil hacer 
		esc("el equipo", *p.nombre," tiene la cantidad de descalificados de ", *p.descalificado)
		p:=*p.prox
	finmientras

    CERRAR(arch)
    CERRAR(arch_sal)
FINACCION
