package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/Kaukov/gopher-translator/handlers"
	"github.com/Kaukov/gopher-translator/middleware"
)

func main() {
	port := flag.Int("port", 3000, "The desired port for running the API")
	logger := log.New(os.Stdout, "[gopher-translator]\t", log.LstdFlags)
	storageMiddleware := middleware.NewTranslatorStorage(logger)
	methodPostMw := middleware.NewMethod(http.MethodPost)
	methodGetMw := middleware.NewMethod(http.MethodGet)
	twHandler := handlers.NewTranslatorWord(logger)
	tsHandler := handlers.NewTranslatorSentence(logger)
	histHandler := handlers.NewTranslatorHistory(logger, middleware.GetStoredData())

	http.Handle("/word", methodPostMw(storageMiddleware(twHandler)))
	http.Handle("/sentence", methodPostMw(storageMiddleware(tsHandler)))
	http.Handle("/history", methodGetMw(histHandler))

	flag.Parse()

	// start the server
	go func() {
		logger.Println("Starting server on port", *port)

		// err := server.ListenAndServe()
		if err := http.ListenAndServe(":"+strconv.Itoa(*port), nil); err != nil {
			logger.Println("Error starting servor starting server: ", err.Error())
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT)

	// Block until a signal is received.
	sig := <-c
	logger.Println("Got signal:", sig)
}
