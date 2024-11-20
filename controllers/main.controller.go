package controllers

import (
	"final_api/database"
	"final_api/models" // Gantilah dengan path yang sesuai untuk model Anda
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// FilterTransactionsByDate untuk memfilter transaksi berdasarkan bulan atau hari
func FilterTransactionsByDate(ctx *gin.Context) {
	var transactions []models.Transaction
	dateFilter := ctx.Query("date") // Ambil filter tanggal dari query parameter

	if dateFilter == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Date parameter is required"})
		return
	}

	// Parse tanggal dari parameter
	parsedDate, err := time.Parse("2006-01-02", dateFilter)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
		return
	}

	// Filter berdasarkan bulan jika hanya bulan dan tahun yang diberikan
	month := ctx.Query("month")
	if month != "" {
		if err := database.DB.Where("MONTH(date) = ? AND YEAR(date) = ?", parsedDate.Month(), parsedDate.Year()).
			Find(&transactions).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"transactions": transactions})
		return
	}

	// Filter berdasarkan hari jika tanggal lengkap diberikan
	if err := database.DB.Where("DATE(date) = ?", parsedDate).
		Find(&transactions).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"transactions": transactions})
}

// Home akan mengembalikan informasi dasar tentang endpoint dan filter transaksi
func Home(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Info Endpoint",
		"Endpoints": []gin.H{
			{
				"method":      "GET",
				"url":         "/users",
				"description": "Mendapatkan semua pengguna",
				"contohInput": nil,
			},
			{
				"method":      "POST",
				"url":         "/users",
				"description": "Membuat pengguna baru",
				"contohInput": gin.H{
					"Username": "johndoe",
					"Email":    "johndoe@example.com",
					"Password": "kataSandiRahasia123",
				},
			},
			{
				"method":      "PUT",
				"url":         "/users/:id",
				"description": "Memperbarui pengguna berdasarkan ID",
				"contohInput": gin.H{
					"Username": "john_doe_diperbarui",
					"Email":    "john_diperbarui@example.com",
					"Password": "kataSandiBaru456",
				},
			},
			{
				"method":      "DELETE",
				"url":         "/users/:id",
				"description": "Menghapus pengguna berdasarkan ID",
				"contohInput": nil,
			},
			{
				"method":      "GET",
				"url":         "/transactions",
				"description": "Mendapatkan semua transaksi berdasarkan filter",
				"contohInput": gin.H{
					"category":  "Makanan",     // Filter berdasarkan kategori
					"type":      "Pengeluaran", // Filter berdasarkan tipe transaksi
					"minAmount": 100000,        // Filter berdasarkan jumlah minimum
					"maxAmount": 1000000,       // Filter berdasarkan jumlah maksimum
				},
			},
			{
				"method":      "POST",
				"url":         "/transactions",
				"description": "Membuat transaksi baru",
				"contohInput": gin.H{
					"UserID":       1,
					"Amount":       500000,
					"Type":         "Pengeluaran",
					"Category":     "Makanan",
					"Descriptions": "Makan siang di restoran (Optional)",
					"Date":         "2024-11-20T10:00:00Z",
				},
			},
			{
				"method":      "PUT",
				"url":         "/transactions/:id",
				"description": "Memperbarui transaksi berdasarkan ID",
				"contohInput": gin.H{
					"UserID":       1,
					"Amount":       450000,
					"Type":         "Pengeluaran",
					"Category":     "Makanan",
					"Descriptions": "Pengeluaran makan siang diperbarui (Optional)",
					"Date":         "2024-11-20T12:00:00Z",
				},
			},
			{
				"method":      "DELETE",
				"url":         "/transactions/:id",
				"description": "Menghapus transaksi berdasarkan ID",
				"contohInput": nil,
			},
		},
	})
}

// GetUsers akan mengembalikan daftar semua pengguna
func GetUsers(ctx *gin.Context) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

// GetTransactions akan mengembalikan daftar semua transaksi
func GetTransactions(ctx *gin.Context) {
	var transactions []models.Transaction

	// Mengambil parameter filter dari query string
	category := ctx.DefaultQuery("category", "")          // Filter berdasarkan kategori (default: kosong)
	transactionType := ctx.DefaultQuery("type", "")       // Filter berdasarkan tipe (default: kosong)
	minAmount := ctx.DefaultQuery("minAmount", "0")       // Filter berdasarkan jumlah minimum (default: 0)
	maxAmount := ctx.DefaultQuery("maxAmount", "1000000") // Filter berdasarkan jumlah maksimum (default: 1 juta)

	// Membuat query GORM untuk transaksi
	query := database.DB.Model(&models.Transaction{})

	// Menambahkan filter berdasarkan kategori jika ada
	if category != "" {
		query = query.Where("category = ?", category)
	}

	// Menambahkan filter berdasarkan tipe jika ada
	if transactionType != "" {
		query = query.Where("type = ?", transactionType)
	}

	// Menambahkan filter berdasarkan jumlah minimum jika ada
	if minAmount != "" {
		query = query.Where("amount >= ?", minAmount)
	}

	// Menambahkan filter berdasarkan jumlah maksimum jika ada
	if maxAmount != "" {
		query = query.Where("amount <= ?", maxAmount)
	}

	// Menjalankan query untuk mengambil transaksi yang sesuai dengan filter
	if err := query.Find(&transactions).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get transactions"})
		return
	}

	// Mengembalikan hasil transaksi yang sudah difilter
	ctx.JSON(http.StatusOK, gin.H{"transactions": transactions})
}

// CreateUser untuk menambahkan user baru
func SignUp(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	if err := database.DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}

// Login untuk autentikasi pengguna
func Login(ctx *gin.Context) {
	var input struct {
		Username string `json:"Username" binding:"required"`
		Password string `json:"Password" binding:"required"`
	}

	// Bind JSON dari request ke struct input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var user models.User
	// Cari user berdasarkan username
	if err := database.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Login berhasil
	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": gin.H{
		"id":       user.UserID,
		"username": user.Username,
		"email":    user.Email,
	}})
}

// CreateTransaction untuk menambahkan transaksi baru
func CreateTransaction(ctx *gin.Context) {
	var transaction models.Transaction
	if err := ctx.ShouldBindJSON(&transaction); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	if err := database.DB.Create(&transaction).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Transaction created successfully", "transaction": transaction})
}

// UpdateUser untuk memperbarui data pengguna berdasarkan ID
func UpdateUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Bind data yang baru dari request
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Update user
	if err := database.DB.Save(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "user": user})
}

// UpdateTransaction untuk memperbarui transaksi berdasarkan ID
func UpdateTransaction(ctx *gin.Context) {
	transactionID := ctx.Param("id")
	var transaction models.Transaction
	if err := database.DB.First(&transaction, transactionID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	// Bind data yang baru dari request
	if err := ctx.ShouldBindJSON(&transaction); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Update transaction
	if err := database.DB.Save(&transaction).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update transaction"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Transaction updated successfully", "transaction": transaction})
}

// DeleteUser untuk menghapus user berdasarkan ID
func DeleteUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Hapus user
	if err := database.DB.Delete(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// DeleteTransaction untuk menghapus transaksi berdasarkan ID
func DeleteTransaction(ctx *gin.Context) {
	transactionID := ctx.Param("id")
	var transaction models.Transaction
	if err := database.DB.First(&transaction, transactionID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	// Hapus transaksi
	if err := database.DB.Delete(&transaction).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete transaction"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Transaction deleted successfully"})
}
