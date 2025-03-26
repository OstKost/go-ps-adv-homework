package main

import (
	"bytes"
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"go-ps-adv-homework/internal/carts"
	"go-ps-adv-homework/internal/products"
	"go-ps-adv-homework/internal/sessions"
	"go-ps-adv-homework/internal/users"
	"go-ps-adv-homework/pkg/di"
	"go-ps-adv-homework/pkg/jwt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func migrateDB(db *gorm.DB) {
	err := db.AutoMigrate(
		&users.User{},
		&sessions.Session{},
		&di.Order{},
		&di.OrderItem{},
		&products.Product{},
		&carts.CartItem{},
		&carts.Cart{},
	)
	if err != nil {
		panic(err)
	}
}

func initDb() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//migrateDB(db)

	return db
}

func createUser(db *gorm.DB, phone string, name string) *users.User {
	testUser := users.User{
		Phone: phone,
		Name:  name,
	}
	db.Create(&testUser)
	return &testUser
}

func createSession(db *gorm.DB, userID uint, phone string, session string, code string) *sessions.Session {
	testSession := sessions.Session{
		UserID:  userID,
		Phone:   phone,
		Session: session,
		Code:    code,
	}
	db.Create(&testSession)
	return &testSession
}

func createProduct(db *gorm.DB, name string, description string, images []string, price float64) products.Product {
	testProduct := products.Product{
		Name:        name,
		Description: description,
		Images:      images,
		Price:       price,
	}
	db.Create(&testProduct)
	return testProduct
}

func initData(db *gorm.DB) {
	testData := &struct {
		Phone   string
		Name    string
		Session string
		Code    string
	}{
		Phone:   "+7911111111",
		Name:    "Test User",
		Session: "0000000000",
		Code:    "1234",
	}
	testProducts := []products.Product{
		{
			Name:        "Test Product 1",
			Description: "Test Product Description 1",
			Images:      []string{"/image1.png", "/image2.png"},
			Price:       100,
		},
		{
			Name:        "Test Product 2",
			Description: "Test Product Description 2",
			Images:      []string{"/image2.png", "/image3.png"},
			Price:       200,
		},
	}
	// User
	testUser := createUser(db, testData.Phone, testData.Name)
	// Session
	createSession(db, testUser.ID, testData.Phone, testData.Session, testData.Code)
	// Products
	for _, testProduct := range testProducts {
		createProduct(db, testProduct.Name, testProduct.Description, testProduct.Images, testProduct.Price)
	}
}

func removeData(db *gorm.DB) {
	// User
	db.Unscoped().
		Where("phone = ?", "+7911111111").
		Delete(&users.User{})
	// Session
	db.Unscoped().
		Where("session = ?", "0000000000").
		Delete(&sessions.Session{})
	// Order Items
	db.Unscoped().
		Where("order_id IS NOT NULL").
		Delete(&di.OrderItem{})
	// Order
	db.Unscoped().
		Where("id IS NOT NULL").
		Delete(&di.Order{})
	// Products
	db.Unscoped().
		Where("name LIKE ?", "Test Product%").
		Delete(&products.Product{})
}

func TestCreateOrderSuccess(t *testing.T) {
	db := initDb()
	initData(db)
	t.Cleanup(func() { removeData(db) })

	ts := httptest.NewServer(App())
	t.Cleanup(func() { ts.Close() })

	var testProducts []products.Product
	db.Find(&testProducts)
	var testItems []di.OrderItem
	for _, testProduct := range testProducts {
		testItems = append(testItems, di.OrderItem{
			ProductID: testProduct.ID,
			Count:     1,
			Price:     testProduct.Price,
		})
	}

	data, _ := json.Marshal(&di.CreateOrderRequest{Items: testItems})

	// JWT Token
	secret := os.Getenv("SECRET")
	token, err := jwt.NewJWT(secret).SignToken(jwt.JWTData{
		Phone:   "+7911111111",
		Session: "0000000000",
	})
	require.NoError(t, err)

	req, err := newRequest("POST", ts.URL+"/orders", data, token)
	require.NoError(t, err)

	client := &http.Client{}
	res, err := client.Do(req)
	require.NoError(t, err)
	t.Cleanup(func() { res.Body.Close() })

	require.Equal(t, http.StatusCreated, res.StatusCode, "unexpected status code")

	body, err := io.ReadAll(res.Body)
	require.NoError(t, err)

	var resData di.Order
	err = json.Unmarshal(body, &resData)
	require.NoError(t, err)

	// Check Order is created
	require.NotEqual(t, 0, resData.ID, "Order ID should not be empty")

	// Check Order Items
	require.Len(t, resData.Items, len(testItems), "unexpected items count")

	// Check total price with order items
	totalPrice := 0.0
	for _, item := range resData.Items {
		totalPrice += item.Price
	}
	require.Equal(t, totalPrice, resData.Total, "total price mismatch")

	// Check total price with test items
	totalPrice = 0.0
	for _, item := range testItems {
		totalPrice += item.Price
	}
	require.Equal(t, totalPrice, resData.Total, "total price mismatch")

	// Check right order items
	for i, item := range resData.Items {
		require.Equal(t, testItems[i].ProductID, item.ProductID, "product ID mismatch")
	}
}

func newRequest(method, url string, data []byte, token string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	return req, nil
}
