package register

import (
	"encoding/json"
	"fmt"
	"main/pkg/db/sqlite"
	Struct "main/pkg/struct"
	responses "main/pkg/untils/Responses"
	Errors "main/pkg/untils/error"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func MiddlewareRegister(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err, msg = Register(w, r)
		if err {
			next.ServeHTTP(w, r)
			return
		}
		fmt.Println("err : ", msg)
		Errors.SendError(w, r, http.StatusOK, msg)
		// Si l'utilisateur n'est pas authentifié, renvoyer une réponse d'erreur 401 Unauthorized
		//http.Error(w, msg, http.StatusUnauthorized)
	})
}

func Register(w http.ResponseWriter, r *http.Request) (bool, string) {
	var newUser Struct.User
	//recuperation des données du register
	decoder := json.NewDecoder(r.Body)
	errs := decoder.Decode(&newUser)
	if errs != nil {
		fmt.Println("Erreur de décodage JSON")
		//http.Error(w, "Erreur de décodage JSON", http.StatusBadRequest)
		return false, "Erreur de decodage JSON"
	}
	Alluser := sqlite.GetAllUser()
	var err, msg = UserUnique(Alluser, newUser.Email)
	if !err {
		fmt.Println("l'erreur est ici")
		return err, msg
	}
	//AJouter l'utilisateur a la BD
	newUser.Password = CryptPassword(newUser.Password)
	if len(newUser.Nickname) == 0 {
		newUser.Nickname = generateRandomUsername(newUser.Firstname + newUser.Lastname)
	}
	sqlite.CreateNewUser(newUser)
	if newUser.Avatar != "" {
		SaveImage(newUser.Avatar, newUser.AvatarData, "avatars")
	}
	fmt.Println(newUser.Nickname," is registered !")
	return err, msg
}
func SaveImage(filename string, img []byte, route string) {
	var src = "../src/assets/images/"+route
	var filepaths = filepath.Join(src, filename)
	err := os.WriteFile(filepaths, img, 0644)
	if err != nil {
		fmt.Println("l'erreur viens de la creation d'image : ", err)
		return
	}
}
func UserUnique(Alluser []Struct.User, Email string) (bool, string) {
	for _, v := range Alluser {
		if v.Email == Email {
			return false, "Email is already used"
		}
	}
	return true, "It's valid"
}

var RegisterHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// Envoyez la réponse JSON
	w.Header().Set("Content-Type", "application/json") // Définir le type de contenu de la réponse comme JSON
	responses.SendResponsesHome(w, r, "Your account is created", nil)
})

func CryptPassword(password string) string {
	var convertByte = []byte(password)
	var Hashed, err = bcrypt.GenerateFromPassword(convertByte, 10)
	if err != nil {
		fmt.Println("Erreur lors de la generation de hachage : ", err)
		return "error"
	}
	hash := string(Hashed)
	return hash
}

// Fonction pour générer un nom d'utilisateur aléatoire avec un nombre aléatoire entre 0 et 99999
func generateRandomUsername(baseUsername string) string {
	// Initialiser le générateur de nombres aléatoires
	rand.Seed(time.Now().UnixNano())

	// Générer un nombre aléatoire entre 0 et 99999
	randomNumber := rand.Intn(100000)

	// Créer le nom d'utilisateur en ajoutant le nombre aléatoire à la base
	randomUsername := fmt.Sprintf("%s%d", baseUsername, randomNumber)

	return randomUsername
}
