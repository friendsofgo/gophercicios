# Quiz Game (Juego de preguntas)

## Autor

Este ejercicio está basado en el primer ejercicio de una serie de ejercicios realizados por [Jon Calhoun](https://twitter.com/joncalhoun)
para [Gophercises](https://gophercises.com/), puedes visitar el [ejercicio original en github](https://github.com/gophercises/quiz).

## Ejercicio

Este ejercicio se compone de dos partes, para intentar simplificar el proceso de explicación
y que resulte más sencillo de resolver.
La segunda parte es algo más difícil que la primera, con lo que si te bloqueas mejor pasar a otro tema
y volver a la parte 2 más tarde.

*Nota: Este ejercicio no se ha dividido en diferentes ejercicios, porque ambos combinados deberían de poder hacerse
en aproximadamente 30 minutos*

### Parte 1

Crear un programa que lea una hoja de preguntas y respuestas en formato [CSV](https://es.wikipedia.org/wiki/Valores_separados_por_comas),
mostraremos dichas preguntas al usuario, y recogeremos sus preguntas almacenando el total de respuestas correctas que conteste.
Responda bien o responda mal deberemos no informar de sus aciertos hasta el final y mostrar la siguiente pregunta inmediatamente después.

El fichero `CSV` deberá ser por defecto `problems.csv`, pero el usuario debe ser capaz de personalizarlo via `flag`.

Nuestro `CSV` tendrá que tener un formato como el que mostramos, donde la primera columna será la pregunta y la siguiente la respuesta.

```csv
5+5,10
7+3,10
1+1,2
8+3,11
1+2,3
8+6,14
3+1,4
1+4,5
5+1,6
2+3,5
3+3,6
2+4,6
5+2,7
```

Puedes asumir que los cuestionarios son relativamente pequeños, menos de 100 preguntas y con una sola palabra o número como respuesta.

Al final del programa debería mostrar el total de preguntas que se hayan respondido
correctamente y cuantas preguntas había en total. 

* **NOTA**: El fichero `CSV` tiene que poder tener preguntas con comas en él, ejemplo: `"¿Cuánto es 2+2, gracias?",4`.
* Os sugerimos no complicaros y utilizar el [paquete CSV del core de Go](https://pkg.go.dev/encoding/csv), en vez de intentar picar uno propio.

### Parte 2

Modificar el programa para añadir un temporizador. El tiempo límite por defecto debería ser **30 segundos**, pero debe poderse
modificar via `flag`.

El programa debe terminar inmediatamente se acabe el tiempo, eso quiere decir que no
esperaremos a que el usuario responda si ha terminado el tiempo. Igualmente no esperaremos a que
se acabe el tiempo si el usuario ha terminado antes de responder al cuestionario.

Antes de empezar a mostrar preguntas debemos de permitir al usuario pulsando cualquier tecla empezar el temporizador, entonces
empezaremos a mostrar las preguntas como hicimos en la primera parte.

Debemos mostrar el resultado final al usuario, tanto si el programa termina por alcanzar el límite de tiempo como
por haber respondido a todas las respuestas.

### Bonus

Por si os parece poco las dos partes, vamos con un bonus:

1. Añadir una limpieza de sting, eliminar espacios no necesarios, mayúsculas, etc. de las respuestas de manera que no se consideren incorrectas. *Pista, echar un ojo al paquete [strings](https://pkg.go.dev/strings)*.
2. Añadir la opción con un nuevo `flag` de ejecutar las preguntas en orden aleatorio.