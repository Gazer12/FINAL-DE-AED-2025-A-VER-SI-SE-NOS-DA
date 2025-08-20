ACCION ES 
	ambiente 
		PARTICIPANTE=regisitro 
			pais:AN(30)
			equipo:AN(30)
			diciplinas:["100m","200m","4x100","4x400"]
			tiempo:R(3)
			descalificado:["si","no"]
		finregistro 
		arch: archivo de PARTICIPANTE
		reg:PARTICIPANTE

		RELEVOS=regsitro 
			pais:AN(30)
			tiempo_total:R(3)
		finregistro
		arch_s:archivo de RELEVOS
		reg_s:relevos 

		res_pais,res_equipos:alfanumerico
		procedimeinto corte_equipos() es 
			esc("el tiempo del equipo ", res_equipos "es de:" reg.total)
			total_pais:=total_pais + total
			total:=0
			si bandera entonces 
			esc("del equipo",res_equipos "se descalifico" , cant_desc )
			cant_desc:=0
			finsi 
			si cant_disi<"160" entonces 
			res_s.pais:=res_pais
			res_s.tiempo:cant_disi
			finsi 
			cant_disi:=0
			res_equipos:=reg.equipo
		finprocedimiento 
		procedimeinto corte_pais() es 
			esc("el pais",res_pais"hizo un tiempo de",total_pais)
			total_general:total_general + total_pais
			total_pais:=0
			res_pais:=reg.res_pais
		finprocedimiento

			
		proceso
		abrir e/(arch)
		abrir s/(arch_s)
		leer(arch,reg)
		total:=0
		cant_desc:=0
		cant_disi:=0
		res_equipos:=reg.equipo
		res_pais:=reg.pais
		Mientras NFDA(arch) hacer 
			si res_pais <> reg.pais entonces
				corte_pais()
			sino 
				si res_equipos <> reg.equipo entonces 
					corte_equipos()
				finsi 
			finsi 
			si reg.descalificado="no" entonces 
				total:reg.tiempo
				si reg.diciplinas= "4x100" o "4x400" entonces 
					cant_disi:=cant_disi + reg.tiempo
				finisi
			sino 
				si reg.descalificado="si" entonces 
					bandera:=verdadero 
					cant_desc:=cant_desc + 1 
				finsi 
			finsi 
		finmientras 
		cerrar(arch)
		cerrar(arch_s)
		esc("el total general del torneo es", total_general)
	finaccion 
	
		