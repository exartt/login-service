package tests

//
//import (
//	"bytes"
//	"encoding/json"
//	"github.com/gofiber/fiber/v2"
//	"golang.org/x/crypto/bcrypt"
//	"gorm.io/gorm"
//	"login-service/internal/handlers"
//	"login-service/internal/domain/model"
//	"login-service/internal/infrastructure"
//	"net/http/httptest"
//	"testing"
//)
//
//func TestSignup(t *testing.T) {
//	// Crie um objeto *fiber.Ctx mock
//	mockCtx := new(fiber.Ctx)
//
//	// Crie um objeto *gorm.DB mock
//	mockDB, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
//	if err != nil {
//		t.Fatalf("Erro ao abrir banco de dados de teste: %s", err)
//	}
//
//	// Substitua o objeto de banco de dados real por um objeto mock
//	infrastructure.DB = mockDB
//
//	// Configure o objeto mock para retornar um corpo JSON de teste
//	body := struct {
//		Email    string `json:"email"`
//		Password string `json:"password"`
//	}{
//		Email:    "test@example.com",
//		Password: "testpassword",
//	}
//
//	bodyBytes, err := json.Marshal(body)
//	if err != nil {
//		t.Fatalf("Erro ao criar corpo de solicitação de teste: %s", err)
//	}
//
//	mockCtx.Request = httptest.NewRequest("POST", "/signup", bytes.NewReader(bodyBytes))
//
//	// Execute a função Signup
//	err = handlers.Signup(mockCtx)
//
//	// Verifique se a função não retornou um erro
//	if err != nil {
//		t.Fatalf("Signup retornou um erro inesperado: %s", err)
//	}
//
//	// Verifique se o objeto *gorm.DB mock foi chamado corretamente
//	var user model.User
//	result := mockDB.First(&user, "email = ?", "test@example.com")
//	if result.Error != nil {
//		t.Fatalf("Erro ao buscar usuário criado: %s", result.Error)
//	}
//
//	if !bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("testpassword")) == nil {
//		t.Fatalf("A senha do usuário não foi criptografada corretamente")
//	}
//
//	// Limpe o objeto mock
//	infrastructure.DB = db
//}
