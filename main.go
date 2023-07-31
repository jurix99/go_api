package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// Fonction qui effectue une tâche et renvoie un résultat
func processTask(taskID int) string {
	// Ici, nous simulons le temps d'exécution d'une tâche en dormant pendant quelques secondes
	time.Sleep(2 * time.Second)
	return fmt.Sprintf("Tâche %d terminée", taskID)
}

func main() {
	// Crée une nouvelle instance du routeur Gin
	r := gin.Default()

	// Route GET pour l'endpoint "/api/tasks"
	r.GET("/api/tasks/parallele", func(c *gin.Context) {
		// Liste des tâches à effectuer
		tasks := []int{1, 2, 3, 4, 5}
		
		println("Parallele started ...")

		// Utilisation d'une WaitGroup pour attendre que toutes les goroutines se terminent
		var wg sync.WaitGroup

		// Crée un channel pour stocker les résultats de chaque goroutine
		results := make(chan string, len(tasks))

		// Démarre une goroutine pour chaque tâche
		for _, taskID := range tasks {
			wg.Add(1)
			go func(id int) {
				// Lorsque la goroutine se termine, signale à la WaitGroup qu'elle est terminée
				defer wg.Done()

				// Appelle la fonction processTask pour effectuer la tâche et envoie le résultat au channel
				result := processTask(id)
				results <- result
			}(taskID)
		}

		// Attends que toutes les goroutines se terminent
		wg.Wait()

		// Ferme le channel des résultats car nous n'ajoutons plus de valeurs
		close(results)

		// Récupère les résultats des goroutines et les ajoute à un slice
		var response []string
		for result := range results {
			response = append(response, result)
		}

		// Renvoie la réponse JSON contenant les résultats des tâches
		c.JSON(http.StatusOK, gin.H{"parallele_tasks": response})
	})

	r.GET("/api/tasks/sequential", func(c *gin.Context) {
		// Liste des tâches à effectuer
		tasks := []int{1, 2, 3, 4, 5}
		
		println("Sequential started ...")
		// Slice pour stocker les résultats des tâches
		var response []string

		// Effectue les tâches séquentiellement
		for _, taskID := range tasks {
			result := processTask(taskID)
			response = append(response, result)
		}

		// Renvoie la réponse JSON contenant les résultats des tâches
		c.JSON(http.StatusOK, gin.H{"sequential_tasks": response})
	})

	// Lance le serveur sur le port 8080
	r.Run(":8080")
}
