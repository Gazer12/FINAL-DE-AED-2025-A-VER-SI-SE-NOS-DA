ACCION ES 
	AMBIENTE 
		fecha=registro 
			dia:N(2)
			mes:N(2)
			año:N(4)
		finregistro  
		ventas=regsitro 
			clave:registro
				codrep:1..20
				codpro:N(30)
			finregsitro 
			cantidad:N(4)
			precio:N(4)
			fechavt:fecha 
		finregistro
		arch_v:archivo de ventas ordenado por clave
		reg_v:ventas 
		repartidores=regsitro
			codrep:1..20
			nombre;AN(30)
			porc_extra:R(4)
			domicilio:AN(30)
		finregistro
		arch_r:archivo indexado de reparidores ordenado por codrep
		reg_r:repartidores
		comisiones=regsitro 
			clave=regsitro 
				codrep:1..20
				codpro:N(30)
			finregsitro
			comision:R(4)
		finregistro 
		arch_c:archivo indexado de comisiones ordenado por clave 
		reg_c:comisiones 
		producto=registro 
			codpr:N(30)
			descripcion:AN(30)
			stock:N(4)
			comisiones:N(4)
			ultcompra:fecha 
		regsitro 
		arch_p:archivo indexado de producto ordenado por codpro
		reg_p:producto
		
		nodo=registro 
			codrep:1..20
			nombre;AN(30)
			comisiones:N(4)
			prox:puntero a nodo 
		finregsitro 
		prim,q,p:puntero a nodo 

		RANKING=registro 
			codrep:1..20
			nombre;AN(30)
			comisiones:N(4)
			prox:puntero a RANKING
		finregistro 
		prim1,q1,p1:puntero a RANKING


		procedimiento inicializar() es 
			abrir e/(arch_v)
			leer(arch_v,reg_v)
			res_ventas:=reg_v.clave.codrep
			prim:=nil 
			prim1:=nil
		finprocedimiento  
		procedimiento corte_ventas() es 
			reg_r:=res_ventas//hago una busqueda directa por calve en el archivo indexado para despues usar el ranking
			leer(arch_r,reg_r)
			si (acum_prod>10000) entonces 
				total_cobrar:=acum_prod + reg_r.porc_extra
			sino 
				total_cobrar:=acum_prod
			finsi 
			esc("el repartidor, "reg_r.nombre " tiene a cobrar", total_cobrar)
			crear(reg_r.codrep,reg_r.nombre,total_cobrar)
			total_cobrar:=0 
			res_ventas:=reg_v.clave.codrep

		finprocedimiento
		procedimiento cargar_raking(q:punteor a nodo) es 
			a:=prim1
			mientras a<>nil y (*q.comision>*p.comision) hacer 
				ult:=a 
				a:=*a.prox 
			finmientras 
			si a=prim entonces 
				prim1:=q 
				*q.prox:=a 
			sino 
				si a=nil entonces 
					*ult.prox:=q 
					*q.prox:=nil 
				sino 
					*ant.prox:=q 
					*q.prox:=a
				finis 
			finsi 
		finprocedimiento
		procedimiento RANKING(q:punteor a nodo,j:AN(30),r:N(4)) es 
			si prim1=nil entoces 
				prim1:=q
				*q.prox:=nil 
			sino 
				p1:=prim 
				mientras p1<>nil y (j<>*p1.nombre) entonces 
					p1:=*p.prox 
				finmientras
				si p1=nil entonces 
					nuevo(q1)
					*q1.codrep:=*q.codrep
					*q1.nombre:=j
					*q1.comision:=r
					cargar_raking(q1)
				sino 
					*p.comision:=*p.comision + r 
				finis 
			finis 

			
			procedimiento cargar(q:puntero a nodo) es 
			si (prim=nil) entonce 
			prim:=q 
			*q.prox:=nil 
			sino
			p:=prim 
			Mientras *p.prox <> nil hacer 
			p:*p.prox 
			FinMientras
			*p.prox:=q 
			*q.prox:=nil 
			finprocedimiento 
			procedimiento crear(a:N(30),x:AN(30,t:N(4))) es
				neuvo(q)
				*q.codrep:=a 
				*q.nombre:=x 
				*q.comision:=t 
				aux:=a 
				cargar(q)
				RANKING(q,x,t)
			finprocedimiento
		proceso 
			inicializar()
			Mientras NFDA(arch_v) hacer 
				si res_ventas<>reg_v.clave.codrep entonces 
					corte_ventas()
				finsi 

				si (reg_v.fechavt.mes=mes_actual) entonces 
					acum_prod:=reg_v.cantidad * reg_v.precio
					reg_c.codrep:=reg_v.codrep
					reg_c.codpro:=reg_v.codpro
					leer(arch_c,reg_c)
					si existe entonces 
						procen_comis:=reg_c.comision
					sino 
						reg_p.codpro:=reg_v.codpro
						leer(arch_p,reg_p)
						si existen entonces 
							procen_comis:=reg_p.comision
						finsi 
					finsi 
				finsi 
				acum_comision:=acum_comision + ((reg_v.cantidad * reg_v.precio)/100)
				leer(arch_v,reg_v)
			FinMientras
			corte_ventas()
			p:=prim
			Mientras p<>nil entonces 
				esc("el reparridor " , *p.nombre "tiene una comisonde ", *p.comision)
				p:=*p.prox 
			finmientras 
		finproceso 
finaccion 

		
