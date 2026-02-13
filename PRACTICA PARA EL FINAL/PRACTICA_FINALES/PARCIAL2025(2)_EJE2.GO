ACCION ES
AMBIENTE
    sec : secuencia de caracter
    v : caracter

    letra = registro
        dato : AN(2)
        prox : puntero a letra
    finregistro

    prim, p, q, ult : puntero a letra

    salida = registro
        dato : AN(30)
        nivel : AN(30)
        prox : puntero a salida
    finregistro

    prim2, p2, q2 : puntero a salida

    ARR : arreglo [1..4] de entero

    c, cl, ce, vocales, consonantes,menor ,rep,total_claves, claves_repetidas: entero
    solo_numero, solo_letra,hay_repetido : booleano
	

	funcion es_entero(x: caracter) -> booleano
		segun x hacer
			="0" : es_entero := verdadero
			="1" : es_entero := verdadero
			="2" : es_entero := verdadero
			="3" : es_entero := verdadero
			="4" : es_entero := verdadero
			="5" : es_entero := verdadero
			="6" : es_entero := verdadero
			="7" : es_entero := verdadero
			="8" : es_entero := verdadero
			="9" : es_entero := verdadero
			=otro : es_entero := falso
		finsegun
	finfuncion

	funcion es_letra_minuscula(c: caracter) -> booleano
		segun c hacer
			="a" : es_letra_minuscula := verdadero
			="b" : es_letra_minuscula := verdadero
			="c" : es_letra_minuscula := verdadero
			="d" : es_letra_minuscula := verdadero
			="e" : es_letra_minuscula := verdadero
			="f" : es_letra_minuscula := verdadero
			="g" : es_letra_minuscula := verdadero
			="h" : es_letra_minuscula := verdadero
			="i" : es_letra_minuscula := verdadero
			="j" : es_letra_minuscula := verdadero
			="k" : es_letra_minuscula := verdadero
			="l" : es_letra_minuscula := verdadero
			="m" : es_letra_minuscula := verdadero
			="n" : es_letra_minuscula := verdadero
			="o" : es_letra_minuscula := verdadero
			="p" : es_letra_minuscula := verdadero
			="q" : es_letra_minuscula := verdadero
			="r" : es_letra_minuscula := verdadero
			="s" : es_letra_minuscula := verdadero
			="t" : es_letra_minuscula := verdadero
			="u" : es_letra_minuscula := verdadero
			="v" : es_letra_minuscula := verdadero
			="w" : es_letra_minuscula := verdadero
			="x" : es_letra_minuscula := verdadero
			="y" : es_letra_minuscula := verdadero
			="z" : es_letra_minuscula := verdadero
			=otro : es_letra_minuscula := falso
		finsegun
	finfuncion

	funcion vocal(x: caracter) -> booleano
		segun x hacer
			="a" : vocal := verdadero
			="e" : vocal := verdadero
			="i" : vocal := verdadero
			="o" : vocal := verdadero
			="u" : vocal := verdadero
			=otro : vocal := falso
		finsegun
	finfuncion

	procedimiento insertar() es
		si prim = nil entonces
			prim := q
			*q.prox := nil
		sino
			p := prim
			mientras p <> nil hacer
				ult := p
				p := *p.prox
			finmientras
			*ult.prox := q
			*q.prox := nil
		finsi
	finprocedimiento

	procedimiento eliminar() es
		p := prim
		mientras p <> nil hacer
			prim := *p.prox
			disponer(p)
			p := prim
		finmientras
	finprocedimiento

	procedimiento crear(x: caracter) es
		nuevo(q)
		*q.dato := x
		c := c + 1
		insertar()
	finprocedimiento

	funcion es_simbolo(c: caracter) -> booleano
		segun c hacer
			="(" : es_simbolo := verdadero
			="#" : es_simbolo := verdadero
			="@" : es_simbolo := verdadero
			="%" : es_simbolo := verdadero
			="&" : es_simbolo := verdadero
			=otro : es_simbolo := falso
		finsegun
	finfuncion


	procedimiento nivel_1() es
		si c < 6 entonces
			p := prim
			solo_numero := verdadero
			solo_letra := verdadero

			mientras p <> nil hacer
				si no es_entero(*p.dato) entonces
					solo_numero := falso
				finsi
				si no es_letra_minuscula(*p.dato) entonces
					solo_letra := falso
				finsi
				p := *p.prox
			finmientras

			si solo_numero o solo_letra entonces
				ARR[1] := ARR[1] + 1
				p := prim
				mientras p <> nil hacer
					nuevo(q2)
					*q2.dato := *p.dato
					*q2.nivel := "nivel 1"
					insertar(q2, prim2, p2)
					p := *p.prox
				finmientras
			finsi
			eliminar()
		finsi
	finprocedimiento

	procedimiento nivel_2() es
		si c >= 6 y c <= 8 entonces
			cl := 0
			ce := 0
			p := prim

			mientras p <> nil hacer
				si es_entero(*p.dato) entonces
					ce := ce + 1
				finsi
				si es_letra_minuscula(*p.dato) entonces
					cl := cl + 1
				finsi
				p := *p.prox
			finmientras

			si cl > ce y ce > 0 entonces
				ARR[2] := ARR[2] + 1
				p := prim
				mientras p <> nil hacer
					nuevo(q2)
					*q2.dato := *p.dato
					*q2.nivel := "nivel 2"
					insertar(q2, prim2, p2)
					p := *p.prox
				finmientras
			finsi
			eliminar()
		finsi
	finprocedimiento

	procedimiento nivel_3() es
		si c >= 8 y c <= 10 entonces
			vocales := 0
			consonantes := 0
			ce := 0
			p := prim

			mientras p <> nil hacer
				si es_entero(*p.dato) entonces
					ce := ce + 1
				finsi
				si es_letra_minuscula(*p.dato) entonces
					si vocal(*p.dato) entonces
						vocales := vocales + 1
					sino
						consonantes := consonantes + 1
					finsi
				finsi
				p := *p.prox
			finmientras

			si vocales > consonantes entonces
				ARR[3] := ARR[3] + 1
				p := prim
				mientras p <> nil hacer
					nuevo(q2)
					*q2.dato := *p.dato
					*q2.nivel := "nivel 3"
					insertar(q2, prim2, p2)
					p := *p.prox
				finmientras
			finsi
			eliminar()
		finsi
	finprocedimiento
	procedimiento nivel_4() es
		si c > 10 entonces
			tiene_letra := falso
			tiene_numero := falso
			tiene_simbolo := falso

			p := prim
			mientras p <> nil hacer
				si es_letra_minuscula(*p.dato) entonces
					tiene_letra := verdadero
				finsi

				si es_entero(*p.dato) entonces
					tiene_numero := verdadero
				finsi

				si es_simbolo(*p.dato) entonces
					tiene_simbolo := verdadero
				finsi

				p := *p.prox
			finmientras

			si tiene_letra y tiene_numero y tiene_simbolo entonces
				ARR[4] := ARR[4] + 1
				p := prim
				mientras p <> nil hacer
					nuevo(q2)
					*q2.dato := *p.dato
					*q2.nivel := "nivel 4"
					insertar(q2, prim2, p2)
					p := *p.prox
				finmientras
			finsi

			eliminar()
		finsi
	finprocedimiento
	procedimiento verificar_repetidos() es
		hay_repetido := falso
		p := prim

		mientras p <> nil y no hay_repetido hacer
			rep := 0
			q := p

			mientras q <> nil hacer
				si *q.dato = *p.dato entonces
					rep := rep + 1
				finsi
				q := *q.prox
			finmientras

			si rep > 3 entonces
				hay_repetido := verdadero
			finsi

			p := *p.prox
		finmientras

		si hay_repetido entonces
			claves_repetidas := claves_repetidas + 1
		finsi
	finprocedimiento
	procedimeinto buscar() es 
		para i:=1 a 4 hacer 
			si ARR[i] < menor entonces 
				max:=ARR[i]
				nivel:=i
			finsi 
		finpara
	finprocedimeinto


	PROCESO
		arr(sec)
		avz(sec, v)

		prim := nil
		prim2 := nil
		c := 0
		menor:=HV	

		mientras NFDS(sec) hacer
			mientras v <> "+" hacer
				crear(v)
				avz(sec, v)
			finmientras
			total_claves := total_claves + 1
			verificar_repetidos()
			nivel_1()
			nivel_2()
			nivel_3()
			nivel_4()
			c := 0
			avz(sec, v)
		finmientras
	FINPROCESO
	buscar()
	porcentaje := (claves_repetidas * 100) / total_claves
	esc("Porcentaje de contraseñas con repetición mayor a 3: ", porcentaje)
	esc("el nivel", nivel ,"tiene la menor cantidad con ", menor)
	p2:=prim2
	mientras p2<> nil hacer 
		esc("clave de nivel", *p2.nivel)
		esc(*p2.dato)
		p2:=*p2.prox
	finmientras
	

FINACCION

