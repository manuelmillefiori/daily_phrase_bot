/**
 * Author: Manuel Millefiori
 * Date: 2023-11-13
 */

package main

// Libs
import (
	"fmt"
	"math/rand"
	"time"
	"os"
	"io/ioutil"
	"strings"
)

func main() {

	// Imposto il timer fino alla prossima mezzanotte
	timer := time.NewTimer(time.Until(nextMidnight()))
	<- timer.C

	// Ciclo infinito
	for {
		// Genero la frase del giorno e la scrivo
		// sul file
		generateDailyPhrase()

		// Sleep per evitare problematiche
		time.Sleep(60)
		
		// Resetto il timer
		timer.Reset(24 * time.Hour)

		// Attendo la prossima mezzanotte
		<- timer.C
	}
}

/**
 * @brief
 * Funzione per la generazione della frase
 * del giorno
 */
func generateDailyPhrase() {
	// Apro il file delle frasi da inviare
	file, err := os.Open("txts/frasi_da_inviare.txt")

	// Error management
	if err != nil {
		fmt.Printf("Errore nell'apertura del file: %s", err)
	} else {

		// Leggo tutto il contenuto del file
		data, err := ioutil.ReadAll(file)

		// Error management
		if err != nil {
			fmt.Printf("Errore nell'apertura del file: %s", err)
		} else {
			
			// Split delle frasi riga per riga
			phrases := strings.Split(string(data), "\n")

			// Seme di generazione randomica
			rand.Seed(time.Now().UnixNano())

			// Preinizalizzo l'indice casuale
			randomIndex := -1

			// Seleziono casualmente una frase escludendo quelle nulle
			for {
				
				// Generazione casuale dell'indice
				randomIndex = rand.Intn(len(phrases))

				// Verifico che la stringa non sia vuota
				if phrases[randomIndex] != "" {
					// Esco dal ciclo
					break;
				}
			}

			// Ottengo la frase casuale
			randomPhrase := phrases[randomIndex]

			// Salvo la frase sul file della frase da inviare
			file, err := os.Create("txts/frase_del_giorno.txt")

			// Error management
			if err != nil {
				fmt.Printf("Errore nell'apertura del file: %s", err)
			} else {

				// Scrivo sul file della frase da inviare
				_, err = file.WriteString(strings.TrimSpace(randomPhrase))

				// Error management
				if err != nil {
					fmt.Printf("Errore nella scrittura nel file: %s", err)
				} else {
					
					// Apro il file delle frasi già inviate
					// in append
					file, err := os.OpenFile("txts/frasi_gia_inviate.txt", os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0644)

					// Error management
					if err != nil {
						fmt.Printf("Errore nell'apertura del file: %s", err)
					} else {
						
						// Scrivo sul file delle frasi già
						// inviate la frase appena sorteggiata
						_, err = file.WriteString(strings.TrimSpace(randomPhrase) + "\n")

						// Riscrivo sul file delle frasi da inviare
						// tutte le frasi senza quella sorteggiata
						file, err := os.Create("txts/frasi_da_inviare.txt")

						// Error manaagement
						if err != nil {
							fmt.Printf("Errore nell'apertura del file: %s", err)
						} else {

							// Scrivo sul file tutte le frasi ancora da inviare
							for i := 0; i < len(phrases); i++ {

								// Verifico che l'indice non corrisponda con quello
								// sorteggiato
								if i != randomIndex {
									if i == len(phrases) - 1 { 
										_, err = file.WriteString(strings.TrimSpace(phrases[i]))

										// Error management
										if err != nil { fmt.Printf("Errore nella scrittura nel file: %s", err) }
									} else {
										_, err = file.WriteString(strings.TrimSpace(phrases[i]) + "\n")

										// Error management
										if err != nil { fmt.Printf("Errore nella scrittura nel file: %s", err) }
									}
								}
							}

							fmt.Println("Frase generata!")
						}
					}
				}
			}
		}
	}
}

/**
 * @brief
 * Funzione per definire la data e il tempo
 * della prossima mezzanotte
 */
func nextMidnight() time.Time {
	// Ottengo il tempo attuale
	now := time.Now()

	// Restituisco la data e l'ora della prossima mezzanotte
	return time.Date(now.Year(), now.Month(), now.Day() + 1, 0, 0, 0, 0, time.Local)
}
