ACCION cifrar_cesar ES
AMBIENTE
    sec, salida : secuencia de caracter
    v : caracter
    k, i : entero

    Nodo = REGISTRO
        dato : caracter
        prox : puntero a Nodo
    FINREGISTRO

    prim, ult, p, q, j : puntero a Nodo
PROCESO
    ESC("Ingrese la clave K")
    LEER(k)

		FUNCION conv_letra (v : caracter) : entero
			='a' : conv_letra := 1
			='b' : conv_letra := 2
			='c' : conv_letra := 3
			='d' : conv_letra := 4
			='e' : conv_letra := 5
			='f' : conv_letra := 6
			='g' : conv_letra := 7
			='h' : conv_letra := 8
			='i' : conv_letra := 9
			='j' : conv_letra := 10
			='k' : conv_letra := 11
			='l' : conv_letra := 12
			='m' : conv_letra := 13
			='n' : conv_letra := 14
			='o' : conv_letra := 15
			='p' : conv_letra := 16
			='q' : conv_letra := 17
			='r' : conv_letra := 18
			='s' : conv_letra := 19
			='t' : conv_letra := 20
			='u' : conv_letra := 21
			='v' : conv_letra := 22
			='w' : conv_letra := 23
			='x' : conv_letra := 24
			='y' : conv_letra := 25
			='z' : conv_letra := 26
		FinFUNCION
		FUNCION conv_numero (v : entero) : caracter
			=1  : conv_numero := 'a'
			=2  : conv_numero := 'b'
			=3  : conv_numero := 'c'
			=4  : conv_numero := 'd'
			=5  : conv_numero := 'e'
			=6  : conv_numero := 'f'
			=7  : conv_numero := 'g'
			=8  : conv_numero := 'h'
			=9  : conv_numero := 'i'
			=10 : conv_numero := 'j'
			=11 : conv_numero := 'k'
			=12 : conv_numero := 'l'
			=13 : conv_numero := 'm'
			=14 : conv_numero := 'n'
			=15 : conv_numero := 'o'
			=16 : conv_numero := 'p'
			=17 : conv_numero := 'q'
			=18 : conv_numero := 'r'
			=19 : conv_numero := 's'
			=20 : conv_numero := 't'
			=21 : conv_numero := 'u'
			=22 : conv_numero := 'v'
			=23 : conv_numero := 'w'
			=24 : conv_numero := 'x'
			=25 : conv_numero := 'y'
			=26 : conv_numero := 'z'
		FinFUNCION
		 // 1. Crear lista circular del abecedario
    prim := NIL

    PARA i := 1 HASTA 26 HACER
        NUEVO(q)
        q.dato := conv_numero(i)
        q.prox := NIL

        SI prim = NIL ENTONCES
            prim := q
            ult := q
        SINO
            ult.prox := q
            ult := q
        FINSI
    FINPARA

    ult.prox := prim   // cierre circular

    // 2. Procesar el texto
    CREAR(salida)
    ARR(sec)

    MIENTRAS NFDS(sec) HACER
        AVZ(sec, v)

        SI ES_LETRA(v) ENTONCES
            // buscar la letra
            j := prim
            MIENTRAS j.dato <> v HACER
                j := j.prox
            FINMIENTRAS

            // avanzar K posiciones
            PARA i := 1 HASTA k HACER
                j := j.prox
            FINPARA

            GRABAR(salida, j.dato)
        SINO
            GRABAR(salida, v)
        FINSI
    FINMIENTRAS

    CERRAR(sec)
    CERRAR(salida)
FINACCION








/*
ACCION cargar_abecedario_circular (REF Prim : puntero a Nodo) ES
AMBIENTE
    p, ult : puntero a Nodo
    letra  : caracter
PROCESO
    Prim := NIL

    PARA letra DESDE 'A' HASTA 'Z' HACER
        NUEVO(p)
        *p.Dato := letra
        *p.Prox := NIL

        SI Prim = NIL ENTONCES
            Prim := p
            ult := p
        SINO
            *ult.Prox := p
            ult := p
        FINSI
    FINPARA

    // 🔴 convertir en lista circular
    *ult.Prox := Prim
FINACCION
aca el ult es el equivalete a usar p 
Prim → A → B → C → ... → Z → NIL
                              ↑
                             ult
A → B → C → ... → Z
↑                 ↓
└─────────────────┘

para buscar dentor de la lista 
Si sabés que la lista tiene t elementos (ej. t = 5):

n := n MOD t


Ejemplo:

9 MOD 5 = 4
avazar un lista circular 
p := Prim        // empiezo en A

PARA i DESDE 1 HASTA 9 HACER
    p := p.Prox
FINPARA

ESC(p.Dato)      // muestro el carácter donde caí



