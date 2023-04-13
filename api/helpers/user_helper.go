package helper

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	supa "github.com/nedpals/supabase-go"
	model "github.com/omjogani/bookmark-manager/models"
)

var supabase *supa.Client

func init() {
	envError := godotenv.Load(".env")
	checkNilError(envError)
	fmt.Println("Connection with Database Established Successfully.")

	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	supabase = supa.CreateClient(supabaseUrl, supabaseKey)
}

func RegisterUserHelper(user model.User) {
	_, err := supabase.Auth.SignUp(context.Background(), supa.UserCredentials{
		Email:    user.Email,
		Password: user.Password,
	})
	checkNilError(err)
}

func LoginUserHelper(user model.User) {
	_, err := supabase.Auth.SignIn(context.Background(), supa.UserCredentials{
		Email:    user.Email,
		Password: user.Password,
	})
	checkNilError(err)
}

func checkNilError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
