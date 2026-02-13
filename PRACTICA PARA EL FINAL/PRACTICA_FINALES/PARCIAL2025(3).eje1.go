accion es 
	ambiente 
		lista=regsitro 
			numero:N(100)
			prox:puntero a lista 
		finregistro
		prim,a:puntero a lista 

		nodo=regsitro
			dato:=N(100)
			prox:punteor a nodo 
		finregistro
		prim1,p,q,ant:puntero a nodo 

		Funcion DigitosCrecientes(n : Entero) : Booleano
			aactual,anterior:entero
			SI n < 10 ENTONCES
				DigitosCrecientes := VERDADERO
			SINO
				actual := n MOD 10
				anterior := (n DIV 10) MOD 10

				SI anterior < actual ENTONCES
					DigitosCrecientes := DigitosCrecientes(n DIV 10)
				SINO
					DigitosCrecientes := FALSO
				FIN_SI
			FIN_SI
		FinFuncion
		procedimiento eliminar() es 
			si prim=a entonces 
				prim:=*a,prox 
				disponer(a)
			sino 
				si a=nil entonces 
					ant:=*a.prox 
					disponer(a)
				sino 
					ant:=*a.prox 
					disponer(a) 
				finsi 
			finsi 

		finprocedimiento
		procedimiento insertar() es 
			si prim=nil entonces 
				prim:=q
				*q.prox:=nil 
			sino 
				a:=prim 
				mientras a<>nil hacer 
					ant:a 
					a:=*a.prox 
				finmientras 
				si prim=a entonces 
					prim:=q
					*q.prox:=a 
				sino 
					*ant.prox:=q 
					*q.prox:=nil  
				finsi 
			finsi 	
		finprocedimiento  
		procedimiento crear(x:entero) es
			nuevo(q)
			*q.dato:=x
			incertar()
		finprocedimiento  

		proceso 
		a:=prim 
		prim1:=nil 
		ant:=nil
		mientras a<>nil hacer 
			si no DigitosCrecientes(*a.numero) entonces 
				crear(*a.numero)
				ant:=p
				a:=*a.prox
				eliminar() 
			sino 
				ant:=a
				a:=*a.prox
			finsi 
		finmientras 
finaccion 
				


 
			

				