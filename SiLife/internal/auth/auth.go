package auth

import (
    "log"
    "os"

    "github.com/joho/godotenv"
    // "github.com/gorilla/sessions"
    "github.com/markbates/goth"
    "github.com/markbates/goth/providers/github"
    // "github.com/shareed2k/goth_fiber"
)

const (
    key = "Kuctayrop]ZafjagEbIotheadraz"
    MaxAge = 60 * 60 * 24 * 30
    IsProd = false
)

func NewAuth() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    githubClientID := os.Getenv("GITHUB_CLIENT_ID")
    githubClientSecret := os.Getenv("GITHUB_CLIENT_SECRET")

    // store := sessions.NewStore([]byte(key))
    // store.MaxAge(MaxAge)
    //
    // store.Options.Path = "/"
    // store.Options.HttpOnly = true
    // store.Options.Secure = IsProd
    //
    // goth_fiber.SessionStore = store

    goth.UseProviders(
        github.New(githubClientID, githubClientSecret, "http://127.0.0.1:4000/auth/github/callback"),
    )
}
