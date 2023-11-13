/**
 * Author: Manuel Millefiori
 * Date: 2023-11-10
 */

package main

// Importo la libreria per utilizzare
// le API di telegram
import (
    "os"
    "io/ioutil"
    "log"
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Main
func main() {

    bot, err := tgbotapi.NewBotAPI("6768818695:AAFWCRWYqI0eEL809sxVEJ25HQa8G2rQ3Nc")
    if err != nil {
        log.Panic(err)
    }

    // Disabilito il debug
    bot.Debug = false

    // DEBUG
    log.Printf("Bot avviato: %s", bot.Self.UserName)

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    // Istanza per restare in attesa degli updates
    // del bot
    updates := bot.GetUpdatesChan(u)

    // Ciclo per la gestione degli updates del bot
    for update := range updates {

        // Se otteniamo un messaggio
        if update.Message != nil {

            // Definisco il comando
            if (update.Message.Text == "/start") {

                // Istanzio il messaggio da inviare
                msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ciao " + update.Message.From.FirstName + " questo è il bot delle frasi del giorno ideato e sviluppato da Doppia M e Lino" + "\U0001F4AB" + "\nUtilizza il comando /frase per mostrare la frase del giorno" + "\U0001F632")

                // Invio il messaggio di benvenuto
                bot.Send(msg)
            } else if (update.Message.Text == "/frase") {

                // Ottengo la frase del giorno

                // Apro il file in modalità di sola lettura
                file, err := os.Open("txts/frase_del_giorno.txt")

                // Gestione degli errori
                if err != nil {
                    log.Printf("Errore nell'apertura del file: ", err)
                } else {

                    // Leggo il contenuto del file
                    data, err := ioutil.ReadAll(file)

                    // Gestione degli errori
                    if err != nil {
                        log.Printf("Errore nella lettura del file: ", err)
                    } else {
                        // Istanzio il messaggio della frase del giorno da inviare
                        msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(data))

                        // Invio la frase del giorno
                        bot.Send(msg)
                    }
                }

                // Chiudo il file
                file.Close();
            }
        }
    }
}
