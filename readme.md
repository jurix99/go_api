# API en Go avec Gin - Comaraison de tâches séquentielles et parallèles

Ce projet est un exemple d'API en Go utilisant le framework Gin pour créer un serveur HTTP qui effectue des tâches de manière séquentielle ou en parallele et renvoie les résultats sous forme de JSON dans le but de démontrer la puissance du golang.

## Sommaire
- [Installation](#installation)
- [Exécution](#exécution)
- [À propos de l'API](#à-propos-de-lapi)
  - [GET /api/tasks/parallele](#get-apitasksparallele)
  - [GET /api/tasks/sequential](#get-apitaskssequential)
- [Personnalisation](#personnalisation)
- [Auteur](#auteur)

## Prérequis
 - Si Go n'est pas installé, vous pouvez le télécharger et l'installer à partir du site officiel de Go : https://golang.org/  
 - Si Git n'est pas installé, vous pouvez le télécharger et l'installer à partir du site officiel de Git : https://git-scm.com/
## Installation

Assurez-vous d'avoir Go installé sur votre système. Vous pouvez vérifier si Go est installé en exécutant la commande suivante dans votre terminal :

```bash
go version
```
Installez le package gin:
```bash
go get -u github.com/gin-gonic/gin
```

Après avoir installé Go, vous pouvez cloner ce dépôt en utilisant la commande git :

```bash
git clone https://github.com/votre-utilisateur/nom-du-depot.git
```

## Exécution

Pour exécuter le serveur, accédez au répertoire du projet et exécutez la commande suivante :

```bash
go run main.go
```

Le serveur démarrera et écoutera sur le port 8080. Vous pouvez accéder à l'API en utilisant votre navigateur ou un outil comme curl :

```
http://localhost:8080/api/tasks
```

Vous verrez les résultats des tâches s'afficher séquentiellement avec un délai d'environ 2 secondes entre chaque résultat.

## À propos de l'API

L'API a deux endpoints :

### `GET /api/tasks/parallele`
 - Exécute 5 taches de 2s en parallele en utilisant des goroutines
### `GET /api/tasks/sequential`
 - Exécute 5 taches de 2s en séquentiel

Ensuite chacun renvoie les résultats sous forme de JSON.

## Personnalisation

Vous pouvez personnaliser ce projet en ajoutant de nouveaux endpoints, en utilisant des goroutines pour des tâches concurrentes, en intégrant une base de données ou en implémentant d'autres fonctionnalités selon vos besoins.

## Auteur

Cet exemple a été développé par Julien Terrier.

