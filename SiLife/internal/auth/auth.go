package auth

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

    store := sessions.NewCookieStore([]byte(key))
    store.MaxAge(MaxAge)

    store.Options.Path = "/"
    store.Options.HttpOnly = true
    store.Options.Secute = IsProd

    gothic.Store = store

    // TODO: Add Provider info
    goth.UseProviders(
        github.New(githubClientID, githubClientSecret, "http://localhost:4000/auth/github/callback"),
    )
}
