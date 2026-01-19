ACCION resumen_ventas ES
AMBIENTE

    fecha = registro
        dia  : N(2)
        mes  : N(2)
        año  : N(4)
    FinRegistro

    ventas = registro
        clave = registro
            codvendedor : 1..10
            codproducto : 1..10
            fecha_venta : fecha
        FinRegistro
        cantidad : N(3)
        precio   : R(6)
    FinRegistro

    comisiones = registro
        clave_c = registro
            codvendedor : 1..10
            codproducto : 1..10
        FinRegistro
        comision : R(4)
    FinRegistro

    arch_v : archivo de ventas ordenado por clave
    reg_v  : ventas

    arch_c : archivo indexado de comisiones
    reg_c  : comisiones

    res_vendedor : entero
    res_producto : entero

    total_producto : real
    total_vendedor : real
    total_comision : real

    mes_actual : N(2)
    año_actual : N(4)

    nodo = registro
        codvendedor : entero
        comision    : real
        prox        : puntero a nodo
    FinRegistro

    Prim, p, q, a : puntero a nodo

--------------------------------------------------

Procedimiento corte_producto() ES
    Escribir("Producto ", res_producto,
             " total vendido ", total_producto)

    total_vendedor := total_vendedor + total_producto
    total_producto := 0
    res_producto   := reg_v.codproducto
FinProcedimiento

--------------------------------------------------

Procedimiento carga_ordenada_ranking(vend : entero, com : real) ES
    Nuevo(p)
    p.codvendedor := vend
    p.comision    := com
    p.prox        := NIL

    a := NIL
    q := Prim

    Mientras (q <> NIL) Y (q.comision > com) hacer
        a := q
        q := q.prox
    FinMientras

    Si a = NIL entonces
        p.prox := Prim
        Prim   := p
    Sino
        p.prox := q
        a.prox := p
    FinSi
FinProcedimiento

--------------------------------------------------

Procedimiento corte_vendedor() ES
    corte_producto()

    Escribir("Vendedor ", res_vendedor,
             " total vendido ", total_vendedor,
             " comision ", total_comision)

    carga_ordenada_ranking(res_vendedor, total_comision)

    total_vendedor := 0
    total_comision := 0
    res_vendedor   := reg_v.codvendedor
FinProcedimiento

--------------------------------------------------

Procedimiento mostrar_ranking() ES
Ambiente
    aux : puntero a nodo
    pos : entero
Proceso
    Escribir("RANKING DE VENDEDORES POR COMISION")
    pos := 1
    aux := Prim

    Mientras aux <> NIL hacer
        Escribir(pos, ") Vendedor ", aux.codvendedor,
                 " Comision ", aux.comision)
        pos := pos + 1
        aux := aux.prox
    FinMientras
FinProcedimiento

--------------------------------------------------

Proceso
    mes_actual := MES_ACTUAL
    año_actual := AÑO_ACTUAL

    Abrir E/(arch_v)
    Leer(arch_v, reg_v)

    total_producto := 0
    total_vendedor := 0
    total_comision := 0

    Prim := NIL

    res_vendedor := reg_v.codvendedor
    res_producto := reg_v.codproducto

    Mientras NFDA(arch_v) hacer

        Si res_vendedor <> reg_v.codvendedor entonces
            corte_vendedor()
        Sino
            Si res_producto <> reg_v.codproducto entonces
                corte_producto()
            FinSi
        FinSi

        Si (reg_v.fecha_venta.mes = mes_actual) Y (reg_v.fecha_venta.año = año_actual) entonces
            total_producto := total_producto + (reg_v.cantidad * reg_v.precio)

            total_comision := total_comision +(reg_v.cantidad * reg_v.precio) * reg_c.comision
        FinSi

        Leer(arch_v, reg_v)
    FinMientras
	
    mostrar_ranking()

    Cerrar(arch_v)
FinProceso
FinAccion
