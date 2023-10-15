# Restaurant API

Este proyecto es una API para gestionar información de restaurantes. Utiliza Docker para su ejecución, lo que hace que sea fácil de configurar y ejecutar en diferentes entornos.

## Requisitos

Asegúrate de tener Docker instalado en tu sistema antes de seguir las instrucciones a continuación. Puedes descargar Docker desde [la página oficial de Docker](https://www.docker.com/get-started).

## Configuración de Variables de Entorno

Antes de ejecutar la aplicación, puedes configurar las siguientes variables de entorno para personalizar su comportamiento:

- `SCOPE`: Define el alcance del entorno de la aplicación. Para producción, establece `SCOPE` en `prod`. Esto puede afectar la configuración de la base de datos y otras características.

- `HOST`: Especifica la dirección IP en la que la aplicación escuchará las solicitudes. El valor predeterminado es `0.0.0.0`, lo que permite que la aplicación escuche en todas las interfaces de red.

- `PORT`: Define el puerto en el que la aplicación escuchará las solicitudes HTTP. El valor predeterminado es `8080`.

Puedes configurar estas variables de entorno antes de construir y ejecutar el contenedor de Docker. Por ejemplo, si deseas cambiar el puerto al 9090, puedes configurar la variable `PORT=9090`.

## Instrucciones para construir y ejecutar

Sigue estos pasos para construir y ejecutar la aplicación:

1. Clona o descarga el repositorio de "restaurant-api" en tu máquina local.

   ```shell
   git clone https://github.com/SocarComunica/restaurant-api.git
   ```

2. Navega al directorio del proyecto.

   ```shell
   cd restaurant-api
   ```

3. Configura las variables de entorno según tus necesidades en un archivo `.env`.

4. Construye la imagen de Docker para el proyecto utilizando el siguiente comando:

   ```shell
   docker build -t restaurant-api .
   ```

5. Una vez que la imagen se haya construido con éxito, puedes ejecutar el contenedor de Docker con el siguiente comando:

   ```shell
   docker run -p 8080:8080 restaurant-api
   ```

   Esto ejecutará la API del restaurante en un contenedor Docker y la hará accesible en `http://localhost:8080`.

## Uso de la API

Puedes acceder a la API del restaurante en `http://localhost:8080`. Asegúrate de verificar la documentación de la API o los puntos finales disponibles en la aplicación para conocer las rutas y funcionalidades específicas que ofrece.

## Detener el contenedor

Para detener el contenedor en ejecución, abre una nueva terminal y ejecuta el siguiente comando:

```shell
docker stop $(docker ps -a -q --filter ancestor=restaurant-api)
```

Esto detendrá el contenedor de la API del restaurante.

## Contribuciones

Si deseas contribuir a este proyecto, por favor consulta nuestras pautas de contribución y envía solicitudes de extracción a través de GitHub.

¡Disfruta utilizando la API del restaurante!