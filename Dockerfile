# Utiliza la imagen base de Golang en la versión 1.21.1.
FROM golang:1.21.1

# Establece el directorio de trabajo dentro del contenedor en /usr/src/app.
WORKDIR /usr/src/app  

# Instala la herramienta "air" para recarga en caliente (Hot Reload).
RUN go install github.com/cosmtrek/air@latest  

# Copia el contenido del directorio actual (donde se encuentra el Dockerfile) al directorio de trabajo del contenedor.
COPY . .  

# Ejecuta "go mod tidy" para gestionar las dependencias del proyecto.
RUN go mod tidy  
