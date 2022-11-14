# Arq-Software2


Desarrollaremos un sistema de publicación de clasificados, mediante el cual las empresas inmobiliarias puedan cargar sus bases de datos con el posteo de un archivo json de la información relacionada a los inmuebles.
Los navegantes pueden buscar esos clasificados desde la home del sitio en base a una oración y traiga los resultados priorizados que permitan ver el detalle de la publicación.
Para esto, se piden desarrollar 3 microservicios:
1. búsqueda
2. publicaciones 
3. frontend

Para el servicio (1) utilizar un motor de búsqueda que permita una indexación y búsqueda de los ítems por sus características (título, descripción, atributos, ciudad, estado, etc), que se nutra mediante notificaciones del servicio de ítems y busque la información de ese servicio.
Para el servicio (2) tener un caché centralizado con los datos del item. Si la información del usuario no está en cache, tiene que solicitarla al servicio correspondiente y almacenarla para que los próximos requests se sirvan del caché. Los datos del item se pueden almacenan de forma persistente no relacionalmente. Lo importante acá es que las imágenes se sirven localmente, por lo que hay que prever una descarga concurrente asíncrona no bloqueante de las imágenes de los items provistas en la carga del catálogo inicial.
La carga de la información de los items, la realiza el administrador del sistema posteando un archivo .json con un arreglo de items, una vez que obtengamos el archivo finaliza el request y el procesamiento se realiza de forma asíncrona
Para el servicio (3) proveer un frontend que tenga la vista de inicio con el input de búsqueda, el listado de publicaciones, el detalle de la publicación.
Requerimientos no funcionales:
 
1) Proveer uno o varios balanceadores de carga para tener múltiples instancias de los servicios.
2) Utilizar github o similar para versionar el código fuente
3) Escribir la documentación en el readme del repositorio
Ej de la estructura de un item: https://api.mercadolibre.com/items/MLA1146082202