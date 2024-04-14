package model
type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
    Role     string  `json:"role"`
    
}

type Product struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    ImageURL string `json:"imageUrl"`
    Price    int    `json:"price"`
    Speccs   string  `json:"speccs"`
}