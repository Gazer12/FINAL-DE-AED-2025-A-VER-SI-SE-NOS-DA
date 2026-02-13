accion es 
	ambiente 
		sec:secuencia de caracter
		v:caracter
		alumnos=regsitro 
			legajo:N(6)
			nomyape:AN(30)
			materia_cursada:N(10)
			año_ingreso:=N(4)
			prox:puntero a alumnos
		finregistro

		regional=regsitro 
			cant_alumno:N(100)
			cant_materia:N(100)
			lista_alumno:puntero a alumno
		finregsitro 

		arr:arreglo [1..10] de regional

		nombre_facu:caracter

		procedimiento inicializar() es 
			arr(sec)
			avz(sec,v)
			para i:=1 a 10 hacer 
				regional[i].cant_alumno:=0
				regional[i].cant_materia:=0
				regional[i].lista_alumno:=nil 
			finpara 

		finprocedimiento