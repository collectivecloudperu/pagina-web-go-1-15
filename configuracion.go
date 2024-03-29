package main
 
import (
	"fmt"
	"net/http"
)

// Función main() 
func main() { 	
    
    /* Rutas */
    http.HandleFunc("/", rootHandler)
    http.HandleFunc("/nosotros", nosotrosHandler)
    http.HandleFunc("/servicios", serviciosHandler)
    http.HandleFunc("/contacto", contactoHandler)

    // Servimos Archivos Estáticos (CSS, Imágenes, JS, etc.)
    http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public")))) 

    fmt.Printf("Servidor corriendo en http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
 

// Función que apunta a la página Home 
func rootHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./public/index.html")
}

// Función que apunta a la página Nosotros 
func nosotrosHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./public/nosotros.html")
}

// Función que apunta a la página Servicios 
func serviciosHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./public/servicios.html")
}

// Función que apunta a la página Contacto 
func contactoHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./public/contacto.html")
}

