package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AnimalInfo struct {
	ASCII   string
	Message string
}

var animals = map[string]AnimalInfo{
	"cat": {
		ASCII: `
 /\_/\
( o.o )
 > ^ <
        `,
		Message: "The cat watches beyond the veil.",
	},
	"dog": {
		ASCII: `
  / \__
 (    @\___
 /         O
/   (_____/
/_____/   U
        `,
		Message: "The dog is loyal to the code.",
	},
	"bird": {
		ASCII: `
  \\
   (o>
\\_//)
 \_/_)
  _|_
        `,
		Message: "The bird sings in frequencies only the divine debugger hears.",
	},
	"fish": {
		ASCII: `
      ><(((('> 
        `,
		Message: "The fish swims through memory like packets in the ocean of RAM.",
	},
	"lion": {
		ASCII: `
   ,w.
 ,YWMMw  ,M  ,
__..---..__   MM
"~~*~~~*~~"   MM
              MM
        `,
		Message: "The lion roars with the authority of root access.",
	},
	"dragon": {
		ASCII: `
              /           / 
     (o )   ( o)    (o )   ( o) 
      \ \   / /      \ \   / /
       \ \_/ /        \ \_/ /
        \   /          \   /
         | |            | |
         | |            | |
        (___)          (___)
        `,
		Message: "The dragon breathes fire into your kernel space.",
	},
	"snake": {
		ASCII: `
         /^\/^\
       _|__|  O|
\/     /~     \_/ \
 \____|__________/  
        \_______/
        `,
		Message: "The snake slithers silently, like shellcode in the heap.",
	},
	"owl": {
		ASCII: `
  ,_,  
 (O,O) 
 (   ) 
  " "
        `,
		Message: "Uhmmmmmm, the owl is just, an owl ig ... ?",
	},
}

type SpellRequest struct {
	Animal string `json:"animal"`
}

func getAnimalsHandler(w http.ResponseWriter, r *http.Request) {
	keys := []string{}
	for k := range animals {
		keys = append(keys, k)
	}
	json.NewEncoder(w).Encode(keys)
}

func spellAnimalHandler(w http.ResponseWriter, r *http.Request) {

	var req SpellRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	animal, exists := animals[req.Animal]
	w.Header().Set("Content-Type", "application/json")

	if exists {
		json.NewEncoder(w).Encode(animal)
	} else {
		resp := AnimalInfo{
			ASCII: `
   (•_•) 
  <)   )╯ This animal is unknown to our servers...
   /   \  
			`,
			Message: fmt.Sprintf("'%s' is not yet in our spellbook, but it will be added soon!", req.Animal),
		}
		json.NewEncoder(w).Encode(resp)
	}
}
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Animal Spellbook Holy server, Use /animals to see available animals and POST to /spellananimal with {\"animal\": \"name\"} to get its sacred spelling.")
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/animals", getAnimalsHandler)
	mux.HandleFunc("POST /spellananimal", spellAnimalHandler)
	mux.HandleFunc("/", welcomeHandler)
	fmt.Println("Like god intended, running on the 8080")
	http.ListenAndServe(":8080", mux)
}
