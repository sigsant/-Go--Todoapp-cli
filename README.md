# [Go] Todo App - CLI

Organizador de tareas en línea de comando que crea, elimina y guarda las notas en formato JSON.

## Modo de uso

Puede iniciar el programa dentro de la carpeta cmd mediante la compilación del programa con:

`go build -o nombredeseadodelfichero.exe`

o ejecutarlo directamente con:

`go run main.go`

En ambos casos utilizamos los siguientes argumentos:

| Argumento |  Ejemplo de uso  | Acción |
| ---- | ---- | ---- |
| -task string | -task "Programar en Go | Crea y agrega la tarea a "todo.json"|
| -list | -list |  Lista las tareas existentes en archivo JSON|
| -complete int | -complete 2 | Marca como completado la tarea del archivo JSON |
| -delete int | -delete 2 | Elimina la tarea del archivo JSON |
| -h | -h | Muestra los argumentos mencionados anteriormente|

Por defecto crea en el directorio un archivo llamado "todo.json" donde se guarda todas las tareas agregadas. Actualmente no hay opción de que el usuario pueda modificar el nombre de éste.