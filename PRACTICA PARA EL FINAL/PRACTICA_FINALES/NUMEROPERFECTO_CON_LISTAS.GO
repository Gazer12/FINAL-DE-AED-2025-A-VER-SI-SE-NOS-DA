accion es 
	ambiente 
		lista=regsitro 
			dato:N(4)
			prox,ant:puntero a lista 
		finregsitro 
		q,p,prim:puntero a lista 


		funcion pepe(x:entero,d:entero,suma:entero)->Booleano 
			si (d=0) entonces 
				si (suma=x) entonces 
					pepe:=verdadero 
				sino 
					pepe:=falso 
				finsi 
			sino 
				si ((x MOD d)= 0) entonce 
					suma:=suma + d 
					pepe(x,d-1,suma)
				sino 
					pepe(x,d-1,suma)
				finsi 
			finsi 
		finfuncion 
	procedimiento eliminar() es 
		si p=prim entonces 
			prim:=*p.prox
			*prim.ant:=nil 
		sino 
			si p=ult entonces 
				ult:=*p.ant 
				*ult.prox:=nil 
			sino 
				*(*p.prox).ant:=*p.ant 
				*(*p.ant).prox:=*p.prox 
			finsi 
		finsi 
		disponer(p)
	finprocedimiento 
	

	proceso 
		p:=prim 
		mientras (p<>nil) hacer 
			si pepe(*p.dato,*p.dato - 1, 0) entonces 
				eliminar()
				p:=*p.prox
			sino 
				p:=*p.prox
			finsi 
		finmientas 
	finproceso 
finaccion

