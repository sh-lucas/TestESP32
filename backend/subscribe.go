package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SherClockHolmes/webpush-go"
)

var latestDoorClose = time.Now()

func CORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "https://seu-frontend.vercel.app")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func Home(w http.ResponseWriter, req *http.Request) {
	CORS(w)

	latestDoorClose = time.Now()
	fmt.Fprintln(w, "door closed!")
}

// chaves geradas pelo /keys/keys.go
const (
	vapidPublicKey  = "BD-SXHOCw74F9xM1crwo66XrEGHsiC6hSSnEJpDC7GOpd9ll_Nslaa8Tt9UP5qS8Rccw8J2hcsRC7hee-B-C5VM"
	vapidPrivateKey = "oNq4q5PXk4RwGDD_wN9QNKRFOUqHPdkSTvbm0haM270"
)

func Subscribe(w http.ResponseWriter, r *http.Request) {
	CORS(w)

	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var sub webpush.Subscription
	if err := json.NewDecoder(r.Body).Decode(&sub); err != nil {
		http.Error(w, "Corpo da requisição inválido", http.StatusBadRequest)
		return
	}

	// Guarda a inscrição para uso posterior
	subscription := &sub
	log.Println("Nova inscrição recebida!")

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Inscrição recebida com sucesso!")

	go func(sub *webpush.Subscription) {
		fmt.Println("Registrando novo inscrito!")
		for {
			// timeout
			if time.Since(latestDoorClose) > time.Second*5 {
				Alarm(sub)
				time.Sleep(time.Second * 15)
			}
			time.Sleep(1 * time.Second)
		}
	}(subscription)
}

func Alarm(sub *webpush.Subscription) {
	if sub == nil {
		panic("A inscrição é nula!!!")
	}

	notificationPayload, _ := json.Marshal(map[string]interface{}{
		"title": "Teste: Alarme!",
		"body":  "Sua porta de casa está aberta a mais de 15 segundos.",
		// "icon":  "https://gopherlabs.dev/img/gopher-reading.png", // Um ícone legal :)
	})

	// Envia a notificação
	resp, err := webpush.SendNotification(notificationPayload, sub, &webpush.Options{
		Subscriber:      "mailto:lnaoedelixo42@gmail.com", // Email para contato
		VAPIDPublicKey:  vapidPublicKey,
		VAPIDPrivateKey: vapidPrivateKey,
		TTL:             30,
	})
	if err != nil {
		panic(fmt.Sprintf("Erro ao enviar notificação: %v", err))
	}
	defer resp.Body.Close()

	log.Printf("Notificação enviada! Status: %s", resp.Status)
}
